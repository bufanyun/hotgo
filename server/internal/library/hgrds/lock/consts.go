// Package lock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package lock

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

const (
	// 加锁脚本
	lockScript = `
	local token = redis.call('get', KEYS[1])
	if token then
		return 0
	else
		local setResult = redis.call('setex', KEYS[1], ARGV[2], ARGV[1])
		return setResult
	end
`

	// 续约脚本
	renewalScript = `
	if redis.call('get',KEYS[1])==ARGV[2] then
		return redis.call('expire',KEYS[1],ARGV[1])
	end
	return 0
`

	// 解锁脚本
	unlockScript = `
	if redis.call("get",KEYS[1]) == ARGV[1] then
		return redis.call("del",KEYS[1])
	else
		return 2
	end
`
)

const (
	// DefaultTTL 锁默认过期时间
	DefaultTTL = time.Second * 10
	// DefaultTryLockInterval 默认重试获取锁间隔时间
	DefaultTryLockInterval = time.Millisecond * 100
)

var (
	// ErrLockFailed 加锁失败
	ErrLockFailed = gerror.New("lock failed")
	// ErrTimeout 加锁超时
	ErrTimeout = gerror.New("timeout")
	// ErrNotCaller 锁持有者不是当前实例
	ErrNotCaller = gerror.New("lock not held by the caller")
	// ErrNotExist 锁不存在
	ErrNotExist = gerror.New("lock does not exist")
)
