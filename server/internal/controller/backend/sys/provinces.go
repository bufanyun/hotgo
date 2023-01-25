// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/provinces"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Provinces = cProvinces{}
)

type cProvinces struct{}

// Tree 关系树选项列表
func (c *cProvinces) Tree(ctx context.Context, req *provinces.TreeReq) (*provinces.TreeRes, error) {
	var (
		res provinces.TreeRes
		err error
	)
	res.List, err = service.SysProvinces().Tree(ctx)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Delete 删除
func (c *cProvinces) Delete(ctx context.Context, req *provinces.DeleteReq) (res *provinces.DeleteRes, err error) {
	var in sysin.ProvincesDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysProvinces().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cProvinces) Edit(ctx context.Context, req *provinces.EditReq) (res *provinces.EditRes, err error) {
	var in sysin.ProvincesEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysProvinces().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// MaxSort 最大排序
func (c *cProvinces) MaxSort(ctx context.Context, req *provinces.MaxSortReq) (res *provinces.MaxSortRes, err error) {
	data, err := service.SysProvinces().MaxSort(ctx, sysin.ProvincesMaxSortInp{})
	if err != nil {
		return nil, err
	}

	res = new(provinces.MaxSortRes)
	res.ProvincesMaxSortModel = data
	return res, nil
}

// View 获取指定信息
func (c *cProvinces) View(ctx context.Context, req *provinces.ViewReq) (*provinces.ViewRes, error) {
	data, err := service.SysProvinces().View(ctx, sysin.ProvincesViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var res provinces.ViewRes
	res.ProvincesViewModel = data
	return &res, nil
}

// List 查看列表
func (c *cProvinces) List(ctx context.Context, req *provinces.ListReq) (*provinces.ListRes, error) {
	var (
		in  sysin.ProvincesListInp
		res provinces.ListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	list, totalCount, err := service.SysProvinces().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// Status 更新部门状态
func (c *cProvinces) Status(ctx context.Context, req *provinces.StatusReq) (res *provinces.StatusRes, err error) {
	var in sysin.ProvincesStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysProvinces().Status(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// ChildrenList 获取省市区下级列表
func (c *cProvinces) ChildrenList(ctx context.Context, req *provinces.ChildrenListReq) (*provinces.ChildrenListRes, error) {
	var (
		in  sysin.ProvincesChildrenListInp
		res provinces.ChildrenListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	list, totalCount, err := service.SysProvinces().ChildrenList(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// UniqueId 地区ID是否唯一
func (c *cProvinces) UniqueId(ctx context.Context, req *provinces.UniqueIdReq) (res *provinces.UniqueIdRes, err error) {
	var in sysin.ProvincesUniqueIdInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	data, err := service.SysProvinces().UniqueId(ctx, in)
	if err != nil {
		return nil, err
	}

	res = new(provinces.UniqueIdRes)
	res.ProvincesUniqueIdModel = data

	return res, nil
}
