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
	"hotgo/internal/queues"
)

var (
	Queue = &gcmd.Command{
		Name:        "queue",
		Brief:       "消息队列",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Infof(ctx, "start queue consumer..")
			queues.Run(ctx)
			return
		},
	}
)
