// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/cron"
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
		return nil, err
	}
	if err = service.SysCronGroup().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cCronGroup) Edit(ctx context.Context, req *cron.GroupEditReq) (res *cron.GroupEditRes, err error) {

	var in sysin.CronGroupEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysCronGroup().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// MaxSort 最大排序
func (c *cCronGroup) MaxSort(ctx context.Context, req *cron.GroupMaxSortReq) (*cron.GroupMaxSortRes, error) {

	data, err := service.SysCronGroup().MaxSort(ctx, sysin.CronGroupMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res cron.GroupMaxSortRes
	res.Sort = data.Sort
	return &res, nil
}

// View 获取指定信息
func (c *cCronGroup) View(ctx context.Context, req *cron.GroupViewReq) (*cron.GroupViewRes, error) {

	data, err := service.SysCronGroup().View(ctx, sysin.CronGroupViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res cron.GroupViewRes
	res.CronGroupViewModel = data
	return &res, nil
}

// List 查看列表
func (c *cCronGroup) List(ctx context.Context, req *cron.GroupListReq) (*cron.GroupListRes, error) {

	var (
		in  sysin.CronGroupListInp
		res cron.GroupListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.SysCronGroup().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// Status 更新部门状态
func (c *cCronGroup) Status(ctx context.Context, req *cron.GroupStatusReq) (res *cron.GroupStatusRes, err error) {

	var in sysin.CronGroupStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysCronGroup().Status(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// Select 选项
func (c *cCronGroup) Select(ctx context.Context, req *cron.GroupSelectReq) (res *cron.GroupSelectRes, err error) {
	list, err := service.SysCronGroup().Select(ctx, sysin.CronGroupSelectInp{})
	if err != nil {
		return nil, err
	}
	res = (*cron.GroupSelectRes)(&list)

	return res, nil
}
