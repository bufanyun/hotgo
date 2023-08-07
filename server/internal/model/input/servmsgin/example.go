// Package servmsgin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package servmsgin

import "github.com/gogf/gf/v2/os/gtime"

// ExampleHelloModel 授权信息
type ExampleHelloModel struct {
	Desc      string      `json:"desc"  description:"描述信息"`
	Timestamp *gtime.Time `json:"timestamp"  description:"服务器时间"`
}
