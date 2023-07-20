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
)

// Server tcp服务器
type Server struct {
	ctx        context.Context  // 上下文
	name       string           // 服务器名称
	addr       string           // 服务器地址
	ln         *gtcp.Server     // tcp服务器
	logger     *glog.Logger     // 日志处理器
	wgLn       sync.WaitGroup   // 流程控制，让tcp服务器按可控流程启动退出
	closeFlag  *gtype.Bool      // 服务关闭标签
	clients    map[string]*Conn // 已登录的认证客户端
	mutexConns sync.Mutex       // 连接锁，主要用于客户端上下线
	taskGo     *grpool.Pool     // 任务协程池
	msgParser  *MsgParser       // 消息处理器
}

// ServerConfig tcp服务器配置
type ServerConfig struct {
	Name string // 服务名称
	Addr string // 监听地址
}

// NewServer 初始一个tcp服务器对象
func NewServer(config *ServerConfig) (server *Server) {
	server = new(Server)
	server.ctx = gctx.New()
	server.logger = g.Log("tcpServer")

	baseErr := gerror.New("TCPServer start fail")
	if config == nil {
		server.logger.Fatal(server.ctx, gerror.Wrap(baseErr, "config is nil"))
		return
	}

	if config.Addr == "" {
		server.logger.Fatal(server.ctx, gerror.Wrap(baseErr, "server address is not set"))
		return
	}

	if config.Name == "" {
		config.Name = simple.AppName(server.ctx)
	}

	server.addr = config.Addr
	server.name = config.Name
	server.ln = gtcp.NewServer(server.addr, server.accept, config.Name)
	server.closeFlag = gtype.NewBool(false)
	server.clients = make(map[string]*Conn)
	server.taskGo = grpool.New(20)
	server.msgParser = NewMsgParser(server.handleRoutineTask)

	server.startCron()
	return
}

// accept
func (server *Server) accept(conn *gtcp.Conn) {
	tcpConn := NewConn(conn, server.logger, server.msgParser)
	server.AddClient(tcpConn)
	go func() {
		if err := tcpConn.Run(); err != nil {
			server.logger.Info(server.ctx, err)
		}

		// cleanup
		tcpConn.Close()
		server.RemoveClient(tcpConn)
	}()
}

// RemoveClient 移除客户端
func (server *Server) RemoveClient(conn *Conn) {
	label := server.ClientLabel(conn.Conn)
	if _, ok := server.clients[label]; ok {
		server.mutexConns.Lock()
		delete(server.clients, label)
		server.mutexConns.Unlock()
	}
}

// AddClient 添加客户端
func (server *Server) AddClient(conn *Conn) {
	server.mutexConns.Lock()
	server.clients[server.ClientLabel(conn.Conn)] = conn
	server.mutexConns.Unlock()
}

// AuthClient 认证客户端
func (server *Server) AuthClient(conn *Conn, auth *AuthMeta) {
	label := server.ClientLabel(conn.Conn)
	client, ok := server.clients[label]
	if !ok {
		server.logger.Debugf(server.ctx, "authClient client does not exist:%v", label)
		return
	}
	client.Auth = auth
}

// ClientLabel 客户端标识
func (server *Server) ClientLabel(conn *gtcp.Conn) string {
	return conn.RemoteAddr().String()
}

// GetClient 获取指定连接
func (server *Server) GetClient(conn *gtcp.Conn) *Conn {
	client, ok := server.clients[server.ClientLabel(conn)]
	if !ok {
		return nil
	}
	return client
}

// GetClients 获取所有连接
func (server *Server) GetClients() map[string]*Conn {
	return server.clients
}

// GetClientById 通过连接ID获取连接
func (server *Server) GetClientById(id int64) *Conn {
	server.mutexConns.Lock()
	defer server.mutexConns.Unlock()

	for _, v := range server.clients {
		if v.CID == id {
			return v
		}
	}
	return nil
}

