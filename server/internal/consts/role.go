// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package consts

// 数据范围
const (
	RoleDataAll = 1 // 全部权限

	// 通过部门划分
	RoleDataNowDept    = 2 // 当前部门
	RoleDataDeptAndSub = 3 // 当前部门及以下部门
	RoleDataDeptCustom = 4 // 自定义部门

	// 通过上下级关系划分
	RoleDataSelf          = 5 // 仅自己
	RoleDataSelfAndSub    = 6 // 自己和直属下级
	RoleDataSelfAndAllSub = 7 // 自己和全部下级
)

var RoleDataNameMap = map[int]string{
	RoleDataAll:           "全部权限",
	RoleDataNowDept:       "当前部门",
	RoleDataDeptAndSub:    "当前及以下部门",
	RoleDataDeptCustom:    "自定义部门",
	RoleDataSelf:          "仅自己",
	RoleDataSelfAndSub:    "自己和直属下级",
	RoleDataSelfAndAllSub: "自己和全部下级",
}
