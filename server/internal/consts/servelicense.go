// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import (
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

func init() {
	dict.RegisterEnums("ServerLicenseGroupOptions", "服务授权分组选项", ServerLicenseGroupOptions)
}

// 授权分组
const (
	LicenseGroupDefault = "default" // 默认组
	LicenseGroupCron    = "cron"    // 定时任务
	LicenseGroupAuth    = "auth"    // 服务授权
)

// ServerLicenseGroupOptions 服务授权分组选项
var ServerLicenseGroupOptions = []*model.Option{
	dict.GenWarningOption(LicenseGroupDefault, "默认组"),
	dict.GenSuccessOption(LicenseGroupCron, "定时任务"),
	dict.GenSuccessOption(LicenseGroupAuth, "服务授权"),
}
