//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package boot

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/hook"
	"github.com/bufanyun/hotgo/app/middleware"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/router"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/protocol/goai"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of HotGo!",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			if _, err := g.Cfg().Get(ctx, "hotgo.debug"); err != nil {
				g.Log().Fatal(ctx, "配置读取异常:", err, "\r\n你确定 config/config.yaml 文件存在且格式正确吗？\r\n")
			}

			s := g.Server()

			// 错误状态码接管
			s.BindStatusHandler(404, func(r *ghttp.Request) {
				r.Response.Writeln("404 - 你似乎来到了没有知识存在的荒原…")
			})
			s.BindStatusHandler(403, func(r *ghttp.Request) {
				r.Response.Writeln("403 - 网站拒绝显示此网页")
			})

			// 请求结束事件回调
			s.BindHookHandler("/*any", ghttp.HookAfterOutput, hook.Instance().GlobalLog)

			s.Group("/", func(group *ghttp.RouterGroup) {

				// 注册全局中间件
				group.Middleware(
					middleware.Instance().Ctx, //必须第一个加载
					middleware.Instance().CORS,
					middleware.Instance().HandlerResponse,
				)

				// 注册默认首页路由
				group.ALL("/", func(r *ghttp.Request) {
					r.Response.Write("hello hotGo!!")
				})

				// 注册后台路由
				router.Admin(ctx, group)

				// 注册API路由
				router.Api(ctx, group)

			})

			// Custom enhance API document.
			enhanceOpenAPIDoc(s)

			// 消息队列
			QueueListen(ctx)

			// Just run the server.
			s.Run()
			return nil
		},
	}
)

//
//  @Title  API document
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   s
//
func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = model.Response{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: consts.OpenAPIName,
			URL:  consts.OpenAPIURL,
		},
	}
}
