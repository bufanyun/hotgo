// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"reflect"
	"sync"
)

// MsgParser 消息处理器
type MsgParser struct {
	mutex        sync.Mutex               // 路由锁
	task         RoutineTask              // 投递协程任务
	rpc          *RPC                     // rpc
	routers      map[string]*RouteHandler // 已注册的路由
	interceptors []Interceptor            // 拦截器
}

// RoutineTask 投递协程任务
type RoutineTask func(ctx context.Context, task func())

// Interceptor 拦截器
type Interceptor func(ctx context.Context, msg *Message) (err error)

// Message 标准消息
type Message struct {
	Router  string      `json:"router"`          // 路由
	TraceId string      `json:"traceId"`         // 链路ID
	Data    interface{} `json:"data"`            // 数据
	MsgId   string      `json:"msgId,omitempty"` // 消息ID，rpc用
	Error   string      `json:"error,omitempty"` // 消息错误，rpc用
}

// NewMsgParser 初始化消息处理器
func NewMsgParser(task RoutineTask) *MsgParser {
	m := new(MsgParser)
	m.task = task
	m.routers = make(map[string]*RouteHandler)
	m.rpc = NewRPC(task)
	return m
}

// RegisterRouter 注册路由
func (m *MsgParser) RegisterRouter(routers ...interface{}) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, router := range routers {
		info, err := ParseRouteHandler(router, false)
		if err != nil {
			return err
		}
		if _, ok := m.routers[info.Id]; ok {
			return gerror.Newf("server router duplicate registration:%v", info.Id)
		}
		m.routers[info.Id] = info
	}
	return
}

// RegisterRPCRouter 注册rpc路由
func (m *MsgParser) RegisterRPCRouter(routers ...interface{}) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, router := range routers {
		info, err := ParseRouteHandler(router, true)
		if err != nil {
			return err
		}
		if _, ok := m.routers[info.Id]; ok {
			return gerror.Newf("server rpc router duplicate registration:%v", info.Id)
		}
		m.routers[info.Id] = info
	}
	return
}

// RegisterInterceptor 注册拦截器
func (m *MsgParser) RegisterInterceptor(interceptors ...Interceptor) {
	m.interceptors = append(interceptors, interceptors...)
}

// Encoding 消息编码
func (m *MsgParser) Encoding(data []byte) (*Message, error) {
	var msg Message
	if err := gconv.Scan(data, &msg); err != nil {
		return nil, gerror.Newf("invalid package struct: %s", err.Error())
	}
	if msg.Router == "" {
		return nil, gerror.Newf("message is not router: %+v", msg)
	}
	return &msg, nil
}

// Decoding 消息解码
func (m *MsgParser) Decoding(ctx context.Context, data interface{}, msgId string) ([]byte, error) {
	message, err := m.doDecoding(ctx, data, msgId)
	if err != nil {
		return nil, err
	}
	return json.Marshal(message)
}

// Decoding 消息解码
func (m *MsgParser) doDecoding(ctx context.Context, data interface{}, msgId string) (*Message, error) {
	msgType := reflect.TypeOf(data)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return nil, gerror.Newf("json message pointer required: %+v", data)
	}

	message := &Message{Router: msgType.Elem().Name(), TraceId: gctx.CtxId(ctx), MsgId: msgId, Data: data}
	return message, nil
}

// handleRouterMsg 处理路由消息
func (m *MsgParser) handleRouterMsg(ctx context.Context, msg *Message) error {
	// rpc消息
	if m.rpc.Response(ctx, msg) {
		return nil
	}

	handler, ok := m.routers[msg.Router]
	if !ok {
		return gerror.NewCodef(gcode.CodeInternalError, "invalid message: %+v, or the router is not registered.", msg)
	}
	return m.doHandleRouterMsg(ctx, handler, msg)
}

// doHandleRouterMsg 处理路由消息
func (m *MsgParser) doHandleRouterMsg(ctx context.Context, handler *RouteHandler, msg *Message) (err error) {
	var input = gutil.Copy(handler.Input.Interface())
	if err = gjson.New(msg.Data).Scan(input); err != nil {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"router scan failed:%v to parse message:%v",
			err, handler.Id)
	}

	args := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(input)}

	m.task(ctx, func() {
		results := handler.Func.Call(args)
		if handler.IsRPC {
			switch len(results) {
			case 2:
				var responseErr error
				if !results[1].IsNil() {
					if err, ok := results[1].Interface().(error); ok {
						responseErr = err
					}
				}

				responseMsg, deErr := m.doDecoding(ctx, results[0].Interface(), msg.MsgId)
				if deErr != nil && responseErr == nil {
					responseErr = deErr
				}

				if responseErr != nil {
					responseMsg.Error = responseErr.Error()
				}

				b, err := json.Marshal(responseMsg)
				if err != nil {
					return
				}

				user := GetCtx(ctx)
				user.Conn.Write(b)
			}
			return
		}
	})
	return
}

// handleInterceptor 处理拦截器
func (m *MsgParser) handleInterceptor(ctx context.Context, msg *Message) (interceptErr error) {
	for _, f := range m.interceptors {
		if interceptErr = f(ctx, msg); interceptErr != nil {
			break
		}
	}

	if interceptErr == nil {
		return
	}

	handler, ok := m.routers[msg.Router]
	if !ok {
		return
	}

	if handler.IsRPC {
		var output = gutil.Copy(handler.Output.Interface())
		response, doerr := m.doDecoding(ctx, output, msg.MsgId)
		if doerr != nil {
			return doerr
		}

		response.Error = interceptErr.Error()
		b, err := json.Marshal(response)
		if err != nil {
			return err
		}
		ConnFromCtx(ctx).Write(b)
	}
	return
}
