// Package lock_test
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package lock_test

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/library/hgrds/lock"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestDefaultLock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		l := lock.Mutex("test")
		err := l.TryLock(context.Background())
		if err != nil {
			t.Error(err)
		}
		t.Log("A加锁成功")
		time.Sleep(lock.DefaultTTL)
		err = l.Unlock(context.Background())
		if err != nil {
			t.Error(err)
		}
		t.Log("A已释放锁")
	}()

	time.Sleep(time.Second)

	go func() {
		defer wg.Done()
		for {
			l := lock.Mutex("test")
			err := l.TryLock(context.Background())
			if err == nil {
				t.Log("B加锁成功")

				// 等待1s，模拟业务
				time.Sleep(time.Second)

				err = l.Unlock(context.Background())
				if err != nil {
					t.Error(err)
				}
				t.Log("B已释放锁")
				break
			}

			if gerror.Is(err, lock.ErrLockFailed) {
				t.Log("B加锁失败，等待1s重试...")
				time.Sleep(time.Second)
			} else {
				t.Error(err)
				return
			}
		}
	}()
	wg.Wait()
}

func TestNewLock(t *testing.T) {
	locker := lock.NewConfig(time.Second*30, time.Second)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		l := locker.Mutex("test")
		err := l.TryLock(context.Background())
		if err != nil {
			t.Error(err)
		}
		time.Sleep(lock.DefaultTTL)
		err = l.Unlock(context.Background())
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(time.Second)

	go func() {
		defer wg.Done()
		l := lock.Mutex("test")
		err := l.TryLock(context.Background())
		if err != nil && !gerror.Is(err, lock.ErrLockFailed) {
			t.Error(err)
		}
	}()
	wg.Wait()
}

func TestNewLock2(t *testing.T) {
	locker := lock.NewConfig(time.Second*30, time.Second)
	var wg sync.WaitGroup
	wg.Add(2)
	count := 0
	times := 1000
	go func() {
		defer wg.Done()
		for i := 0; i < times; i++ {
			l := locker.Mutex("test")
			err := l.Lock(context.Background())
			if err != nil {
				t.Error(err)
			}
			count++
			err = l.Unlock(context.Background())
			if err != nil {
				t.Error(err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < times; i++ {
			l := lock.Mutex("test")
			err := l.Lock(context.Background())
			if err != nil {
				t.Error(err)
			}
			count++
			err = l.Unlock(context.Background())
			if err != nil {
				t.Error(err)
			}
		}
	}()
	wg.Wait()
	if count != times*2 {
		t.Errorf("count = %d", count)
	}
}

func Test_Fix_watchDogMemoryLeak(t *testing.T) {
	i := 0
	for i < 5 {
		TestDefaultLock(t)
		t.Log("current goroutine num:", runtime.NumGoroutine())
		i++
	}
}
