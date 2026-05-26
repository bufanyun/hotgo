// Package form
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package form

import "hotgo/internal/consts"

// StatusReq 通用状态查询
type StatusReq struct {
	Status int `json:"status" v:"in:-1,0,1,2,3#输入的状态是无效的" dc:"状态"`
}

// SwitchReq 更新开关状态
type SwitchReq struct {
	Key   string `json:"key" v:"required#测试ID不能为空" dc:"开关字段"`
	Value int    `json:"value" v:"in:1,2#输入的开关值是无效的" dc:"更新值"`
}

// AvatarGroup 头像组
type AvatarGroup struct {
	Name string `json:"name" dc:"姓名"`
	Src  string `json:"src" dc:"头像地址"`
}

// DefaultMaxSort 默认最大排序
func DefaultMaxSort(baseSort int) int {
	return baseSort + consts.MaxSortIncr
}
