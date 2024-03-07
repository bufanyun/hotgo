// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// HomeAuth 前台页面鉴权中间件
func (s *sMiddleware) HomeAuth(r *ghttp.Request) {
	// 鉴权
	// ...

	r.Middleware.Next()
}
