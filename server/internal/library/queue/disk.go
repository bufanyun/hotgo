// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"

	"hotgo/internal/library/queue/disk"
)

const (
	defaultDiskRetry     = 5
	defaultDiskRetryBase = time.Second
)

// Disk 磁盘队列

type DiskProducerMq struct {
	config    *disk.Config
	producers map[string]*disk.Queue
	sync.Mutex
}

type DiskConsumerMq struct {
	config    *disk.Config
	maxRetry  int
	retryBase time.Duration
	stop      context.Context
}

// RegisterDiskMqConsumer 注册磁盘队列消费者。
func RegisterDiskMqConsumer(config *disk.Config) (client MqConsumer, err error) {
	maxRetry := defaultDiskRetry
	base := defaultDiskRetryBase
	if config != nil && config.Retry > 0 {
		maxRetry = config.Retry
	}
	if config != nil && config.RetryInterval > 0 {
		base = time.Duration(config.RetryInterval) * time.Second
	}
	return &DiskConsumerMq{
		config:    config,
		maxRetry:  maxRetry,
		retryBase: base,
	}, nil
}

// ListenReceiveMsgDo 消费数据。处理成功才从队列消失；失败按次数退避后重试，超过次数进入死信。
// 坏消息同样进入死信，不堵住同主题后续消息。收到退出信号后循环结束。
func (q *DiskConsumerMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg) error) (err error) {
	if topic == "" {
		return gerror.New("disk.ListenReceiveMsgDo topic is empty")
	}
	ctx := q.stop
	release := func() {}
	if ctx == nil {
		ctx, release = beginConsumer()
	}
	defer release()

	queue := NewDiskQueue(topic, q.config)
	if queue == nil {
		return gerror.New("disk queue init failed")
	}
	defer queue.Close()

	for {
		if ctx.Err() != nil {
			return nil
		}
		index, offset, data, readErr := queue.Read()
		if readErr != nil {
			msg, ok, popErr := delayPop(q.config, topic, time.Now())
			if popErr != nil {
				Logger().Warningf(ctx, "disk delay pop err:%+v", popErr)
			}
			if ok {
				q.finish(topic, msg, receiveDo(msg), nil)
				continue
			}
			if !q.sleep(ctx, 30*time.Millisecond) {
				return nil
			}
			continue
		}

		var mqMsg MqMsg
		if unmarshalErr := json.Unmarshal(data, &mqMsg); unmarshalErr != nil || mqMsg.MsgId == "" {
			_ = appendDeadLetter(q.config, topic, data, "bad message")
			queue.Commit(index, offset)
			continue
		}
		if !mqMsg.ExecuteAt.IsZero() && mqMsg.ExecuteAt.After(time.Now()) {
			_ = delayAdd(q.config, topic, mqMsg)
			queue.Commit(index, offset)
			continue
		}
		commit := func() { queue.Commit(index, offset) }
		q.finish(topic, mqMsg, receiveDo(mqMsg), commit)
	}
}

// finish 成功后提交；失败按次数退避，超过次数或坏消息进入死信。commit 为空表示消息已经离开磁盘日志。
func (q *DiskConsumerMq) finish(topic string, mqMsg MqMsg, handleErr error, commit func()) {
	if commit == nil {
		commit = func() {}
	}
	if handleErr == nil {
		commit()
		return
	}
	mqMsg.Retry++
	if mqMsg.Retry > q.maxRetry {
		body, _ := json.Marshal(mqMsg)
		_ = appendDeadLetter(q.config, topic, body, handleErr.Error())
		commit()
		return
	}
	mqMsg.ExecuteAt = time.Now().Add(q.backoff(mqMsg.Retry))
	_ = delayAdd(q.config, topic, mqMsg)
	commit()
}

// backoff 按已重试次数指数增加等待时间，最长 5 分钟。
func (q *DiskConsumerMq) backoff(retry int) time.Duration {
	base := q.retryBase
	if base <= 0 {
		base = defaultDiskRetryBase
	}
	d := base
	for i := 1; i < retry; i++ {
		d *= 2
		if d > 5*time.Minute {
			return 5 * time.Minute
		}
	}
	return d
}

// sleep 等待一段时间。收到退出信号时返回 false。
func (q *DiskConsumerMq) sleep(ctx context.Context, d time.Duration) bool {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

// RegisterDiskMqProducer 注册磁盘队列生产者。
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
		return mqMsg, gerror.New(fmt.Sprint("queue disk 生产者解析json消息失败:", err))
	}

	queue := d.getProducer(topic)
	if queue == nil {
		return mqMsg, gerror.New("queue disk 生产者初始化失败")
	}
	if err = queue.Write(mqMsgJson); err != nil {
		return mqMsg, gerror.New(fmt.Sprint("queue disk 生产者添加消息失败:", err))
	}
	return
}

// SendDelayMsg 延迟消息先落到延迟文件，到点后才投递到可消费队列。
func (d *DiskProducerMq) SendDelayMsg(topic string, body string, delaySecond int64) (mqMsg MqMsg, err error) {
	if topic == "" {
		return mqMsg, gerror.New("DiskMq topic is empty")
	}
	if delaySecond < 0 {
		delaySecond = 0
	}
	mqMsg = MqMsg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      []byte(body),
		Timestamp: time.Now(),
		ExecuteAt: time.Now().Add(time.Duration(delaySecond) * time.Second),
	}
	if err = delayAdd(d.config, topic, mqMsg); err != nil {
		return mqMsg, gerror.Wrap(err, "queue disk 延迟消息写入失败")
	}
	return
}

// getProducer 按主题获取生产者。整个过程持锁，避免并发读写 map。
func (d *DiskProducerMq) getProducer(topic string) *disk.Queue {
	d.Lock()
	defer d.Unlock()
	if queue, ok := d.producers[topic]; ok {
		return queue
	}
	queue := NewDiskQueue(topic, d.config)
	d.producers[topic] = queue
	return queue
}

// NewDiskQueue 按主题打开磁盘队列目录。
func NewDiskQueue(topic string, config *disk.Config) *disk.Queue {
	if config == nil {
		return nil
	}
	conf := &disk.Config{
		Path:         config.Path + "/" + config.GroupName + "/" + topic,
		BatchSize:    config.BatchSize,
		BatchTime:    config.BatchTime * time.Second,
		SegmentSize:  config.SegmentSize,
		SegmentLimit: config.SegmentLimit,
	}

	if !gfile.Exists(conf.Path) {
		if err := gfile.Mkdir(conf.Path); err != nil {
			Logger().Errorf(ctx, "NewDiskQueue Failed to create the cache directory. Procedure, err:%+v", err)
			return nil
		}
	}

	queue, err := disk.New(conf)
	if err != nil {
		Logger().Errorf(ctx, "NewDiskQueue err:%v", err)
		return nil
	}
	return queue
}
