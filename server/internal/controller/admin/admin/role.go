// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"hotgo/api/admin/role"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/service"
)

var (
	Role = cRole{}
)

type cRole struct{}

// List 获取列表
func (c *cRole) List(ctx context.Context, req *role.ListReq) (res *role.ListRes, err error) {
	list, totalCount, err := service.AdminRole().List(ctx, &req.RoleListInp)
	if err != nil {
		return
	}

	res = new(role.ListRes)
	res.RoleListModel = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Edit 修改角色
func (c *cRole) Edit(ctx context.Context, req *role.EditReq) (res *role.EditRes, err error) {
	err = service.AdminRole().Edit(ctx, &req.RoleEditInp)
	return
}

// Delete 删除
func (c *cRole) Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error) {
	err = service.AdminRole().Delete(ctx, &req.RoleDeleteInp)
	return
}

// Dynamic 动态路由
func (c *cRole) Dynamic(ctx context.Context, _ *role.DynamicReq) (res *role.DynamicRes, err error) {
	return service.AdminMenu().GetMenuList(ctx, contexts.GetUserId(ctx))
}

// GetPermissions 获取指定角色权限
func (c *cRole) GetPermissions(ctx context.Context, req *role.GetPermissionsReq) (res *role.GetPermissionsRes, err error) {
	data, err := service.AdminRole().GetPermissions(ctx, &req.GetPermissionsInp)
	if err != nil {
		return
	}

	res = new(role.GetPermissionsRes)
	res.GetPermissionsModel = data
	return
}

// UpdatePermissions 修改角色菜单权限
func (c *cRole) UpdatePermissions(ctx context.Context, req *role.UpdatePermissionsReq) (res *role.UpdatePermissionsRes, err error) {
	err = service.AdminRole().UpdatePermissions(ctx, &req.UpdatePermissionsInp)
	return
}

// DataScopeSelect 获取数据权限选项
func (c *cRole) DataScopeSelect(_ context.Context, _ *role.DataScopeSelectReq) (res *role.DataScopeSelectRes, err error) {
	res = new(role.DataScopeSelectRes)
	res.List = consts.DataScopeSelect //service.AdminRole().DataScopeSelect()
	return
}

// DataScopeEdit 获取数据权限选项
func (c *cRole) DataScopeEdit(ctx context.Context, req *role.DataScopeEditReq) (res *role.DataScopeEditRes, err error) {
	err = service.AdminRole().DataScopeEdit(ctx, &req.DataScopeEditInp)
	return
}
