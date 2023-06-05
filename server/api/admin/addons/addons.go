// Package addons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package addons

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/addons/list" method:"get" tags:"插件管理" summary:"获取插件列表"`
	sysin.AddonsListInp
}

type ListRes struct {
	List []*sysin.AddonsListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

type SelectsReq struct {
	g.Meta `path:"/addons/selects" method:"get" tags:"插件管理" summary:"生成入口选项"`
}

type SelectsRes struct {
	*sysin.AddonsSelectsModel
}

// BuildReq 提交生成
type BuildReq struct {
	g.Meta `path:"/addons/build" method:"post" tags:"插件管理" summary:"提交生成"`
	sysin.AddonsBuildInp
}

type BuildRes struct {
}

// InstallReq 安装模块
type InstallReq struct {
	g.Meta `path:"/addons/install" method:"post" tags:"插件管理" summary:"安装模块"`
	sysin.AddonsInstallInp
}

type InstallRes struct {
}

// UpgradeReq 更新模块
type UpgradeReq struct {
	g.Meta `path:"/addons/upgrade" method:"post" tags:"插件管理" summary:"更新模块"`
	sysin.AddonsUpgradeInp
}

type UpgradeRes struct {
}

// UnInstallReq 卸载模块
type UnInstallReq struct {
	g.Meta `path:"/addons/uninstall" method:"post" tags:"插件管理" summary:"卸载模块"`
	sysin.AddonsUnInstallInp
}

type UnInstallRes struct {
}
