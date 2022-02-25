package queue

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bufanyun/hotgo/app/utils"
	"github.com/gogf/gf/v2/errors/gerror"
	"sync"
	"time"
)

type KafkaMq struct {
	endPoints   []string
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

var wg sync.WaitGroup

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
		return mqMsg, gerror.New("queue kafka producerIns is nil")
	}

	r.producerIns.Input() <- msg
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

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
	case <-ctx.Done():
		return mqMsg, gerror.New("send mqMst timeout")
	}

	return mqMsg, nil
}

// ListenReceiveMsgDo 消费数据
func (r *KafkaMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg)) (err error) {
	if r.consumerIns == nil {
		return gerror.New("queue kafka consumer not register")
	}

	consumer := Consumer{
		ready:        make(chan bool),
		receiveDoFun: receiveDo,
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {

		for {
			if err := r.consumerIns.Consume(ctx, []string{topic}, &consumer); err != nil {
				FatalLog(ctx, "kafka Error from consumer", err)
			}

			if ctx.Err() != nil {
				Log(ctx, fmt.Sprint("kafka consoumer stop : %v", ctx.Err()))
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	Log(ctx, "kafka consumer up and running!...")

	utils.Signal.AppDefer(func() {
		Log(ctx, "kafka consumer close...")
		cancel()
		if err = r.consumerIns.Close(); err != nil {
			FatalLog(ctx, "kafka Error closing client", err)
		}
	})

	return
}

// RegisterRedisMqConsumerMust 注册消费者
func RegisterKafkaMqConsumerMust(connOpt KafkaConfig) (client MqConsumer) {
	mqIns := &KafkaMq{}

	kfkVersion, _ := sarama.ParseKafkaVersion(connOpt.Version)
	if validateVersion(kfkVersion) == false {
		kfkVersion = sarama.V2_4_0_0
	}

	brokers := connOpt.Brokers
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = kfkVersion
	if connOpt.UserName != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = connOpt.UserName
		config.Net.SASL.Password = connOpt.Password
	}

	// 默认按随机方式消费
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.AutoCommit.Interval = 10 * time.Millisecond
	config.ClientID = connOpt.ClientId

	consumerClient, err := sarama.NewConsumerGroup(brokers, connOpt.GroupID, config)
	if err != nil {
		panic(err)
	}
	mqIns.consumerIns = consumerClient
	return mqIns
}

// RegisterKafkaProducerMust 注册并启动生产者接口实现
func RegisterKafkaProducerMust(connOpt KafkaConfig) (client MqProducer) {
	mqIns := &KafkaMq{}

	connOpt.ClientId = "HOTGO-Producer"
	RegisterKafkaProducer(connOpt, mqIns) //这里如果使用go程需要处理chan同步问题

	return mqIns
}

// RegisterKafkaProducerAsync 注册同步类型实例
func RegisterKafkaProducer(connOpt KafkaConfig, mqIns *KafkaMq) {
	kfkVersion, _ := sarama.ParseKafkaVersion(connOpt.Version)
	if validateVersion(kfkVersion) == false {
		kfkVersion = sarama.V2_4_0_0
	}

	brokers := connOpt.Brokers
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true

	config.Producer.Return.Errors = true
	config.Producer.Compression = sarama.CompressionNone
	config.ClientID = connOpt.ClientId

	config.Version = kfkVersion
	if connOpt.UserName != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = connOpt.UserName
		config.Net.SASL.Password = connOpt.Password
	}

	var err error
	mqIns.producerIns, err = sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	utils.Signal.AppDefer(func() {
		Log(ctx, "kafka producer AsyncClose...")
		mqIns.producerIns.AsyncClose()
	})
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

type Consumer struct {
	ready        chan bool
	receiveDoFun func(mqMsg MqMsg)
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

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
