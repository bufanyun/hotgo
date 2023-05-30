// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"hotgo/internal/library/addons"
	"hotgo/internal/library/casbin"
	"hotgo/internal/router"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
)

var (
	Http = &gcmd.Command{
		Name:  "http",
		Usage: "http",
		Brief: "HTTP服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 加载权限
			casbin.InitEnforcer(ctx)

			s := g.Server()

			// 错误状态码接管
			s.BindStatusHandler(404, func(r *ghttp.Request) {
				r.Response.Writeln("404 - 你似乎来到了没有知识存在的荒原…")
			})
			s.BindStatusHandler(403, func(r *ghttp.Request) {
				r.Response.Writeln("403 - 网站拒绝显示此网页")
			})

			// 初始化请求前回调
			s.BindHookHandler("/*any", ghttp.HookBeforeServe, service.Hook().BeforeServe)

			// 请求响应结束后回调
			s.BindHookHandler("/*any", ghttp.HookAfterOutput, service.Hook().AfterOutput)

			s.Group("/", func(group *ghttp.RouterGroup) {

				// 注册全局中间件
				group.Middleware(
					service.Middleware().Ctx, //必须第一个加载
					service.Middleware().CORS,
					service.Middleware().Blacklist,
					service.Middleware().DemoLimit,
					service.Middleware().ResponseHandler,
				)

				// 注册后台路由
				router.Admin(ctx, group)

				// 注册Api路由
				router.Api(ctx, group)

				// 注册websocket路由
				router.WebSocket(ctx, group)

				// 注册前台页面路由
				router.Home(ctx, group)

				// 注册插件路由
				addons.RegisterModulesRouter(ctx, group)
			})

			// 启动tcp服务
			service.TCPServer().Start(ctx)

			// https
			setSSL(ctx, s)

			serverWg.Add(1)

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			go func() {
				<-serverCloseSignal
				websocket.Stop()
				service.TCPServer().Stop(ctx)
				_ = s.Shutdown() // 主服务建议放在最后一个关闭
				g.Log().Debug(ctx, "http successfully closed ..")
				serverWg.Done()
			}()

			// Just run the server.
			s.Run()
			return
		},
	}
)

func setSSL(ctx context.Context, s *ghttp.Server) {
	config, err := service.SysConfig().GetLoadSSL(ctx)
	if err != nil {
		g.Log().Fatal(ctx, "ssl配置获取失败：err:%+v", err)
	}
	if config != nil && config.Switch {
		s.EnableHTTPS(config.CrtPath, config.KeyPath)
	}
}
