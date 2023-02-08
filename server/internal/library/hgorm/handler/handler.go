// Package handler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package handler

// handler.
import (
	"github.com/gogf/gf/v2/database/gdb"
)

// Option 预处理选项
type Option struct {
	FilterAuth bool // 过滤权限
	ForceCache bool // 强制缓存
}

// DefaultOption 默认预处理选项
var DefaultOption = &Option{
	FilterAuth: true,
}

func Model(m *gdb.Model, opt ...*Option) *gdb.Model {
	var option *Option
	if len(opt) > 0 {
		option = opt[0]
	} else {
		option = DefaultOption
	}
	if option.FilterAuth {
		m = m.Handler(FilterAuth)
	}
	if option.ForceCache {
		m = m.Handler(ForceCache)
	}
	return m
}
