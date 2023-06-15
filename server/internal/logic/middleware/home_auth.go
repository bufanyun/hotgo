package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// HomeAuth 前台页面鉴权中间件
func (s *sMiddleware) HomeAuth(r *ghttp.Request) {
	r.Response.Header().Set("Content-Type", "text/html")

	// 鉴权
	// ...

	r.Middleware.Next()
}
