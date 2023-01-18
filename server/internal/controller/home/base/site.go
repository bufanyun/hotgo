// Package base
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package base

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/home/base"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"hotgo/internal/service"
)

// Site 基础
var Site = cSite{}

type cSite struct{}

func (a *cSite) Index(ctx context.Context, req *base.SiteIndexReq) (res *base.SiteIndexRes, err error) {
	service.View().Render(ctx, model.View{Data: g.Map{
		"name":    "HotGo",
		"version": consts.VersionApp,
		"debug":   g.Cfg().MustGet(ctx, "hotgo.debug", true),
	}})
	return
}
