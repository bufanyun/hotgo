//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package router

import (
	"context"
	"github.com/bufanyun/hotgo/app/controller/adminController"
	"github.com/bufanyun/hotgo/app/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//
//  @Title  后台路由配置
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   group
//
func Admin(ctx context.Context, group *ghttp.RouterGroup) {

	routerPrefix, _ := g.Cfg().Get(ctx, "router.admin.prefix", "/admin")

	group.Group(routerPrefix.String(), func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Instance().AdminAuth)
		group.Bind(
			adminController.Login,  // 登录
			adminController.Role,   // 路由
			adminController.Member, // 会员
			adminController.Menu,   // 菜单
			adminController.Log,    // 日志
			adminController.Dict,   // 字典
			adminController.Post,   // 字典
			adminController.Dept,   // 部门
			adminController.Config, // 配置
			adminController.Notice, // 公告
		)
	})
}
