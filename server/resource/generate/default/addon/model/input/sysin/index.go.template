// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

// IndexTestInp 测试
type IndexTestInp struct {
	Name string `json:"name" d:"HotGo" dc:"名称"`
}

func (in *IndexTestInp) Filter(ctx context.Context) (err error) {
	return
}

type IndexTestModel struct {
	Name   string      `json:"name" dc:"名称"`
	Module string      `json:"module" dc:"当前插件模块"`
	Time   *gtime.Time `json:"time" dc:"当前时间"`
}
