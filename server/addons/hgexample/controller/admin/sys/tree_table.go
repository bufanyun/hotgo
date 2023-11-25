// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/addons/hgexample/api/admin/treetable"
	"hotgo/addons/hgexample/service"
)

var (
	TreeTable = cTreeTable{}
)

type cTreeTable struct{}

// List 查看列表
func (c *cTreeTable) List(ctx context.Context, req *treetable.ListReq) (res *treetable.ListRes, err error) {
	list, totalCount, err := service.SysTreeTable().List(ctx, &req.TreeTableListInp)
	if err != nil {
		return
	}

	res = new(treetable.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出列表
func (c *cTreeTable) Export(ctx context.Context, req *treetable.ExportReq) (res *treetable.ExportRes, err error) {
	err = service.SysTable().Export(ctx, &req.TableListInp)
	return
}

// Edit 更新
func (c *cTreeTable) Edit(ctx context.Context, req *treetable.EditReq) (res *treetable.EditRes, err error) {
	err = service.SysTreeTable().Edit(ctx, &req.TableEditInp)
	return
}

// MaxSort 最大排序
func (c *cTreeTable) MaxSort(ctx context.Context, req *treetable.MaxSortReq) (res *treetable.MaxSortRes, err error) {
	data, err := service.SysTable().MaxSort(ctx, &req.TableMaxSortInp)
	if err != nil {
		return
	}

	res = new(treetable.MaxSortRes)
	res.TableMaxSortModel = data
	return
}

// View 获取指定信息
func (c *cTreeTable) View(ctx context.Context, req *treetable.ViewReq) (res *treetable.ViewRes, err error) {
	data, err := service.SysTable().View(ctx, &req.TableViewInp)
	if err != nil {
		return
	}

	res = new(treetable.ViewRes)
	res.TableViewModel = data
	return
}

// Delete 删除
func (c *cTreeTable) Delete(ctx context.Context, req *treetable.DeleteReq) (res *treetable.DeleteRes, err error) {
	err = service.SysTreeTable().Delete(ctx, &req.TableDeleteInp)
	return
}

// Status 更新状态
func (c *cTreeTable) Status(ctx context.Context, req *treetable.StatusReq) (res *treetable.StatusRes, err error) {
	err = service.SysTable().Status(ctx, &req.TableStatusInp)
	return
}

// Switch 更新开关状态
func (c *cTreeTable) Switch(ctx context.Context, req *treetable.SwitchReq) (res *treetable.SwitchRes, err error) {
	err = service.SysTable().Switch(ctx, &req.TableSwitchInp)
	return
}

// Select 树形选项
func (c *cTreeTable) Select(ctx context.Context, _ *treetable.SelectReq) (res *treetable.SelectRes, err error) {
	res = new(treetable.SelectRes)
	res.List, err = service.SysTreeTable().Select(ctx)
	return
}
