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
