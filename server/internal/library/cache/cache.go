// Package cache
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package cache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

func New() *gcache.Cache {
	c := gcache.New()

	//redis
	adapter := gcache.NewAdapterRedis(g.Redis())

	//内存
	//adapter := gcache.NewAdapterMemory()
	c.SetAdapter(adapter)
	return c
}
