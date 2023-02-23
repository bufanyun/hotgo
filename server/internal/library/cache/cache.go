// Package cache
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cache

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"hotgo/internal/library/cache/file"
	"hotgo/internal/model"
	"hotgo/internal/service"
)

// cache 缓存驱动
var cache *gcache.Cache

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		panic("cache uninitialized.")
	}
	return cache
}

// SetAdapter 设置缓存适配器
func SetAdapter(ctx context.Context) {
	var adapter gcache.Adapter
	conf, err := service.SysConfig().GetLoadCache(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "cache init err:%+v", err)
		return
	}

	if conf == nil {
		conf = new(model.CacheConfig)
		g.Log().Infof(ctx, "no cache driver is configured. default memory cache is used.")
	}

	switch conf.Adapter {
	case "redis":
		adapter = gcache.NewAdapterRedis(g.Redis())
	case "file":
		if conf.FileDir == "" {
			g.Log().Fatalf(ctx, "file path must be configured for file caching.")
			return
		}

		if !gfile.Exists(conf.FileDir) {
			if err := gfile.Mkdir(conf.FileDir); err != nil {
				g.Log().Fatalf(ctx, "Failed to create the cache directory. Procedure, err:%+v", err)
				return
			}
		}
		adapter = file.NewAdapterFile(conf.FileDir)
	default:
		adapter = gcache.NewAdapterMemory()
	}

	// 数据库缓存，默认和通用缓冲驱动一致，如果你不想使用默认的，可以自行调整
	g.DB().GetCache().SetAdapter(adapter)

	// 通用缓存
	cache = gcache.New()
	cache.SetAdapter(adapter)
	return
}
