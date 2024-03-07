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
	"github.com/gogf/gf/v2/util/gmode"
	"hotgo/internal/library/addons"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/hggen"
	"hotgo/internal/router"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
)

var (
	Http = &gcmd.Command{
		Name:  "http",
		Usage: "http",
		Brief: "HTTP服务，也可以称为主服务，包含http、websocket、tcpserver多个可对外服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化http服务
			s := g.Server()

			// 初始化请求前回调
			s.BindHookHandler("/*any", ghttp.HookBeforeServe, service.Hook().BeforeServe)

			// 请求响应结束后回调
			s.BindHookHandler("/*any", ghttp.HookAfterOutput, service.Hook().AfterOutput)

			// 注册全局中间件
			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				service.Middleware().Ctx,             // 初始化请求上下文，一般需要第一个进行加载，后续中间件存在依赖关系
				service.Middleware().CORS,            // 跨域中间件，自动处理跨域问题
				service.Middleware().Blacklist,       // IP黑名单中间件，如果请求IP被后台拉黑，所有请求将被拒绝
				service.Middleware().DemoLimit,       // 演示系統操作限制，当开启演示模式时，所有POST请求将被拒绝
				service.Middleware().PreFilter,       // 请求输入预处理，api使用gf规范路由并且XxxReq结构体实现了validate.Filter接口即可隐式预处理
				service.Middleware().ResponseHandler, // HTTP响应预处理，在业务处理完成后，对响应结果进行格式化和错误过滤，将处理后的数据发送给请求方
			}...)

			s.Group("/", func(group *ghttp.RouterGroup) {
				// 注册后台路由
				router.Admin(ctx, group)

				// 注册Api路由
				router.Api(ctx, group)

				// 注册websocket路由
				router.WebSocket(ctx, group)

				// 注册前台页面路由
				router.Home(ctx, group)
			})

			// 启动插件
			if err = addons.StartModules(ctx, &addons.Option{Server: s}); err != nil {
				return err
			}

			// 初始化casbin权限
			casbin.InitEnforcer(ctx)

			// 初始化生成代码配置
			if !gmode.IsProduct() {
				hggen.InIt(ctx)
			}

			// 启动tcp服务
			service.TCPServer().Start(ctx)

			// 启动服务监控
			service.AdminMonitor().StartMonitor(ctx)

			// 加载ip访问黑名单
			service.SysBlacklist().Load(ctx)

			// 注册支付成功回调方法
			service.Pay().RegisterNotifyCall()

			serverWg.Add(1)

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			go func() {
				<-serverCloseSignal
				websocket.Stop()              // 关闭websocket
				service.TCPServer().Stop(ctx) // 关闭tcp服务器
				addons.StopModules(ctx)       // 停止插件
				_ = s.Shutdown()              // 关闭http服务，主服务建议放在最后一个关闭
				g.Log().Debug(ctx, "http successfully closed ..")
				serverWg.Done()
			}()

			// Just run the server.
			s.Run()
			return
		},
	}
)
