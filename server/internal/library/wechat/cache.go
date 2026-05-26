// Package wechat
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package wechat

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"hotgo/internal/library/cache"
	"time"
)

type Cache struct {
	ctx   context.Context
	cache *gcache.Cache
}

// NewCache 实例化
func NewCache(ctx context.Context, name ...*gcache.Cache) *Cache {
	var defaultCache = cache.Instance()
	if len(name) > 0 {
		defaultCache = name[0]
	}
	return &Cache{ctx: ctx, cache: defaultCache}
}

// SetCache 设置缓存驱动
func (r *Cache) SetCache(cache *gcache.Cache) {
	r.cache = cache
}

// SetCtx 设置 ctx 参数
func (r *Cache) SetCtx(ctx context.Context) {
	r.ctx = ctx
}

// Get 获取一个值
func (r *Cache) Get(key string) interface{} {
	get, err := r.cache.Get(r.ctx, key)
	if err != nil || get.IsNil() || get.IsEmpty() {
		return nil
	}
	return get.Interface()
}

// Set 设置一个值
func (r *Cache) Set(key string, val interface{}, timeout time.Duration) error {
	return r.cache.Set(r.ctx, key, val, timeout)
}

// IsExist 判断key是否存在
func (r *Cache) IsExist(key string) bool {
	contains, err := r.cache.Contains(r.ctx, key)
	if err != nil {
		return false
	}
	return contains
}

// Delete 删除
func (r *Cache) Delete(key string) error {
	_, err := r.cache.Remove(r.ctx, key)
	if err != nil {
		return err
	}
	return nil
}
