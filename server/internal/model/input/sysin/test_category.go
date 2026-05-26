// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.15.7
package sysin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestCategoryUpdateFields 修改测试分类字段过滤
type TestCategoryUpdateFields struct {
	Name        string `json:"name"        dc:"分类名称"`
	ShortName   string `json:"shortName"   dc:"简称"`
	Description string `json:"description" dc:"描述"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态"`
	Remark      string `json:"remark"      dc:"备注"`
}

// TestCategoryInsertFields 新增测试分类字段过滤
type TestCategoryInsertFields struct {
	Name        string `json:"name"        dc:"分类名称"`
	ShortName   string `json:"shortName"   dc:"简称"`
	Description string `json:"description" dc:"描述"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态"`
	Remark      string `json:"remark"      dc:"备注"`
}

// TestCategoryEditInp 修改/新增测试分类
type TestCategoryEditInp struct {
	entity.TestCategory
}

func (in *TestCategoryEditInp) Filter(ctx context.Context) (err error) {
	// 验证分类名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("分类名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type TestCategoryEditModel struct{}

// TestCategoryDeleteInp 删除测试分类
type TestCategoryDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *TestCategoryDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TestCategoryDeleteModel struct{}

// TestCategoryViewInp 获取指定测试分类信息
type TestCategoryViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *TestCategoryViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TestCategoryViewModel struct {
	entity.TestCategory
}

// TestCategoryListInp 获取测试分类列表
type TestCategoryListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	Name      string        `json:"name"      dc:"分类名称"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TestCategoryListInp) Filter(ctx context.Context) (err error) {
	return
}

type TestCategoryListModel struct {
	Id          int64       `json:"id"          dc:"分类ID"`
	Name        string      `json:"name"        dc:"分类名称"`
	ShortName   string      `json:"shortName"   dc:"简称"`
	Description string      `json:"description" dc:"描述"`
	Status      int         `json:"status"      dc:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
}

// TestCategoryMaxSortInp 获取测试分类最大排序
type TestCategoryMaxSortInp struct{}

func (in *TestCategoryMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type TestCategoryMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// TestCategoryStatusInp 更新测试分类状态
type TestCategoryStatusInp struct {
	Id     int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *TestCategoryStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("分类ID不能为空")
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

type TestCategoryStatusModel struct{}
