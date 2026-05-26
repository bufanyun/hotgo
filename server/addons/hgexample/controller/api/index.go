// Package api
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package api

import (
	"context"
	"hotgo/addons/hgexample/api/api/index"
	"hotgo/addons/hgexample/service"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

// Test 测试
func (c *cIndex) Test(ctx context.Context, req *index.TestReq) (res *index.TestRes, err error) {
	data, err := service.SysIndex().Test(ctx, &req.IndexTestInp)
	if err != nil {
		return
	}

	res = new(index.TestRes)
	res.IndexTestModel = data
	return
}
