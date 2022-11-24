// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package global

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/library/location"
	"os"
)

func Init(ctx context.Context) {
	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		fmt.Printf("时区设置异常err：%v \r\n", err)
		return
	}

	RootPtah, _ = os.Getwd()
	fmt.Printf("欢迎使用HotGo！\r\n当前运行环境：%v, 运行根路径为：%v \r\n", SysType, RootPtah)
	loadMonitor(ctx)
}

func loadMonitor(ctx context.Context) {
	err := grpool.AddWithRecover(ctx, func(ctx context.Context) {
		MonitorData.STartTime = gtime.Now()
		MonitorData.IntranetIP, _ = location.GetLocalIP()
		MonitorData.PublicIP, _ = location.GetPublicIP()

	})
	if err != nil {
		g.Log().Fatal(ctx, "global loadMonitor Fatal:", err)
		return
	}
}
