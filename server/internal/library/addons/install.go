// Package addons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package addons

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
)

// InstallRecord 安装记录
type InstallRecord struct {
	Id        int64       `json:"id"        description:"安装ID"`
	Version   string      `json:"version"   description:"安装版本"`
	Status    int         `json:"status"    description:"安装状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}

func GetModel(ctx context.Context) *gdb.Model {
	return g.Model("sys_addons_install").Ctx(ctx)
}

func ScanInstall(m Module) (record *InstallRecord, err error) {
	err = GetModel(m.Ctx()).Where("name", m.GetSkeleton().Name).Scan(&record)
	return
}

// IsInstall 模块是否已安装
func IsInstall(m Module) bool {
	record, err := ScanInstall(m)
	if err != nil {
		g.Log().Debugf(m.Ctx(), "addons.IsInstall err:%+v", err)
		return false
	}
	if record == nil {
		return false
	}
	return record.Status == consts.AddonsInstallStatusOk
}

// Install 安装模块
func Install(m Module) (err error) {
	record, err := ScanInstall(m)
	if err != nil {
		return
	}
	if record != nil && record.Status == consts.AddonsInstallStatusOk {
		return gerror.New("插件已安装，无需重复操作！")
	}

	data := g.Map{
		"name":    m.GetSkeleton().Name,
		"version": m.GetSkeleton().Version,
		"status":  consts.AddonsInstallStatusOk,
	}
	return g.DB().Transaction(m.Ctx(), func(ctx context.Context, tx gdb.TX) error {
		if record != nil {
			_, _ = GetModel(m.Ctx()).Where("id", record.Id).Delete()
		}

		if _, err = GetModel(m.Ctx()).Data(data).Insert(); err != nil {
			return err
		}
		return m.Install(ctx)
	})
}

// Upgrade 更新模块
func Upgrade(m Module) (err error) {
	record, err := ScanInstall(m)
	if err != nil {
		return
	}

	if record == nil || record.Status != consts.AddonsInstallStatusOk {
		return gerror.New("插件未安装，请安装后操作！")
	}

	data := g.Map{
		"version": m.GetSkeleton().Version,
	}
	return g.DB().Transaction(m.Ctx(), func(ctx context.Context, tx gdb.TX) error {
		if _, err = GetModel(m.Ctx()).Where("id", record.Id).Data(data).Update(); err != nil {
			return err
		}
		return m.Upgrade(ctx)
	})
}

// UnInstall 卸载模块
func UnInstall(m Module) (err error) {
	record, err := ScanInstall(m)
	if err != nil {
		return
	}

	if record == nil || record.Status != consts.AddonsInstallStatusOk {
		return gerror.New("插件未安装，请安装后操作！")
	}

	data := g.Map{
		"version": m.GetSkeleton().Version,
		"status":  consts.AddonsInstallStatusUn,
	}
	return g.DB().Transaction(m.Ctx(), func(ctx context.Context, tx gdb.TX) error {
		if _, err = GetModel(m.Ctx()).Where("id", record.Id).Data(data).Update(); err != nil {
			return err
		}
		return m.UnInstall(ctx)
	})
}
