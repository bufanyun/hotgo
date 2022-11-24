// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package queue

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/utility/charset"
	"sync"
	"time"
)

type MqProducer interface {
	SendMsg(topic string, body string) (mqMsg MqMsg, err error)
	SendByteMsg(topic string, body []byte) (mqMsg MqMsg, err error)
}

type MqConsumer interface {
	ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg)) (err error)
}

const (
	_ = iota
	SendMsg
	ReceiveMsg
)

type Config struct {
	Switch        bool   `json:"switch"`
	Driver        string `json:"driver"`
	Retry         int    `json:"retry"`
	MultiComsumer bool   `json:"multiComsumer"`
	GroupName     string `json:"groupName"`
	Redis         RedisConf
	Rocketmq      RocketmqConf
	Kafka         KafkaConf
}

type RedisConf struct {
	Address string `json:"address"`
	Db      int    `json:"db"`
	Pass    string `json:"pass"`
	Timeout int    `json:"timeout"`
}
type RocketmqConf struct {
	Address  []string `json:"address"`
	LogLevel string   `json:"logLevel"`
}

type KafkaConf struct {
	Address    []string `json:"address"`
	Version    string   `json:"version"`
	RandClient bool     `json:"randClient"`
}

type MqMsg struct {
	RunType   int       `json:"run_type"`
	Topic     string    `json:"topic"`
	MsgId     string    `json:"msg_id"`
	Offset    int64     `json:"offset"`
	Partition int32     `json:"partition"`
	Timestamp time.Time `json:"timestamp"`
	Body      []byte    `json:"body"`
}

var (
	ctx                   = gctx.New()
	mqProducerInstanceMap map[string]MqProducer
	mqConsumerInstanceMap map[string]MqConsumer
	mutex                 sync.Mutex
	config                Config
)

func init() {
	mqProducerInstanceMap = make(map[string]MqProducer)
	mqConsumerInstanceMap = make(map[string]MqConsumer)
	get, err := g.Cfg().Get(ctx, "queue")
	if err != nil {
		g.Log().Fatalf(ctx, "queue config load fail, err .%v", err)
		return
	}
	get.Scan(&config)
}

// InstanceConsumer 实例化消费者
func InstanceConsumer() (mqClient MqConsumer, err error) {
	return NewConsumer(config.GroupName)
}

// InstanceProducer 实例化生产者
func InstanceProducer() (mqClient MqProducer, err error) {
	return NewProducer(config.GroupName)
}

// NewProducer 新建一个生产者实例
func NewProducer(groupName string) (mqClient MqProducer, err error) {
	if item, ok := mqProducerInstanceMap[groupName]; ok {
		return item, nil
	}

	if groupName == "" {
		return mqClient, gerror.New("mq groupName is empty.")
	}

	switch config.Driver {
	case "rocketmq":
		if len(config.Rocketmq.Address) == 0 {
			g.Log().Fatal(ctx, "queue rocketmq address is not support")
		}
		mqClient = RegisterRocketProducerMust(config.Rocketmq.Address, groupName, config.Retry)
	case "kafka":
		if len(config.Kafka.Address) == 0 {
			g.Log().Fatal(ctx, "queue kafka address is not support")
		}
		mqClient = RegisterKafkaProducerMust(KafkaConfig{
			Brokers: config.Kafka.Address,
			GroupID: groupName,
			Version: config.Kafka.Version,
		})
	case "redis":
		address, _ := g.Cfg().Get(ctx, "queue.redis.address", nil)
		if len(address.String()) == 0 {
			g.Log().Fatal(ctx, "queue redis address is not support")
		}
		mqClient = RegisterRedisMqProducerMust(RedisOption{
			Addr:    config.Redis.Address,
			Passwd:  config.Redis.Pass,
			DBnum:   config.Redis.Db,
			Timeout: config.Redis.Timeout,
		}, PoolOption{
			5, 50, 5,
		}, groupName, config.Retry)

	default:
		g.Log().Fatal(ctx, "queue driver is not support")
	}

	mutex.Lock()
	defer mutex.Unlock()
	mqProducerInstanceMap[groupName] = mqClient

	return mqClient, nil
}

// NewConsumer 新建一个消费者实例
func NewConsumer(groupName string) (mqClient MqConsumer, err error) {
	randTag := string(charset.RandomCreateBytes(6))

	// 是否支持创建多个消费者
	if config.MultiComsumer == false {
		randTag = "001"
	}

	if item, ok := mqConsumerInstanceMap[groupName+"-"+randTag]; ok {
		return item, nil
	}

	if groupName == "" {
		return mqClient, gerror.New("mq groupName is empty.")
	}

	switch config.Driver {
	case "rocketmq":
		if len(config.Rocketmq.Address) == 0 {
			return nil, gerror.New("queue.rocketmq.address is empty.")
		}
		mqClient = RegisterRocketConsumerMust(config.Rocketmq.Address, groupName)
	case "kafka":
		if len(config.Kafka.Address) == 0 {
			g.Log().Fatal(ctx, "queue kafka address is not support")
		}

		clientId := "HOTGO-Consumer-" + groupName
		if config.Kafka.RandClient {
			clientId += "-" + randTag
		}

		mqClient = RegisterKafkaMqConsumerMust(KafkaConfig{
			Brokers:  config.Kafka.Address,
			GroupID:  groupName,
			Version:  config.Kafka.Version,
			ClientId: clientId,
		})
	case "redis":
		if len(config.Redis.Address) == 0 {
			g.Log().Fatal(ctx, "queue redis address is not support")
		}

		mqClient = RegisterRedisMqConsumerMust(RedisOption{
			Addr:    config.Redis.Address,
			Passwd:  config.Redis.Pass,
			DBnum:   config.Redis.Db,
			Timeout: config.Redis.Timeout,
		}, PoolOption{
			5, 50, 5,
		}, groupName)
	default:
		g.Log().Fatal(ctx, "queue driver is not support")
	}

	mutex.Lock()
	defer mutex.Unlock()
	mqConsumerInstanceMap[groupName] = mqClient

	return mqClient, nil
}

// BodyString 返回消息体
func (m *MqMsg) BodyString() string {
	return string(m.Body)
}
