package handler

import "github.com/gogf/gf/v2/database/gdb"

// ForceCache 强制缓存
func ForceCache(m *gdb.Model) *gdb.Model {
	return m.Cache(gdb.CacheOption{Duration: -1, Force: true})
}
