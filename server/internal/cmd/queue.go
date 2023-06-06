// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
	"hotgo/internal/global"
	"hotgo/internal/library/queue"
	"hotgo/utility/simple"
)

var (
	Queue = &gcmd.Command{
		Name:        "queue",
		Brief:       "消息队列",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 服务日志处理
			queue.Logger().SetHandlers(global.LoggingServeLogHandler)

			simple.SafeGo(ctx, func(ctx context.Context) {
				queue.Logger().Debug(ctx, "start queue consumer..")
				queue.StartConsumersListener(ctx)
				queue.Logger().Debug(ctx, "start queue consumer success..")
			})

			serverWg.Add(1)

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			<-serverCloseSignal
			queue.Logger().Debug(ctx, "queue successfully closed ..")
			serverWg.Done()
			return
		},
	}
)
