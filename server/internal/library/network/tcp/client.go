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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/utility/simple"
	"reflect"
	"sync"
	"time"
)

// ClientConfig 客户端配置
type ClientConfig struct {
	Addr            string
	Auth            *AuthMeta
	Timeout         time.Duration
	ConnectInterval time.Duration
	MaxConnectCount uint
	ConnectCount    uint
	AutoReconnect   bool
	LoginEvent      CallbackEvent
	CloseEvent      CallbackEvent
}

// Client 客户端
type Client struct {
	Ctx             context.Context
	Logger          *glog.Logger
	IsLogin         bool // 是否已登录
	addr            string
	auth            *AuthMeta
	rpc             *Rpc
	timeout         time.Duration
	connectInterval time.Duration
	maxConnectCount uint
	connectCount    uint
	autoReconnect   bool
	loginEvent      CallbackEvent
	closeEvent      CallbackEvent
	sync.Mutex
	heartbeat int64
	routers   map[string]RouterHandler
	conn      *gtcp.Conn
	wg        sync.WaitGroup
	closeFlag bool // 关闭标签，关闭以后可以重连
	stopFlag  bool // 停止标签，停止以后不能重连
}

func NewClient(config *ClientConfig) (client *Client, err error) {
	client = new(Client)

	if config == nil {
		err = gerror.New("config is nil")
		return
	}

	if config.Addr == "" {
		err = gerror.New("client address is not set")
		return
	}

	if config.Auth == nil {
		err = gerror.New("client auth cannot be empty")
		return
	}

	if config.Auth.Group == "" || config.Auth.Name == "" {
		err = gerror.New("Auth.Group or Auth.Group is nil")
		return
	}

	client.Ctx = gctx.New()
	client.autoReconnect = true
	client.addr = config.Addr
	client.auth = config.Auth
	client.loginEvent = config.LoginEvent
	client.closeEvent = config.CloseEvent

	logger := glog.New()
	path := g.Cfg().MustGet(client.Ctx, "logger.path", "logs/logger").String()
	if err = logger.SetPath(fmt.Sprintf("%s/tcp.client/%s.%s", path, config.Auth.Group, config.Auth.Name)); err != nil {
		return
	}
	client.Logger = logger

	if config.ConnectInterval <= 0 {
		client.connectInterval = 5 * time.Second
		client.Logger.Debugf(client.Ctx, "invalid connectInterval, reset to %v", client.connectInterval)
	} else {
		client.connectInterval = config.ConnectInterval
	}

	if config.Timeout <= 0 {
		client.timeout = 10 * time.Second
		client.Logger.Debugf(client.Ctx, "invalid timeout, reset to %v", client.timeout)
	} else {
		client.timeout = config.Timeout
	}

	client.rpc = NewRpc(client.Ctx)
	return
}

// Start 启动
func (client *Client) Start() (err error) {
	client.Lock()
	defer client.Unlock()

	if client.stopFlag {
		err = gerror.New("client is stop")
		return
	}

	if client.conn != nil {
		return gerror.New("client is running")
	}

	client.IsLogin = false
	client.connectCount = 0
	client.closeFlag = false
	client.stopFlag = false

	client.wg.Add(1)
	simple.SafeGo(client.Ctx, func(ctx context.Context) {
		client.connect()
	})

	return
}

// RegisterRouter 注册路由
func (client *Client) RegisterRouter(routers map[string]RouterHandler) (err error) {
	if client.conn != nil {
		return gerror.New("client is running")
	}

	client.Lock()
	defer client.Unlock()

	if client.routers == nil {
		client.routers = make(map[string]RouterHandler)
		// 默认路由
		client.routers = map[string]RouterHandler{
			"ResponseServerHeartbeat": client.onResponseServerHeartbeat,
			"ResponseServerLogin":     client.onResponseServerLogin,
		}
	}

	for i, router := range routers {
		_, ok := client.routers[i]
		if ok {
			return gerror.Newf("client route duplicate registration:%v", i)
		}
		client.routers[i] = router
	}
	return
}

func (client *Client) dial() *gtcp.Conn {
	for {
		conn, err := gtcp.NewConn(client.addr, client.timeout)
		if err == nil || client.closeFlag {
			return conn
		}

		if client.maxConnectCount > 0 {
			if client.connectCount < client.maxConnectCount {
				client.connectCount += 1
			} else {
				return nil
			}
		}

		client.Logger.Debugf(client.Ctx, "connect to %v error: %v", client.addr, err)
		time.Sleep(client.connectInterval)
		continue
	}
}

