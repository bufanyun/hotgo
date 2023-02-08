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
	"hotgo/api/backend/emslog"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	EmsLog = cEmsLog{}
)

type cEmsLog struct{}

// Delete 删除
func (c *cEmsLog) Delete(ctx context.Context, req *emslog.DeleteReq) (res *emslog.DeleteRes, err error) {
	var in sysin.EmsLogDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysEmsLog().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cEmsLog) Edit(ctx context.Context, req *emslog.EditReq) (res *emslog.EditRes, err error) {
	var in sysin.EmsLogEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysEmsLog().Edit(ctx, in)
	return
}

// View 获取指定信息
func (c *cEmsLog) View(ctx context.Context, req *emslog.ViewReq) (res *emslog.ViewRes, err error) {
	data, err := service.SysEmsLog().View(ctx, sysin.EmsLogViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(emslog.ViewRes)
	res.EmsLogViewModel = data
	return
}

// List 查看列表
func (c *cEmsLog) List(ctx context.Context, req *emslog.ListReq) (res *emslog.ListRes, err error) {
	var in sysin.EmsLogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysEmsLog().List(ctx, in)
	if err != nil {
		return
	}

	res = new(emslog.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新部门状态
func (c *cEmsLog) Status(ctx context.Context, req *emslog.StatusReq) (res *emslog.StatusRes, err error) {
	var in sysin.EmsLogStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysEmsLog().Status(ctx, in)
	return
}
