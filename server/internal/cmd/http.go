// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	baseApi "hotgo/api/home/base"
	"hotgo/internal/controller/home/base"
	"hotgo/internal/library/casbin"
	"hotgo/internal/model"
	"hotgo/internal/router"
	"hotgo/internal/service"
)

var (
	Http = &gcmd.Command{
		Name:  "http",
		Usage: "http",
		Brief: "HTTP服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			if _, err := g.Cfg().Get(ctx, "hotgo.debug"); err != nil {
				g.Log().Fatal(ctx, "配置读取异常:", err, "\r\n你确定 config/config.yaml 文件存在且格式正确吗？\r\n")
			}

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

			// 请求结束事件回调
			s.BindHookHandler("/*any", ghttp.HookAfterOutput, service.Hook().GlobalLog)

			s.Group("/", func(group *ghttp.RouterGroup) {

				// 注册全局中间件
				group.Middleware(
					service.Middleware().Ctx, //必须第一个加载
					service.Middleware().CORS,
					service.Middleware().DemoLimit,
					service.Middleware().ResponseHandler,
				)

				// 注册默认首页路由
				group.ALL("/", func(r *ghttp.Request) {
					_, _ = base.Site.Index(r.Context(), &baseApi.SiteIndexReq{})
					return
				})

				group.ALL("/login", func(r *ghttp.Request) {
					r.Response.RedirectTo("/admin")
				})

				// 注册后台路由
				router.Admin(ctx, group)

				// 注册前台路由
				router.Api(ctx, group)

				// 注册websocket路由
				router.WebSocket(ctx, group)

				// 注册前台页面路由
				router.Home(ctx, group)
			})

			// 启动定时任务
			service.SysCron().StartCron(ctx)

			// 信号监听
			signalListen(ctx, signalHandlerForCron, signalHandlerForWebSocket)

			// 开启https访问
			var (
				sSLConfig *model.SSLConfig
				ssl, _    = g.Cfg().Get(ctx, "hotgo.ssl")
			)
			if err := ssl.Struct(&sSLConfig); err != nil {
				g.Log().Fatalf(ctx, "hotgo启动失败, ssl err:", err)
				return err
			}
			if sSLConfig != nil && sSLConfig.Switch {
				s.EnableHTTPS(sSLConfig.CrtPath, sSLConfig.KeyPath)
			}

			// Just run the server.
			s.Run()

			return nil
		},
	}
)
