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
	if err := g.Cfg().MustGet(ctx, "queue").Scan(&config); err != nil {
		g.Log().Infof(ctx, "queue init err:%+v", err)
	}
}

// InstanceConsumer 实例化消费者
func InstanceConsumer() (mqClient MqConsumer, err error) {
	return NewConsumer(config.GroupName)
}

// InstanceProducer 实例化生产者
func InstanceProducer() (mqClient MqProducer, err error) {
	return NewProducer(config.GroupName)
}

// NewProducer 初始化生产者实例
func NewProducer(groupName string) (mqClient MqProducer, err error) {
	if item, ok := mqProducerInstanceMap[groupName]; ok {
		return item, nil
	}

	if groupName == "" {
		err = gerror.New("mq groupName is empty.")
		return
	}

	switch config.Driver {
	case "rocketmq":
		if len(config.Rocketmq.Address) == 0 {
			err = gerror.New("queue rocketmq address is not support")
			return
		}
		mqClient, err = RegisterRocketProducer(config.Rocketmq.Address, groupName, config.Retry)
	case "kafka":
		if len(config.Kafka.Address) == 0 {
			err = gerror.New("queue kafka address is not support")
			return
		}
		mqClient, err = RegisterKafkaProducer(KafkaConfig{
			Brokers: config.Kafka.Address,
			GroupID: groupName,
			Version: config.Kafka.Version,
		})
	case "redis":
		address := g.Cfg().MustGet(ctx, "queue.redis.address", nil).String()
		if len(address) == 0 {
			err = gerror.New("queue redis address is not support")
			return
		}
		mqClient, err = RegisterRedisMqProducer(RedisOption{
			Addr:    config.Redis.Address,
			Passwd:  config.Redis.Pass,
			DBnum:   config.Redis.Db,
			Timeout: config.Redis.Timeout,
		}, PoolOption{
			5, 50, 5,
		}, groupName, config.Retry)

	default:
		err = gerror.New("queue driver is not support")
	}

	if err != nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	mqProducerInstanceMap[groupName] = mqClient

	return
}

// NewConsumer 初始化消费者实例
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
		err = gerror.New("mq groupName is empty.")
		return
	}

	switch config.Driver {
	case "rocketmq":
		if len(config.Rocketmq.Address) == 0 {
			err = gerror.New("queue.rocketmq.address is empty.")
			return
		}
		mqClient, err = RegisterRocketConsumer(config.Rocketmq.Address, groupName)
	case "kafka":
		if len(config.Kafka.Address) == 0 {
			err = gerror.New("queue kafka address is not support")
			return
		}

		clientId := "HOTGO-Consumer-" + groupName
		if config.Kafka.RandClient {
			clientId += "-" + randTag
		}

		mqClient, err = RegisterKafkaMqConsumer(KafkaConfig{
			Brokers:  config.Kafka.Address,
			GroupID:  groupName,
			Version:  config.Kafka.Version,
			ClientId: clientId,
		})
	case "redis":
		if len(config.Redis.Address) == 0 {
			err = gerror.New("queue redis address is not support")
			return
		}

		mqClient, err = RegisterRedisMqConsumer(RedisOption{
			Addr:    config.Redis.Address,
			Passwd:  config.Redis.Pass,
			DBnum:   config.Redis.Db,
			Timeout: config.Redis.Timeout,
		}, PoolOption{
			5, 50, 5,
		}, groupName)
	default:
		err = gerror.New("queue driver is not support")
	}

	if err != nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	mqConsumerInstanceMap[groupName] = mqClient

	return
}

// BodyString 返回消息体
func (m *MqMsg) BodyString() string {
	return string(m.Body)
}
