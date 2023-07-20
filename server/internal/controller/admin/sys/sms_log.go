// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/smslog"
	"hotgo/internal/service"
)

var (
	SmsLog = cSmsLog{}
)

type cSmsLog struct{}

// Delete 删除
func (c *cSmsLog) Delete(ctx context.Context, req *smslog.DeleteReq) (res *smslog.DeleteRes, err error) {
	err = service.SysSmsLog().Delete(ctx, &req.SmsLogDeleteInp)
	return
}

// Edit 更新
func (c *cSmsLog) Edit(ctx context.Context, req *smslog.EditReq) (res *smslog.EditRes, err error) {
	err = service.SysSmsLog().Edit(ctx, &req.SmsLogEditInp)
	return
}

// View 获取指定信息
func (c *cSmsLog) View(ctx context.Context, req *smslog.ViewReq) (res *smslog.ViewRes, err error) {
	data, err := service.SysSmsLog().View(ctx, &req.SmsLogViewInp)
	if err != nil {
		return
	}

	res = new(smslog.ViewRes)
	res.SmsLogViewModel = data
	return
}

// List 查看列表
func (c *cSmsLog) List(ctx context.Context, req *smslog.ListReq) (res *smslog.ListRes, err error) {
	list, totalCount, err := service.SysSmsLog().List(ctx, &req.SmsLogListInp)
	if err != nil {
		return
	}

	res = new(smslog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Status 更新状态
func (c *cSmsLog) Status(ctx context.Context, req *smslog.StatusReq) (res *smslog.StatusRes, err error) {
	err = service.SysSmsLog().Status(ctx, &req.SmsLogStatusInp)
	return
}
