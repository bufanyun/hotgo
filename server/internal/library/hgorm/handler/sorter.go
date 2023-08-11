package handler

import (
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/internal/model/input/form"
)

// ISorter 排序器接口，实现该接口即可使用Handler匹配排序，支持多字段排序
type ISorter interface {
	GetSorters() []form.Sorter
}

// Sorter 排序器
func Sorter(in ISorter) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		hasSort := false
		sorters := in.GetSorters()
		for _, sorter := range sorters {
			if len(sorter.ColumnKey) == 0 || !sorter.Sorter {
				continue
			}

			switch sorter.Order {
			case "descend":
				hasSort = true
				m = m.OrderDesc(sorter.ColumnKey)
			case "ascend":
				hasSort = true
				m = m.OrderAsc(sorter.ColumnKey)
			default:
				continue
			}
		}

		// 不存在排序条件
		if !hasSort {
			// ...
		}
		return m
	}
}
