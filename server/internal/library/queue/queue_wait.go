package queue

import (
	"context"
	"sync"
	"time"
)

// queueIdleRecheck 没有待处理消息时的兜底检查间隔。
// 同进程写入会立刻唤醒；这个间隔只用来发现别的进程写入、或错过的通知。
const queueIdleRecheck = time.Minute

var queueSlots sync.Map // topic -> chan struct{}

// queueSlot 每个主题一个缓冲为 1 的通知。忙的时候信号会留着，空闲时合并成一次唤醒。
func queueSlot(topic string) chan struct{} {
	if v, ok := queueSlots.Load(topic); ok {
		return v.(chan struct{})
	}
	ch := make(chan struct{}, 1)
	actual, _ := queueSlots.LoadOrStore(topic, ch)
	return actual.(chan struct{})
}

// publishQueue 通知本进程里正在等这个主题的消费者。
func publishQueue(topic string) {
	ch := queueSlot(topic)
	select {
	case ch <- struct{}{}:
	default:
	}
}

// waitQueue 等到截止时间、新消息或退出。返回 false 表示该停止消费。
func waitQueue(ctx context.Context, topic string, until time.Time) bool {
	timer := time.NewTimer(time.Until(until))
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	case <-queueSlot(topic):
		return true
	}
}
