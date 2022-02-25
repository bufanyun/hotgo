//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package queue

import (
	"github.com/bufanyun/hotgo/app/utils"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"sync"
	"time"
)

//
//  MqProducer
//  @Description 
//
type MqProducer interface {
	SendMsg(topic string, body string) (mqMsg MqMsg, err error)
	SendByteMsg(topic string, body []byte) (mqMsg MqMsg, err error)
}

//
//  MqConsumer
//  @Description 
//
type MqConsumer interface {
	ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg)) (err error)
}

const (
	_ = iota
	SendMsg
	ReceiveMsg
)

type MqMsg struct {
	RunType   int       `json:"run_type"`
	Topic     string    `json:"topic"`
	MsgId     string    `json:"msg_id"`
	Offset    int64     `json:"offset"`
	Partition int32     `json:"partition"`
	Timestamp time.Time `json:"timestamp"`

	Body []byte `json:"body"`
}

var (
	ctx                   = gctx.New()
	mqProducerInstanceMap map[string]MqProducer
	mqConsumerInstanceMap map[string]MqConsumer
	mutex                 sync.Mutex
)

func init() {
	mqProducerInstanceMap = make(map[string]MqProducer)
	mqConsumerInstanceMap = make(map[string]MqConsumer)
}

//
//  @Title  实例化消费者
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Return  mqClient
//  @Return  err
//
func InstanceConsumer() (mqClient MqConsumer, err error) {
	groupName, _ := g.Cfg().Get(ctx, "queue.groupName", "hotgo")
	return NewConsumer(groupName.String())
}

//
//  @Title  实例化生产者
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Return  mqClient
//  @Return  err
//
func InstanceProducer() (mqClient MqProducer, err error) {
	groupName, _ := g.Cfg().Get(ctx, "queue.groupName", "hotgo")
	return NewProducer(groupName.String())
}

//
//  @Title  新建一个生产者实例
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   groupName
//  @Return  mqClient
//  @Return  err
//
func NewProducer(groupName string) (mqClient MqProducer, err error) {
	if item, ok := mqProducerInstanceMap[groupName]; ok {
		return item, nil
	}

	if groupName == "" {
		return mqClient, gerror.New("mq groupName is empty.")
	}

	// 驱动
	driver, _ := g.Cfg().Get(ctx, "queue.driver", "")

	// 重试次数
	retryCount, _ := g.Cfg().Get(ctx, "queue.retry", 2)
	retry := retryCount.Int()

	switch driver.String() {
	case "rocketmq":
		address, _ := g.Cfg().Get(ctx, "queue.rocketmq.address", nil)
		if len(address.Strings()) == 0 {
			panic("queue rocketmq address is not support")
		}
		mqClient = RegisterRocketProducerMust(address.Strings(), groupName, retry)
	case "kafka":
		address, _ := g.Cfg().Get(ctx, "queue.kafka.address", nil)
		if len(address.Strings()) == 0 {
			panic("queue kafka address is not support")
		}
		version, _ := g.Cfg().Get(ctx, "queue.kafka.version", "2.0.0")
		mqClient = RegisterKafkaProducerMust(KafkaConfig{
			Brokers: address.Strings(),
			GroupID: groupName,
			Version: version.String(),
		})
	case "redis":
		address, _ := g.Cfg().Get(ctx, "queue.redis.address", nil)
		if len(address.String()) == 0 {
			panic("queue redis address is not support")
		}
		db, _ := g.Cfg().Get(ctx, "queue.redis.db", 0)
		pass, _ := g.Cfg().Get(ctx, "queue.redis.pass", "")
		timeout, _ := g.Cfg().Get(ctx, "queue.redis.timeout", 0)

		mqClient = RegisterRedisMqProducerMust(RedisOption{
			Addr:    address.String(),
			Passwd:  pass.String(),
			DBnum:   db.Int(),
			Timeout: timeout.Int(),
		}, PoolOption{
			5, 50, 5,
		}, groupName, retry)

	default:
		panic("queue driver is not support")
	}

	mutex.Lock()
	defer mutex.Unlock()
	mqProducerInstanceMap[groupName] = mqClient

	return mqClient, nil
}

//
//  @Title  新建一个消费者实例
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   groupName
//  @Return  mqClient
//  @Return  err
//
func NewConsumer(groupName string) (mqClient MqConsumer, err error) {
	// 是否支持创建多个消费者
	multiComsumer, _ := g.Cfg().Get(ctx, "queue.multiComsumer", true)
	randTag := string(utils.Charset.RandomCreateBytes(6))
	if multiComsumer.Bool() == false {
		randTag = "001"
	}

	if item, ok := mqConsumerInstanceMap[groupName+"-"+randTag]; ok {
		return item, nil
	}

	driver, _ := g.Cfg().Get(ctx, "queue.driver", "")

	if groupName == "" {
		return mqClient, gerror.New("mq groupName is empty.")
	}

	switch driver.String() {
	case "rocketmq":
		address, _ := g.Cfg().Get(ctx, "queue.rocketmq.address", nil)
		if address == nil {
			return nil, gerror.New("queue.rocketmq.address is empty.")
		}

		mqClient = RegisterRocketConsumerMust(address.Strings(), groupName)
	case "kafka":
		address, _ := g.Cfg().Get(ctx, "queue.kafka.address", nil)
		if len(address.Strings()) == 0 {
			panic("queue kafka address is not support")
		}
		version, _ := g.Cfg().Get(ctx, "queue.kafka.version", "2.0.0")

		clientId := "HOTGO-Consumer-" + groupName
		randClient, _ := g.Cfg().Get(ctx, "queue.kafka.randClient", true)
		if randClient.Bool() {
			clientId += "-" + randTag
		}

		mqClient = RegisterKafkaMqConsumerMust(KafkaConfig{
			Brokers:  address.Strings(),
			GroupID:  groupName,
			Version:  version.String(),
			ClientId: clientId,
		})
	case "redis":
		address, _ := g.Cfg().Get(ctx, "queue.redis.address", nil)
		if len(address.String()) == 0 {
			panic("queue redis address is not support")
		}
		db, _ := g.Cfg().Get(ctx, "queue.redis.db", 0)
		pass, _ := g.Cfg().Get(ctx, "queue.redis.pass", "")
		timeout, _ := g.Cfg().Get(ctx, "queue.redis.pass", 0)

		mqClient = RegisterRedisMqConsumerMust(RedisOption{
			Addr:    address.String(),
			Passwd:  pass.String(),
			DBnum:   db.Int(),
			Timeout: timeout.Int(),
		}, PoolOption{
			5, 50, 5,
		}, groupName)
	default:
		panic("queue driver is not support")
	}

	mutex.Lock()
	defer mutex.Unlock()
	mqConsumerInstanceMap[groupName] = mqClient

	return mqClient, nil
}

//
//  @Title  返回消息体
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Return  string
//
func (m *MqMsg) BodyString() string {
	return string(m.Body)
}
