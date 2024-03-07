// Package dict
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package dict

import (
	"context"
	"fmt"
	"hotgo/internal/model"
	"sync"
)

// FuncDict 方法字典，实现本接口即可使用内置方法字典
type FuncDict func(ctx context.Context) (res []*model.Option, err error)

type FuncOption struct {
	Id    int64           // 字典ID，由系统自动生成
	Key   string          // 字典选项key
	Label string          // 字典选项标签名称
	Fun   FuncDict        // 字典方法
	Cache bool            // 是否缓存数据选项
	Opts  []*model.Option // 缓存的数据选项
	sync.Mutex
}

var (
	funcOptions = make(map[string]*FuncOption)
	fLock       sync.Mutex
)

// GetAllFunc 获取所有方法字典
func GetAllFunc() map[string]*FuncOption {
	return funcOptions
}

// RegisterFunc 注册方法字典选项
func RegisterFunc(key, label string, fun FuncDict, cache ...bool) {
	fLock.Lock()
	defer fLock.Unlock()

	if len(key) == 0 {
		panic("字典key不能为空")
	}

	if _, ok := funcOptions[key]; ok {
		panic(fmt.Sprintf("重复注册方法选项:%v", key))
	}

	isCache := false
	if len(cache) > 0 {
		isCache = cache[0]
	}

	funcOptions[key] = &FuncOption{
		Id:    GenIdHash(key, FuncId),
		Key:   key,
		Label: label,
		Fun:   fun,
		Cache: isCache,
		Opts:  nil,
	}
}

// SaveFunc 更新方法字典选项
func SaveFunc(key, label string, fun FuncDict, cache ...bool) {
	fLock.Lock()
	defer fLock.Unlock()
	if _, ok := funcOptions[key]; ok {
		delete(funcOptions, key)
	}
	RegisterFunc(key, label, fun, cache...)
}

// ClearFuncCache 清理指定方法缓存选项
func ClearFuncCache(key string) (err error) {
	fun, ok := funcOptions[key]
	if !ok {
		err = NotExistKeyError
		return
	}

	fun.Lock()
	defer fun.Unlock()

	if fun.Opts != nil {
		fun.Opts = nil
	}
	return
}

// GetFuncOptions 获取指定方法字典的数据选项
func GetFuncOptions(ctx context.Context, key string) (res []*model.Option, err error) {
	fun, ok := funcOptions[key]
	if !ok {
		err = NotExistKeyError
		return
	}
	return LoadFuncOptions(ctx, fun)
}

// LoadFuncOptions 加载指定方法字典的数据选项
func LoadFuncOptions(ctx context.Context, fun *FuncOption) (res []*model.Option, err error) {
	if fun.Cache && fun.Opts != nil {
		res = fun.Opts
		return
	}

	fun.Lock()
	defer fun.Unlock()

	if fun.Cache && fun.Opts != nil {
		res = fun.Opts
		return
	}

	res, err = fun.Fun(ctx)
	if err != nil {
		return nil, err
	}

	for k := range res {
		res[k].Type = fun.Key
	}

	if fun.Cache {
		fun.Opts = res
	}
	return
}
