// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/controller/home/base"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// Home 前台页面路由
func Home(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().HomeAuth)
		// 允许通过根地址访问的路由可以加到这里，访问地址：http://127.0.0.1:8000
		group.Bind(
			base.Site, // 基础
		)

		// 默认访问地址：http://127.0.0.1:8000/home
		group.Group(simple.RouterPrefix(ctx, consts.AppHome), func(group *ghttp.RouterGroup) {
			group.Bind(
				base.Site, // 基础
			)
		})
	})
}
