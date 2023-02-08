// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/gencodes"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	GenCodes = cGenCodes{}
)

type cGenCodes struct{}

// Delete 删除
func (c *cGenCodes) Delete(ctx context.Context, req *gencodes.DeleteReq) (res *gencodes.DeleteRes, err error) {
	var in sysin.GenCodesDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysGenCodes().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cGenCodes) Edit(ctx context.Context, req *gencodes.EditReq) (res *gencodes.EditRes, err error) {
	var in sysin.GenCodesEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	data, err := service.SysGenCodes().Edit(ctx, in)
	if err != nil {
		return
	}

	res = new(gencodes.EditRes)
	res.GenCodesEditModel = data
	return
}

// MaxSort 最大排序
func (c *cGenCodes) MaxSort(ctx context.Context, req *gencodes.MaxSortReq) (res *gencodes.MaxSortRes, err error) {
	data, err := service.SysGenCodes().MaxSort(ctx, sysin.GenCodesMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(gencodes.MaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cGenCodes) View(ctx context.Context, req *gencodes.ViewReq) (res *gencodes.ViewRes, err error) {
	data, err := service.SysGenCodes().View(ctx, sysin.GenCodesViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(gencodes.ViewRes)
	res.GenCodesViewModel = data
	return
}

// List 查看列表
func (c *cGenCodes) List(ctx context.Context, req *gencodes.ListReq) (res *gencodes.ListRes, err error) {
	var in sysin.GenCodesListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysGenCodes().List(ctx, in)
	if err != nil {
		return
	}

	res = new(gencodes.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新部门状态
func (c *cGenCodes) Status(ctx context.Context, req *gencodes.StatusReq) (res *gencodes.StatusRes, err error) {
	var in sysin.GenCodesStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysGenCodes().Status(ctx, in)
	return
}

// Selects 获取指定信息
func (c *cGenCodes) Selects(ctx context.Context, req *gencodes.SelectsReq) (res *gencodes.SelectsRes, err error) {
	data, err := service.SysGenCodes().Selects(ctx, sysin.GenCodesSelectsInp{})
	if err != nil {
		return
	}

	res = new(gencodes.SelectsRes)
	res.GenCodesSelectsModel = data
	return
}

// TableSelect 数据库表选项
func (c *cGenCodes) TableSelect(ctx context.Context, req *gencodes.TableSelectReq) (res *gencodes.TableSelectRes, err error) {
	data, err := service.SysGenCodes().TableSelect(ctx, sysin.GenCodesTableSelectInp{Name: req.Name})
	if err != nil {
		return
	}

	res = (*gencodes.TableSelectRes)(&data)
	return
}

// ColumnSelect 表字段选项
func (c *cGenCodes) ColumnSelect(ctx context.Context, req *gencodes.ColumnSelectReq) (res *gencodes.ColumnSelectRes, err error) {
	data, err := service.SysGenCodes().ColumnSelect(ctx, sysin.GenCodesColumnSelectInp{Name: req.Name, Table: req.Table})
	if err != nil {
		return
	}

	res = (*gencodes.ColumnSelectRes)(&data)
	return
}

// ColumnList 表字段列表
func (c *cGenCodes) ColumnList(ctx context.Context, req *gencodes.ColumnListReq) (res *gencodes.ColumnListRes, err error) {
	var in sysin.GenCodesColumnListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	data, err := service.SysGenCodes().ColumnList(ctx, in)
	if err != nil {
		return
	}

	res = (*gencodes.ColumnListRes)(&data)
	return
}

// Preview 生成预览
func (c *cGenCodes) Preview(ctx context.Context, req *gencodes.PreviewReq) (res *gencodes.PreviewRes, err error) {
	var in sysin.GenCodesPreviewInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	data, err := service.SysGenCodes().Preview(ctx, in)
	if err != nil {
		return
	}

	res = new(gencodes.PreviewRes)
	res.GenCodesPreviewModel = data
	return
}

// Build 生成预览
func (c *cGenCodes) Build(ctx context.Context, req *gencodes.BuildReq) (res *gencodes.BuildRes, err error) {
	var in sysin.GenCodesBuildInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysGenCodes().Build(ctx, in)
	return
}
