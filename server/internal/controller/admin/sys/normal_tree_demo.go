// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.15.7
package sys

import (
	"context"
	"hotgo/api/admin/normaltreedemo"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	NormalTreeDemo = cNormalTreeDemo{}
)

type cNormalTreeDemo struct{}

// List 查看普通树表列表
func (c *cNormalTreeDemo) List(ctx context.Context, req *normaltreedemo.ListReq) (res *normaltreedemo.ListRes, err error) {
	list, totalCount, err := service.SysNormalTreeDemo().List(ctx, &req.NormalTreeDemoListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.NormalTreeDemoListModel{}
	}

	res = new(normaltreedemo.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Edit 更新普通树表
func (c *cNormalTreeDemo) Edit(ctx context.Context, req *normaltreedemo.EditReq) (res *normaltreedemo.EditRes, err error) {
	err = service.SysNormalTreeDemo().Edit(ctx, &req.NormalTreeDemoEditInp)
	return
}

// MaxSort 获取普通树表最大排序
func (c *cNormalTreeDemo) MaxSort(ctx context.Context, req *normaltreedemo.MaxSortReq) (res *normaltreedemo.MaxSortRes, err error) {
	data, err := service.SysNormalTreeDemo().MaxSort(ctx, &req.NormalTreeDemoMaxSortInp)
	if err != nil {
		return
	}

	res = new(normaltreedemo.MaxSortRes)
	res.NormalTreeDemoMaxSortModel = data
	return
}

// View 获取指定普通树表信息
func (c *cNormalTreeDemo) View(ctx context.Context, req *normaltreedemo.ViewReq) (res *normaltreedemo.ViewRes, err error) {
	data, err := service.SysNormalTreeDemo().View(ctx, &req.NormalTreeDemoViewInp)
	if err != nil {
		return
	}

	res = new(normaltreedemo.ViewRes)
	res.NormalTreeDemoViewModel = data
	return
}

// Delete 删除普通树表
func (c *cNormalTreeDemo) Delete(ctx context.Context, req *normaltreedemo.DeleteReq) (res *normaltreedemo.DeleteRes, err error) {
	err = service.SysNormalTreeDemo().Delete(ctx, &req.NormalTreeDemoDeleteInp)
	return
}

// TreeOption 获取普通树表关系树选项
func (c *cNormalTreeDemo) TreeOption(ctx context.Context, req *normaltreedemo.TreeOptionReq) (res *normaltreedemo.TreeOptionRes, err error) {
	data, err := service.SysNormalTreeDemo().TreeOption(ctx)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		res = (*normaltreedemo.TreeOptionRes)(&data)
	} else {
		temp := make(normaltreedemo.TreeOptionRes, 0)
		res = &temp
	}
	return
}
