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
	"hotgo/internal/controller/api/member"
	"hotgo/internal/controller/api/user"
	"hotgo/internal/service"
)

// Api 前台路由
func Api(ctx context.Context, group *ghttp.RouterGroup) {

	routerPrefix, _ := g.Cfg().Get(ctx, "router.api.prefix", "/api")

	group.Group(routerPrefix.String(), func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().ApiAuth)
		group.Bind(
			user.Hello,
			member.Member, // 管理员
		)
	})
}
