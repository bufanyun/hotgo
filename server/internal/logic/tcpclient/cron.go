package tcpclient

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// tcp客户端
type sCronClient struct {
	client *tcp.Client
}

func init() {
	service.RegisterCronClient(newCronClient())
}

func newCronClient() *sCronClient {
	return &sCronClient{}
}

// Start 启动服务
func (s *sCronClient) Start(ctx context.Context) {
	g.Log().Debug(ctx, "CronClient start..")

	config, err := service.SysConfig().GetLoadTCP(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "CronClient start fail:%+v", err)
		return
	}

	if config == nil || config.Client == nil || config.Client.Cron == nil {
		g.Log().Errorf(ctx, "CronClient config is invalid")
		return
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		client, err := tcp.NewClient(&tcp.ClientConfig{
			Addr: config.Client.Cron.Address,
			Auth: &tcp.AuthMeta{
				Group:     config.Client.Cron.Group,
				Name:      config.Client.Cron.Name,
				AppId:     config.Client.Cron.AppId,
				SecretKey: config.Client.Cron.SecretKey,
			},
			LoginEvent: s.onLoginEvent,
			CloseEvent: s.onCloseEvent,
		})
		if err != nil {
			g.Log().Errorf(ctx, "CronClient NewClient fail：%+v", err)
			return
		}

		s.client = client

		err = s.client.RegisterRouter(map[string]tcp.RouterHandler{
			"CronDelete":     s.OnCronDelete,     // 删除任务
			"CronEdit":       s.OnCronEdit,       // 编辑任务
			"CronStatus":     s.OnCronStatus,     // 修改任务状态
			"CronOnlineExec": s.OnCronOnlineExec, // 执行一次任务
		})

		if err != nil {
			g.Log().Errorf(ctx, "CronClient RegisterRouter fail：%+v", err)
			return
		}

		if err = s.client.Start(); err != nil {
			g.Log().Errorf(ctx, "CronClient Start fail：%+v", err)
			return
		}
	})

}

// Stop 停止服务
func (s *sCronClient) Stop(ctx context.Context) {
	if s.client != nil {
		s.client.Stop()
		g.Log().Debug(ctx, "CronClient stop..")
	}
}

// IsLogin 是否已登录认证
func (s *sCronClient) IsLogin() bool {
	if s.client == nil {
		return false
	}
	return s.client.IsLogin
}

// onLoginEvent 登录认证成功事件
func (s *sCronClient) onLoginEvent() {
	// ...
}

// onCloseEvent 连接关闭回调事件
func (s *sCronClient) onCloseEvent() {
	// ...
}
