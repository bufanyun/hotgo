// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.7.6
package sys

import (
	"context"
	"hotgo/api/admin/servelicense"
	"hotgo/internal/service"
)

var (
	ServeLicense = cServeLicense{}
)

type cServeLicense struct{}

// List 查看服务授权许可列表
func (c *cServeLicense) List(ctx context.Context, req *servelicense.ListReq) (res *servelicense.ListRes, err error) {
	list, totalCount, err := service.SysServeLicense().List(ctx, &req.ServeLicenseListInp)
	if err != nil {
		return
	}

	res = new(servelicense.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出服务授权许可列表
func (c *cServeLicense) Export(ctx context.Context, req *servelicense.ExportReq) (res *servelicense.ExportRes, err error) {
	err = service.SysServeLicense().Export(ctx, &req.ServeLicenseListInp)
	return
}

// Edit 更新服务授权许可
func (c *cServeLicense) Edit(ctx context.Context, req *servelicense.EditReq) (res *servelicense.EditRes, err error) {
	err = service.SysServeLicense().Edit(ctx, &req.ServeLicenseEditInp)
	return
}

// View 获取指定服务授权许可信息
func (c *cServeLicense) View(ctx context.Context, req *servelicense.ViewReq) (res *servelicense.ViewRes, err error) {
	data, err := service.SysServeLicense().View(ctx, &req.ServeLicenseViewInp)
	if err != nil {
		return
	}

	res = new(servelicense.ViewRes)
	res.ServeLicenseViewModel = data
	return
}

// Delete 删除服务授权许可
func (c *cServeLicense) Delete(ctx context.Context, req *servelicense.DeleteReq) (res *servelicense.DeleteRes, err error) {
	err = service.SysServeLicense().Delete(ctx, &req.ServeLicenseDeleteInp)
	return
}

// Status 更新服务授权许可状态
func (c *cServeLicense) Status(ctx context.Context, req *servelicense.StatusReq) (res *servelicense.StatusRes, err error) {
	err = service.SysServeLicense().Status(ctx, &req.ServeLicenseStatusInp)
	return
}

// AssignRouter 分配服务授权许可路由
func (c *cServeLicense) AssignRouter(ctx context.Context, req *servelicense.AssignRouterReq) (res *servelicense.AssignRouterRes, err error) {
	err = service.SysServeLicense().AssignRouter(ctx, &req.ServeLicenseAssignRouterInp)
	return
}
