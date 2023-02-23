// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/addons"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var (
	Addons = cAddons{}
)

type cAddons struct{}

// List 查看列表
func (c *cAddons) List(ctx context.Context, req *addons.ListReq) (res *addons.ListRes, err error) {
	var in sysin.AddonsListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysAddons().List(ctx, in)
	if err != nil {
		return
	}

	res = new(addons.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Selects 获取指定信息
func (c *cAddons) Selects(ctx context.Context, req *addons.SelectsReq) (res *addons.SelectsRes, err error) {
	data, err := service.SysAddons().Selects(ctx, sysin.AddonsSelectsInp{})
	if err != nil {
		return
	}

	res = new(addons.SelectsRes)
	res.AddonsSelectsModel = data
	return
}

// Build 生成预览
func (c *cAddons) Build(ctx context.Context, req *addons.BuildReq) (res *addons.BuildRes, err error) {
	var in sysin.AddonsBuildInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysAddons().Build(ctx, in)
	return
}

// Install 安装模块
func (c *cAddons) Install(ctx context.Context, req *addons.InstallReq) (res *addons.InstallRes, err error) {
	var in sysin.AddonsInstallInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = service.SysAddons().Install(ctx, in); err != nil {
		return
	}
	return
}

// Upgrade 更新模块
func (c *cAddons) Upgrade(ctx context.Context, req *addons.UpgradeReq) (res *addons.UpgradeRes, err error) {
	var in sysin.AddonsUpgradeInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = service.SysAddons().Upgrade(ctx, in); err != nil {
		return
	}
	return
}

// UnInstall 卸载模块
func (c *cAddons) UnInstall(ctx context.Context, req *addons.UnInstallReq) (res *addons.UnInstallRes, err error) {
	var in sysin.AddonsUnInstallInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = service.SysAddons().UnInstall(ctx, in); err != nil {
		return
	}
	return
}
