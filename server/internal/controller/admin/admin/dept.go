// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"hotgo/api/admin/dept"
	"hotgo/internal/service"
)

var (
	Dept = cDept{}
)

type cDept struct{}

// Delete 删除
func (c *cDept) Delete(ctx context.Context, req *dept.DeleteReq) (res *dept.DeleteRes, err error) {
	err = service.AdminDept().Delete(ctx, &req.DeptDeleteInp)
	return
}

// Edit 更新
func (c *cDept) Edit(ctx context.Context, req *dept.EditReq) (res *dept.EditRes, err error) {
	err = service.AdminDept().Edit(ctx, &req.DeptEditInp)
	return
}

// MaxSort 最大排序
func (c *cDept) MaxSort(ctx context.Context, req *dept.MaxSortReq) (res *dept.MaxSortRes, err error) {
	res = new(dept.MaxSortRes)
	res.DeptMaxSortModel, err = service.AdminDept().MaxSort(ctx, &req.DeptMaxSortInp)
	return
}

// View 获取指定信息
func (c *cDept) View(ctx context.Context, req *dept.ViewReq) (res *dept.ViewRes, err error) {
	res = new(dept.ViewRes)
	res.DeptViewModel, err = service.AdminDept().View(ctx, &req.DeptViewInp)
	return
}

// List 查看列表
func (c *cDept) List(ctx context.Context, req *dept.ListReq) (res *dept.ListRes, err error) {
	data, err := service.AdminDept().List(ctx, &req.DeptListInp)
	if err != nil || data == nil {
		return
	}

	res = (*dept.ListRes)(&data)
	return
}

// Option 获取部门选项树
func (c *cDept) Option(ctx context.Context, req *dept.OptionReq) (res *dept.OptionRes, err error) {
	list, totalCount, err := service.AdminDept().Option(ctx, &req.DeptOptionInp)
	if err != nil {
		return
	}

	res = new(dept.OptionRes)
	res.DeptOptionModel = list
	res.PageRes.Pack(req, totalCount)
	return
}

// TreeOption 获取部门管理关系树选项
func (c *cDept) TreeOption(ctx context.Context, req *dept.TreeOptionReq) (res *dept.TreeOptionRes, err error) {
	data, err := service.AdminDept().TreeOption(ctx)
	res = (*dept.TreeOptionRes)(&data)
	return
}
