package queue

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"hotgo/internal/library/queue/disk"
	"sync"
	"time"
)

type DiskProducerMq struct {
	config    *disk.Config
	producers map[string]*disk.Queue
	sync.Mutex
}

type DiskConsumerMq struct {
	config *disk.Config
}

func RegisterDiskMqConsumer(config *disk.Config) (client MqConsumer, err error) {
	return &DiskConsumerMq{
		config: config,
	}, nil
}

// ListenReceiveMsgDo 消费数据
func (q *DiskConsumerMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg)) (err error) {
	if topic == "" {
		return gerror.New("disk.ListenReceiveMsgDo topic is empty")
	}

	var (
		queue = NewDiskQueue(topic, q.config)
		sleep = time.Second
	)

	go func() {
		for {
			if index, offset, data, err := queue.Read(); err == nil {
				var mqMsg MqMsg
				if err = json.Unmarshal(data, &mqMsg); err != nil {
					g.Log().Warningf(ctx, "disk.ListenReceiveMsgDo Unmarshal err:%+v, topic：%v, data:%+v .", err, topic, string(data))
					continue
				}
				if mqMsg.MsgId != "" {
					receiveDo(mqMsg)
					queue.Commit(index, offset)
					sleep = time.Millisecond * 1
				}
			} else {
				sleep = time.Second
			}

			time.Sleep(sleep)
		}
	}()

	select {}
}

func RegisterDiskMqProducer(config *disk.Config) (client MqProducer, err error) {
	return &DiskProducerMq{
		config:    config,
		producers: make(map[string]*disk.Queue),
	}, nil
}

// SendMsg 按字符串类型生产数据
func (d *DiskProducerMq) SendMsg(topic string, body string) (mqMsg MqMsg, err error) {
	return d.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (d *DiskProducerMq) SendByteMsg(topic string, body []byte) (mqMsg MqMsg, err error) {
	if topic == "" {
		return mqMsg, gerror.New("DiskMq topic is empty")
	}

	mqMsg = MqMsg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      body,
		Timestamp: time.Now(),
	}

	mqMsgJson, err := json.Marshal(mqMsg)
	if err != nil {
		return mqMsg, gerror.New(fmt.Sprint("queue redis 生产者解析json消息失败:", err))
	}

	queue := d.getProducer(topic)
	if err = queue.Write(mqMsgJson); err != nil {
		return mqMsg, gerror.New(fmt.Sprint("queue disk 生产者添加消息失败:", err))
	}
	return
}

func (d *DiskProducerMq) getProducer(topic string) *disk.Queue {
	queue, ok := d.producers[topic]
	if ok {
		return queue
	}
	queue = NewDiskQueue(topic, d.config)
	d.Lock()
	defer d.Unlock()
	d.producers[topic] = queue
	return queue
}

func NewDiskQueue(topic string, config *disk.Config) *disk.Queue {
	conf := &disk.Config{
		Path:         fmt.Sprintf(config.Path + "/" + config.GroupName + "/" + topic),
		BatchSize:    config.BatchSize,
		BatchTime:    config.BatchTime * time.Second,
		SegmentSize:  config.SegmentSize,
		SegmentLimit: config.SegmentLimit,
	}

	if !gfile.Exists(conf.Path) {
		if err := gfile.Mkdir(conf.Path); err != nil {
			g.Log().Errorf(ctx, "NewDiskQueue Failed to create the cache directory. Procedure, err:%+v", err)
			return nil
		}
	}

	queue, err := disk.New(conf)
	if err != nil {
		g.Log().Errorf(ctx, "NewDiskQueue err:%v", err)
		return nil
	}
	return queue
}
