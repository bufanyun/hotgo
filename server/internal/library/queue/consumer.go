package queue

import (
	"context"
	"sync"
)

// Consumer 消费者接口，实现该接口即可加入到消费队列中
type Consumer interface {
	GetTopic() string                                    // 获取消费主题
	Handle(ctx context.Context, mqMsg MqMsg) (err error) // 处理消息的方法
}

// consumerManager 消费者管理
type consumerManager struct {
	sync.Mutex
	list map[string]Consumer // 维护的消费者列表
}

var consumers = &consumerManager{
	list: make(map[string]Consumer),
}

// RegisterConsumer 注册任务到消费者队列
func RegisterConsumer(cs Consumer) {
	consumers.Lock()
	defer consumers.Unlock()
	topic := cs.GetTopic()
	if _, ok := consumers.list[topic]; ok {
		Logger().Debugf(ctx, "queue.RegisterConsumer topic:%v duplicate registration.", topic)
		return
	}
	consumers.list[topic] = cs
}

// StartConsumersListener 启动所有已注册的消费者监听
func StartConsumersListener(ctx context.Context) {
	for _, c := range consumers.list {
		go func(c Consumer) {
			consumerListen(ctx, c)
		}(c)
	}
}

// consumerListen 消费者监听
func consumerListen(ctx context.Context, job Consumer) {
	var (
		topic  = job.GetTopic()
		c, err = InstanceConsumer()
	)

	if err != nil {
		Logger().Fatalf(ctx, "InstanceConsumer %s err:%+v", topic, err)
		return
	}

	if listenErr := c.ListenReceiveMsgDo(topic, func(mqMsg MqMsg) {
		err = job.Handle(ctx, mqMsg)

		// if err != nil {
		//	// 遇到错误，重新加入到队列
		//	//queue.Push(topic, mqMsg.Body)
		// }

		// 记录消费队列日志
		ConsumerLog(ctx, topic, mqMsg, err)
	}); listenErr != nil {
		Logger().Fatalf(ctx, "消费队列：%s 监听失败, err:%+v", topic, listenErr)
	}
}
