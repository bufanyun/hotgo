// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
	"hotgo/utility/simple"
	"sync"
	"time"
)

// Client tcp客户端
type Client struct {
	ctx       context.Context // 上下文
	conn      *Conn           // 连接对象
	config    *ClientConfig   // 配置
	msgParser *MsgParser      // 消息处理器
	logger    *glog.Logger    // 日志处理器
	isLogin   *gtype.Bool     // 是否已登录
	taskGo    *grpool.Pool    // 任务协程池
	closeFlag *gtype.Bool     // 关闭标签，关闭以后可以重连
	stopFlag  *gtype.Bool     // 停止标签，停止以后不能重连
	wg        sync.WaitGroup  // 流程控制
	mutex     sync.Mutex      // 状态锁
}

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

// CallbackEvent 回调事件
type CallbackEvent func()

// NewClient 初始化一个tcp客户端
func NewClient(config *ClientConfig) (client *Client) {
	client = new(Client)
	client.ctx = gctx.New()
	client.logger = g.Log("tcpClient")

	baseErr := gerror.New("NewClient fail")

	if config == nil {
		client.logger.Fatal(client.ctx, gerror.Wrap(baseErr, "config is nil"))
		return
	}

	if config.Addr == "" {
		client.logger.Fatal(client.ctx, gerror.Wrap(baseErr, "client address is not set"))
		return
	}

	client.config = config
	if config.ConnectInterval <= 0 {
		client.config.ConnectInterval = 5 * time.Second
	} else {
		client.config.ConnectInterval = config.ConnectInterval
	}

	if config.Timeout <= 0 {
		client.config.Timeout = 10 * time.Second
	} else {
		client.config.Timeout = config.Timeout
	}

	client.isLogin = gtype.NewBool(false)
	client.closeFlag = gtype.NewBool(false)
	client.stopFlag = gtype.NewBool(false)

	client.taskGo = grpool.New(5)
	client.msgParser = NewMsgParser(client.handleRoutineTask)
	client.registerDefaultRouter()
	return
}

// Start 启动tcp连接
func (client *Client) Start() (err error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if client.stopFlag.Val() {
		err = gerror.New("client is stop")
		return
	}

	if client.conn != nil {
		return gerror.New("client is running")
	}

	client.isLogin.Set(false)
	client.config.ConnectCount = 0
	client.closeFlag.Set(false)
	client.stopFlag.Set(false)

	client.wg.Add(1)
	simple.SafeGo(client.ctx, func(ctx context.Context) {
		client.connect()
	})
	return
}

// registerDefaultRouter 注册默认路由
func (client *Client) registerDefaultRouter() {
	var routers = []interface{}{
		client.onResponseServerLogin,     // 服务登录
		client.onResponseServerHeartbeat, // 心跳
	}
	client.RegisterRouter(routers...)
}

// RegisterRouter 注册路由
func (client *Client) RegisterRouter(routers ...interface{}) {
	err := client.msgParser.RegisterRouter(routers...)
	if err != nil {
		client.logger.Fatal(client.ctx, err)
	}
}

// RegisterRPCRouter 注册RPC路由
func (client *Client) RegisterRPCRouter(routers ...interface{}) {
	err := client.msgParser.RegisterRPCRouter(routers...)
	if err != nil {
		client.logger.Fatal(client.ctx, err)
	}
}

// RegisterInterceptor 注册拦截器
func (client *Client) RegisterInterceptor(interceptors ...Interceptor) {
	client.msgParser.RegisterInterceptor(interceptors...)
}

// dial
func (client *Client) dial() *gtcp.Conn {
	for {
		conn, err := gtcp.NewConn(client.config.Addr, client.config.Timeout)
		if err == nil || client.closeFlag.Val() {
			return conn
		}

		if client.config.MaxConnectCount > 0 {
			if client.config.ConnectCount < client.config.MaxConnectCount {
				client.config.ConnectCount += 1
			} else {
				return nil
			}
		}

		client.logger.Debugf(client.ctx, "connect to %v error: %v", client.config.Addr, err)
		time.Sleep(client.config.ConnectInterval)
		continue
	}
}

func (client *Client) connect() {
	defer client.wg.Done()

	goto reconnect
reconnect:
	conn := client.dial()
	if conn == nil {
		if !client.stopFlag.Val() {
			client.logger.Debugf(client.ctx, "client dial failed")
		}
		return
	}

	client.mutex.Lock()
	if client.closeFlag.Val() {
		client.mutex.Unlock()
		conn.Close()
		client.logger.Debugf(client.ctx, "client connect but closeFlag is true")
		return
	}

	client.conn = NewConn(conn, client.logger, client.msgParser)
	client.config.ConnectCount = 0
	client.read()
	client.mutex.Unlock()

	client.serverLogin()
	client.startCron()
}

// read
func (client *Client) read() {
	go func() {
		defer func() {
			client.Close()
			client.logger.Debugf(client.ctx, "client are about to be reconnected..")
			time.Sleep(client.config.ConnectInterval)
			client.Start()
		}()

		if err := client.conn.Run(); err != nil {
			client.logger.Debug(client.ctx, err)
		}
	}()
}

// Close 关闭同服务器的链接
func (client *Client) Close() {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	client.isLogin.Set(false)
	client.closeFlag.Set(true)
	if client.conn != nil {
		client.conn.Close()
		client.conn = nil
	}

	if client.config.CloseEvent != nil {
		client.config.CloseEvent()
	}
	client.wg.Wait()
}

// Destroy 销毁当前连接
func (client *Client) Destroy() {
	client.stopCron()
	if client.conn != nil {
		client.conn.Close()
		client.conn = nil
	}
}

// Stop 停止服务
func (client *Client) Stop() {
	if client.stopFlag.Val() {
		return
	}
	client.stopFlag.Set(true)
	client.stopCron()
	client.Close()
}

// IsStop 是否已停止
func (client *Client) IsStop() bool {
	return client.stopFlag.Val()
}

// IsLogin 是否已登录成功
func (client *Client) IsLogin() bool {
	return client.isLogin.Val()
}

func (client *Client) handleRoutineTask(ctx context.Context, task func()) {
	ctx, cancel := context.WithCancel(ctx)
	err := client.taskGo.AddWithRecover(ctx,
		func(ctx context.Context) {
			task()
			cancel()
		},
		func(ctx context.Context, err error) {
			client.logger.Warningf(ctx, "routineTask exec err:%+v", err)
			cancel()
		},
	)
	if err != nil {
		client.logger.Warningf(ctx, "routineTask add err:%+v", err)
	}
}

// Conn 获取当前连接
func (client *Client) Conn() *Conn {
	return client.conn
}

// Send 发送消息
func (client *Client) Send(ctx context.Context, data interface{}) error {
	if client.conn == nil {
		return gerror.New("conn is nil")
	}
	return client.conn.Send(ctx, data)
}

// Request 发送消息并等待响应结果
func (client *Client) Request(ctx context.Context, data interface{}) (interface{}, error) {
	return client.conn.Request(ctx, data)
}

// RequestScan 发送消息并等待响应结果，将结果保存在response中
func (client *Client) RequestScan(ctx context.Context, data, response interface{}) error {
	return client.conn.RequestScan(ctx, data, response)
}
