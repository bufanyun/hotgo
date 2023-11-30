// Package cron
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cron

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"strings"
)

// GenCronSn 生成任务序列号
func GenCronSn(sysCron *entity.SysCron) string {
	return fmt.Sprintf("%s@%d", sysCron.Name, sysCron.Id)
}

// GenCronCtx 生成任务上下文
func GenCronCtx(sysCron *entity.SysCron) (ctx context.Context) {
	ctx = context.WithValue(gctx.New(), consts.ContextKeyCronArgs, strings.Split(sysCron.Params, consts.CronSplitStr))
	ctx = context.WithValue(ctx, consts.ContextKeyCronSn, GenCronSn(sysCron))
	return ctx
}

func GenLoggerByCtx(ctx context.Context) *glog.Logger {
	sn, ok := ctx.Value(consts.ContextKeyCronSn).(string)
	if !ok {
		Logger().Panic(ctx, "获取定时任务序列号失败!")
	}

	logger, ok := crons.loggers[sn]
	if ok {
		return logger
	}

	logger = glog.New()
	if err := logger.SetConfig(Logger().GetConfig()); err != nil {
		Logger().Panic(ctx, err)
	}

	logger.SetFlags(glog.F_TIME_STD | glog.F_FILE_SHORT)
	// 设置子路径
	if err := logger.SetPath(fmt.Sprintf("%v/%v", logger.GetPath(), sn)); err != nil {
		Logger().Panic(ctx, err)
	}

	crons.Lock()
	defer crons.Unlock()
	crons.loggers[sn] = logger
	return logger
}

// GenExecuteFun 生成执行过程
func GenExecuteFun(fun func(ctx context.Context, parser *Parser) (err error)) func(ctx context.Context) {
	return func(ctx context.Context) {
		args, ok := ctx.Value(consts.ContextKeyCronArgs).([]string)
		if !ok {
			Logger().Panic(ctx, "执行定时任务时，参数解析失败!")
			return
		}
		parser := new(Parser)
		parser.Args = args
		parser.Logger = GenLoggerByCtx(ctx)

		st := gtime.Now()
		err := g.Try(ctx, func(ctx context.Context) {
			if err := fun(ctx, parser); err != nil {
				panic(err)
			}
		})

		milliseconds := gtime.Now().Sub(st).Milliseconds() // 执行耗时
		if err != nil {
			parser.Logger.Errorf(ctx, "execute failed, took %vms, err:%+v", milliseconds, err)
			return
		}
		parser.Logger.Infof(ctx, "execute success, took %vms.", milliseconds)
	}
}
