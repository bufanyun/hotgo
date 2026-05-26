// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.13.1
package sys

import (
	"context"
	"hotgo/addons/hgexample/api/admin/tenantorder"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
)

var (
	TenantOrder = cTenantOrder{}
)

type cTenantOrder struct{}

// List 查看多租户功能演示列表
func (c *cTenantOrder) List(ctx context.Context, req *tenantorder.ListReq) (res *tenantorder.ListRes, err error) {
	list, totalCount, err := service.SysTenantOrder().List(ctx, &req.TenantOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.TenantOrderListModel{}
	}

	res = new(tenantorder.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出多租户功能演示列表
func (c *cTenantOrder) Export(ctx context.Context, req *tenantorder.ExportReq) (res *tenantorder.ExportRes, err error) {
	err = service.SysTenantOrder().Export(ctx, &req.TenantOrderListInp)
	return
}

// Edit 更新多租户功能演示
func (c *cTenantOrder) Edit(ctx context.Context, req *tenantorder.EditReq) (res *tenantorder.EditRes, err error) {
	err = service.SysTenantOrder().Edit(ctx, &req.TenantOrderEditInp)
	return
}

// View 获取指定多租户功能演示信息
func (c *cTenantOrder) View(ctx context.Context, req *tenantorder.ViewReq) (res *tenantorder.ViewRes, err error) {
	data, err := service.SysTenantOrder().View(ctx, &req.TenantOrderViewInp)
	if err != nil {
		return
	}

	res = new(tenantorder.ViewRes)
	res.TenantOrderViewModel = data
	return
}

// Delete 删除多租户功能演示
func (c *cTenantOrder) Delete(ctx context.Context, req *tenantorder.DeleteReq) (res *tenantorder.DeleteRes, err error) {
	err = service.SysTenantOrder().Delete(ctx, &req.TenantOrderDeleteInp)
	return
}