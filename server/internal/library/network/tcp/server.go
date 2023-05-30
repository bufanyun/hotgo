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
	"hotgo/internal/consts"
	"reflect"
	"sync"
	"time"
)

type ClientConn struct {
	Conn      *gtcp.Conn
	Auth      *AuthMeta
	heartbeat int64
}

type ServerConfig struct {
	Name string // 服务名称
	Addr string // 监听地址
}

type Server struct {
	Ctx          context.Context
	Logger       *glog.Logger
	addr         string
	name         string
	rpc          *Rpc
	ln           *gtcp.Server
	wgLn         sync.WaitGroup
	mutex        sync.Mutex
	closeFlag    bool
	clients      map[string]*ClientConn // 已登录的认证客户端
	mutexConns   sync.Mutex
	wgConns      sync.WaitGroup
	cronRouters  map[string]RouterHandler // 路由
	queueRouters map[string]RouterHandler
	authRouters  map[string]RouterHandler
}

func NewServer(config *ServerConfig) (server *Server, err error) {
	if config == nil {
		err = gerror.New("config is nil")
		return
	}

	if config.Addr == "" {
		err = gerror.New("server address is not set")
		return
	}

	if config.Name == "" {
		config.Name = "hotgo"
	}

	server = new(Server)
	server.Ctx = gctx.New()
	server.addr = config.Addr
	server.name = config.Name
	server.ln = gtcp.NewServer(server.addr, server.accept, config.Name)
	server.clients = make(map[string]*ClientConn)
	server.closeFlag = false

	logger := glog.New()
	path := g.Cfg().MustGet(server.Ctx, "logger.path", "logs/logger").String()
	if err = logger.SetPath(fmt.Sprintf("%s/tcp.server/%s", path, config.Name)); err != nil {
		return
	}
	server.Logger = logger
	server.rpc = NewRpc(server.Ctx)

	server.startCron()

	return
}

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
			ctx, cancel := initCtx(gctx.New(), &Context{
				Conn: conn,
			})
			doHandleRouterMsg(server.onServerLogin, ctx, cancel, msg.Data)
		case "ServerHeartbeat": // 心跳
			if client == nil {
				server.Logger.Infof(server.Ctx, "conn not connected, ignore the heartbeat, msg:%+v", msg)
				continue
			}
			// 初始化上下文
			ctx, cancel := initCtx(gctx.New(), &Context{})
			doHandleRouterMsg(server.onServerHeartbeat, ctx, cancel, msg.Data, client)
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
	ctx, cancel := initCtx(gctx.New(), &Context{
		Conn:    client.Conn,
		Auth:    client.Auth,
		TraceID: in.TraceID,
	})

	// 响应rpc消息
	if server.rpc.HandleMsg(ctx, cancel, msg.Data) {
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

		doHandleRouterMsg(f, ctx, cancel, msg.Data)
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
	server.wgConns.Wait()

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
