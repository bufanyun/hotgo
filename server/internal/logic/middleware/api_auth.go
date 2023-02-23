// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
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
	"hotgo/internal/library/response"
	"hotgo/utility/auth"
)

// ApiAuth API鉴权中间件
func (s *sMiddleware) ApiAuth(r *ghttp.Request) {

	var (
		ctx = r.Context()
	)

	// 替换掉模块前缀
	routerPrefix := g.Cfg().MustGet(ctx, "router.api.prefix", "/api")
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

	//// 验证路由访问权限
	//verify := service.AdminRole().Verify(ctx, path, r.Method)
	//if !verify {
	//	response.JsonExit(r, consts.CodeSecurityReason, "你没有访问权限！")
	//	return
	//}

	r.Middleware.Next()
}
