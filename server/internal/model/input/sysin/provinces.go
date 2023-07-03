// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
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

func (in *ProvincesEditInp) Filter(ctx context.Context) (err error) {
	if in.Title == "" {
		err = gerror.New("标题不能为空")
		return
	}

	if in.Id <= 0 {
		err = gerror.New("地区Id必须大于0")
		return
	}

	return
}

type ProvincesEditModel struct{}

// ProvincesUpdateFields 修改数据字段过滤
type ProvincesUpdateFields struct {
	Id     int64  `json:"id"        description:"省市区ID"`
	Title  string `json:"title"     description:"栏目名称"`
	Pinyin string `json:"pinyin"    description:"拼音"`
	Lng    string `json:"lng"       description:"经度"`
	Lat    string `json:"lat"       description:"纬度"`
	Pid    int64  `json:"pid"       description:"父栏目"`
	Level  int    `json:"level"     description:"关系树等级"`
	Tree   string `json:"tree"      description:"关系"`
	Sort   int    `json:"sort"      description:"排序"`
	Status int    `json:"status"    description:"状态"`
}

// ProvincesInsertFields 新增数据字段过滤
type ProvincesInsertFields struct {
	Title  string `json:"title"     description:"栏目名称"`
	Pinyin string `json:"pinyin"    description:"拼音"`
	Lng    string `json:"lng"       description:"经度"`
	Lat    string `json:"lat"       description:"纬度"`
	Pid    int64  `json:"pid"       description:"父栏目"`
	Level  int    `json:"level"     description:"关系树等级"`
	Tree   string `json:"tree"      description:"关系"`
	Sort   int    `json:"sort"      description:"排序"`
	Status int    `json:"status"    description:"状态"`
}

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

func (in *ProvincesStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
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

// ProvincesTree 树
type ProvincesTree struct {
	entity.SysProvinces
	Key      int64            `json:"key"       dc:"key"`
	Label    string           `json:"label"     dc:"标签"`
	Value    int64            `json:"value"     dc:"键值"`
	Children []*ProvincesTree `json:"children"`
}
