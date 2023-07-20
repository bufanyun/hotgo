// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/cron"
	"hotgo/internal/service"
)

var (
	CronGroup = cCronGroup{}
)

type cCronGroup struct{}

// Delete 删除
func (c *cCronGroup) Delete(ctx context.Context, req *cron.GroupDeleteReq) (res *cron.GroupDeleteRes, err error) {
	err = service.SysCronGroup().Delete(ctx, &req.CronGroupDeleteInp)
	return
}

// Edit 更新
func (c *cCronGroup) Edit(ctx context.Context, req *cron.GroupEditReq) (res *cron.GroupEditRes, err error) {
	err = service.SysCronGroup().Edit(ctx, &req.CronGroupEditInp)
	return
}

// MaxSort 最大排序
func (c *cCronGroup) MaxSort(ctx context.Context, req *cron.GroupMaxSortReq) (res *cron.GroupMaxSortRes, err error) {
	data, err := service.SysCronGroup().MaxSort(ctx, &req.CronGroupMaxSortInp)
	if err != nil {
		return
	}

	res = new(cron.GroupMaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cCronGroup) View(ctx context.Context, req *cron.GroupViewReq) (res *cron.GroupViewRes, err error) {
	data, err := service.SysCronGroup().View(ctx, &req.CronGroupViewInp)
	if err != nil {
		return
	}

	res = new(cron.GroupViewRes)
	res.CronGroupViewModel = data
	return
}

// List 查看列表
func (c *cCronGroup) List(ctx context.Context, req *cron.GroupListReq) (res *cron.GroupListRes, err error) {
	list, totalCount, err := service.SysCronGroup().List(ctx, &req.CronGroupListInp)
	if err != nil {
		return
	}

	res = new(cron.GroupListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Status 更新状态
func (c *cCronGroup) Status(ctx context.Context, req *cron.GroupStatusReq) (res *cron.GroupStatusRes, err error) {
	err = service.SysCronGroup().Status(ctx, &req.CronGroupStatusInp)
	return
}

// Select 选项
func (c *cCronGroup) Select(ctx context.Context, req *cron.GroupSelectReq) (res *cron.GroupSelectRes, err error) {
	data, err := service.SysCronGroup().Select(ctx, &req.CronGroupSelectInp)
	if err != nil {
		return
	}

	res = new(cron.GroupSelectRes)
	res.CronGroupSelectModel = data
	return
}
