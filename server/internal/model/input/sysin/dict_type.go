// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
)

// DictTypeEditInp 修改/新增字典数据
type DictTypeEditInp struct {
	entity.SysDictType
}
type DictTypeEditModel struct{}

// DictTypeDeleteInp 删除字典类型
type DictTypeDeleteInp struct {
	Id interface{}
}
type DictTypeDeleteModel struct{}

// DictTreeSelectInp 获取类型关系树选项
type DictTreeSelectInp struct {
}

type DictTreeSelectModel []g.Map

type DictTypeTree struct {
	entity.SysDictType
	Disabled bool            `json:"disabled"  dc:"是否禁用"`
	Label    string          `json:"label"     dc:"标签"`
	Value    int64           `json:"value"     dc:"键值"`
	Key      int64           `json:"key"       dc:"键名"`
	Children []*DictTypeTree `json:"children"  dc:"子级"`
}
