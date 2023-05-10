package tcpclient

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/msgin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// tcp客户端
type sAuthClient struct {
	client  *tcp.Client
	summary *msgin.AuthSummaryData
}

func init() {
	service.RegisterAuthClient(newAuthClient())
}

func newAuthClient() *sAuthClient {
	return &sAuthClient{}
}

// Start 启动服务
func (s *sAuthClient) Start(ctx context.Context) {
	g.Log().Debug(ctx, "AuthClient start..")

	config, err := service.SysConfig().GetLoadTCP(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "AuthClient start fail:%+v", err)
		return
	}

	if config.Client == nil || config.Client.Auth == nil {
		g.Log().Errorf(ctx, "AuthClient config is invalid")
		return
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		s.client, err = tcp.NewClient(&tcp.ClientConfig{
			Addr: config.Client.Auth.Address,
			Auth: &tcp.AuthMeta{
				Group:     config.Client.Auth.Group,
				Name:      config.Client.Auth.Name,
				AppId:     config.Client.Auth.AppId,
				SecretKey: config.Client.Auth.SecretKey,
			},
			LoginEvent: s.onLoginEvent,
			CloseEvent: s.onCloseEvent,
		})
		if err != nil {
			g.Log().Errorf(ctx, "AuthClient NewClient fail：%+v", err)
			return
		}

		err = s.client.RegisterRouter(map[string]tcp.RouterHandler{
			"ResponseAuthSummary": s.OnResponseAuthSummary,
		})

		if err != nil {
			g.Log().Errorf(ctx, "AuthClient RegisterRouter fail：%+v", err)
			return
		}

		if err = s.client.Start(); err != nil {
			g.Log().Errorf(ctx, "AuthClient Start fail：%+v", err)
			return
		}
	})

}

// Stop 停止服务
func (s *sAuthClient) Stop(ctx context.Context) {
	if s.client != nil {
		s.client.Stop()
		g.Log().Debug(ctx, "AuthClient stop..")
	}
}

// IsLogin 是否已登录认证
func (s *sAuthClient) IsLogin() bool {
	if s.client == nil {
		return false
	}
	return s.client.IsLogin
}

// onLoginEvent 登录认证成功事件
func (s *sAuthClient) onLoginEvent() {

	// 获取授权数据
	s.client.Send(s.client.Ctx, &msgin.AuthSummary{})
}

// onCloseEvent 连接关闭回调事件
func (s *sAuthClient) onCloseEvent() {
	// ...
}
