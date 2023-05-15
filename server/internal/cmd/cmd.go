// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"hotgo/utility/simple"
)

var (
	serverCloseSignal chan struct{}
	Main              = &gcmd.Command{
		Description: `默认启动所有服务`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return All.Func(ctx, parser)
		},
	}

	Help = &gcmd.Command{
		Name:  "help",
		Brief: "查看帮助",
		Description: `
		命令提示符
		---------------------------------------------------------------------------------
		启动服务
		>> 所有服务  [go run main.go]   热编译  [gf run main.go]
		>> HTTP服务  [go run main.go http]
		>> 消息队列  [go run main.go queue]
		>> 定时任务  [go run main.go cron]
		>> 查看帮助  [go run main.go help]

		---------------------------------------------------------------------------------
		工具
		>> 释放casbin权限，用于清理无效的权限设置  [go run main.go tools -m=casbin -a1=refresh]

		---------------------------------------------------------------------------------
		更多
       	github地址：https://github.com/bufanyun/hotgo
		文档地址：https://github.com/bufanyun/hotgo/tree/v2.0/docs/guide-zh-CN	
		HotGo框架交流1群：190966648
    `,
	}

	All = &gcmd.Command{
		Name:        "all",
		Brief:       "start all server",
		Description: "this is the command entry for starting all server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "starting all server")

			simple.SafeGo(ctx, func(ctx context.Context) {
				if err := Queue.Func(ctx, parser); err != nil {
					g.Log().Fatal(ctx, "queue consumer start fail:", err)
				}
			})

			simple.SafeGo(ctx, func(ctx context.Context) {
				if err := Cron.Func(ctx, parser); err != nil {
					g.Log().Fatal(ctx, "cron start fail:", err)
				}
			})

			simple.SafeGo(ctx, func(ctx context.Context) {
				if err := Http.Func(ctx, parser); err != nil {
					g.Log().Fatal(ctx, "http server start fail:", err)
				}
			})

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			select {
			case <-serverCloseSignal:
				// ...
			}

			g.Log().Debug(ctx, "service successfully closed ..")
			return
		},
	}
)

func init() {
	if err := Main.AddCommand(All, Http, Queue, Cron, Auth, Tools, Help); err != nil {
		panic(err)
	}
	serverCloseSignal = make(chan struct{}, 1)
}
