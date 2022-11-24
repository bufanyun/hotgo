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
		return nil, err
	}
	if err = service.SysBlacklist().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cBlacklist) Edit(ctx context.Context, req *blacklist.EditReq) (res *blacklist.EditRes, err error) {

	var in sysin.BlacklistEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysBlacklist().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// MaxSort 最大排序
func (c *cBlacklist) MaxSort(ctx context.Context, req *blacklist.MaxSortReq) (*blacklist.MaxSortRes, error) {

	data, err := service.SysBlacklist().MaxSort(ctx, sysin.BlacklistMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res blacklist.MaxSortRes
	res.Sort = data.Sort
	return &res, nil
}

// View 获取指定信息
func (c *cBlacklist) View(ctx context.Context, req *blacklist.ViewReq) (*blacklist.ViewRes, error) {

	data, err := service.SysBlacklist().View(ctx, sysin.BlacklistViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res blacklist.ViewRes
	res.BlacklistViewModel = data
	return &res, nil
}

// List 查看列表
func (c *cBlacklist) List(ctx context.Context, req *blacklist.ListReq) (*blacklist.ListRes, error) {

	var (
		in  sysin.BlacklistListInp
		res blacklist.ListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.SysBlacklist().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// Status 更新部门状态
func (c *cBlacklist) Status(ctx context.Context, req *blacklist.StatusReq) (res *blacklist.StatusRes, err error) {

	var in sysin.BlacklistStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysBlacklist().Status(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}