// GetAppIdClients 获取指定appid的所有连接
func (server *Server) GetAppIdClients(appid string) (list []*Conn) {
	server.mutexConns.Lock()
	defer server.mutexConns.Unlock()

	for _, v := range server.clients {
		if v.Auth != nil && v.Auth.AppId == appid {
			list = append(list, v)
		}
	}
	return
}

// GetGroupClients 获取指定分组的所有连接
func (server *Server) GetGroupClients(group string) (list []*Conn) {
	server.mutexConns.Lock()
	defer server.mutexConns.Unlock()

	for _, v := range server.clients {
		if v.Auth != nil && v.Auth.Group == group {
			list = append(list, v)
		}
	}
	return
}

// GetAppIdOnline 获取指定appid的在线数量
func (server *Server) GetAppIdOnline(appid string) int {
	return len(server.GetAppIdClients(appid))
}

// GetAllOnline 获取所有在线数量
func (server *Server) GetAllOnline() int {
	return len(server.clients)
}

// GetAuthOnline 获取所有已登录认证在线数量
func (server *Server) GetAuthOnline() int {
	server.mutexConns.Lock()
	defer server.mutexConns.Unlock()

	online := 0
	for _, v := range server.clients {
		if v.Auth != nil {
			online++
		}
	}
	return online
}

// RegisterRouter 注册路由
func (server *Server) RegisterRouter(routers ...interface{}) {
	err := server.msgParser.RegisterRouter(routers...)
	if err != nil {
		server.logger.Fatal(server.ctx, err)
	}
	return
}

// RegisterRPCRouter 注册RPC路由
func (server *Server) RegisterRPCRouter(routers ...interface{}) {
	err := server.msgParser.RegisterRPCRouter(routers...)
	if err != nil {
		server.logger.Fatal(server.ctx, err)
	}
	return
}

// RegisterInterceptor 注册拦截器
func (server *Server) RegisterInterceptor(interceptors ...Interceptor) {
	server.msgParser.RegisterInterceptor(interceptors...)
}

// Listen 监听服务
func (server *Server) Listen() (err error) {
	server.wgLn.Add(1)
	defer server.wgLn.Done()
	return server.ln.Run()
}

// Close 关闭服务
func (server *Server) Close() {
	if server.closeFlag.Val() {
		return
	}

	server.closeFlag.Set(true)
	server.stopCron()

	server.mutexConns.Lock()
	for _, client := range server.clients {
		client.Conn.Close()
	}
	server.clients = nil
	server.mutexConns.Unlock()

	if server.ln != nil {
		server.ln.Close()
	}
	server.wgLn.Wait()
}

// IsClose 服务是否关闭
func (server *Server) IsClose() bool {
	return server.closeFlag.Val()
}

// handleRoutineTask 处理协程任务
func (server *Server) handleRoutineTask(ctx context.Context, task func()) {
	ctx, cancel := context.WithCancel(ctx)
	err := server.taskGo.AddWithRecover(ctx,
		func(ctx context.Context) {
			task()
			cancel()
		},
		func(ctx context.Context, err error) {
			server.logger.Warningf(ctx, "routineTask exec err:%+v", err)
			cancel()
		},
	)

	if err != nil {
		server.logger.Warningf(ctx, "routineTask add err:%+v", err)
	}
}

// GetRoutes 获取所有路由
func (server *Server) GetRoutes() (routes []RouteHandler) {
	if server.msgParser.routers == nil {
		return
	}

	for _, v := range server.msgParser.routers {
		routes = append(routes, *v)
	}
	return
}

// Request 向指定客户端发送消息并等待响应结果
func (server *Server) Request(ctx context.Context, client *Conn, data interface{}) (interface{}, error) {
	return client.Request(ctx, data)
}

// RequestScan 向指定客户端发送消息并等待响应结果，将结果保存在response中
func (server *Server) RequestScan(ctx context.Context, client *Conn, data, response interface{}) error {
	return client.RequestScan(ctx, data, response)
}
