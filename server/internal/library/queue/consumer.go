package queue

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

// consumerStrategy 消费者策略，实现该接口即可加入到消费队列中
type consumerStrategy interface {
	GetTopic() string                                    // 获取消费主题
	Handle(ctx context.Context, mqMsg MqMsg) (err error) // 处理消息
}

// consumerManager 消费者管理
type consumerManager struct {
	sync.Mutex
	list map[string]consumerStrategy // 维护的消费者列表
}

var consumers = &consumerManager{
	list: make(map[string]consumerStrategy),
}

// RegisterConsumer 注册任务到消费者队列
func RegisterConsumer(cs consumerStrategy) {
	consumers.Lock()
	defer consumers.Unlock()
	topic := cs.GetTopic()
	if _, ok := consumers.list[topic]; ok {
		g.Log().Debugf(ctx, "queue.RegisterConsumer topic:%v duplicate registration.", topic)
		return
	}
	consumers.list[topic] = cs
}

// StartConsumersListener 启动所有已注册的消费者监听
func StartConsumersListener(ctx context.Context) {
	for _, consumer := range consumers.list {
		go func(consumer consumerStrategy) {
			consumerListen(ctx, consumer)
		}(consumer)
	}
}

// consumerListen 消费者监听
func consumerListen(ctx context.Context, job consumerStrategy) {
	var (
		topic         = job.GetTopic()
		consumer, err = InstanceConsumer()
	)

	if err != nil {
		g.Log().Fatalf(ctx, "InstanceConsumer %s err:%+v", topic, err)
		return
	}

	if listenErr := consumer.ListenReceiveMsgDo(topic, func(mqMsg MqMsg) {
		err = job.Handle(ctx, mqMsg)

		if err != nil {
			// 遇到错误，重新加入到队列
			//queue.Push(topic, mqMsg.Body)
		}

		// 记录消费队列日志
		ConsumerLog(ctx, topic, mqMsg, err)

	}); listenErr != nil {
		g.Log().Fatalf(ctx, "消费队列：%s 监听失败, err:%+v", topic, listenErr)
	}

}
