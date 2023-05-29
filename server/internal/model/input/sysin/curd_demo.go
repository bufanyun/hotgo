// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.7.3
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

// CurdDemoUpdateFields 修改生成演示字段过滤
type CurdDemoUpdateFields struct {
	CategoryId  int64  `json:"categoryId"  dc:"分类ID"`
	Title       string `json:"title"       dc:"标题"`
	Description string `json:"description" dc:"描述"`
	Content     string `json:"content"     dc:"内容"`
	Image       string `json:"image"       dc:"单图"`
	Attachfile  string `json:"attachfile"  dc:"附件"`
	CityId      int64  `json:"cityId"      dc:"所在城市"`
	Switch      int    `json:"switch"      dc:"显示开关"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态"`
	UpdatedBy   int64  `json:"updatedBy"   dc:"更新者"`
}

// CurdDemoInsertFields 新增生成演示字段过滤
type CurdDemoInsertFields struct {
	CategoryId  int64  `json:"categoryId"  dc:"分类ID"`
	Title       string `json:"title"       dc:"标题"`
	Description string `json:"description" dc:"描述"`
	Content     string `json:"content"     dc:"内容"`
	Image       string `json:"image"       dc:"单图"`
	Attachfile  string `json:"attachfile"  dc:"附件"`
	CityId      int64  `json:"cityId"      dc:"所在城市"`
	Switch      int    `json:"switch"      dc:"显示开关"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态"`
	CreatedBy   int64  `json:"createdBy"   dc:"创建者"`
}

// CurdDemoEditInp 修改/新增生成演示
type CurdDemoEditInp struct {
	entity.SysGenCurdDemo
}

func (in *CurdDemoEditInp) Filter(ctx context.Context) (err error) {
	// 验证分类ID
	if err := g.Validator().Rules("required").Data(in.CategoryId).Messages("分类ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证描述
	if err := g.Validator().Rules("required").Data(in.Description).Messages("描述不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证内容
	if err := g.Validator().Rules("required").Data(in.Content).Messages("内容不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type CurdDemoEditModel struct{}

// CurdDemoDeleteInp 删除生成演示
type CurdDemoDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *CurdDemoDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoDeleteModel struct{}

// CurdDemoViewInp 获取指定生成演示信息
type CurdDemoViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *CurdDemoViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoViewModel struct {
	entity.SysGenCurdDemo
}

// CurdDemoListInp 获取生成演示列表
type CurdDemoListInp struct {
	form.PageReq
	Id               int64         `json:"id"               dc:"ID"`
	Status           int           `json:"status"           dc:"状态"`
	CreatedAt        []*gtime.Time `json:"createdAt"        dc:"创建时间"`
	TestCategoryName string        `json:"testCategoryName" dc:"分类名称"`
}

func (in *CurdDemoListInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoListModel struct {
	Id               int64       `json:"id"               dc:"ID"`
	CategoryId       int64       `json:"categoryId"       dc:"分类ID"`
	Title            string      `json:"title"            dc:"标题"`
	Description      string      `json:"description"      dc:"描述"`
	Image            string      `json:"image"            dc:"单图"`
	Attachfile       string      `json:"attachfile"       dc:"附件"`
	CityId           int64       `json:"cityId"           dc:"所在城市"`
	Switch           int         `json:"switch"           dc:"显示开关"`
	Sort             int         `json:"sort"             dc:"排序"`
	Status           int         `json:"status"           dc:"状态"`
	CreatedBy        int64       `json:"createdBy"        dc:"创建者"`
	UpdatedBy        int64       `json:"updatedBy"        dc:"更新者"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        dc:"修改时间"`
	TestCategoryName string      `json:"testCategoryName" dc:"分类名称"`
}

// CurdDemoExportModel 导出生成演示
type CurdDemoExportModel struct {
	Id               int64       `json:"id"               dc:"ID"`
	CategoryId       int64       `json:"categoryId"       dc:"分类ID"`
	Title            string      `json:"title"            dc:"标题"`
	Description      string      `json:"description"      dc:"描述"`
	Image            string      `json:"image"            dc:"单图"`
	Attachfile       string      `json:"attachfile"       dc:"附件"`
	CityId           int64       `json:"cityId"           dc:"所在城市"`
	Switch           int         `json:"switch"           dc:"显示开关"`
	Sort             int         `json:"sort"             dc:"排序"`
	Status           int         `json:"status"           dc:"状态"`
	CreatedBy        int64       `json:"createdBy"        dc:"创建者"`
	UpdatedBy        int64       `json:"updatedBy"        dc:"更新者"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        dc:"修改时间"`
	TestCategoryName string      `json:"testCategoryName" dc:"分类名称"`
}

// CurdDemoMaxSortInp 获取生成演示最大排序
type CurdDemoMaxSortInp struct{}

func (in *CurdDemoMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// CurdDemoStatusInp 更新生成演示状态
type CurdDemoStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *CurdDemoStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSliceInt(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type CurdDemoStatusModel struct{}

// CurdDemoSwitchInp 更新生成演示开关状态
type CurdDemoSwitchInp struct {
	form.SwitchReq
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *CurdDemoSwitchInp) Filter(ctx context.Context) (err error) {
	return
}

type CurdDemoSwitchModel struct{}
