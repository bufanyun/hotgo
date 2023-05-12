// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
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
)

// AdminAuth 后台鉴权中间件
func (s *sMiddleware) AdminAuth(r *ghttp.Request) {
	var (
		ctx    = r.Context()
		prefix = g.Cfg().MustGet(ctx, "router.admin.prefix", "/admin").String()
		path   = gstr.Replace(r.URL.Path, prefix, "", 1)
	)

	// 不需要验证登录的路由地址
	if isExceptLogin(ctx, consts.AppAdmin, path) {
		r.Middleware.Next()
		return
	}

	// 将用户信息传递到上下文中
	if err := deliverUserContext(r); err != nil {
		g.Log().Warningf(ctx, "deliverUserContext err:%+v", err)
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}

	// 不需要验证权限的路由地址
	if isExceptAuth(ctx, consts.AppAdmin, path) {
		r.Middleware.Next()
		return
	}

	// 验证路由访问权限
	if !service.AdminRole().Verify(ctx, path, r.Method) {
		g.Log().Debugf(ctx, "AdminAuth fail path:%+v, GetRoleKey:%+v, r.Method:%+v", path, contexts.GetRoleKey(ctx), r.Method)
		response.JsonExit(r, gcode.CodeSecurityReason.Code(), "你没有访问权限！")
		return
	}

	r.Middleware.Next()
}
