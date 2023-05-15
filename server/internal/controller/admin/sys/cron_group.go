// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/cron"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	CronGroup = cCronGroup{}
)

type cCronGroup struct{}

// Delete 删除
func (c *cCronGroup) Delete(ctx context.Context, req *cron.GroupDeleteReq) (res *cron.GroupDeleteRes, err error) {
	var in sysin.CronGroupDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysCronGroup().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cCronGroup) Edit(ctx context.Context, req *cron.GroupEditReq) (res *cron.GroupEditRes, err error) {
	var in sysin.CronGroupEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysCronGroup().Edit(ctx, in)
	return res, nil
}

// MaxSort 最大排序
func (c *cCronGroup) MaxSort(ctx context.Context, req *cron.GroupMaxSortReq) (res *cron.GroupMaxSortRes, err error) {
	data, err := service.SysCronGroup().MaxSort(ctx, sysin.CronGroupMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(cron.GroupMaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cCronGroup) View(ctx context.Context, req *cron.GroupViewReq) (res *cron.GroupViewRes, err error) {
	data, err := service.SysCronGroup().View(ctx, sysin.CronGroupViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(cron.GroupViewRes)
	res.CronGroupViewModel = data
	return
}

// List 查看列表
func (c *cCronGroup) List(ctx context.Context, req *cron.GroupListReq) (res *cron.GroupListRes, err error) {
	var in sysin.CronGroupListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysCronGroup().List(ctx, in)
	if err != nil {
		return
	}

	res = new(cron.GroupListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新状态
func (c *cCronGroup) Status(ctx context.Context, req *cron.GroupStatusReq) (res *cron.GroupStatusRes, err error) {
	var in sysin.CronGroupStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysCronGroup().Status(ctx, in)
	return res, nil
}

// Select 选项
func (c *cCronGroup) Select(ctx context.Context, _ *cron.GroupSelectReq) (res *cron.GroupSelectRes, err error) {
	data, err := service.SysCronGroup().Select(ctx, sysin.CronGroupSelectInp{})
	if err != nil {
		return
	}

	res = new(cron.GroupSelectRes)
	res.CronGroupSelectModel = data
	return
}
