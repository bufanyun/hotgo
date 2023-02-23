// Package home
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package home

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/addons/hgexample/api/home/index"
	"hotgo/addons/hgexample/global"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
	"hotgo/internal/model"
	isc "hotgo/internal/service"
	"hotgo/utility/validate"
)

// Index 基础
var Index = cIndex{}

type cIndex struct{}

func (a *cIndex) Index(ctx context.Context, req *index.TestReq) (res *index.TestRes, err error) {
	var in sysin.IndexTestInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	data, err := service.SysIndex().Test(ctx, in)
	if err != nil {
		return
	}

	isc.View().RenderTpl(ctx, global.Tpl("home/index.html"), model.View{Data: g.Map{
		"name":   data.Name,
		"module": data.Module,
		"time":   data.Time,
	}})
	return
}
