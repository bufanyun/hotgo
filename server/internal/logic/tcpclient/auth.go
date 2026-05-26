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
	"hotgo/api/servmsg"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/servmsgin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// tcp客户端
type sAuthClient struct {
	client  *tcp.Client
	summary *servmsgin.AuthSummaryModel
}

func init() {
	service.RegisterAuthClient(newAuthClient())
}

func newAuthClient() *sAuthClient {
	return &sAuthClient{}
}

// Instance 获取实例
func (s *sAuthClient) Instance() *tcp.Client {
	return s.client
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

	// 创建客户端配置
	clientConfig := &tcp.ClientConfig{
		Addr:          config.Client.Auth.Address,
		AutoReconnect: true,
		Auth: &tcp.AuthMeta{
			Name: config.Client.Auth.Name,
			Extra: g.Map{
				"test": 13,
			},
			Group:     config.Client.Auth.Group,
			AppId:     config.Client.Auth.AppId,
			SecretKey: config.Client.Auth.SecretKey,
		},
		LoginEvent: s.onLoginEvent,
		CloseEvent: s.onCloseEvent,
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		s.client = tcp.NewClient(clientConfig)

		// 注册路由
		s.client.RegisterRouter(
			s.OnResponseAuthSummary,  // 响应授权信息
			s.OnResponseExampleHello, // 一个tcp请求例子
		)

		if err = s.client.Start(); err != nil {
			g.Log().Errorf(ctx, "AuthClient Start fail：%+v", err)
			return
		}
	})
}

// Stop 停止服务
func (s *sAuthClient) Stop(ctx context.Context) {
	if s.client != nil && !s.client.IsStop() {
		s.client.Stop()
		g.Log().Debug(ctx, "AuthClient stop..")
	}
}

// onLoginEvent 登录认证成功事件
func (s *sAuthClient) onLoginEvent() {
	ctx := gctx.New()

	// 获取授权信息
	_ = s.client.Send(ctx, &servmsg.AuthSummaryReq{})

	// 测试例子，实际使用时可以注释掉
	s.testExample(ctx)

	g.Log().Debug(ctx, "AuthClient login succeed.")
}

// onCloseEvent 连接关闭回调事件
func (s *sAuthClient) onCloseEvent() {
	g.Log().Debug(gctx.New(), "AuthClient closed.")
}

// Summary 获取授权信息
func (s *sAuthClient) Summary() *servmsgin.AuthSummaryModel {
	if s.summary == nil {
		s.summary = new(servmsgin.AuthSummaryModel)
	}
	return s.summary
}
