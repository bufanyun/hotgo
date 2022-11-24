// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package global

import (
	"hotgo/internal/model"
	"runtime"
)

var (
	// RootPtah 运行根路径
	RootPtah string
	// SysType 操作系统类型  windows | linux
	SysType = runtime.GOOS
	// MonitorData 监控数据
	MonitorData model.MonitorData
)
