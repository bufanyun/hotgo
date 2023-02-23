// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
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
	TypeID int64  `json:"typeId"`
	Key    string `json:"key"`
	entity.SysDictData
}

// DataSelectInp 获取指定字典选项
type DataSelectInp struct {
	Type string
}

type DataSelectModel []*SelectData

type SelectData struct {
	Key       interface{} `json:"key"`
	Label     string      `json:"label"     description:"字典标签"`
	Value     interface{} `json:"value"     description:"字典键值"`
	ValueType string      `json:"valueType" description:"键值数据类型：string,int,uint,bool,datetime,date"`
	Type      string      `json:"type"      description:"字典类型"`
	ListClass string      `json:"listClass" description:"表格回显样式"`
}
