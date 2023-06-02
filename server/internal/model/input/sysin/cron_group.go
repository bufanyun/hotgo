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

// CronGroupMaxSortInp 最大排序
type CronGroupMaxSortInp struct {
	Id int64
}

type CronGroupMaxSortModel struct {
	Sort int
}

// CronGroupEditInp 修改/新增字典数据
type CronGroupEditInp struct {
	entity.SysCronGroup
}

func (in *CronGroupEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}

	if in.Id > 0 && in.Id == in.Pid {
		err = gerror.New("上级分组不能是自己")
		return
	}

	return
}

type CronGroupEditModel struct{}

// CronGroupUpdateFields 修改数据字段过滤
type CronGroupUpdateFields struct {
	Id     int64  `json:"id"        description:"任务分组ID"`
	Pid    int64  `json:"pid"       description:"父类字典类型ID"`
	Name   string `json:"name"      description:"字典类型名称"`
	Type   string `json:"type"      description:"字典类型"`
	Sort   int    `json:"sort"      description:"排序"`
	Remark string `json:"remark"    description:"备注"`
	Status int    `json:"status"    description:"字典类型状态"`
}

// CronGroupInsertFields 新增数据字段过滤
type CronGroupInsertFields struct {
	Pid       int64  `json:"pid"       description:"父类任务分组ID"`
	Name      string `json:"name"      description:"分组名称"`
	IsDefault int    `json:"isDefault" description:"是否默认"`
	Sort      int    `json:"sort"      description:"排序"`
	Remark    string `json:"remark"    description:"备注"`
	Status    int    `json:"status"    description:"分组状态"`
}

// CronGroupDeleteInp 删除字典类型
type CronGroupDeleteInp struct {
	Id interface{}
}
type CronGroupDeleteModel struct{}

// CronGroupViewInp 获取信息
type CronGroupViewInp struct {
	Id int64
}

type CronGroupViewModel struct {
	entity.SysCronGroup
}

// CronGroupListInp 获取列表
type CronGroupListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name string
}

type CronGroupListModel struct {
	entity.SysCronGroup
}

// CronGroupStatusInp 更新状态
type CronGroupStatusInp struct {
	entity.SysCronGroup
}

func (in *CronGroupStatusInp) Filter(ctx context.Context) (err error) {
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

type CronGroupStatusModel struct{}

// CronGroupSelectInp 选项
type CronGroupSelectInp struct {
}

type CronGroupSelectModel struct {
	List []*CronGroupTree `json:"list"`
}

type CronGroupTree struct {
	entity.SysCronGroup
	Disabled bool             `json:"disabled"  dc:"是否禁用"`
	Label    string           `json:"label"     dc:"标签"`
	Value    int64            `json:"value"     dc:"键值"`
	Key      int64            `json:"key"       dc:"键名"`
	Children []*CronGroupTree `json:"children"  dc:"子级"`
}
