// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queue

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
	"sync"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
)

type RocketMq struct {
	producerIns rocketmq.Producer
	consumerIns rocketmq.PushConsumer
}

type RocketManager struct {
	Producer *RocketMq
	Consumer *RocketMq
	pMutex   sync.Mutex
	cMutex   sync.Mutex
	goPool   *grpool.Pool
}

var rocketManager = &RocketManager{}

func init() {
	setRocketCloseEvent()
}

func setRocketCloseEvent() {
	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		if rocketManager == nil {
			return
		}

		if rocketManager.Producer != nil {
			err := rocketManager.Producer.producerIns.Shutdown()
			if err != nil {
				Logger().Warningf(ctx, "rocketmq producer close err:%v", err)
				return
			}
			Logger().Debug(ctx, "rocketmq producer close...")
		}

		if rocketManager.Consumer != nil {
			err := rocketManager.Consumer.consumerIns.Shutdown()
			if err != nil {
				Logger().Warningf(ctx, "rocketmq consumer close err:%v", err)
				return
			}
			Logger().Debug(ctx, "rocketmq consumer close...")
		}

		for rocketManager.goPool != nil && rocketManager.goPool.Size() != 0 {
			g.Log().Debugf(ctx, "waiting for eocketmq consumer to complete execution[%v][%v]...", rocketManager.goPool.Size(), rocketManager.goPool.Jobs())
			time.Sleep(time.Second)
		}
	})
}

func GetRocketManager() *RocketManager {
	return rocketManager
}

// RegisterRocketProducer 注册并启动生产者接口实现
func RegisterRocketProducer() (client MqProducer, err error) {
	return RegisterRocketMqProducer()
}

// RegisterRocketConsumer 注册消费者
func RegisterRocketConsumer() (client MqConsumer, err error) {
	return RegisterRocketMqConsumer()
}

// createTopicIfNotExists 主题不存在就自动创建
func (r *RocketMq) createTopicIfNotExists(topic string) (err error) {
	if len(config.Rocketmq.BrokerAddr) == 0 {
		return
	}

	client, err := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(config.Rocketmq.NameSrvAdders)),
		admin.WithCredentials(primitive.Credentials{
			AccessKey: config.Rocketmq.AccessKey,
			SecretKey: config.Rocketmq.SecretKey,
		}),
	)
	if err != nil {
		return err
	}
	defer client.Close()

	result, err := client.FetchAllTopicList(ctx)
	if err != nil {
		return err
	}

	if validate.InSlice(result.TopicList, topic) {
		return
	}

	Logger().Debugf(ctx, "create topic:%v", topic)
	err = client.CreateTopic(ctx, admin.WithTopicCreate(topic), admin.WithBrokerAddrCreate(config.Rocketmq.BrokerAddr))
	return
}

// SendMsg 按字符串类型生产数据
func (r *RocketMq) SendMsg(topic string, body string) (mqMsg MqMsg, err error) {
	return r.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (r *RocketMq) SendByteMsg(topic string, body []byte) (mqMsg MqMsg, err error) {
	if r.producerIns == nil {
		return mqMsg, gerror.New("rocketMq producer not register")
	}

	result, err := r.producerIns.SendSync(ctx, &primitive.Message{
		Topic: topic,
		Body:  body,
	})

	if err != nil {
		return
	}
	if result.Status != primitive.SendOK {
		return mqMsg, gerror.Newf("rocketMq producer send msg error status:%v", result.Status)
	}

	mqMsg = MqMsg{
		RunType: SendMsg,
		Topic:   topic,
		MsgId:   result.MsgID,
		Body:    body,
	}
	return mqMsg, nil
}

func (r *RocketMq) SendDelayMsg(topic string, body string, delayTimeLevel int64) (mqMsg MqMsg, err error) {
	if r.producerIns == nil {
		return mqMsg, gerror.New("rocketMq producer not register")
	}

	msg := primitive.NewMessage(topic, []byte(body))
	msg.WithDelayTimeLevel(int(delayTimeLevel))

	result, err := r.producerIns.SendSync(ctx, msg)
	if err != nil {
		return
	}
	if result.Status != primitive.SendOK {
		return mqMsg, gerror.Newf("rocketMq producer send msg error status:%v", result.Status)
	}

	mqMsg = MqMsg{
		RunType: SendMsg,
		Topic:   topic,
		MsgId:   result.MsgID,
		Body:    []byte(body),
	}
	return mqMsg, nil
}

// ListenReceiveMsgDo 消费数据
func (r *RocketMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg)) (err error) {
	if r.consumerIns == nil {
		return gerror.New("rocketMq consumer not register")
	}

	rocketManager.cMutex.Lock()
	defer rocketManager.cMutex.Unlock()

	if err = r.createTopicIfNotExists(topic); err != nil {
		return err
	}

	err = r.consumerIns.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, item := range msgs {
			_ = rocketManager.goPool.Add(ctx, func(ctx context.Context) {
				receiveDo(MqMsg{
					RunType: ReceiveMsg,
					Topic:   item.Topic,
					MsgId:   item.MsgId,
					Body:    item.Body,
				})
			})
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		return
	}

	if err = r.consumerIns.Start(); err != nil {
		_ = r.consumerIns.Unsubscribe(topic)
		return
	}
	return
}

