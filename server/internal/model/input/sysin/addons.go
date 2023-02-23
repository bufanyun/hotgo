// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
	"hotgo/internal/library/addons"
	"hotgo/internal/model/input/form"
)

// AddonsListInp 获取列表
type AddonsListInp struct {
	form.PageReq
	Name   string `json:"name"`
	Group  int    `json:"group"`
	Status int    `json:"status"`
}

type AddonsListModel struct {
	addons.Skeleton
	GroupName      string `json:"groupName" dc:"分组名称"`
	InstallVersion string `json:"installVersion" dc:"安装版本"`
	InstallStatus  int    `json:"installStatus" dc:"安装状态"`
	CanSave        bool   `json:"canSave" dc:"是否可以更新"`
}

// AddonsSelectsInp 选项
type AddonsSelectsInp struct {
}
type AddonsSelectsModel struct {
	GroupType form.Selects `json:"groupType" dc:"分组类型"`
	Status    form.Selects `json:"status" dc:"安装状态"`
}

// AddonsBuildInp 提交生成
type AddonsBuildInp struct {
	addons.Skeleton
}

func (in *AddonsBuildInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("插件名称不能为空")
		return
	}

	if !gregex.IsMatchString(`^[a-zA-Z]{1}\w{1,23}$`, in.Name) {
		err = gerror.New("插件名称格式不正确，字母开头，只能包含字母、数字和下划线，长度在2~24之间")
		return
	}

	return
}

// AddonsInstallInp 安装模块
type AddonsInstallInp struct {
	addons.Skeleton
}

// AddonsUpgradeInp 更新模块
type AddonsUpgradeInp struct {
	addons.Skeleton
}

// AddonsUnInstallInp 卸载模块
type AddonsUnInstallInp struct {
	addons.Skeleton
}
