// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/response"
	"hotgo/internal/service"
	"hotgo/utility/auth"
)

// AdminAuth 后台鉴权中间件
func (s *sMiddleware) AdminAuth(r *ghttp.Request) {

	var (
		ctx = r.Context()
	)

	// 替换掉模块前缀
	routerPrefix, _ := g.Cfg().Get(ctx, "router.admin.prefix", "/admin")
	path := gstr.Replace(r.URL.Path, routerPrefix.String(), "", 1)

	/// 不需要验证登录的路由地址
	if auth.IsExceptLogin(ctx, path) {
		r.Middleware.Next()
		return
	}

	if err := inspectAuth(r, consts.AppAdmin); err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}

	// 验证路由访问权限
	if !service.AdminRole().Verify(ctx, path, r.Method) {
		g.Log().Warningf(ctx, "AdminAuth fail path:%+v, GetRoleKey:%+v, r.Method:%+v", path, contexts.GetRoleKey(ctx), r.Method)
		response.JsonExit(r, gcode.CodeSecurityReason.Code(), "你没有访问权限！")
		return
	}

	r.Middleware.Next()
}
