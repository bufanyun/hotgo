//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package router

import (
	"context"
	"github.com/bufanyun/hotgo/app/controller/apiController"
	"github.com/bufanyun/hotgo/app/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//
//  @Title  接口路由配置
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   group
//
func Api(ctx context.Context, group *ghttp.RouterGroup) {

	routerPrefix, _ := g.Cfg().Get(ctx, "router.api.prefix", "/api")

	group.Group(routerPrefix.String(), func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Instance().ApiAuth)
		group.Bind(
			apiController.Login,  // 登录
			apiController.Base,   // 基础
			apiController.Member, // 会员
			apiController.Dict,   // 字典
			apiController.Log,    // 日志
		)
	})
}
