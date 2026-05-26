// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hook

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/library/contexts"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	"strings"
)

// 忽略的请求方式
var ignoredRequestMethods = []string{"HEAD", "PRI"}

// accessLog 访问日志
func (s *sHook) accessLog(r *ghttp.Request) {
	if s.isIgnoredRequest(r) {
		return
	}

	var ctx = r.Context()
	if contexts.Get(ctx) == nil {
		return
	}

	contexts.SetDataMap(ctx, g.Map{
		"request.takeUpTime": gtime.Now().Sub(gtime.New(r.EnterTime)).Milliseconds(),
		// ...
	})

	simple.SafeGo(ctx, func(ctx context.Context) {
		if err := service.SysLog().AutoLog(ctx); err != nil {
			g.Log().Infof(ctx, "hook accessLog err:%+v", err)
		}
	})
}

// isIgnoredRequest 是否忽略请求
func (s *sHook) isIgnoredRequest(r *ghttp.Request) bool {
	if r.IsFileRequest() {
		return true
	}

	if gstr.InArray(ignoredRequestMethods, strings.ToUpper(r.Method)) {
		return true
	}
	return false
}
