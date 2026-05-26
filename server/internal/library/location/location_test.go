// Package location_test
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package location_test

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"hotgo/internal/library/location"
	"sync"
	"testing"
)

var ip = "120.12.151.65"

func Test_GetLocation(t *testing.T) {
	ctx := gctx.New()
	data, err := location.GetLocation(ctx, ip)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("data:%v", gjson.New(data).String())
}

func Test_ParallelGetLocation(t *testing.T) {
	ctx := gctx.New()
	start := gtime.Now()

	t.Log("start")
	for i := 0; i < 10; i++ {
		data, err := location.GetLocation(ctx, ip)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("index:%v, data:%v", i, gjson.New(data).String())
	}

	t.Logf("总耗时:%vs", gtime.Now().Sub(start).Seconds())
}

func Test_ConcurrentGetLocation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()
		start := gtime.Now()
		var wg sync.WaitGroup

		t.Log("start")
		for i := 0; i < 10; i++ {
			wg.Add(1)
			index := i
			go func() {
				defer wg.Done()

				data, err := location.GetLocation(ctx, ip)
				if err != nil {
					t.Error(err)
					return
				}
				t.Logf("index:%v, data:%v", index, gjson.New(data).String())
			}()
		}

		wg.Wait()
		t.Logf("总耗时:%vs", gtime.Now().Sub(start).Seconds())
	})
}
