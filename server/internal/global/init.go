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
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/library/hggen"
	"hotgo/internal/library/location"
	"hotgo/utility/simple"
	"os"
)

func Init(ctx context.Context) {
	if _, err := g.Cfg().Get(ctx, "hotgo.debug"); err != nil {
		g.Log().Fatal(ctx, "配置读取异常:", err, "\r\n你确定 config/config.yaml 文件存在且格式正确吗？\r\n")
	}
	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常err：%+v", err)
		return
	}

	RootPtah, _ = os.Getwd()
	fmt.Printf("欢迎使用HotGo！\r\n当前运行环境：%v, 运行根路径为：%v \r\nHotGo版本：v%v, gf版本：%v \n", SysType, RootPtah, consts.VersionApp, gf.VERSION)

	setOrmCacheAdapter()

	startMonitor(ctx)

	hggen.InIt(ctx)

}

func startMonitor(ctx context.Context) {
	simple.SafeGo(ctx, func(ctx context.Context) {
		MonitorData.STartTime = gtime.Now()
		intranetIP, err := location.GetLocalIP()
		if err != nil {
			g.Log().Warningf(ctx, "parse intranetIP err:%+v", err)
		}
		MonitorData.IntranetIP = intranetIP

		publicIP, err := location.GetPublicIP(ctx)
		if err != nil {
			g.Log().Warningf(ctx, "parse publicIP err:%+v", err)
		}
		MonitorData.PublicIP = publicIP
	})
}

func setOrmCacheAdapter() {
	redisCache := gcache.NewAdapterRedis(g.Redis())
	g.DB().GetCache().SetAdapter(redisCache)
}
