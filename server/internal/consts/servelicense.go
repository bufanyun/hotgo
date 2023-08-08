// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

// 授权分组
const (
	LicenseGroupDefault = "default" // 默认组
	LicenseGroupCron    = "cron"    // 定时任务
	LicenseGroupAuth    = "auth"    // 服务授权
)

var LicenseGroupNameMap = map[string]string{
	LicenseGroupDefault: "默认组",
	LicenseGroupCron:    "定时任务",
	LicenseGroupAuth:    "服务授权",
}