func (client *Client) connect() {
	defer client.wg.Done()

	goto reconnect
reconnect:
	conn := client.dial()
	if conn == nil {
		if !client.stopFlag {
			client.Logger.Debugf(client.Ctx, "client dial failed")
		}
		return
	}

	client.Lock()
	if client.closeFlag {
		client.Unlock()
		conn.Close()
		client.Logger.Debugf(client.Ctx, "client connect but closeFlag is true")
		return
	}

	client.conn = conn
	client.connectCount = 0
	client.heartbeat = gtime.Timestamp()

	client.read()
	client.Unlock()

	client.serverLogin()
	client.startCron()
}

func (client *Client) read() {
	simple.SafeGo(client.Ctx, func(ctx context.Context) {
		defer func() {
			client.Close()
			client.Logger.Debugf(client.Ctx, "client are about to be reconnected..")
			time.Sleep(client.connectInterval)
			client.Start()
		}()

		for {
			if client.conn == nil {
				client.Logger.Debugf(client.Ctx, "client client.conn is nil, server closed")
				break
			}

			msg, err := RecvPkg(client.conn)
			if err != nil {
				client.Logger.Debugf(client.Ctx, "client RecvPkg err:%+v, server closed", err)
				break
			}

			if client.routers == nil {
				client.Logger.Debugf(client.Ctx, "client RecvPkg routers is nil")
				break
			}

			if msg == nil {
				client.Logger.Debugf(client.Ctx, "client RecvPkg msg is nil")
				break
			}

			f, ok := client.routers[msg.Router]
			if !ok {
				client.Logger.Debugf(client.Ctx, "client RecvPkg invalid message: %+v", msg)
				continue
			}

			switch msg.Router {
			case "ResponseServerLogin", "ResponseServerHeartbeat": // 服务登录、心跳无需验证签名
				ctx, cancel := initCtx(gctx.New(), &Context{})
				doHandleRouterMsg(f, ctx, cancel, msg.Data)
			default: // 通用路由消息处理
				in, err := VerifySign(msg.Data, client.auth.AppId, client.auth.SecretKey)
				if err != nil {
					client.Logger.Warningf(client.Ctx, "client read VerifySign err:%+v message: %+v", err, msg)
					continue
				}

				ctx, cancel := initCtx(gctx.New(), &Context{
					Conn:    client.conn,
					Auth:    client.auth,
					TraceID: in.TraceID,
				})

				// 响应rpc消息
				if client.rpc.HandleMsg(ctx, cancel, msg.Data) {
					return
				}

				doHandleRouterMsg(f, ctx, cancel, msg.Data)
			}
		}
	})
}

// Close 关闭同服务器的链接
func (client *Client) Close() {
	client.Lock()
	defer client.Unlock()

	client.IsLogin = false
	client.closeFlag = true
	if client.conn != nil {
		client.conn.Close()
		client.conn = nil
	}

	if client.closeEvent != nil {
		client.closeEvent()
	}
	client.wg.Wait()
}

// Stop 停止服务
func (client *Client) Stop() {
	if client.stopFlag {
		return
	}
	client.stopFlag = true
	client.stopCron()
	client.Close()
}

// IsStop 是否已停止
func (client *Client) IsStop() bool {
	return client.stopFlag
}

// Destroy 销毁当前连接
func (client *Client) Destroy() {
	client.stopCron()
	if client.conn != nil {
		client.conn.Close()
		client.conn = nil
	}
}

// Write
func (client *Client) Write(data interface{}) error {
	client.Lock()
	defer client.Unlock()

	if client.conn == nil {
		return gerror.New("client conn is nil")
	}

	if client.closeFlag {
		return gerror.New("client conn is closed")
	}

	if data == nil {
		return gerror.New("client Write message is nil")
	}

	msgType := reflect.TypeOf(data)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return gerror.Newf("client json message pointer required: %+v", data)
	}
	msg := &Message{Router: msgType.Elem().Name(), Data: data}

	return SendPkg(client.conn, msg)
}

// Send 发送消息
func (client *Client) Send(ctx context.Context, data interface{}) error {
	MsgPkg(data, client.auth, gctx.CtxId(ctx))
	return client.Write(data)
}

// Reply 回复消息
func (client *Client) Reply(ctx context.Context, data interface{}) (err error) {
	user := GetCtx(ctx)
	if user == nil {
		err = gerror.New("获取回复用户信息失败")
		return
	}
	MsgPkg(data, client.auth, user.TraceID)
	return client.Write(data)
}

// RpcRequest 发送消息并等待响应结果
func (client *Client) RpcRequest(ctx context.Context, data interface{}) (res interface{}, err error) {
	var (
		traceID = MsgPkg(data, client.auth, gctx.CtxId(ctx))
		key     = client.rpc.GetCallId(client.conn, traceID)
	)

	if traceID == "" {
		err = gerror.New("traceID is required")
		return
	}

	return client.rpc.Request(key, func() {
		client.Write(data)
	})
}
