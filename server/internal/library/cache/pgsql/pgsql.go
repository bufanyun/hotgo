// Package pgsql 把 gcache 数据放到 PostgreSQL。
package pgsql

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

const cacheTable = "hg_sys_cache"

// Adapter PostgreSQL 缓存驱动
type Adapter struct {
	db gdb.DB
	mu sync.Mutex
}

// New 用当前 PostgreSQL 连接构造缓存适配器。
func New(db gdb.DB) gcache.Adapter {
	return &Adapter{db: db}
}

// Set 写入缓存。duration 小于 0 或值为空时删除；0 表示不过期。
func (a *Adapter) Set(ctx context.Context, key any, value any, duration time.Duration) error {
	fileKey := gconv.String(key)
	if value == nil || duration < 0 {
		_, err := a.Remove(ctx, fileKey)
		return err
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	expireAt := int64(0)
	if duration > 0 {
		expireAt = time.Now().Add(duration).Unix()
		if expireAt <= time.Now().Unix() {
			expireAt = time.Now().Unix() + 1
		}
	}
	_, err = a.db.Exec(ctx, `INSERT INTO `+cacheTable+`(cache_key, cache_value, expire_at) VALUES(?,?,?)
ON CONFLICT (cache_key) DO UPDATE SET cache_value=excluded.cache_value, expire_at=excluded.expire_at`, fileKey, string(raw), expireAt)
	return err
}

// SetMap 批量写入，过期时间相同。
func (a *Adapter) SetMap(ctx context.Context, data map[any]any, duration time.Duration) error {
	for k, v := range data {
		if err := a.Set(ctx, k, v, duration); err != nil {
			return err
		}
	}
	return nil
}

// SetIfNotExist 键不存在时才写入。
func (a *Adapter) SetIfNotExist(ctx context.Context, key any, value any, duration time.Duration) (bool, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	found, err := a.Get(ctx, key)
	if err != nil {
		return false, err
	}
	if found != nil && !found.IsNil() {
		return false, nil
	}
	return true, a.Set(ctx, key, value, duration)
}

// SetIfNotExistFunc 键不存在时执行函数并写入。
func (a *Adapter) SetIfNotExistFunc(ctx context.Context, key any, f gcache.Func, duration time.Duration) (bool, error) {
	return a.SetIfNotExistFuncLock(ctx, key, f, duration)
}

// SetIfNotExistFuncLock 与 SetIfNotExistFunc 相同，并在本进程内加锁。
func (a *Adapter) SetIfNotExistFuncLock(ctx context.Context, key any, f gcache.Func, duration time.Duration) (bool, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	found, err := a.Get(ctx, key)
	if err != nil {
		return false, err
	}
	if found != nil && !found.IsNil() {
		return false, nil
	}
	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	if value == nil {
		return false, nil
	}
	return true, a.Set(ctx, key, value, duration)
}

// Get 读取缓存。过期数据会删掉并当作未命中。
func (a *Adapter) Get(ctx context.Context, key any) (*gvar.Var, error) {
	one, err := a.db.GetOne(ctx, "SELECT cache_value, expire_at FROM "+cacheTable+" WHERE cache_key=?", gconv.String(key))
	if err != nil || one.IsEmpty() {
		return gvar.New(nil), err
	}
	expireAt := one["expire_at"].Int64()
	if expireAt > 0 && expireAt <= time.Now().Unix() {
		_, _ = a.db.Exec(ctx, "DELETE FROM "+cacheTable+" WHERE cache_key=?", gconv.String(key))
		return gvar.New(nil), nil
	}
	var decoded any
	if err = json.Unmarshal([]byte(one["cache_value"].String()), &decoded); err != nil {
		return nil, err
	}
	return gvar.New(decoded), nil
}

// GetOrSet 读到就返回，没有则写入后返回。
func (a *Adapter) GetOrSet(ctx context.Context, key any, value any, duration time.Duration) (*gvar.Var, error) {
	v, err := a.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v == nil || v.IsNil() {
		return gvar.New(value), a.Set(ctx, key, value, duration)
	}
	return v, nil
}

// GetOrSetFunc 没有缓存时执行函数并写入。
func (a *Adapter) GetOrSetFunc(ctx context.Context, key any, f gcache.Func, duration time.Duration) (*gvar.Var, error) {
	return a.GetOrSetFuncLock(ctx, key, f, duration)
}

// GetOrSetFuncLock 与 GetOrSetFunc 相同，并在本进程内加锁。
func (a *Adapter) GetOrSetFuncLock(ctx context.Context, key any, f gcache.Func, duration time.Duration) (*gvar.Var, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	v, err := a.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v != nil && !v.IsNil() {
		return v, nil
	}
	value, err := f(ctx)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, nil
	}
	return gvar.New(value), a.Set(ctx, key, value, duration)
}

