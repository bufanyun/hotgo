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
	"hotgo/internal/library/queue"
	"hotgo/utility/simple"
)

var (
	Queue = &gcmd.Command{
		Name:        "queue",
		Brief:       "消息队列",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			simple.SafeGo(ctx, func(ctx context.Context) {
				g.Log().Debug(ctx, "start queue consumer..")
				queue.StartConsumersListener(ctx)
				g.Log().Debug(ctx, "start queue consumer success..")
			})

			serverWg.Add(1)

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			select {
			case <-serverCloseSignal:
				serverWg.Done()
			}

			g.Log().Debug(ctx, "queue successfully closed ..")
			return
		},
	}
)
