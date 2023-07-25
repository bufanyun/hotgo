// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/blacklist"
	"hotgo/internal/service"
)

var (
	Blacklist = cBlacklist{}
)

type cBlacklist struct{}

// Delete 删除
func (c *cBlacklist) Delete(ctx context.Context, req *blacklist.DeleteReq) (res *blacklist.DeleteRes, err error) {
	err = service.SysBlacklist().Delete(ctx, &req.BlacklistDeleteInp)
	return
}

// Edit 更新
func (c *cBlacklist) Edit(ctx context.Context, req *blacklist.EditReq) (res *blacklist.EditRes, err error) {
	err = service.SysBlacklist().Edit(ctx, &req.BlacklistEditInp)
	return
}

// View 获取指定信息
func (c *cBlacklist) View(ctx context.Context, req *blacklist.ViewReq) (res *blacklist.ViewRes, err error) {
	data, err := service.SysBlacklist().View(ctx, &req.BlacklistViewInp)
	if err != nil {
		return
	}

	res = new(blacklist.ViewRes)
	res.BlacklistViewModel = data
	return
}

// List 查看列表
func (c *cBlacklist) List(ctx context.Context, req *blacklist.ListReq) (res *blacklist.ListRes, err error) {
	list, totalCount, err := service.SysBlacklist().List(ctx, &req.BlacklistListInp)
	if err != nil {
		return
	}

	res = new(blacklist.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Status 更新黑名单状态
func (c *cBlacklist) Status(ctx context.Context, req *blacklist.StatusReq) (res *blacklist.StatusRes, err error) {
	err = service.SysBlacklist().Status(ctx, &req.BlacklistStatusInp)
	return
}
