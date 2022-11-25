// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
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
		Description: `
		欢迎使用HotGo!
		---------------------------------------------------------------------------------
		启动服务
		>> HTTP服务  [go run main.go http]
		>> 消息队列  [go run main.go queue]
		>> 所有服务  [go run main.go all]

		---------------------------------------------------------------------------------
		工具
		>> 释放casbin权限，用于清理无效的权限设置  [go run main.go tools -m=casbin -a1=refresh]
`,
	}

	Help = &gcmd.Command{
		Name:  "help",
		Brief: "查看帮助",
		Description: `
       欢迎使用 HotGo
       当前版本:v2.0.0
    `,
	}

	All = &gcmd.Command{
		Name:        "all",
		Brief:       "start all server",
		Description: "this is the command entry for starting all server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "start all server")

			simple.SafeGo(ctx, func(ctx context.Context) {
				if err := Http.Func(ctx, parser); err != nil {
					g.Log().Fatal(ctx, "http server start fail:", err)
				}
			})

			simple.SafeGo(ctx, func(ctx context.Context) {
				if err := Queue.Func(ctx, parser); err != nil {
					g.Log().Fatal(ctx, "queue consumer start fail:", err)
				}
			})

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			select {
			case <-serverCloseSignal:
				// ...
			}

			g.Log().Info(ctx, "service successfully closed ..")
			return
		},
	}
)

func init() {
	if err := Main.AddCommand(Http, Queue, Tools, All, Help); err != nil {
		panic(err)
	}
	serverCloseSignal = make(chan struct{}, 1)
}
