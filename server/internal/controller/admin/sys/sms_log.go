// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/smslog"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	SmsLog = cSmsLog{}
)

type cSmsLog struct{}

// Delete 删除
func (c *cSmsLog) Delete(ctx context.Context, req *smslog.DeleteReq) (res *smslog.DeleteRes, err error) {
	var in sysin.SmsLogDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysSmsLog().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cSmsLog) Edit(ctx context.Context, req *smslog.EditReq) (res *smslog.EditRes, err error) {
	var in sysin.SmsLogEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysSmsLog().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cSmsLog) MaxSort(ctx context.Context, req *smslog.MaxSortReq) (res *smslog.MaxSortRes, err error) {
	data, err := service.SysSmsLog().MaxSort(ctx, sysin.SmsLogMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(smslog.MaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cSmsLog) View(ctx context.Context, req *smslog.ViewReq) (res *smslog.ViewRes, err error) {
	data, err := service.SysSmsLog().View(ctx, sysin.SmsLogViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(smslog.ViewRes)
	res.SmsLogViewModel = data
	return
}

// List 查看列表
func (c *cSmsLog) List(ctx context.Context, req *smslog.ListReq) (res *smslog.ListRes, err error) {
	var in sysin.SmsLogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysSmsLog().List(ctx, in)
	if err != nil {
		return
	}

	res = new(smslog.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新部门状态
func (c *cSmsLog) Status(ctx context.Context, req *smslog.StatusReq) (res *smslog.StatusRes, err error) {
	var in sysin.SmsLogStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysSmsLog().Status(ctx, in)
	return
}
