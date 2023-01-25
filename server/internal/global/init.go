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
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/library/hggen"
	"hotgo/internal/library/location"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
	"os"
)

func Init(ctx context.Context) {
	_, err := g.Cfg().Get(ctx, "hotgo.debug")
	if err != nil {
		g.Log().Fatal(ctx, "配置读取异常:", err, "\r\n你确定 config/config.yaml 文件存在且格式正确吗？\r\n")
		return
	}
	//g.SetDebug(debug.Bool())

	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常err：%+v", err)
		return
	}

	RootPtah, _ = os.Getwd()
	fmt.Printf("欢迎使用HotGo！\r\n当前运行环境：%v, 运行根路径为：%v \r\nHotGo版本：v%v, gf版本：%v \n", SysType, RootPtah, consts.VersionApp, gf.VERSION)

	g.Log().SetHandlers(LoggingServeLogHandler)

	setOrmCacheAdapter()

	service.SysBlacklist().Load(ctx)

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

func LoggingServeLogHandler(ctx context.Context, in *glog.HandlerInput) {
	in.Next(ctx)

	conf, err := service.SysConfig().GetLoadServeLog(ctx)
	if err != nil {
		return
	}

	if !conf.Switch {
		return
	}

	if in.LevelFormat == "" || !gstr.InArray(conf.LevelFormat, in.LevelFormat) {
		return
	}

	var data entity.SysServeLog
	data.TraceId = gctx.CtxId(ctx)
	data.LevelFormat = in.LevelFormat
	data.Content = in.Content
	data.Stack = gjson.New(charset.ParseStack(in.Stack))
	data.Line = in.CallerPath
	data.TriggerNs = in.Time.UnixNano()
	data.Status = consts.StatusEnabled

	if data.Stack.IsNil() {
		data.Stack = gjson.New(consts.NilJsonToString)
	}

	if gstr.Contains(in.Content, `exception recovered`) {
		data.LevelFormat = "PANI"
	}

	if conf.Queue {
		err = queue.Push(consts.QueueServeLogTopic, data)
	} else {
		err = service.SysServeLog().RealWrite(ctx, data)
	}

	if err != nil {
		g.Log().Printf(ctx, "LoggingServeLogHandler err:%+v", err)
	}
}
