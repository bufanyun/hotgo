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
		return nil, err
	}
	if err = service.SysGenCodes().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cGenCodes) Edit(ctx context.Context, req *gencodes.EditReq) (res *gencodes.EditRes, err error) {

	var in sysin.GenCodesEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	data, err := service.SysGenCodes().Edit(ctx, in)
	if err != nil {
		return nil, err
	}
	res = new(gencodes.EditRes)
	res.GenCodesEditModel = data
	return res, nil
}

// MaxSort 最大排序
func (c *cGenCodes) MaxSort(ctx context.Context, req *gencodes.MaxSortReq) (*gencodes.MaxSortRes, error) {

	data, err := service.SysGenCodes().MaxSort(ctx, sysin.GenCodesMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res gencodes.MaxSortRes
	res.Sort = data.Sort
	return &res, nil
}

// View 获取指定信息
func (c *cGenCodes) View(ctx context.Context, req *gencodes.ViewReq) (*gencodes.ViewRes, error) {

	data, err := service.SysGenCodes().View(ctx, sysin.GenCodesViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res gencodes.ViewRes
	res.GenCodesViewModel = data
	return &res, nil
}

// List 查看列表
func (c *cGenCodes) List(ctx context.Context, req *gencodes.ListReq) (*gencodes.ListRes, error) {

	var (
		in  sysin.GenCodesListInp
		res gencodes.ListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.SysGenCodes().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// Status 更新部门状态
func (c *cGenCodes) Status(ctx context.Context, req *gencodes.StatusReq) (res *gencodes.StatusRes, err error) {

	var in sysin.GenCodesStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysGenCodes().Status(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// Selects 获取指定信息
func (c *cGenCodes) Selects(ctx context.Context, req *gencodes.SelectsReq) (*gencodes.SelectsRes, error) {
	data, err := service.SysGenCodes().Selects(ctx, sysin.GenCodesSelectsInp{})
	if err != nil {
		return nil, err
	}

	var res gencodes.SelectsRes
	res.GenCodesSelectsModel = data
	return &res, nil
}

// TableSelect 数据库表选项
func (c *cGenCodes) TableSelect(ctx context.Context, req *gencodes.TableSelectReq) (*gencodes.TableSelectRes, error) {
	data, err := service.SysGenCodes().TableSelect(ctx, sysin.GenCodesTableSelectInp{Name: req.Name})
	if err != nil {
		return nil, err
	}

	var res gencodes.TableSelectRes
	res = data
	return &res, nil
}

// ColumnSelect 表字段选项
func (c *cGenCodes) ColumnSelect(ctx context.Context, req *gencodes.ColumnSelectReq) (*gencodes.ColumnSelectRes, error) {
	data, err := service.SysGenCodes().ColumnSelect(ctx, sysin.GenCodesColumnSelectInp{Name: req.Name, Table: req.Table})
	if err != nil {
		return nil, err
	}

	var res gencodes.ColumnSelectRes
	res = data
	return &res, nil
}

// ColumnList 表字段列表
func (c *cGenCodes) ColumnList(ctx context.Context, req *gencodes.ColumnListReq) (*gencodes.ColumnListRes, error) {
	var (
		in  sysin.GenCodesColumnListInp
		err error
	)
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	data, err := service.SysGenCodes().ColumnList(ctx, in)
	if err != nil {
		return nil, err
	}

	var res gencodes.ColumnListRes
	res = data
	return &res, nil
}

// Preview 生成预览
func (c *cGenCodes) Preview(ctx context.Context, req *gencodes.PreviewReq) (*gencodes.PreviewRes, error) {
	var (
		in  sysin.GenCodesPreviewInp
		err error
	)
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	data, err := service.SysGenCodes().Preview(ctx, in)
	if err != nil {
		return nil, err
	}

	res := new(gencodes.PreviewRes)
	res.GenCodesPreviewModel = data
	return res, nil
}

// Build 生成预览
func (c *cGenCodes) Build(ctx context.Context, req *gencodes.BuildReq) (*gencodes.BuildRes, error) {
	var (
		in  sysin.GenCodesBuildInp
		err error
	)
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = service.SysGenCodes().Build(ctx, in); err != nil {
		return nil, err
	}

	return nil, nil
}
