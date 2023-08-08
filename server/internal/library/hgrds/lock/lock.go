// Package lock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package lock

// 分布式锁

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"sync"
	"time"
)

// Config 锁配置
type Config struct {
	ttl             time.Duration // 过期时间
	tryLockInterval time.Duration // 重新获取锁间隔
}

// Lock 一把锁 不可重复使用
type Lock struct {
	resource        string        // 锁定的资源
	randomValue     string        // 随机值
	watchDog        chan struct{} // 看门狗
	ttl             time.Duration // 过期时间
	tryLockInterval time.Duration // 重新获取锁间隔
	wg              sync.WaitGroup
}

// NewConfig 初始化一个锁配置
func NewConfig(ttl, tryLockInterval time.Duration) *Config {
	return &Config{
		ttl:             ttl,
		tryLockInterval: tryLockInterval,
	}
}

// Mutex 根据配置创建一把锁
func (lc *Config) Mutex(resource string) *Lock {
	return &Lock{
		resource:        resource,
		randomValue:     guid.S(),
		watchDog:        make(chan struct{}),
		ttl:             lc.ttl,
		tryLockInterval: lc.tryLockInterval,
	}
}

// Lock 阻塞加锁
func (l *Lock) Lock(ctx context.Context) error {
	// 尝试加锁
	err := l.TryLock(ctx)
	if err == nil {
		return nil
	}
	if !gerror.Is(err, ErrLockFailed) {
		return err
	}
	// 加锁失败，不断尝试
	ticker := time.NewTicker(l.tryLockInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			// 超时
			return ErrTimeout
		case <-ticker.C:
			// 重新尝试加锁
			err = l.TryLock(ctx)
			if err == nil {
				return nil
			}
			if !gerror.Is(err, ErrLockFailed) {
				return err
			}
		}
	}
}

// TryLock 尝试加锁，如果失败立即返回错误，而不会阻塞等待锁
func (l *Lock) TryLock(ctx context.Context) error {
	var args = []interface{}{l.randomValue, l.ttl.Seconds()}
	eval, err := g.Redis().GroupScript().Eval(ctx, lockScript, 1, []string{l.resource}, args)
	if err != nil {
		return err
	}

	if eval.String() != "OK" {
		return ErrLockFailed
	}

	go l.startWatchDog()
	return nil
}

// Unlock 解锁
func (l *Lock) Unlock(ctx context.Context) error {
	var args []interface{}
	args = append(args, l.randomValue)
	eval, err := g.Redis().GroupScript().Eval(ctx, unlockScript, 1, []string{l.resource}, args)

	if eval.Int() == 2 {
		return ErrNotCaller
	}

	if eval.Int() == 0 {
		return ErrNotExist
	}

	close(l.watchDog)
	return err
}

func (l *Lock) LockFunc(ctx context.Context, f func()) error {
	if err := l.Lock(ctx); err != nil {
		return err
	}
	defer func() {
		_ = l.Unlock(ctx)
	}()
	f()
	return nil
}

// TryLockFunc tries locking the mutex for writing with given callback function `f`.
// it returns true immediately if success, or if there's a lock on the mutex,
// it returns error immediately.
//
// It releases the lock after `f` is executed.
func (l *Lock) TryLockFunc(ctx context.Context, f func()) error {
	err := l.TryLock(ctx)
	if err != nil {
		defer func() {
			_ = l.Unlock(ctx)
		}()
		f()
	}
	return err
}

// startWatchDog 看门狗
func (l *Lock) startWatchDog() {
	resetTTLInterval := l.ttl / 3
	ticker := time.NewTicker(resetTTLInterval)
	defer ticker.Stop()

	l.wg.Add(1)
	defer l.wg.Wait()

	conn := g.Redis()
	for {
		select {
		case <-ticker.C:
			// 延长锁的过期时间
			ctx, cancel := context.WithTimeout(context.Background(), resetTTLInterval)
			var args = []interface{}{l.ttl.Seconds(), l.randomValue}
			eval, err := conn.GroupScript().Eval(ctx, renewalScript, 1, []string{l.resource}, args)
			cancel()

			// 异常或锁已经不存在则不再续期
			if err != nil || eval.Int() < 1 {
				return
			}
		case <-l.watchDog:
			// 已经解锁
			return
		}
	}
}

// defaultMutex 一个默认配置锁
var defaultMutex = NewConfig(DefaultTTL, DefaultTryLockInterval)

// Mutex 获取默认锁
func Mutex(resource string) *Lock {
	return defaultMutex.Mutex(resource)
}
