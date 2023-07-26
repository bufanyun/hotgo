// Package lock_test
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package lock_test

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/library/hgrds/lock"
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
