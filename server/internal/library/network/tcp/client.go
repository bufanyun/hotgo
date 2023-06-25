// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/utility/simple"
	"reflect"
	"sync"
	"time"
)

// ClientConfig 客户端配置
type ClientConfig struct {
	Addr            string        // 连接地址
	Auth            *AuthMeta     // 认证元数据
	Timeout         time.Duration // 连接超时时间
	ConnectInterval time.Duration // 重连时间间隔
	MaxConnectCount uint          // 最大重连次数，0不限次数
	ConnectCount    uint          // 已重连次数
	AutoReconnect   bool          // 是否开启自动重连
	LoginEvent      CallbackEvent // 登录成功事件
	CloseEvent      CallbackEvent // 连接关闭事件
}

// Client 客户端
type Client struct {
	Ctx             context.Context          // 上下文
	Logger          *glog.Logger             // 日志处理器
	IsLogin         bool                     // 是否已登录
	addr            string                   // 连接地址
	auth            *AuthMeta                // 认证元数据
	rpc             *Rpc                     // rpc协议支持
	timeout         time.Duration            // 连接超时时间
	connectInterval time.Duration            // 重连时间间隔
	maxConnectCount uint                     // 最大重连次数，0不限次数
	connectCount    uint                     // 已重连次数
	autoReconnect   bool                     // 是否开启自动重连
	loginEvent      CallbackEvent            // 登录成功事件
	closeEvent      CallbackEvent            // 连接关闭事件
	sync.Mutex                               // 状态锁
	heartbeat       int64                    // 心跳
	msgGo           *grpool.Pool             // 消息处理协程池
	routers         map[string]RouterHandler // 已注册的路由
	conn            *gtcp.Conn               // 连接对象
	wg              sync.WaitGroup           // 状态控制
	closeFlag       bool                     // 关闭标签，关闭以后可以重连
	stopFlag        bool                     // 停止标签，停止以后不能重连
}

// NewClient 初始化一个tcp客户端
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
	client.Logger = g.Log("tcpClient")

	if config.ConnectInterval <= 0 {
		client.connectInterval = 5 * time.Second
	} else {
		client.connectInterval = config.ConnectInterval
	}

	if config.Timeout <= 0 {
		client.timeout = 10 * time.Second
	} else {
		client.timeout = config.Timeout
	}

	client.msgGo = grpool.New(5)
	client.rpc = NewRpc(client.Ctx, client.msgGo, client.Logger)
	return
}

// Start 启动tcp连接
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

// dial
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
		_ = conn.Close()
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

// read
func (client *Client) read() {
	simple.SafeGo(client.Ctx, func(ctx context.Context) {
		defer func() {
			client.Close()
			client.Logger.Debugf(client.Ctx, "client are about to be reconnected..")
			time.Sleep(client.connectInterval)
			_ = client.Start()
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
				client.doHandleRouterMsg(initCtx(gctx.New(), &Context{}), f, msg.Data)
			default: // 通用路由消息处理
				in, err := VerifySign(msg.Data, client.auth.AppId, client.auth.SecretKey)
				if err != nil {
					client.Logger.Warningf(client.Ctx, "client read VerifySign err:%+v message: %+v", err, msg)
					continue
				}

				ctx := initCtx(gctx.New(), &Context{
					Conn:    client.conn,
					Auth:    client.auth,
					TraceID: in.TraceID,
				})

				// 响应rpc消息
				if client.rpc.HandleMsg(ctx, msg.Data) {
					return
				}

				client.doHandleRouterMsg(ctx, f, msg.Data)
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
		_ = client.Write(data)
	})
}

// doHandleRouterMsg 处理路由消息
func (client *Client) doHandleRouterMsg(ctx context.Context, fun RouterHandler, args ...interface{}) {
	ctx, cancel := context.WithCancel(ctx)
	err := client.msgGo.AddWithRecover(ctx,
		func(ctx context.Context) {
			fun(ctx, args...)
			cancel()
		},
		func(ctx context.Context, err error) {
			client.Logger.Warningf(ctx, "doHandleRouterMsg msgGo exec err:%+v", err)
			cancel()
		},
	)

	if err != nil {
		client.Logger.Warningf(ctx, "doHandleRouterMsg msgGo Add err:%+v", err)
		return
	}
}
