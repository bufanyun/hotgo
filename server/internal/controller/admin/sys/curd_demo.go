// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.1.4
// @AutoGenerate Date 2023-02-20 16:41:58
//
package sys

import (
	"context"
	"hotgo/api/admin/curddemo"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/util/gconv"
)

var (
	CurdDemo = cCurdDemo{}
)

type cCurdDemo struct{}

// List 查看生成演示列表
func (c *cCurdDemo) List(ctx context.Context, req *curddemo.ListReq) (res *curddemo.ListRes, err error) {
	var in sysin.CurdDemoListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysCurdDemo().List(ctx, in)
	if err != nil {
		return
	}

	res = new(curddemo.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Export 导出生成演示列表
func (c *cCurdDemo) Export(ctx context.Context, req *curddemo.ExportReq) (res *curddemo.ExportRes, err error) {
	var in sysin.CurdDemoListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysCurdDemo().Export(ctx, in)
	return
}

// Edit 更新生成演示
func (c *cCurdDemo) Edit(ctx context.Context, req *curddemo.EditReq) (res *curddemo.EditRes, err error) {
	var in sysin.CurdDemoEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysCurdDemo().Edit(ctx, in)
	return
}

// MaxSort 获取生成演示最大排序
func (c *cCurdDemo) MaxSort(ctx context.Context, req *curddemo.MaxSortReq) (res *curddemo.MaxSortRes, err error) {
	data, err := service.SysCurdDemo().MaxSort(ctx, sysin.CurdDemoMaxSortInp{})
	if err != nil {
		return
	}

	res = new(curddemo.MaxSortRes)
	res.CurdDemoMaxSortModel = data
	return
}

// View 获取指定生成演示信息
func (c *cCurdDemo) View(ctx context.Context, req *curddemo.ViewReq) (res *curddemo.ViewRes, err error) {
	var in sysin.CurdDemoViewInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	data, err := service.SysCurdDemo().View(ctx, in)
	if err != nil {
		return
	}

	res = new(curddemo.ViewRes)
	res.CurdDemoViewModel = data
	return
}

// Delete 删除生成演示
func (c *cCurdDemo) Delete(ctx context.Context, req *curddemo.DeleteReq) (res *curddemo.DeleteRes, err error) {
	var in sysin.CurdDemoDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysCurdDemo().Delete(ctx, in)
	return
}

// Status 更新生成演示状态
func (c *cCurdDemo) Status(ctx context.Context, req *curddemo.StatusReq) (res *curddemo.StatusRes, err error) {
	var in sysin.CurdDemoStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysCurdDemo().Status(ctx, in)
	return
}

// Switch 更新生成演示开关状态
func (c *cCurdDemo) Switch(ctx context.Context, req *curddemo.SwitchReq) (res *curddemo.SwitchRes, err error) {
	var in sysin.CurdDemoSwitchInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysCurdDemo().Switch(ctx, in)
	return
}
