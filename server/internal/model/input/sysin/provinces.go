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

// ProvincesMaxSortInp 最大排序
type ProvincesMaxSortInp struct {
}

type ProvincesMaxSortModel struct {
	Sort int `json:"sort" dc:"排序"`
}

// ProvincesEditInp 修改/新增字典数据
type ProvincesEditInp struct {
	entity.SysProvinces
}
type ProvincesEditModel struct{}

// ProvincesDeleteInp 删除字典类型
type ProvincesDeleteInp struct {
	Id interface{} `json:"id" v:"required#省市区ID不能为空" dc:"省市区ID"`
}
type ProvincesDeleteModel struct{}

// ProvincesViewInp 获取信息
type ProvincesViewInp struct {
	Id int64
}

type ProvincesViewModel struct {
	entity.SysProvinces
}

// ProvincesListInp 获取列表
type ProvincesListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Title   string
	Content string
}

type ProvincesListModel struct {
	entity.SysProvinces
}

// ProvincesStatusInp 更新状态
type ProvincesStatusInp struct {
	entity.SysProvinces
}
type ProvincesStatusModel struct{}

// ProvincesChildrenListInp 获取省市区下级列表
type ProvincesChildrenListInp struct {
	form.PageReq
	form.StatusReq
	Pid   int64  `json:"pid" dc:"上级ID"`
	Id    int64  `json:"id" dc:"地区ID"`
	Title string `json:"title" dc:"地区名称"`
}

type ProvincesChildrenListModel struct {
	entity.SysProvinces
}

// ProvincesUniqueIdInp 获取省市区下级列表
type ProvincesUniqueIdInp struct {
	OldId int64 `json:"oldId" dc:"原始ID"`
	NewId int64 `json:"newId" dc:"新的ID"`
}

type ProvincesUniqueIdModel struct {
	IsUnique bool `json:"unique" dc:"是否唯一"`
}

type ProvincesSelectInp struct {
	DataType string `json:"dataType" v:"required#数据类型不能为空"  dc:"数据类型"`
	Value    int64  `json:"value" dc:"上级ID"`
}

type ProvincesSelectModel struct {
	List []*ProvincesSelectData `json:"list" dc:"数据列表"`
}

type ProvincesSelectData struct {
	Label  string `json:"label"     description:"地区名称"`
	Value  int64  `json:"value"     description:"地区ID"`
	Level  int    `json:"level"     description:"地区等级"`
	IsLeaf bool   `json:"isLeaf"    description:"是否还有下一级"`
}

// ProvincesCityLabelInp 获取指定城市标签
type ProvincesCityLabelInp struct {
	Id    int64  `json:"oldId" dc:"城市ID"`
	Spilt string `json:"spilt" dc:"分隔符"`
}

type ProvincesCityLabelModel string
