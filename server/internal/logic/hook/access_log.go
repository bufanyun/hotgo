// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package hook

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/library/contexts"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// AccessLog 访问日志
func (s *sHook) AccessLog(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)

	// 没有上下文的请求不记录，如：doc、favicon.ico等非功能类业务
	modelCtx := contexts.Get(ctx)
	if modelCtx == nil {
		return
	}

	// 计算运行耗时
	contexts.SetTakeUpTime(ctx, gtime.TimestampMilli()-r.EnterTime)

	simple.SafeGo(ctx, func(ctx context.Context) {
		if err := service.SysLog().AutoLog(ctx); err != nil {
			g.Log().Warningf(ctx, "hook AccessLog err:%+v", err)
		}
	})
}
