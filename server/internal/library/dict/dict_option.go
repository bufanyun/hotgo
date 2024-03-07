// Package dict
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package dict

import (
	"github.com/gogf/gf/v2/util/gconv"
	"hash/fnv"
	"hotgo/internal/model"
)

// GenDefaultOption 生成默认表格回显样式
func GenDefaultOption(key interface{}, label string) *model.Option {
	return &model.Option{
		Key:       key,
		Label:     label,
		Value:     key,
		ListClass: "default",
	}
}

func GenSuccessOption(key interface{}, label string) *model.Option {
	return &model.Option{
		Key:       key,
		Label:     label,
		Value:     key,
		ListClass: "success",
	}
}

func GenWarningOption(key interface{}, label string) *model.Option {
	return &model.Option{
		Key:       key,
		Label:     label,
		Value:     key,
		ListClass: "warning",
	}
}

func GenErrorOption(key interface{}, label string) *model.Option {
	return &model.Option{
		Key:       key,
		Label:     label,
		Value:     key,
		ListClass: "error",
	}
}

func GenInfoOption(key interface{}, label string) *model.Option {
	return &model.Option{
		Key:       key,
		Label:     label,
		Value:     key,
		ListClass: "info",
	}
}

// GenCustomOption 生成自定义表格回显样式
func GenCustomOption(key interface{}, label string, custom string) *model.Option {
	return &model.Option{
		Key:       key,
		Label:     label,
		Value:     key,
		ListClass: custom,
	}
}

// GenHashOption 根据不同label以hash算法生成表格回显样式
func GenHashOption(key interface{}, label string) *model.Option {
	strings := []string{"default", "primary", "info", "success", "warning", "error"}
	hash := fnv.New32()

	tag := "default"
	if _, err := hash.Write(gconv.Bytes(label)); err == nil {
		index := int(hash.Sum32()) % len(strings)
		tag = strings[index]
	}
	return &model.Option{
		Key:       key,
		Label:     label,
		Value:     key,
		ListClass: tag,
	}
}

// GetOptionLabel 通过key找到label
func GetOptionLabel(ses []*model.Option, key interface{}) string {
	for _, v := range ses {
		if gconv.String(v.Key) == gconv.String(key) {
			return v.Label
		}
	}
	return `Unknown`
}

// HasOptionKey 是否存在指定key
func HasOptionKey(ses []*model.Option, key interface{}) bool {
	for _, v := range ses {
		if gconv.String(v.Key) == gconv.String(key) {
			return true
		}
	}
	return false
}
