// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package router

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/controller/home/base"
)

// Home 前台页面路由
func Home(ctx context.Context, group *ghttp.RouterGroup) {
	routerPrefix, _ := g.Cfg().Get(ctx, "router.home.prefix", "/home")

	group.Group(routerPrefix.String(), func(group *ghttp.RouterGroup) {
		group.Bind(
			base.Site, // 基础
		)

	})
}
