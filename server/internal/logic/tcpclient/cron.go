// Package tcpclient
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpclient

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
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

// Instance 获取实例
func (s *sCronClient) Instance() *tcp.Client {
	return s.client
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

	// 创建客户端配置
	clientConfig := &tcp.ClientConfig{
		Addr:          config.Client.Cron.Address,
		AutoReconnect: true,
		Auth: &tcp.AuthMeta{
			Name:      config.Client.Cron.Name,
			Group:     config.Client.Cron.Group,
			AppId:     config.Client.Cron.AppId,
			SecretKey: config.Client.Cron.SecretKey,
		},
		LoginEvent: s.onLoginEvent,
		CloseEvent: s.onCloseEvent,
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		s.client = tcp.NewClient(clientConfig)

		// 注册RPC路由
		s.client.RegisterRPCRouter(
			s.OnCronDelete,      // 删除任务
			s.OnCronEdit,        // 编辑任务
			s.OnCronStatus,      // 修改任务状态
			s.OnCronOnlineExec,  // 执行一次任务
			s.OnCronDispatchLog, // 查看调度日志
		)

		// 注册拦截器
		s.client.RegisterInterceptor(s.DefaultInterceptor)

		if err = s.client.Start(); err != nil {
			g.Log().Errorf(ctx, "CronClient Start fail：%+v", err)
			return
		}
	})
}

// Stop 停止服务
func (s *sCronClient) Stop(ctx context.Context) {
	if s.client != nil && !s.client.IsStop() {
		s.client.Stop()
		g.Log().Debug(ctx, "CronClient stop..")
	}
}

// onLoginEvent 登录认证成功事件
func (s *sCronClient) onLoginEvent() {
	g.Log().Debug(gctx.New(), "CronClient login succeed.")
}

// onCloseEvent 连接关闭回调事件
func (s *sCronClient) onCloseEvent() {
	g.Log().Debug(gctx.New(), "CronClient closed.")
}
