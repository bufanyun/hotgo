package tcpclient

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// tcp授权
type sTCPAuth struct {
	client *tcp.Client
}

func init() {
	service.RegisterTCPAuth(newTCPAuth())
}

func newTCPAuth() *sTCPAuth {
	return &sTCPAuth{}
}

// Start 启动服务
func (s *sTCPAuth) Start(ctx context.Context) {
	g.Log().Debug(ctx, "TCPAuth start..")
	simple.SafeGo(ctx, func(ctx context.Context) {
		client, err := tcp.NewClient(&tcp.ClientConfig{
			Addr: "127.0.0.1:8099",
			Auth: &tcp.AuthMeta{
				Group:     "auth",
				Name:      "auth1",
				AppId:     "mengshuai",
				SecretKey: "123456",
			},
			LoginEvent: s.onLoginEvent,
			CloseEvent: s.onCloseEvent,
		})
		if err != nil {
			g.Log().Infof(ctx, "TCPAuth NewClient fail：%+v", err)
			return
		}

		s.client = client

		err = s.client.RegisterRouter(map[string]tcp.RouterHandler{
			// ...
		})

		if err != nil {
			g.Log().Infof(ctx, "TCPAuth RegisterRouter fail：%+v", err)
			return
		}

		if err = s.client.Start(); err != nil {
			g.Log().Infof(ctx, "TCPAuth Start fail：%+v", err)
			return
		}
	})
}

// Stop 停止服务
func (s *sTCPAuth) Stop(ctx context.Context) {
	if s.client != nil {
		s.client.Stop()
		g.Log().Debug(ctx, "TCPAuth stop..")
	}
}

// IsLogin 是否已登录认证
func (s *sTCPAuth) IsLogin() bool {
	return s.client.IsLogin
}

// onLoginEvent 登录认证成功事件
func (s *sTCPAuth) onLoginEvent() {
	// ...
}

// onCloseEvent 连接关闭回调事件
func (s *sTCPAuth) onCloseEvent() {
	// ...
}
