// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package sys

import (
	"context"
	"hotgo/api/admin/curddemo"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	CurdDemo = cCurdDemo{}
)

type cCurdDemo struct{}

// List 查看CURD列表列表
func (c *cCurdDemo) List(ctx context.Context, req *curddemo.ListReq) (res *curddemo.ListRes, err error) {
	list, totalCount, err := service.SysCurdDemo().List(ctx, &req.CurdDemoListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.CurdDemoListModel{}
	}

	res = new(curddemo.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出CURD列表列表
func (c *cCurdDemo) Export(ctx context.Context, req *curddemo.ExportReq) (res *curddemo.ExportRes, err error) {
	err = service.SysCurdDemo().Export(ctx, &req.CurdDemoListInp)
	return
}

// Edit 更新CURD列表
func (c *cCurdDemo) Edit(ctx context.Context, req *curddemo.EditReq) (res *curddemo.EditRes, err error) {
	err = service.SysCurdDemo().Edit(ctx, &req.CurdDemoEditInp)
	return
}

// MaxSort 获取CURD列表最大排序
func (c *cCurdDemo) MaxSort(ctx context.Context, req *curddemo.MaxSortReq) (res *curddemo.MaxSortRes, err error) {
	data, err := service.SysCurdDemo().MaxSort(ctx, &req.CurdDemoMaxSortInp)
	if err != nil {
		return
	}

	res = new(curddemo.MaxSortRes)
	res.CurdDemoMaxSortModel = data
	return
}

// View 获取指定CURD列表信息
func (c *cCurdDemo) View(ctx context.Context, req *curddemo.ViewReq) (res *curddemo.ViewRes, err error) {
	data, err := service.SysCurdDemo().View(ctx, &req.CurdDemoViewInp)
	if err != nil {
		return
	}

	res = new(curddemo.ViewRes)
	res.CurdDemoViewModel = data
	return
}

// Delete 删除CURD列表
func (c *cCurdDemo) Delete(ctx context.Context, req *curddemo.DeleteReq) (res *curddemo.DeleteRes, err error) {
	err = service.SysCurdDemo().Delete(ctx, &req.CurdDemoDeleteInp)
	return
}

// Status 更新CURD列表状态
func (c *cCurdDemo) Status(ctx context.Context, req *curddemo.StatusReq) (res *curddemo.StatusRes, err error) {
	err = service.SysCurdDemo().Status(ctx, &req.CurdDemoStatusInp)
	return
}

// Switch 更新CURD列表开关状态
func (c *cCurdDemo) Switch(ctx context.Context, req *curddemo.SwitchReq) (res *curddemo.SwitchRes, err error) {
	err = service.SysCurdDemo().Switch(ctx, &req.CurdDemoSwitchInp)
	return
}