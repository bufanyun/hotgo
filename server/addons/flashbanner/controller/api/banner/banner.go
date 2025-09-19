// Package banner
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package banner

import (
	"context"
	"hotgo/addons/flashbanner/api/admin/banner"
	"hotgo/addons/flashbanner/service"
)

var (
	BannerIndex = cBanner{}
)

type cBanner struct{}

// List 获取轮播图列表
func (c *cBanner) List(ctx context.Context, req *banner.ListReq) (res *banner.ListRes, err error) {
	list, totalCount, err := service.SysBanner().List(ctx, &req.BannerListInp)
	if err != nil {
		return nil, err
	}
	res = &banner.ListRes{
		List: list,
	}
	res.PageRes.Pack(&req.BannerListInp, totalCount)
	return res, nil
}

// View 获取指定轮播图信息
func (c *cBanner) View(ctx context.Context, req *banner.ViewReq) (res *banner.ViewRes, err error) {
	model, err := service.SysBanner().View(ctx, &req.BannerViewInp)
	if err != nil {
		return nil, err
	}
	res = &banner.ViewRes{model}
	return res, nil
}
