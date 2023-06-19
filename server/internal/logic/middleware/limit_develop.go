package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/library/location"
	"hotgo/internal/library/response"
)

// Develop 开发工具白名单过滤
func (s *sMiddleware) Develop(r *ghttp.Request) {
	ips := g.Cfg().MustGet(r.Context(), "hggen.allowedIPs").Strings()
	if len(ips) == 0 {
		response.JsonExit(r, gcode.CodeNotSupported.Code(), "请配置生成白名单！")
		return
	}

	if !gstr.InArray(ips, "*") {
		clientIp := location.GetClientIp(r)
		ok := false
		for _, ip := range ips {
			if ip == clientIp {
				ok = true
				break
			}
		}

		if !ok {
			response.JsonExit(r, gcode.CodeNotSupported.Code(), fmt.Sprintf("当前IP[%s]没有配置生成白名单！", clientIp))
			return
		}
	}

	r.Middleware.Next()
}
