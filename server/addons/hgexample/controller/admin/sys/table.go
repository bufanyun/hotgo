// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/addons/hgexample/api/admin/table"
	"hotgo/addons/hgexample/service"
)

var (
	Table = cTable{}
)

type cTable struct{}

// List 查看列表
func (c *cTable) List(ctx context.Context, req *table.ListReq) (res *table.ListRes, err error) {
	list, totalCount, err := service.SysTable().List(ctx, &req.TableListInp)
	if err != nil {
		return
	}

	res = new(table.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出列表
func (c *cTable) Export(ctx context.Context, req *table.ExportReq) (res *table.ExportRes, err error) {
	err = service.SysTable().Export(ctx, &req.TableListInp)
	return
}

// Edit 更新
func (c *cTable) Edit(ctx context.Context, req *table.EditReq) (res *table.EditRes, err error) {
	err = service.SysTable().Edit(ctx, &req.TableEditInp)
	return
}

// MaxSort 最大排序
func (c *cTable) MaxSort(ctx context.Context, req *table.MaxSortReq) (res *table.MaxSortRes, err error) {
	data, err := service.SysTable().MaxSort(ctx, &req.TableMaxSortInp)
	if err != nil {
		return
	}

	res = new(table.MaxSortRes)
	res.TableMaxSortModel = data
	return
}

// View 获取指定信息
func (c *cTable) View(ctx context.Context, req *table.ViewReq) (res *table.ViewRes, err error) {
	data, err := service.SysTable().View(ctx, &req.TableViewInp)
	if err != nil {
		return
	}

	res = new(table.ViewRes)
	res.TableViewModel = data
	return
}

// Delete 删除
func (c *cTable) Delete(ctx context.Context, req *table.DeleteReq) (res *table.DeleteRes, err error) {
	err = service.SysTable().Delete(ctx, &req.TableDeleteInp)
	return
}

// Status 更新状态
func (c *cTable) Status(ctx context.Context, req *table.StatusReq) (res *table.StatusRes, err error) {
	err = service.SysTable().Status(ctx, &req.TableStatusInp)
	return
}

// Switch 更新开关状态
func (c *cTable) Switch(ctx context.Context, req *table.SwitchReq) (res *table.SwitchRes, err error) {
	err = service.SysTable().Switch(ctx, &req.TableSwitchInp)
	return
}
