// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.1.1
// @AutoGenerate Date 2023-01-19 16:57:33
package sys

import (
	"context"
	"hotgo/api/admin/loginlog"
	"hotgo/internal/service"
)

var (
	LoginLog = cLoginLog{}
)

type cLoginLog struct{}

// List 查看登录日志列表
func (c *cLoginLog) List(ctx context.Context, req *loginlog.ListReq) (res *loginlog.ListRes, err error) {
	list, totalCount, err := service.SysLoginLog().List(ctx, &req.LoginLogListInp)
	if err != nil {
		return
	}

	res = new(loginlog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出登录日志列表
func (c *cLoginLog) Export(ctx context.Context, req *loginlog.ExportReq) (res *loginlog.ExportRes, err error) {
	err = service.SysLoginLog().Export(ctx, &req.LoginLogListInp)
	return
}

// View 获取指定登录日志信息
func (c *cLoginLog) View(ctx context.Context, req *loginlog.ViewReq) (res *loginlog.ViewRes, err error) {
	data, err := service.SysLoginLog().View(ctx, &req.LoginLogViewInp)
	if err != nil {
		return
	}

	res = new(loginlog.ViewRes)
	res.LoginLogViewModel = data
	return
}

// Delete 删除登录日志
func (c *cLoginLog) Delete(ctx context.Context, req *loginlog.DeleteReq) (res *loginlog.DeleteRes, err error) {
	err = service.SysLoginLog().Delete(ctx, &req.LoginLogDeleteInp)
	return
}
