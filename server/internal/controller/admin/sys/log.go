// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/log"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

// Log 日志
var Log = sLog{}

type sLog struct{}

// Clear 清空日志
func (c *sLog) Clear(ctx context.Context, req *log.ClearReq) (res *log.ClearRes, err error) {
	err = gerror.New("暂时考虑到安全问题，请到数据库清空")
	return
}

// Export 导出
func (c *sLog) Export(ctx context.Context, req *log.ExportReq) (res *log.ExportRes, err error) {
	var in sysin.LogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysLog().Export(ctx, in)
	return
}

// List 获取访问日志列表
func (c *sLog) List(ctx context.Context, req *log.ListReq) (res *log.ListRes, err error) {
	var in sysin.LogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysLog().List(ctx, in)
	if err != nil {
		return
	}

	res = new(log.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// View 获取指定信息
func (c *sLog) View(ctx context.Context, req *log.ViewReq) (res *log.ViewRes, err error) {
	data, err := service.SysLog().View(ctx, sysin.LogViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(log.ViewRes)
	res.LogViewModel = data
	return
}

// Delete 删除
func (c *sLog) Delete(ctx context.Context, req *log.DeleteReq) (res *log.DeleteRes, err error) {
	var in sysin.LogDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysLog().Delete(ctx, in)
	return
}
