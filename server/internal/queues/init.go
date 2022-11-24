// Package queues
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package queues

import (
	"context"
	"hotgo/internal/library/queue"
)

type jobStrategy interface {
	Listen(ctx context.Context)
	handle(ctx context.Context, mqMsg queue.MqMsg) (err error)
}

var (
	jobList = []jobStrategy{
		SysLog,
	}
)

func Run(ctx context.Context) {
	for _, job := range jobList {
		job.Listen(ctx)
	}
}
