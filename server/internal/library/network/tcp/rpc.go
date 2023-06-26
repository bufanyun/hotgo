// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
	"hotgo/internal/consts"
	"hotgo/utility/simple"
	"sync"
	"time"
)

type Rpc struct {
	ctx       context.Context
	mutex     sync.Mutex
	callbacks map[string]RpcRespFunc
	msgGo     *grpool.Pool // 消息处理协程池
	logger    *glog.Logger // 日志处理器
}

// RpcResp 响应结构
type RpcResp struct {
	res interface{}
	err error
}

type RpcRespFunc func(resp interface{}, err error)

// NewRpc 初始化一个rpc协议
func NewRpc(ctx context.Context, msgGo *grpool.Pool, logger *glog.Logger) *Rpc {
	return &Rpc{
		ctx:       ctx,
		callbacks: make(map[string]RpcRespFunc),
		msgGo:     msgGo,
		logger:    logger,
	}
}

// GetCallId 获取回调id
func (r *Rpc) GetCallId(client *gtcp.Conn, traceID string) string {
	return fmt.Sprintf("%v.%v", client.LocalAddr().String(), traceID)
}

// HandleMsg 处理rpc消息
func (r *Rpc) HandleMsg(ctx context.Context, data interface{}) bool {
	user := GetCtx(ctx)
	callId := r.GetCallId(user.Conn, user.TraceID)

	if call, ok := r.callbacks[callId]; ok {
		r.mutex.Lock()
		delete(r.callbacks, callId)
		r.mutex.Unlock()

		ctx, cancel := context.WithCancel(ctx)

		err := r.msgGo.AddWithRecover(ctx, func(ctx context.Context) {
			call(data, nil)
			cancel()
		}, func(ctx context.Context, err error) {
			r.logger.Warningf(ctx, "rpc HandleMsg msgGo exec err:%+v", err)
			cancel()
		})

		if err != nil {
			r.logger.Warningf(ctx, "rpc HandleMsg msgGo Add err:%+v", err)
		}
		return true
	}
	return false
}

// Request 发起rpc请求
func (r *Rpc) Request(callId string, send func()) (res interface{}, err error) {
	var (
		waitCh  = make(chan struct{})
		resCh   = make(chan RpcResp, 1)
		isClose = false
	)

	defer func() {
		isClose = true
		close(resCh)

		// 移除消息
		if _, ok := r.callbacks[callId]; ok {
			r.mutex.Lock()
			delete(r.callbacks, callId)
			r.mutex.Unlock()
		}
	}()

	simple.SafeGo(r.ctx, func(ctx context.Context) {
		close(waitCh)

		// 加入回调
		r.mutex.Lock()
		r.callbacks[callId] = func(res interface{}, err error) {
			if !isClose {
				resCh <- RpcResp{res: res, err: err}
			}
		}
		r.mutex.Unlock()

		// 发送消息
		send()
	})

	<-waitCh
	select {
	case <-time.After(time.Second * consts.TCPRpcTimeout):
		err = gerror.New("rpc response timeout")
		return
	case got := <-resCh:
		return got.res, got.err
	}
}
