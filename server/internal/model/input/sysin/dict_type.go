// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
)

// DictTypeEditInp 修改/新增字典数据
type DictTypeEditInp struct {
	entity.SysDictType
}

func (in *DictTypeEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}

	if in.Id > 0 && in.Id == in.Pid {
		err = gerror.New("上级字典不能是自己")
		return
	}

	return
}

type DictTypeEditModel struct{}

// DictTypeUpdateFields 修改数据字段过滤
type DictTypeUpdateFields struct {
	Id     int64  `json:"id"        description:"字典类型ID"`
	Pid    int64  `json:"pid"       description:"父类字典类型ID"`
	Name   string `json:"name"      description:"字典类型名称"`
	Type   string `json:"type"      description:"字典类型"`
	Sort   int    `json:"sort"      description:"排序"`
	Remark string `json:"remark"    description:"备注"`
	Status int    `json:"status"    description:"字典类型状态"`
}

// DictTypeInsertFields 新增数据字段过滤
type DictTypeInsertFields struct {
	Pid    int64  `json:"pid"       description:"父类字典类型ID"`
	Name   string `json:"name"      description:"字典类型名称"`
	Type   string `json:"type"      description:"字典类型"`
	Sort   int    `json:"sort"      description:"排序"`
	Remark string `json:"remark"    description:"备注"`
	Status int    `json:"status"    description:"字典类型状态"`
}

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
