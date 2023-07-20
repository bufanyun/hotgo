// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/servelog"
	"hotgo/internal/service"
)

var (
	ServeLog = cServeLog{}
)

type cServeLog struct{}

// List 查看服务日志列表
func (c *cServeLog) List(ctx context.Context, req *servelog.ListReq) (res *servelog.ListRes, err error) {
	list, totalCount, err := service.SysServeLog().List(ctx, &req.ServeLogListInp)
	if err != nil {
		return
	}

	res = new(servelog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出服务日志列表
func (c *cServeLog) Export(ctx context.Context, req *servelog.ExportReq) (res *servelog.ExportRes, err error) {
	err = service.SysServeLog().Export(ctx, &req.ServeLogListInp)
	return
}

// View 获取指定服务日志信息
func (c *cServeLog) View(ctx context.Context, req *servelog.ViewReq) (res *servelog.ViewRes, err error) {
	data, err := service.SysServeLog().View(ctx, &req.ServeLogViewInp)
	if err != nil {
		return
	}

	res = new(servelog.ViewRes)
	res.ServeLogViewModel = data
	return
}

// Delete 删除服务日志
func (c *cServeLog) Delete(ctx context.Context, req *servelog.DeleteReq) (res *servelog.DeleteRes, err error) {
	err = service.SysServeLog().Delete(ctx, &req.ServeLogDeleteInp)
	return
}
