// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/test"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var (
	Test = cTest{}
)

type cTest struct{}

// List 查看列表
func (c *cTest) List(ctx context.Context, req *test.ListReq) (res *test.ListRes, err error) {
	var in adminin.TestListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	list, totalCount, err := service.AdminTest().List(ctx, in)
	if err != nil {
		return
	}

	res = new(test.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Export 导出列表
func (c *cTest) Export(ctx context.Context, req *test.ExportReq) (res *test.ExportRes, err error) {
	var in adminin.TestListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminTest().Export(ctx, in)
	return
}

// Edit 更新
func (c *cTest) Edit(ctx context.Context, req *test.EditReq) (res *test.EditRes, err error) {
	var in adminin.TestEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminTest().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cTest) MaxSort(ctx context.Context, req *test.MaxSortReq) (res *test.MaxSortRes, err error) {
	data, err := service.AdminTest().MaxSort(ctx, adminin.TestMaxSortInp{})
	if err != nil {
		return
	}

	res = new(test.MaxSortRes)
	res.TestMaxSortModel = data
	return
}

// View 获取指定信息
func (c *cTest) View(ctx context.Context, req *test.ViewReq) (res *test.ViewRes, err error) {
	var in adminin.TestViewInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	data, err := service.AdminTest().View(ctx, in)
	if err != nil {
		return
	}

	res = new(test.ViewRes)
	res.TestViewModel = data
	return
}

// Delete 删除
func (c *cTest) Delete(ctx context.Context, req *test.DeleteReq) (res *test.DeleteRes, err error) {
	var in adminin.TestDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminTest().Delete(ctx, in)
	return
}

// Status 更新状态
func (c *cTest) Status(ctx context.Context, req *test.StatusReq) (res *test.StatusRes, err error) {
	var in adminin.TestStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminTest().Status(ctx, in)
	return
}

// Switch 更新开关状态
func (c *cTest) Switch(ctx context.Context, req *test.SwitchReq) (res *test.SwitchRes, err error) {
	var in adminin.TestSwitchInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminTest().Switch(ctx, in)
	return
}