// RegisterRocketMqProducer 注册rocketmq生产者
func RegisterRocketMqProducer() (mqIns *RocketMq, err error) {
	if rocketManager.Producer != nil {
		return rocketManager.Producer, nil
	}

	rocketManager.pMutex.Lock()
	defer rocketManager.pMutex.Unlock()

	if rocketManager.Producer != nil {
		return rocketManager.Producer, nil
	}

	mqIns = new(RocketMq)
	retry := config.Rocketmq.Retry
	if retry <= 0 {
		retry = 0
	}

	mqIns.producerIns, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(config.Rocketmq.NameSrvAdders)),
		producer.WithRetry(retry),
		producer.WithGroupName(config.GroupName),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: config.Rocketmq.AccessKey,
			SecretKey: config.Rocketmq.SecretKey,
		}),
	)

	if err != nil {
		return nil, err
	}

	if err = mqIns.producerIns.Start(); err != nil {
		return nil, err
	}

	_, err = mqIns.producerIns.SendSync(ctx, primitive.NewMessage("hotgo-ping", []byte("1")))
	if err != nil {
		err = gerror.Newf("连通性测试不通过，请检查`queue.rocketmq.nameSrvAdders`或权限配置是否有误。err:%+v", err.Error())
		return nil, err
	}

	SetRLogLevel()
	rocketManager.Producer = mqIns
	return rocketManager.Producer, nil
}

// RegisterRocketMqConsumer 注册rocketmq消费者
func RegisterRocketMqConsumer() (mqIns *RocketMq, err error) {
	if rocketManager.Consumer != nil {
		return rocketManager.Consumer, nil
	}

	rocketManager.cMutex.Lock()
	defer rocketManager.cMutex.Unlock()

	if rocketManager.Consumer != nil {
		return rocketManager.Consumer, nil
	}

	// 利用生产者检查一下连通性
	if _, err = RegisterRocketMqProducer(); err != nil {
		return nil, err
	}

	mqIns = new(RocketMq)
	mqIns.consumerIns, err = rocketmq.NewPushConsumer(
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(config.Rocketmq.NameSrvAdders)),
		consumer.WithGroupName(config.GroupName),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: config.Rocketmq.AccessKey,
			SecretKey: config.Rocketmq.SecretKey,
		}),
	)

	if err != nil {
		return nil, err
	}

	// 开多携程处理消费任务，可以根据业务实际情况调整该配置
	rocketManager.goPool = grpool.New(5)

	SetRLogLevel()
	rocketManager.Consumer = mqIns
	return rocketManager.Consumer, nil
}

// SetRLogLevel 设置rocketmq日志输出等级
func SetRLogLevel() {
	level := g.Cfg().MustGet(ctx, "queue.rocketmq.logLevel", "all").String()
	rlog.SetLogLevel(level)
}
