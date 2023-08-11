// Package form
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package form

// Sorter 排序器，兼容naiveUI
type Sorter struct {
	ColumnKey string      `json:"columnKey"  dc:"排序字段"`
	Sorter    bool        `json:"sorter"     dc:"是否需要排序"`
	Order     interface{} `json:"order"      dc:"排序方式 descend|ascend|false"`
}

type Sorters struct {
	Sorters []Sorter `json:"sorters"  dc:"排序器"`
}

func (s *Sorters) GetSorters() []Sorter {
	return s.Sorters
}
