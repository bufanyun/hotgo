// Package base
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package base

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/home/base"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// Site 基础
var Site = cSite{}

type cSite struct{}

func (a *cSite) Index(ctx context.Context, _ *base.SiteIndexReq) (res *base.SiteIndexRes, err error) {
	service.View().Render(ctx, model.View{Data: g.Map{
		"name":    simple.AppName(ctx),
		"version": consts.VersionApp,
	}})

	//err = gerror.New("这是一个测试错误")
	//return

	//err = gerror.NewCode(gcode.New(10000, "这是一个测试自定义错误码错误", nil))
	//return

	//service.View().Error(ctx, gerror.New("这是一个允许被自定义格式的错误，默认和通用错误格式一致，你可以修改它"))
	//return
	return
}
