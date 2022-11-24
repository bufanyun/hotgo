// Package queues
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package queues

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/service"
)

// SysLog 系统日志
var SysLog = &qSysLog{topic: consts.QueueLogTopic}

type qSysLog struct {
	topic string
}

// handle 处理消息
func (q *qSysLog) handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	return service.SysLog().QueueJob(ctx, mqMsg)
}

// Listen 监听
func (q *qSysLog) Listen(ctx context.Context) {
	consumer, err := queue.InstanceConsumer()
	if err != nil {
		queue.FatalLog(ctx, "InstanceConsumer "+q.topic+"异常:", err)
		return
	}

	// 全局日志
	if listenErr := consumer.ListenReceiveMsgDo(q.topic, func(mqMsg queue.MqMsg) {
		err = q.handle(ctx, mqMsg)

		// 记录队列日志
		queue.ConsumerLog(ctx, q.topic, mqMsg, err)

	}); listenErr != nil {
		queue.FatalLog(ctx, "队列："+q.topic+" 监听失败", listenErr)
	}

}
