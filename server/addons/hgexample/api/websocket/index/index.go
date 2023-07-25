// Package index
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package index

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/hgexample/model/input/sysin"
)

// TestReq 测试
type TestReq struct {
	g.Meta `path:"/index/test" method:"get" tags:"功能案例" summary:"测试"`
	sysin.IndexTestInp
}

type TestRes struct {
	*sysin.IndexTestModel
}
