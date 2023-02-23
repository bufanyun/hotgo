// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/addons/hgexample/api/admin/table"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
)

var (
	Table = cTable{}
)

type cTable struct{}

// List 查看列表
func (c *cTable) List(ctx context.Context, req *table.ListReq) (res *table.ListRes, err error) {
	var in sysin.TableListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysTable().List(ctx, in)
	if err != nil {
		return
	}

	res = new(table.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Export 导出列表
func (c *cTable) Export(ctx context.Context, req *table.ExportReq) (res *table.ExportRes, err error) {
	var in sysin.TableListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysTable().Export(ctx, in)
	return
}

// Edit 更新
func (c *cTable) Edit(ctx context.Context, req *table.EditReq) (res *table.EditRes, err error) {
	var in sysin.TableEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysTable().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cTable) MaxSort(ctx context.Context, req *table.MaxSortReq) (res *table.MaxSortRes, err error) {
	data, err := service.SysTable().MaxSort(ctx, sysin.TableMaxSortInp{})
	if err != nil {
		return
	}

	res = new(table.MaxSortRes)
	res.TableMaxSortModel = data
	return
}

// View 获取指定信息
func (c *cTable) View(ctx context.Context, req *table.ViewReq) (res *table.ViewRes, err error) {
	var in sysin.TableViewInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	data, err := service.SysTable().View(ctx, in)
	if err != nil {
		return
	}

	res = new(table.ViewRes)
	res.TableViewModel = data
	return
}

// Delete 删除
func (c *cTable) Delete(ctx context.Context, req *table.DeleteReq) (res *table.DeleteRes, err error) {
	var in sysin.TableDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysTable().Delete(ctx, in)
	return
}

// Status 更新状态
func (c *cTable) Status(ctx context.Context, req *table.StatusReq) (res *table.StatusRes, err error) {
	var in sysin.TableStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysTable().Status(ctx, in)
	return
}

// Switch 更新开关状态
func (c *cTable) Switch(ctx context.Context, req *table.SwitchReq) (res *table.SwitchRes, err error) {
	var in sysin.TableSwitchInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysTable().Switch(ctx, in)
	return
}
