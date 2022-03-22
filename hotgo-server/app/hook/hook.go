//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package hook

import (
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/service/sysService"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	// sHook is service struct of module Hook.
	sHook struct{}
)

var (
	// insHook is the instance of service Hook.
	insHook = sHook{}
)

// Hook returns the interface of Hook service.
func Instance() *sHook {
	return &insHook
}

//
//  @Title  全局日志
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   r
//
func (s *sHook) GlobalLog(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)

	// 没有上下文的请求不记录，如：doc、favicon.ico等非服务类业务
	modelCtx := com.Context.Get(ctx)
	if modelCtx == nil {
		return
	}

	// 计算运行耗时
	com.Context.SetTakeUpTime(ctx, gtime.TimestampMilli()-r.EnterTime)

	go sysService.Log.AutoLog(ctx)
}
