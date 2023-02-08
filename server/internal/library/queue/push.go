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
	ProducerLog(ctx, topic, mqMsg.MsgId, err)
	return err
}
