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
	"hotgo/api/backend/servelog"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var (
	ServeLog = cServeLog{}
)

type cServeLog struct{}

// List 查看服务日志列表
func (c *cServeLog) List(ctx context.Context, req *servelog.ListReq) (res *servelog.ListRes, err error) {
	var in sysin.ServeLogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.SysServeLog().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res = new(servelog.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return res, nil
}

// Export 导出服务日志列表
func (c *cServeLog) Export(ctx context.Context, req *servelog.ExportReq) (res *servelog.ExportRes, err error) {
	var in sysin.ServeLogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	if err = service.SysServeLog().Export(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// View 获取指定服务日志信息
func (c *cServeLog) View(ctx context.Context, req *servelog.ViewReq) (res *servelog.ViewRes, err error) {
	var in sysin.ServeLogViewInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	data, err := service.SysServeLog().View(ctx, in)
	if err != nil {
		return nil, err
	}

	res = new(servelog.ViewRes)
	res.ServeLogViewModel = data
	return res, nil
}

// Delete 删除服务日志
func (c *cServeLog) Delete(ctx context.Context, req *servelog.DeleteReq) (res *servelog.DeleteRes, err error) {
	var in sysin.ServeLogDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	if err = service.SysServeLog().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}
