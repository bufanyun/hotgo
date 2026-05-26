// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.15.7
package sys

import (
	"context"
	"hotgo/api/admin/testcategory"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	TestCategory = cTestCategory{}
)

type cTestCategory struct{}

// List 查看测试分类列表
func (c *cTestCategory) List(ctx context.Context, req *testcategory.ListReq) (res *testcategory.ListRes, err error) {
	list, totalCount, err := service.SysTestCategory().List(ctx, &req.TestCategoryListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.TestCategoryListModel{}
	}

	res = new(testcategory.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Edit 更新测试分类
func (c *cTestCategory) Edit(ctx context.Context, req *testcategory.EditReq) (res *testcategory.EditRes, err error) {
	err = service.SysTestCategory().Edit(ctx, &req.TestCategoryEditInp)
	return
}

// MaxSort 获取测试分类最大排序
func (c *cTestCategory) MaxSort(ctx context.Context, req *testcategory.MaxSortReq) (res *testcategory.MaxSortRes, err error) {
	data, err := service.SysTestCategory().MaxSort(ctx, &req.TestCategoryMaxSortInp)
	if err != nil {
		return
	}

	res = new(testcategory.MaxSortRes)
	res.TestCategoryMaxSortModel = data
	return
}

// View 获取指定测试分类信息
func (c *cTestCategory) View(ctx context.Context, req *testcategory.ViewReq) (res *testcategory.ViewRes, err error) {
	data, err := service.SysTestCategory().View(ctx, &req.TestCategoryViewInp)
	if err != nil {
		return
	}

	res = new(testcategory.ViewRes)
	res.TestCategoryViewModel = data
	return
}

// Delete 删除测试分类
func (c *cTestCategory) Delete(ctx context.Context, req *testcategory.DeleteReq) (res *testcategory.DeleteRes, err error) {
	err = service.SysTestCategory().Delete(ctx, &req.TestCategoryDeleteInp)
	return
}

// Status 更新测试分类状态
func (c *cTestCategory) Status(ctx context.Context, req *testcategory.StatusReq) (res *testcategory.StatusRes, err error) {
	err = service.SysTestCategory().Status(ctx, &req.TestCategoryStatusInp)
	return
}
