// Package handler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package handler

import (
	"github.com/gogf/gf/v2/database/gdb"
)

// ForceCache 强制缓存
func ForceCache(m *gdb.Model) *gdb.Model {
	return m.Cache(gdb.CacheOption{Duration: 0, Force: true})
}
