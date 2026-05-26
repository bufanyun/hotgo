// Package form
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package form

import "github.com/gogf/gf/v2/util/gconv"

// Selects 选项
type Selects []*Select

type Select struct {
	Value    interface{} `json:"value"`
	Label    string      `json:"label"`
	Name     string      `json:"name"`
	Disabled bool        `json:"disabled"`
}

func (p Selects) Len() int {
	return len(p)
}

func (p Selects) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Selects) Less(i, j int) bool {
	return gconv.Int64(p[j].Value) > gconv.Int64(p[i].Value)
}
