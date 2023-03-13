package tcpserver

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

type sTCPServer struct {
	serv *tcp.Server
}

func init() {
	service.RegisterTCPServer(newTCPServer())
}

func newTCPServer() *sTCPServer {
	return &sTCPServer{}
}

// Start 启动服务
func (s *sTCPServer) Start(ctx context.Context) {
	simple.SafeGo(ctx, func(ctx context.Context) {
		g.Log().Debug(ctx, "TCPServer start..")

		server, err := tcp.NewServer(&tcp.ServerConfig{
			Name: "hotgo",
			Addr: g.Cfg().MustGet(ctx, "tcpServe.address").String(),
		})

		if err != nil {
			g.Log().Warningf(ctx, "TCPServer start fail：%+v", err)
			return
		}

		s.serv = server

		// 消息队列路由
		s.serv.RegisterQueueRouter(map[string]tcp.RouterHandler{
			// ...
		})

		// 定时任务路由
		s.serv.RegisterCronRouter(map[string]tcp.RouterHandler{
			// ...
		})

		// 授权服务路由
		s.serv.RegisterAuthRouter(map[string]tcp.RouterHandler{
			"AuthSummary": s.onAuthSummary, // 获取授权信息
		})

		// 服务监听
		if err := s.serv.Listen(); err != nil {
			g.Log().Warningf(ctx, "TCPServer Listen err:%v", err)
		}
	})
}

// Stop 关闭服务
func (s *sTCPServer) Stop(ctx context.Context) {
	if s.serv != nil {
		s.serv.Close()
		g.Log().Debug(ctx, "TCPServer stop..")
	}
}
