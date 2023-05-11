// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/cron"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/msgin"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Cron = cCron{}
)

type cCron struct{}

// Delete 删除
func (c *cCron) Delete(ctx context.Context, req *cron.DeleteReq) (res *cron.DeleteRes, err error) {
	var in = new(msgin.CronDelete)
	if err = gconv.Scan(req, &in.CronDeleteInp); err != nil {
		return
	}

	err = service.TCPServer().CronDelete(ctx, in)
	return
}

// Edit 更新
func (c *cCron) Edit(ctx context.Context, req *cron.EditReq) (res *cron.EditRes, err error) {
	var in = new(msgin.CronEdit)
	if err = gconv.Scan(req, &in.CronEditInp); err != nil {
		return
	}

	err = service.TCPServer().CronEdit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cCron) MaxSort(ctx context.Context, req *cron.MaxSortReq) (res *cron.MaxSortRes, err error) {
	data, err := service.SysCron().MaxSort(ctx, sysin.CronMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(cron.MaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cCron) View(ctx context.Context, req *cron.ViewReq) (res *cron.ViewRes, err error) {
	data, err := service.SysCron().View(ctx, sysin.CronViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(cron.ViewRes)
	res.CronViewModel = data
	return
}

// List 查看列表
func (c *cCron) List(ctx context.Context, req *cron.ListReq) (res *cron.ListRes, err error) {
	var in sysin.CronListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysCron().List(ctx, in)
	if err != nil {
		return
	}

	res = new(cron.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新部门状态
func (c *cCron) Status(ctx context.Context, req *cron.StatusReq) (res *cron.StatusRes, err error) {
	var in sysin.CronStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysCron().Status(ctx, in)
	return
}

// OnlineExec 在线执行
func (c *cCron) OnlineExec(ctx context.Context, req *cron.OnlineExecReq) (res *cron.OnlineExecRes, err error) {
	if req.Id <= 0 {
		return nil, gerror.New("定时任务ID不能为空")
	}

	var in = new(msgin.CronOnlineExec)
	if err = gconv.Scan(req, &in.OnlineExecInp); err != nil {
		return
	}

	err = service.TCPServer().CronOnlineExec(ctx, in)
	return
}
