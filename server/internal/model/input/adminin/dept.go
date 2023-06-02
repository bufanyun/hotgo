// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
)

// DeptMaxSortInp 最大排序
type DeptMaxSortInp struct {
	Id int64
}

type DeptMaxSortModel struct {
	Sort int
}

// DeptEditInp 修改/新增部门数据
type DeptEditInp struct {
	entity.AdminDept
}

func (in *DeptEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}

	if in.Id > 0 && in.Id == in.Pid {
		err = gerror.New("上级部门不能是自己")
		return
	}

	return
}

type DeptEditModel struct{}

// DeptUpdateFields 修改数据字段过滤
type DeptUpdateFields struct {
	Id     int64  `json:"id"        description:"部门ID"`
	Pid    int64  `json:"pid"       description:"父部门ID"`
	Name   string `json:"name"      description:"部门名称"`
	Code   string `json:"code"      description:"部门编码"`
	Type   string `json:"type"      description:"部门类型"`
	Leader string `json:"leader"    description:"负责人"`
	Phone  string `json:"phone"     description:"联系电话"`
	Email  string `json:"email"     description:"邮箱"`
	Level  int    `json:"level"     description:"关系树等级"`
	Tree   string `json:"tree"      description:"关系树"`
	Sort   int    `json:"sort"      description:"排序"`
	Status int    `json:"status"    description:"部门状态"`
}

// DeptInsertFields 新增数据字段过滤
type DeptInsertFields struct {
	Pid    int64  `json:"pid"       description:"父部门ID"`
	Name   string `json:"name"      description:"部门名称"`
	Code   string `json:"code"      description:"部门编码"`
	Type   string `json:"type"      description:"部门类型"`
	Leader string `json:"leader"    description:"负责人"`
	Phone  string `json:"phone"     description:"联系电话"`
	Email  string `json:"email"     description:"邮箱"`
	Level  int    `json:"level"     description:"关系树等级"`
	Tree   string `json:"tree"      description:"关系树"`
	Sort   int    `json:"sort"      description:"排序"`
	Status int    `json:"status"    description:"部门状态"`
}

// DeptDeleteInp 删除部门类型
type DeptDeleteInp struct {
	Id interface{}
}
type DeptDeleteModel struct{}

// DeptViewInp 获取信息
type DeptViewInp struct {
	Id int64
}

type DeptViewModel struct {
	entity.AdminDept
}

// DeptListInp 获取列表
type DeptListInp struct {
	Name string
	Code string
}

// DeptTree 树
type DeptTree struct {
	entity.AdminDept
	Label    string      `json:"label"     dc:"标签"`
	Value    int64       `json:"value"     dc:"键值"`
	Children []*DeptTree `json:"children"`
}

type DeptListModel struct {
	List []*DeptTree `json:"list"`
}

// DeptStatusInp 更新部门状态
type DeptStatusInp struct {
	entity.AdminDept
}

func (in *DeptStatusInp) Filter(ctx context.Context) (err error) {
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

type DeptStatusModel struct{}

type DeptOptionInp struct {
	form.PageReq
}

type DeptOptionModel struct {
	List []*DeptTree `json:"list"`
}
