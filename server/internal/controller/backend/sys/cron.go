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
	Cron = cCron{}
)

type cCron struct{}

// Delete 删除
func (c *cCron) Delete(ctx context.Context, req *cron.DeleteReq) (res *cron.DeleteRes, err error) {
	var in sysin.CronDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysCron().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cCron) Edit(ctx context.Context, req *cron.EditReq) (res *cron.EditRes, err error) {

	var in sysin.CronEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysCron().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// MaxSort 最大排序
func (c *cCron) MaxSort(ctx context.Context, req *cron.MaxSortReq) (*cron.MaxSortRes, error) {

	data, err := service.SysCron().MaxSort(ctx, sysin.CronMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res cron.MaxSortRes
	res.Sort = data.Sort
	return &res, nil
}

// View 获取指定信息
func (c *cCron) View(ctx context.Context, req *cron.ViewReq) (*cron.ViewRes, error) {

	data, err := service.SysCron().View(ctx, sysin.CronViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res cron.ViewRes
	res.CronViewModel = data
	return &res, nil
}

// List 查看列表
func (c *cCron) List(ctx context.Context, req *cron.ListReq) (*cron.ListRes, error) {

	var (
		in  sysin.CronListInp
		res cron.ListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.SysCron().List(ctx, in)
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
func (c *cCron) Status(ctx context.Context, req *cron.StatusReq) (res *cron.StatusRes, err error) {

	var in sysin.CronStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysCron().Status(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}
