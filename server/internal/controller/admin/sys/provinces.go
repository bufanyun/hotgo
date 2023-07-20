// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/provinces"
	"hotgo/internal/library/location"
	"hotgo/internal/service"
)

var (
	Provinces = cProvinces{}
)

type cProvinces struct{}

// Tree 关系树选项列表
func (c *cProvinces) Tree(ctx context.Context, _ *provinces.TreeReq) (res *provinces.TreeRes, err error) {
	res = new(provinces.TreeRes)
	res.List, err = service.SysProvinces().Tree(ctx)
	return
}

// Delete 删除
func (c *cProvinces) Delete(ctx context.Context, req *provinces.DeleteReq) (res *provinces.DeleteRes, err error) {
	err = service.SysProvinces().Delete(ctx, &req.ProvincesDeleteInp)
	return
}

// Edit 更新
func (c *cProvinces) Edit(ctx context.Context, req *provinces.EditReq) (res *provinces.EditRes, err error) {
	err = service.SysProvinces().Edit(ctx, &req.ProvincesEditInp)
	return
}

// MaxSort 最大排序
func (c *cProvinces) MaxSort(ctx context.Context, req *provinces.MaxSortReq) (res *provinces.MaxSortRes, err error) {
	data, err := service.SysProvinces().MaxSort(ctx, &req.ProvincesMaxSortInp)
	if err != nil {
		return
	}

	res = new(provinces.MaxSortRes)
	res.ProvincesMaxSortModel = data
	return
}

// View 获取指定信息
func (c *cProvinces) View(ctx context.Context, req *provinces.ViewReq) (res *provinces.ViewRes, err error) {
	data, err := service.SysProvinces().View(ctx, &req.ProvincesViewInp)
	if err != nil {
		return
	}

	res = new(provinces.ViewRes)
	res.ProvincesViewModel = data
	return
}

// List 查看列表
func (c *cProvinces) List(ctx context.Context, req *provinces.ListReq) (res *provinces.ListRes, err error) {
	list, totalCount, err := service.SysProvinces().List(ctx, &req.ProvincesListInp)
	if err != nil {
		return
	}

	res = new(provinces.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Status 更新省市区状态
func (c *cProvinces) Status(ctx context.Context, req *provinces.StatusReq) (res *provinces.StatusRes, err error) {
	err = service.SysProvinces().Status(ctx, &req.ProvincesStatusInp)
	return
}

// ChildrenList 获取省市区下级列表
func (c *cProvinces) ChildrenList(ctx context.Context, req *provinces.ChildrenListReq) (res *provinces.ChildrenListRes, err error) {
	list, totalCount, err := service.SysProvinces().ChildrenList(ctx, &req.ProvincesChildrenListInp)
	if err != nil {
		return
	}

	res = new(provinces.ChildrenListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// UniqueId 地区ID是否唯一
func (c *cProvinces) UniqueId(ctx context.Context, req *provinces.UniqueIdReq) (res *provinces.UniqueIdRes, err error) {
	data, err := service.SysProvinces().UniqueId(ctx, &req.ProvincesUniqueIdInp)
	if err != nil {
		return
	}

	res = new(provinces.UniqueIdRes)
	res.ProvincesUniqueIdModel = data
	return
}

// Select 省市区选项
func (c *cProvinces) Select(ctx context.Context, req *provinces.SelectReq) (res *provinces.SelectRes, err error) {
	data, err := service.SysProvinces().Select(ctx, &req.ProvincesSelectInp)
	if err != nil {
		return
	}

	res = new(provinces.SelectRes)
	res.ProvincesSelectModel = data
	return
}

// CityLabel 省市区选项
func (c *cProvinces) CityLabel(ctx context.Context, req *provinces.CityLabelReq) (res *provinces.CityLabelRes, err error) {
	cityLabel, err := location.ParseSimpleRegion(ctx, req.Id, req.Spilt)
	if err != nil {
		return
	}

	res = (*provinces.CityLabelRes)(&cityLabel)
	return
}