// Contains 判断键是否存在且未过期。
func (a *Adapter) Contains(ctx context.Context, key any) (bool, error) {
	v, err := a.Get(ctx, key)
	if err != nil || v == nil {
		return false, err
	}
	return !v.IsNil(), nil
}

// Size 返回表中的记录数，含已过期但还没被读到的行。
func (a *Adapter) Size(ctx context.Context) (int, error) {
	v, err := a.db.GetValue(ctx, "SELECT COUNT(1) FROM "+cacheTable)
	if err != nil {
		return 0, err
	}
	return v.Int(), nil
}

// Data 返回全部未过期的键值。
func (a *Adapter) Data(ctx context.Context) (map[any]any, error) {
	keys, err := a.Keys(ctx)
	if err != nil {
		return nil, err
	}
	out := make(map[any]any, len(keys))
	for _, key := range keys {
		v, err := a.Get(ctx, key)
		if err != nil {
			return nil, err
		}
		out[key] = v.Val()
	}
	return out, nil
}

// Keys 返回全部未过期的键。
func (a *Adapter) Keys(ctx context.Context) ([]any, error) {
	all, err := a.db.GetAll(ctx, "SELECT cache_key, expire_at FROM "+cacheTable)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	keys := make([]any, 0, len(all))
	for _, row := range all {
		expireAt := row["expire_at"].Int64()
		if expireAt > 0 && expireAt <= now {
			continue
		}
		keys = append(keys, row["cache_key"].String())
	}
	return keys, nil
}

// Values 返回全部未过期的值。
func (a *Adapter) Values(ctx context.Context) ([]any, error) {
	data, err := a.Data(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(data))
	for _, v := range data {
		out = append(out, v)
	}
	return out, nil
}

// Update 只更新已存在的值，并保留剩余过期时间。
func (a *Adapter) Update(ctx context.Context, key any, value any) (*gvar.Var, bool, error) {
	old, err := a.Get(ctx, key)
	if err != nil {
		return nil, false, err
	}
	if old == nil || old.IsNil() {
		return nil, false, nil
	}
	expire, err := a.GetExpire(ctx, key)
	if err != nil {
		return nil, false, err
	}
	return old, true, a.Set(ctx, key, value, expire)
}

// UpdateExpire 修改已存在键的过期时间，返回修改前的剩余时间。
func (a *Adapter) UpdateExpire(ctx context.Context, key any, duration time.Duration) (time.Duration, error) {
	old, err := a.GetExpire(ctx, key)
	if err != nil {
		return 0, err
	}
	v, err := a.Get(ctx, key)
	if err != nil {
		return 0, err
	}
	if v == nil || v.IsNil() {
		return -1, nil
	}
	return old, a.Set(ctx, key, v.Val(), duration)
}

// GetExpire 返回剩余过期时间。0 表示不过期，-1 表示不存在或已过期。
func (a *Adapter) GetExpire(ctx context.Context, key any) (time.Duration, error) {
	one, err := a.db.GetOne(ctx, "SELECT expire_at FROM "+cacheTable+" WHERE cache_key=?", gconv.String(key))
	if err != nil || one.IsEmpty() {
		return -1, err
	}
	expireAt := one["expire_at"].Int64()
	if expireAt == 0 {
		return 0, nil
	}
	remain := time.Until(time.Unix(expireAt, 0))
	if remain <= 0 {
		return -1, nil
	}
	return remain, nil
}

// Remove 删除指定键，返回最后一个键删除前的值。
func (a *Adapter) Remove(ctx context.Context, keys ...any) (*gvar.Var, error) {
	if len(keys) == 0 {
		return nil, nil
	}
	last, err := a.Get(ctx, keys[len(keys)-1])
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		if _, err = a.db.Exec(ctx, "DELETE FROM "+cacheTable+" WHERE cache_key=?", gconv.String(key)); err != nil {
			return nil, err
		}
	}
	return last, nil
}

// Clear 清空缓存表。
func (a *Adapter) Clear(ctx context.Context) error {
	_, err := a.db.Exec(ctx, "DELETE FROM "+cacheTable)
	return err
}

// Close 数据库连接由框架管理，这里不关闭。
func (a *Adapter) Close(ctx context.Context) error { return nil }
