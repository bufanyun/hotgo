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
	"hotgo/internal/consts"
	"hotgo/utility/simple"
	"reflect"
	"sync"
	"time"
)

// ClientConn 连接到tcp服务器的客户端对象
type ClientConn struct {
	Conn      *gtcp.Conn // 连接对象
	Auth      *AuthMeta  // 认证元数据
	heartbeat int64      // 心跳
}

// ServerConfig tcp服务器配置
type ServerConfig struct {
	Name string // 服务名称
	Addr string // 监听地址
}

// Server tcp服务器对象结构
type Server struct {
	Ctx          context.Context          // 上下文
	Logger       *glog.Logger             // 日志处理器
	addr         string                   // 连接地址
	name         string                   // 服务器名称
	rpc          *Rpc                     // rpc协议
	ln           *gtcp.Server             // tcp服务器
	wgLn         sync.WaitGroup           // 状态控制，主要用于tcp服务器能够按流程启动退出
	mutex        sync.Mutex               // 服务器状态锁
	closeFlag    bool                     // 服务关闭标签
	clients      map[string]*ClientConn   // 已登录的认证客户端
	mutexConns   sync.Mutex               // 连接锁，主要用于客户端上下线
	msgGo        *grpool.Pool             // 消息处理协程池
	cronRouters  map[string]RouterHandler // 定时任务路由
	queueRouters map[string]RouterHandler // 队列路由
	authRouters  map[string]RouterHandler // 任务路由
}

// NewServer 初始一个tcp服务器对象
func NewServer(config *ServerConfig) (server *Server, err error) {
	if config == nil {
		err = gerror.New("config is nil")
		return
	}

	if config.Addr == "" {
		err = gerror.New("server address is not set")
		return
	}

	server = new(Server)
	server.Ctx = gctx.New()

	if config.Name == "" {
		config.Name = simple.AppName(server.Ctx)
	}

	server.addr = config.Addr
	server.name = config.Name
	server.ln = gtcp.NewServer(server.addr, server.accept, config.Name)
	server.clients = make(map[string]*ClientConn)
	server.closeFlag = false
	server.Logger = g.Log("tcpServer")
	server.msgGo = grpool.New(20)
	server.rpc = NewRpc(server.Ctx, server.msgGo, server.Logger)
	server.startCron()
	return
}

// accept
func (server *Server) accept(conn *gtcp.Conn) {
	defer func() {
		server.mutexConns.Lock()
		_ = conn.Close()
		// 从登录列表中移除
		delete(server.clients, conn.RemoteAddr().String())
		server.mutexConns.Unlock()
	}()

	for {
		msg, err := RecvPkg(conn)
		if err != nil {
			server.Logger.Debugf(server.Ctx, "RecvPkg err:%+v, client closed.", err)
			break
		}

		client := server.getLoginConn(conn)

		switch msg.Router {
		case "ServerLogin": // 服务登录
			// 初始化上下文
			ctx := initCtx(gctx.New(), &Context{
				Conn: conn,
			})
			server.doHandleRouterMsg(ctx, server.onServerLogin, msg.Data)
		case "ServerHeartbeat": // 心跳
			if client == nil {
				server.Logger.Infof(server.Ctx, "conn not connected, ignore the heartbeat, msg:%+v", msg)
				continue
			}
			// 初始化上下文
			ctx := initCtx(gctx.New(), &Context{})
			server.doHandleRouterMsg(ctx, server.onServerHeartbeat, msg.Data, client)
		default: // 通用路由消息处理
			if client == nil {
				server.Logger.Warningf(server.Ctx, "conn is not logged in but sends a routing message. actively conn disconnect, msg:%+v", msg)
				time.Sleep(time.Second)
				conn.Close()
				return
			}
			server.handleRouterMsg(msg, client)
		}
	}
}

// handleRouterMsg 处理路由消息
func (server *Server) handleRouterMsg(msg *Message, client *ClientConn) {
	// 验证签名
	in, err := VerifySign(msg.Data, client.Auth.AppId, client.Auth.SecretKey)
	if err != nil {
		server.Logger.Warningf(server.Ctx, "handleRouterMsg VerifySign err:%+v message: %+v", err, msg)
		return
	}

	// 初始化上下文
	ctx := initCtx(gctx.New(), &Context{
		Conn:    client.Conn,
		Auth:    client.Auth,
		TraceID: in.TraceID,
	})

	// 响应rpc消息
	if server.rpc.HandleMsg(ctx, msg.Data) {
		return
	}

	handle := func(routers map[string]RouterHandler, group string) {
		if routers == nil {
			server.Logger.Debugf(server.Ctx, "handleRouterMsg route is not initialized %v message: %+v", group, msg)
			return
		}
		f, ok := routers[msg.Router]
		if !ok {
			server.Logger.Debugf(server.Ctx, "handleRouterMsg invalid %v message: %+v", group, msg)
			return
		}

		server.doHandleRouterMsg(ctx, f, msg.Data)
	}

	switch client.Auth.Group {
	case consts.TCPClientGroupCron:
		handle(server.cronRouters, client.Auth.Group)
	case consts.TCPClientGroupQueue:
		handle(server.queueRouters, client.Auth.Group)
	case consts.TCPClientGroupAuth:
		handle(server.authRouters, client.Auth.Group)
	default:
		server.Logger.Warningf(server.Ctx, "group is not registered: %+v", client.Auth.Group)
	}
}

