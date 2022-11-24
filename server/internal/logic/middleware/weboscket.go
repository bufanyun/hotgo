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
	"hotgo/internal/library/response"
	"hotgo/utility/auth"
)

// WebSocketToken 检查ws连接token
func (s *sMiddleware) WebSocketToken(r *ghttp.Request) {

	var (
		ctx = r.Context()
	)

	// 替换掉模块前缀
	routerPrefix, _ := g.Cfg().Get(ctx, "router.ws.prefix", "/socket")
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

	r.Middleware.Next()
}
