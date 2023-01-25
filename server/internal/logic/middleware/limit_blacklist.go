package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/global"
	"hotgo/internal/library/location"
	"hotgo/internal/library/response"
)

// Blacklist IP黑名单限制中间件
func (s *sMiddleware) Blacklist(r *ghttp.Request) {
	if global.Blacklists != nil {
		if _, ok := global.Blacklists[location.GetClientIp(r)]; ok {
			response.JsonExit(r, gcode.CodeServerBusy.Code(), "请求异常，已被封禁，如有疑问请联系管理员！")
		}
	} else {
		g.Log().Warningf(r.Context(), "blacklists uninitialized")
	}

	r.Middleware.Next()
}
