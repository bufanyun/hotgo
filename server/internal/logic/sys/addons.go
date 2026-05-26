// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/library/addons"
	"hotgo/internal/library/dict"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

type sSysAddons struct{}

func NewSysAddons() *sSysAddons {
	return &sSysAddons{}
}

func init() {
	service.RegisterSysAddons(NewSysAddons())
}

// List 获取列表
func (s *sSysAddons) List(ctx context.Context, in *sysin.AddonsListInp) (list []*sysin.AddonsListModel, totalCount int, err error) {
	sks := addons.GetSkeletons()
	if len(sks) == 0 {
		return
	}

	var (
		i                  int
		_, perPage, offset = form.CalPage(in.Page, in.PerPage)
	)

	for k, skeleton := range sks {
		ok := k >= offset && i <= perPage
		if !ok {
			break
		}
		row := new(sysin.AddonsListModel)
		row.Skeleton = *skeleton

		if in.Group > 0 {
			if row.Skeleton.Group != in.Group {
				continue
			}
		}

		if in.Name != "" {
			if row.Skeleton.Label != in.Name && row.Skeleton.Name != in.Name {
				continue
			}
		}

		install, err := addons.ScanInstall(row.Skeleton.GetModule())
		if err != nil {
			continue
		}

		if install == nil {
			row.InstallStatus = consts.AddonsInstallStatusNo
			row.InstallVersion = "v0.0.0"
		} else {
			row.InstallStatus = install.Status
			row.InstallVersion = install.Version
			row.CanSave = gstr.CompareVersion(row.Skeleton.Version, install.Version) > 0
		}

		if in.Status > 0 {
			if row.InstallStatus != in.Status {
				continue
			}
		}

		if row.Skeleton.Logo == "" {
			row.Skeleton.Logo = consts.AddonsGroupIconMap[row.Skeleton.Group]
		}

		row.GroupName = dict.GetOptionLabel(consts.AddonsGroupOptions, row.Skeleton.Group)
		list = append(list, row)
		i++
	}

	totalCount = len(sks)
	return
}

// Build 提交生成
func (s *sSysAddons) Build(ctx context.Context, in *sysin.AddonsBuildInp) (err error) {
	config, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return
	}

	if config == nil || config.Addon == nil {
		err = gerror.New("没有找到有效的生成或插件配置，请检查配置文件是否正常")
		return
	}

	option := new(addons.BuildOption)
	option.Config = config.Addon
	option.Skeleton = in.Skeleton
	option.Extend = in.Extend
	return addons.Build(ctx, option)
}

// Install 安装模块
func (s *sSysAddons) Install(ctx context.Context, in *sysin.AddonsInstallInp) (err error) {
	return addons.Install(in.GetModule())
}

// Upgrade 更新模块
func (s *sSysAddons) Upgrade(ctx context.Context, in *sysin.AddonsUpgradeInp) (err error) {
	return addons.Upgrade(in.GetModule())
}

// UnInstall 卸载模块
func (s *sSysAddons) UnInstall(ctx context.Context, in *sysin.AddonsUnInstallInp) (err error) {
	return addons.UnInstall(in.GetModule())
}
