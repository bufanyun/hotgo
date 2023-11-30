// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
	"runtime"
	"strings"
)

func Init(ctx context.Context) {
	// 设置gf运行模式
	SetGFMode(ctx)

	// 设置服务日志处理
	glog.SetDefaultHandler(LoggingServeLogHandler)

	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err：%+v", err)
		return
	}

	fmt.Printf("欢迎使用HotGo！\r\n当前运行环境：%v, 运行根路径为：%v \r\nHotGo版本：v%v, gf版本：%v \n", runtime.GOOS, gfile.Pwd(), consts.VersionApp, gf.VERSION)

	// 初始化链路追踪
	InitTrace(ctx)

	// 设置缓存适配器
	cache.SetAdapter(ctx)

	// 初始化功能库配置
	service.SysConfig().InitConfig(ctx)

	// 加载超管数据
	service.AdminMember().LoadSuperAdmin(ctx)

	// 订阅集群同步
	SubscribeClusterSync(ctx)
}

// LoggingServeLogHandler 服务日志处理
// 需要将异常日志保存到服务日志时可以通过SetHandlers设置此方法
func LoggingServeLogHandler(ctx context.Context, in *glog.HandlerInput) {
	in.Next(ctx)

	err := g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				panic(err)
			}
		}()

		// web服务日志不做记录，因为会导致重复记录
		r := g.RequestFromCtx(ctx)
		if r != nil && r.Server != nil && in.Logger.GetConfig().Path == r.Server.Logger().GetConfig().Path {
			return
		}

		conf, err := service.SysConfig().GetLoadServeLog(ctx)
		if err != nil {
			return
		}

		if conf == nil {
			return
		}

		if !conf.Switch {
			return
		}

		if in.LevelFormat == "" || !gstr.InArray(conf.LevelFormat, in.LevelFormat) {
			return
		}

		if in.Stack == "" {
			in.Stack = in.Logger.GetStack()
		}

		if len(in.Content) == 0 {
			in.Content = gstr.StrLimit(gvar.New(in.Values).String(), consts.MaxServeLogContentLen)
		}

		var data entity.SysServeLog
		data.TraceId = gctx.CtxId(ctx)
		data.LevelFormat = in.LevelFormat
		data.Content = in.Content
		data.Stack = gjson.New(charset.ParseStack(in.Stack))
		data.Line = strings.TrimRight(in.CallerPath, ":")
		data.TriggerNs = in.Time.UnixNano()
		data.Status = consts.StatusEnabled

		if gstr.Contains(in.Content, `exception recovered`) {
			data.LevelFormat = "PANI"
		}

		if data.Stack.IsNil() {
			data.Stack = gjson.New(consts.NilJsonToString)
		}

		if conf.Queue {
			err = queue.Push(consts.QueueServeLogTopic, data)
		} else {
			err = service.SysServeLog().RealWrite(ctx, data)
		}
	})

	if err != nil {
		g.Dump("LoggingServeLogHandler err:", err)
	}
}

// InitTrace 初始化链路追踪
func InitTrace(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, "jaeger.switch").Bool() {
		return
	}

	tp, err := jaeger.Init(simple.AppName(ctx), g.Cfg().MustGet(ctx, "jaeger.endpoint").String())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		_ = tp.Shutdown(ctx)
		g.Log().Debug(ctx, "jaeger closed ..")
	})
}

// SetGFMode 设置gf运行模式
func SetGFMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "hotgo.mode").String()
	if len(mode) == 0 {
		mode = gmode.NOT_SET
	}

	var modes = []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}

	// 如果是有效的运行模式，就进行设置
	if validate.InSlice(modes, mode) {
		gmode.Set(mode)
	}
}
