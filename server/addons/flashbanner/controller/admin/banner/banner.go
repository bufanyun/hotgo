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
	Banner = cBanner{}
)

type cBanner struct{}

// Create 创建轮播图
func (c *cBanner) Create(ctx context.Context, req *banner.CreateReq) (res *banner.CreateRes, err error) {
	err = service.SysBanner().Create(ctx, &req.BannerCreateInp)
	if err != nil {
		return nil, err
	}
	res = &banner.CreateRes{}
	return res, nil
}

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

// Edit 修改/新增轮播图
func (c *cBanner) Edit(ctx context.Context, req *banner.EditReq) (res *banner.EditRes, err error) {
	err = service.SysBanner().Edit(ctx, &req.BannerEditInp)
	if err != nil {
		return nil, err
	}
	res = &banner.EditRes{}
	return res, nil
}

// Delete 删除轮播图
func (c *cBanner) Delete(ctx context.Context, req *banner.DeleteReq) (res *banner.DeleteRes, err error) {
	err = service.SysBanner().Delete(ctx, &req.BannerDeleteInp)
	if err != nil {
		return nil, err
	}
	res = &banner.DeleteRes{}
	return res, nil
}

// Status 更新轮播图状态
func (c *cBanner) Status(ctx context.Context, req *banner.StatusReq) (res *banner.StatusRes, err error) {
	err = service.SysBanner().Status(ctx, &req.BannerStatusInp)
	if err != nil {
		return nil, err
	}
	res = &banner.StatusRes{}
	return res, nil
} 