// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	api "hotgo/api/home/base"
	"hotgo/internal/consts"
	"hotgo/internal/controller/home/base"
	"hotgo/utility/simple"
)

// Home 前台页面路由
func Home(ctx context.Context, group *ghttp.RouterGroup) {
	// 注册首页路由
	group.ALL("/", func(r *ghttp.Request) {
		_, _ = base.Site.Index(r.Context(), &api.SiteIndexReq{})
	})

	group.Group(simple.RouterPrefix(ctx, consts.AppHome), func(group *ghttp.RouterGroup) {
		group.Bind(
			base.Site, // 基础
		)
	})
}
