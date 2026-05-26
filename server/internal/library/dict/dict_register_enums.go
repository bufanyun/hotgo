// Package dict
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package dict

import (
	"fmt"
	"hotgo/internal/model"
	"sync"
)

type EnumsOption struct {
	Id    int64           // 字典ID，由系统自动生成
	Key   string          // 字典选项key
	Label string          // 字典选项标签名称
	Opts  []*model.Option // 数据选项
}

var (
	enumsOptions = make(map[string]*EnumsOption)
	eLock        sync.Mutex
)

// GetAllEnums 获取所有枚举字典
func GetAllEnums() map[string]*EnumsOption {
	return enumsOptions
}

// RegisterEnums 注册枚举字典选项
func RegisterEnums(key, label string, opts []*model.Option) {
	eLock.Lock()
	defer eLock.Unlock()

	if len(key) == 0 {
		panic("字典key不能为空")
	}

	if _, ok := enumsOptions[key]; ok {
		panic(fmt.Sprintf("重复注册枚举字典选项:%v", key))
	}

	for _, v := range opts {
		v.Type = key
	}
	enumsOptions[key] = &EnumsOption{
		Id:    GenIdHash(key, EnumsId),
		Key:   key,
		Label: label,
		Opts:  opts,
	}
}

// SaveEnums 更新枚举字典选项
func SaveEnums(key, label string, opts []*model.Option) {
	eLock.Lock()
	defer eLock.Unlock()
	delete(enumsOptions, key)
	RegisterEnums(key, label, opts)
}

// GetEnumsOptions 获取指定枚举字典的数据选项
func GetEnumsOptions(key string) []*model.Option {
	enums, ok := enumsOptions[key]
	if !ok {
		return nil
	}
	return enums.Opts
}
