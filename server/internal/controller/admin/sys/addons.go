// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/addons"
	"hotgo/internal/service"
)

var (
	Addons = cAddons{}
)

type cAddons struct{}

// List 查看列表
func (c *cAddons) List(ctx context.Context, req *addons.ListReq) (res *addons.ListRes, err error) {
	list, totalCount, err := service.SysAddons().List(ctx, &req.AddonsListInp)
	if err != nil {
		return
	}

	res = new(addons.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Build 生成预览
func (c *cAddons) Build(ctx context.Context, req *addons.BuildReq) (res *addons.BuildRes, err error) {
	err = service.SysAddons().Build(ctx, &req.AddonsBuildInp)
	return
}

// Install 安装模块
func (c *cAddons) Install(ctx context.Context, req *addons.InstallReq) (res *addons.InstallRes, err error) {
	if err = service.SysAddons().Install(ctx, &req.AddonsInstallInp); err != nil {
		return
	}
	return
}

// Upgrade 更新模块
func (c *cAddons) Upgrade(ctx context.Context, req *addons.UpgradeReq) (res *addons.UpgradeRes, err error) {
	if err = service.SysAddons().Upgrade(ctx, &req.AddonsUpgradeInp); err != nil {
		return
	}
	return
}

// UnInstall 卸载模块
func (c *cAddons) UnInstall(ctx context.Context, req *addons.UnInstallReq) (res *addons.UnInstallRes, err error) {
	if err = service.SysAddons().UnInstall(ctx, &req.AddonsUnInstallInp); err != nil {
		return
	}
	return
}
