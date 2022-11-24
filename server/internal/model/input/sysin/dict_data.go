// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// DictDataEditInp 修改/新增字典数据
type DictDataEditInp struct {
	entity.SysDictData
	TypeID int64
}
type DictDataEditModel struct{}

// DictDataDeleteInp 删除字典数据
type DictDataDeleteInp struct {
	Id interface{}
}
type DictDataDeleteModel struct{}

// DictDataListInp 获取列表
type DictDataListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	TypeID int64
	Type   string
	Label  string
}

type DictDataListModel struct {
	TypeID int64 `json:"typeId"`
	entity.SysDictData
}
