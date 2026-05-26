// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queue

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/utility/simple"
	"time"
)

type KafkaMq struct {
	Partitions  int32
	producerIns sarama.AsyncProducer
	consumerIns sarama.ConsumerGroup
}

type KafkaConfig struct {
	ClientId    string
	Brokers     []string
	GroupID     string
	Partitions  int32
	Replication int16
	Version     string
	UserName    string
	Password    string
}

// SendMsg 按字符串类型生产数据
func (r *KafkaMq) SendMsg(topic string, body string) (mqMsg MqMsg, err error) {
	return r.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (r *KafkaMq) SendByteMsg(topic string, body []byte) (mqMsg MqMsg, err error) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(body),
		Timestamp: time.Now(),
	}

	if r.producerIns == nil {
		err = gerror.New("queue kafka producerIns is nil")
		return
	}

	r.producerIns.Input() <- msg
	sendCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case info := <-r.producerIns.Successes():
		return MqMsg{
			RunType:   SendMsg,
			Topic:     info.Topic,
			Offset:    info.Offset,
			Partition: info.Partition,
			Timestamp: info.Timestamp,
		}, nil
	case fail := <-r.producerIns.Errors():
		if nil != fail {
			return mqMsg, fail.Err
		}
	case <-sendCtx.Done():
		return mqMsg, gerror.New("send mqMst timeout")
	}
	return mqMsg, nil
}

func (r *KafkaMq) SendDelayMsg(topic string, body string, delaySecond int64) (mqMsg MqMsg, err error) {
	err = gerror.New("implement me")
	return
}

// ListenReceiveMsgDo 消费数据
func (r *KafkaMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg)) (err error) {
	if r.consumerIns == nil {
		return gerror.New("queue kafka consumer not register")
	}

	consumer := KaConsumer{
		ready:        make(chan bool),
		receiveDoFun: receiveDo,
	}

	consumerCtx, cancel := context.WithCancel(context.Background())
	go func(consumerCtx context.Context) {
		for {
			if err = r.consumerIns.Consume(consumerCtx, []string{topic}, &consumer); err != nil {
				Logger().Fatalf(ctx, "kafka Error from consumer, err%+v", err)
			}

			if consumerCtx.Err() != nil {
				Logger().Debugf(ctx, fmt.Sprintf("kafka consoumer stop : %v", consumerCtx.Err()))
				return
			}
			consumer.ready = make(chan bool)
		}
	}(consumerCtx)

	// await till the consumer has been set up
	<-consumer.ready
	Logger().Debug(ctx, "kafka consumer up and running!...")

	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		Logger().Debug(ctx, "kafka consumer close...")
		cancel()
		if err = r.consumerIns.Close(); err != nil {
			Logger().Fatalf(ctx, "kafka Error closing client, err:%+v", err)
		}
	})
	return
}

// RegisterKafkaMqConsumer 注册消费者
func RegisterKafkaMqConsumer(connOpt KafkaConfig) (client MqConsumer, err error) {
	mqIns := &KafkaMq{}
	kfkVersion, err := sarama.ParseKafkaVersion(connOpt.Version)
	if err != nil {
		return
	}
	if !validateVersion(kfkVersion) {
		kfkVersion = sarama.V2_4_0_0
	}

	brokers := connOpt.Brokers
	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true
	conf.Version = kfkVersion
	if connOpt.UserName != "" {
		conf.Net.SASL.Enable = true
		conf.Net.SASL.User = connOpt.UserName
		conf.Net.SASL.Password = connOpt.Password
	}

	// 默认按随机方式消费
	conf.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	conf.Consumer.Offsets.AutoCommit.Interval = 10 * time.Millisecond
	conf.ClientID = connOpt.ClientId

	consumerClient, err := sarama.NewConsumerGroup(brokers, connOpt.GroupID, conf)
	if err != nil {
		return
	}
	mqIns.consumerIns = consumerClient
	return mqIns, err
}

// RegisterKafkaProducer 注册并启动生产者接口实现
func RegisterKafkaProducer(connOpt KafkaConfig) (client MqProducer, err error) {
	mqIns := &KafkaMq{}
	connOpt.ClientId = "HOTGO-Producer"

	// 这里如果使用go程需要处理chan同步问题
	if err = doRegisterKafkaProducer(connOpt, mqIns); err != nil {
		return nil, err
	}

	return mqIns, nil
}

// doRegisterKafkaProducer 注册同步类型实例
func doRegisterKafkaProducer(connOpt KafkaConfig, mqIns *KafkaMq) (err error) {
	kfkVersion, err := sarama.ParseKafkaVersion(connOpt.Version)
	if err != nil {
		return
	}
	if !validateVersion(kfkVersion) {
		kfkVersion = sarama.V2_4_0_0
	}

	brokers := connOpt.Brokers
	conf := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	conf.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向partition发送消息
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	conf.Producer.Return.Successes = true

	conf.Producer.Return.Errors = true
	conf.Producer.Compression = sarama.CompressionNone
	conf.ClientID = connOpt.ClientId

	conf.Version = kfkVersion
	if connOpt.UserName != "" {
		conf.Net.SASL.Enable = true
		conf.Net.SASL.User = connOpt.UserName
		conf.Net.SASL.Password = connOpt.Password
	}

	mqIns.producerIns, err = sarama.NewAsyncProducer(brokers, conf)
	if err != nil {
		return
	}

	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		g.Log().Debug(ctx, "kafka producer AsyncClose...")
		mqIns.producerIns.AsyncClose()
	})
	return
}

// validateVersion 验证版本是否有效
func validateVersion(version sarama.KafkaVersion) bool {
	for _, item := range sarama.SupportedVersions {
		if version.String() == item.String() {
			return true
		}
	}
	return false
}

type KaConsumer struct {
	ready        chan bool
	receiveDoFun func(mqMsg MqMsg)
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *KaConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *KaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *KaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	// `ConsumeClaim` 方法已经是 goroutine 调用 不要在该方法内进行 goroutine
	for message := range claim.Messages() {
		consumer.receiveDoFun(MqMsg{
			RunType:   ReceiveMsg,
			Topic:     message.Topic,
			Body:      message.Value,
			Offset:    message.Offset,
			Timestamp: message.Timestamp,
			Partition: message.Partition,
		})
		session.MarkMessage(message, "")
	}
	return nil
}
