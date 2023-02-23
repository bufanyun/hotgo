// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
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
	prefix := g.Cfg().MustGet(ctx, "router.api.prefix", "/api")
	group.Group(prefix.String(), func(group *ghttp.RouterGroup) {
		group.Bind(
			user.Hello,
		)
		group.Middleware(service.Middleware().ApiAuth)
		group.Bind(
			member.Member, // 管理员
		)
	})
}
