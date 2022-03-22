//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package com

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

// 缓存
var Cache = new(cache)

type cache struct{}

func (component *cache) New() *gcache.Cache {
	c := gcache.New()

	//redis
	adapter := gcache.NewAdapterRedis(g.Redis())

	//内存
	//adapter := gcache.NewAdapterMemory()
	c.SetAdapter(adapter)
	return c
}
