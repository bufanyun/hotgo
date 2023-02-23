// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/dept"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
)

var (
	Dept = cDept{}
)

type cDept struct{}

// NameUnique 名称是否唯一
func (c *cDept) NameUnique(ctx context.Context, req *dept.NameUniqueReq) (res *dept.NameUniqueRes, err error) {
	data, err := service.AdminDept().NameUnique(ctx, adminin.DeptNameUniqueInp{Id: req.Id, Name: req.Name})
	if err != nil {
		return
	}

	res = new(dept.NameUniqueRes)
	res.IsUnique = data.IsUnique
	return
}

// Delete 删除
func (c *cDept) Delete(ctx context.Context, req *dept.DeleteReq) (res *dept.DeleteRes, err error) {
	var in adminin.DeptDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminDept().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cDept) Edit(ctx context.Context, req *dept.EditReq) (res *dept.EditRes, err error) {
	var in adminin.DeptEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	err = service.AdminDept().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cDept) MaxSort(ctx context.Context, req *dept.MaxSortReq) (res *dept.MaxSortRes, err error) {
	data, err := service.AdminDept().MaxSort(ctx, adminin.DeptMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(dept.MaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cDept) View(ctx context.Context, req *dept.ViewReq) (res *dept.ViewRes, err error) {
	data, err := service.AdminDept().View(ctx, adminin.DeptViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(dept.ViewRes)
	res.DeptViewModel = data
	return
}

// List 查看列表
func (c *cDept) List(ctx context.Context, req *dept.ListReq) (res *dept.ListRes, err error) {
	var in adminin.DeptListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	data, err := service.AdminDept().List(ctx, in)
	if err != nil {
		return
	}

	res = (*dept.ListRes)(&data)
	return
}

// ListTree 查看列表树
func (c *cDept) ListTree(ctx context.Context, req *dept.ListTreeReq) (res *dept.ListTreeRes, err error) {
	var in adminin.DeptListTreeInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	data, err := service.AdminDept().ListTree(ctx, in)
	if err != nil {
		return
	}

	res = (*dept.ListTreeRes)(&data)
	return
}

// Status 更新部门状态
func (c *cDept) Status(ctx context.Context, req *dept.StatusReq) (res *dept.StatusRes, err error) {
	var in adminin.DeptStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminDept().Status(ctx, in)
	return
}
