// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

// 状态码
const (
	StatusALL     int = -1 // 全部状态
	StatusEnabled int = 1  // 启用
	StatusDisable int = 2  // 禁用
	StatusDelete  int = 3  // 已删除
)

var StatusSlice = []int{StatusALL, StatusEnabled, StatusDisable, StatusDelete}
