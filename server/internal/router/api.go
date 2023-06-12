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
	"hotgo/internal/controller/api/member"
	"hotgo/internal/controller/api/pay"
	"hotgo/internal/controller/api/user"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// Api 前台路由
func Api(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.AppApi), func(group *ghttp.RouterGroup) {
		group.Bind(
			user.Hello,
			pay.Notify, // 支付异步通知
		)
		group.Middleware(service.Middleware().ApiAuth)
		group.Bind(
			member.Member, // 管理员
		)
	})
}
