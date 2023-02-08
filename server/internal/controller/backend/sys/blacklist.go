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
	"hotgo/api/backend/blacklist"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Blacklist = cBlacklist{}
)

type cBlacklist struct{}

// Delete 删除
func (c *cBlacklist) Delete(ctx context.Context, req *blacklist.DeleteReq) (res *blacklist.DeleteRes, err error) {
	var in sysin.BlacklistDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysBlacklist().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cBlacklist) Edit(ctx context.Context, req *blacklist.EditReq) (res *blacklist.EditRes, err error) {
	var in sysin.BlacklistEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysBlacklist().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cBlacklist) MaxSort(ctx context.Context, req *blacklist.MaxSortReq) (res *blacklist.MaxSortRes, err error) {
	data, err := service.SysBlacklist().MaxSort(ctx, sysin.BlacklistMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(blacklist.MaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cBlacklist) View(ctx context.Context, req *blacklist.ViewReq) (res *blacklist.ViewRes, err error) {
	data, err := service.SysBlacklist().View(ctx, sysin.BlacklistViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(blacklist.ViewRes)
	res.BlacklistViewModel = data
	return
}

// List 查看列表
func (c *cBlacklist) List(ctx context.Context, req *blacklist.ListReq) (res *blacklist.ListRes, err error) {
	var in sysin.BlacklistListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysBlacklist().List(ctx, in)
	if err != nil {
		return
	}

	res = new(blacklist.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新部门状态
func (c *cBlacklist) Status(ctx context.Context, req *blacklist.StatusReq) (res *blacklist.StatusRes, err error) {
	var in sysin.BlacklistStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysBlacklist().Status(ctx, in)
	return
}
