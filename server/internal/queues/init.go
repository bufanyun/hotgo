// Package queues
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package queues

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/library/queue"
)

type jobStrategy interface {
	getTopic() string
	handle(ctx context.Context, mqMsg queue.MqMsg) (err error)
}

var jobList []jobStrategy

func Run(ctx context.Context) {
	for _, job := range uniqueJob(jobList) {
		go func(job jobStrategy) {
			listen(ctx, job)
		}(job)
	}
}

func listen(ctx context.Context, job jobStrategy) {
	var (
		topic         = job.getTopic()
		consumer, err = queue.InstanceConsumer()
	)

	if err != nil {
		g.Log().Fatalf(ctx, "InstanceConsumer %s err:%+v", topic, err)
		return
	}

	// 访问日志
	if listenErr := consumer.ListenReceiveMsgDo(topic, func(mqMsg queue.MqMsg) {
		err = job.handle(ctx, mqMsg)

		if err != nil {
			// 遇到错误，重新加入到队列
			//queue.Push(topic, mqMsg.Body)
		}

		// 记录队列日志
		queue.ConsumerLog(ctx, topic, mqMsg, err)

	}); listenErr != nil {
		g.Log().Fatalf(ctx, "队列：%s 监听失败, err:%+v", topic, listenErr)
	}

}

// uniqueJob 去重
func uniqueJob(languages []jobStrategy) []jobStrategy {
	result := make([]jobStrategy, 0, len(languages))
	temp := map[jobStrategy]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
