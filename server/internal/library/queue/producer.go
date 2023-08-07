// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queue

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// Push 推送队列
func Push(topic string, data interface{}) (err error) {
	q, err := InstanceProducer()
	if err != nil {
		return
	}
	mqMsg, err := q.SendMsg(topic, gconv.String(data))
	ProducerLog(ctx, topic, mqMsg, err)
	return
}

// DelayPush 推送延迟队列
func DelayPush(topic string, data interface{}, second int64) (err error) {
	q, err := InstanceProducer()
	if err != nil {
		return
	}
	mqMsg, err := q.SendDelayMsg(topic, gconv.String(data), second)
	ProducerLog(ctx, topic, mqMsg, err)
	return
}
