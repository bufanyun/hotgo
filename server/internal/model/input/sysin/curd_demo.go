// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.1.0
// @AutoGenerate Date 2023-01-18 15:19:42
//
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// CurdDemoEditInp 修改/新增生成演示
type CurdDemoEditInp struct {
	entity.Test
}

func (in *CurdDemoEditInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoEditModel struct{}

// CurdDemoDeleteInp 删除生成演示
type CurdDemoDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type CurdDemoDeleteModel struct{}

// CurdDemoViewInp 获取指定生成演示信息
type CurdDemoViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type CurdDemoViewModel struct {
	entity.Test
}

// CurdDemoListInp 获取生成演示列表
type CurdDemoListInp struct {
	form.PageReq
	Id               int64         `json:"id"               dc:"ID"`
	Status           int           `json:"status"           dc:"状态"`
	CreatedAt        []*gtime.Time `json:"createdAt"        dc:"创建时间"`
	TestCategoryName string        `json:"testCategoryName" dc:"分类名称"`
}

type CurdDemoListModel struct {
	Id               int64       `json:"id"               dc:"ID"`
	CategoryId       int64       `json:"categoryId"       dc:"分类ID"`
	Title            string      `json:"title"            dc:"标题"`
	Description      string      `json:"description"      dc:"描述"`
	Image            string      `json:"image"            dc:"单图"`
	Attachfile       string      `json:"attachfile"       dc:"附件"`
	Switch           int         `json:"switch"           dc:"显示开关"`
	Sort             int         `json:"sort"             dc:"排序"`
	Status           int         `json:"status"           dc:"状态"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	TestCategoryName string      `json:"testCategoryName" dc:"分类名称"`
}

func (in *CurdDemoListInp) Filter(ctx context.Context) (err error) {
	return
}

// CurdDemoExportModel 导出生成演示
type CurdDemoExportModel struct {
	Id               int64       `json:"id"               dc:"ID"`
	CategoryId       int64       `json:"categoryId"       dc:"分类ID"`
	Title            string      `json:"title"            dc:"标题"`
	Description      string      `json:"description"      dc:"描述"`
	Image            string      `json:"image"            dc:"单图"`
	Attachfile       string      `json:"attachfile"       dc:"附件"`
	Switch           int         `json:"switch"           dc:"显示开关"`
	Sort             int         `json:"sort"             dc:"排序"`
	Status           int         `json:"status"           dc:"状态"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	TestCategoryName string      `json:"testCategoryName" dc:"分类名称"`
}

// CurdDemoMaxSortInp 获取生成演示最大排序
type CurdDemoMaxSortInp struct{}

type CurdDemoMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// CurdDemoStatusInp 更新生成演示状态
type CurdDemoStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

type CurdDemoStatusModel struct{}

// CurdDemoSwitchInp 更新生成演示开关状态
type CurdDemoSwitchInp struct {
	form.SwitchReq
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type CurdDemoSwitchModel struct{}