// doHandleRouterMsg 处理路由消息
func (server *Server) doHandleRouterMsg(ctx context.Context, fun RouterHandler, args ...interface{}) {
	ctx, cancel := context.WithCancel(ctx)
	err := server.msgGo.AddWithRecover(ctx,
		func(ctx context.Context) {
			fun(ctx, args...)
			cancel()
		},
		func(ctx context.Context, err error) {
			server.Logger.Warningf(ctx, "doHandleRouterMsg msgGo exec err:%+v", err)
			cancel()
		},
	)

	if err != nil {
		server.Logger.Warningf(ctx, "doHandleRouterMsg msgGo Add err:%+v", err)
		return
	}
}

// getLoginConn 获取指定已登录的连接
func (server *Server) getLoginConn(conn *gtcp.Conn) *ClientConn {
	client, ok := server.clients[conn.RemoteAddr().String()]
	if !ok {
		return nil
	}
	return client
}

// getLoginConn 获取指定appid的所有连接
func (server *Server) getAppIdClients(appid string) (list []*ClientConn) {
	for _, v := range server.clients {
		if v.Auth.AppId == appid {
			list = append(list, v)
		}
	}
	return
}

// GetGroupClients 获取指定分组的所有连接
func (server *Server) GetGroupClients(group string) (list []*ClientConn) {
	for _, v := range server.clients {
		if v.Auth.Group == group {
			list = append(list, v)
		}
	}
	return
}

// RegisterAuthRouter 注册授权路由
func (server *Server) RegisterAuthRouter(routers map[string]RouterHandler) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if server.authRouters == nil {
		server.authRouters = make(map[string]RouterHandler)
	}

	for i, router := range routers {
		_, ok := server.authRouters[i]
		if ok {
			server.Logger.Debugf(server.Ctx, "server authRouters duplicate registration:%v", i)
			continue
		}
		server.authRouters[i] = router
	}
}

// RegisterCronRouter 注册任务路由
func (server *Server) RegisterCronRouter(routers map[string]RouterHandler) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if server.cronRouters == nil {
		server.cronRouters = make(map[string]RouterHandler)
	}

	for i, router := range routers {
		_, ok := server.cronRouters[i]
		if ok {
			server.Logger.Debugf(server.Ctx, "server cronRouters duplicate registration:%v", i)
			continue
		}
		server.cronRouters[i] = router
	}
}

// RegisterQueueRouter 注册队列路由
func (server *Server) RegisterQueueRouter(routers map[string]RouterHandler) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if server.queueRouters == nil {
		server.queueRouters = make(map[string]RouterHandler)
	}

	for i, router := range routers {
		_, ok := server.queueRouters[i]
		if ok {
			server.Logger.Debugf(server.Ctx, "server queueRouters duplicate registration:%v", i)
			continue
		}
		server.queueRouters[i] = router
	}
}

// Listen 监听服务
func (server *Server) Listen() (err error) {
	server.wgLn.Add(1)
	defer server.wgLn.Done()
	return server.ln.Run()
}

// Close 关闭服务
func (server *Server) Close() {
	if server.closeFlag {
		return
	}
	server.closeFlag = true

	server.stopCron()

	server.mutexConns.Lock()
	for _, client := range server.clients {
		_ = client.Conn.Close()
	}
	server.clients = nil
	server.mutexConns.Unlock()

	if server.ln != nil {
		_ = server.ln.Close()
	}
	server.wgLn.Wait()
}

// IsClose 服务是否关闭
func (server *Server) IsClose() bool {
	return server.closeFlag
}

// Write 向指定客户端发送消息
func (server *Server) Write(conn *gtcp.Conn, data interface{}) (err error) {
	if server.closeFlag {
		return gerror.New("service is down")
	}

	msgType := reflect.TypeOf(data)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return gerror.Newf("json message pointer required: %+v", data)
	}

	msg := &Message{Router: msgType.Elem().Name(), Data: data}

	return SendPkg(conn, msg)
}

// Send 发送消息
func (server *Server) Send(ctx context.Context, client *ClientConn, data interface{}) (err error) {
	MsgPkg(data, client.Auth, gctx.CtxId(ctx))
	return server.Write(client.Conn, data)
}

// Reply 回复消息
func (server *Server) Reply(ctx context.Context, data interface{}) (err error) {
	user := GetCtx(ctx)
	if user == nil {
		err = gerror.New("获取回复用户信息失败")
		return
	}
	MsgPkg(data, user.Auth, user.TraceID)
	return server.Write(user.Conn, data)
}

// RpcRequest 向指定客户端发送消息并等待响应结果
func (server *Server) RpcRequest(ctx context.Context, client *ClientConn, data interface{}) (res interface{}, err error) {
	var (
		traceID = MsgPkg(data, client.Auth, gctx.CtxId(ctx))
		key     = server.rpc.GetCallId(client.Conn, traceID)
	)

	if traceID == "" {
		err = gerror.New("traceID is required")
		return
	}

	return server.rpc.Request(key, func() {
		_ = server.Write(client.Conn, data)
	})
}
