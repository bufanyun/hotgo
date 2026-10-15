package queue

import (
	"context"
	"sync"
	"time"
)

var (
	rootCtx, rootCancel = context.WithCancel(context.Background())
	consumerWg          sync.WaitGroup
)

// beginConsumer 登记一个消费循环。返回的函数在循环退出时调用。
func beginConsumer() (context.Context, func()) {
	ctx, cancel := context.WithCancel(rootCtx)
	consumerWg.Add(1)
	return ctx, func() {
		cancel()
		consumerWg.Done()
	}
}

// StopConsumers 通知全部消费循环退出，并等待它们结束。
func StopConsumers() {
	rootCancel()
	done := make(chan struct{})
	go func() {
		consumerWg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(8 * time.Second):
		Logger().Warning(context.Background(), "停止队列消费者超时")
	}
}
