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
	"hotgo/internal/library/response"
)

// WebSocketAuth websocket鉴权中间件
func (s *sMiddleware) WebSocketAuth(r *ghttp.Request) {
	var (
		ctx    = r.Context()
		prefix = g.Cfg().MustGet(ctx, "router.websocket.prefix", "/websocket").String()
		path   = gstr.Replace(r.URL.Path, prefix, "", 1)
	)

	// 不需要验证登录的路由地址
	if isExceptLogin(ctx, consts.AppWebSocket, path) {
		r.Middleware.Next()
		return
	}

	// 将用户信息传递到上下文中
	if err := deliverUserContext(r); err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}

	r.Middleware.Next()
}
