// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	"runtime"
)

var (
	// RootPtah 运行根路径
	RootPtah string
	// SysType 操作系统类型  windows | linux
	SysType = runtime.GOOS
	// JaegerSwitch 链路追踪开关
	JaegerSwitch bool
)
