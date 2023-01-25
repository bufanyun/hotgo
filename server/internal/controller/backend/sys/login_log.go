// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.1.1
// @AutoGenerate Date 2023-01-19 16:57:33
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/loginlog"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var (
	LoginLog = cLoginLog{}
)

type cLoginLog struct{}

// List 查看登录日志列表
func (c *cLoginLog) List(ctx context.Context, req *loginlog.ListReq) (res *loginlog.ListRes, err error) {
	var in sysin.LoginLogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.SysLoginLog().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res = new(loginlog.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return res, nil
}

// Export 导出登录日志列表
func (c *cLoginLog) Export(ctx context.Context, req *loginlog.ExportReq) (res *loginlog.ExportRes, err error) {
	var in sysin.LoginLogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	if err = service.SysLoginLog().Export(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// View 获取指定登录日志信息
func (c *cLoginLog) View(ctx context.Context, req *loginlog.ViewReq) (res *loginlog.ViewRes, err error) {
	var in sysin.LoginLogViewInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	data, err := service.SysLoginLog().View(ctx, in)
	if err != nil {
		return nil, err
	}

	res = new(loginlog.ViewRes)
	res.LoginLogViewModel = data
	return res, nil
}

// Delete 删除登录日志
func (c *cLoginLog) Delete(ctx context.Context, req *loginlog.DeleteReq) (res *loginlog.DeleteRes, err error) {
	var in sysin.LoginLogDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	if err = service.SysLoginLog().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}
