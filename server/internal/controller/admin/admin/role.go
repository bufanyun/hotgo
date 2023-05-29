// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/role"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var (
	Role = cRole{}
)

type cRole struct{}

// List 获取列表
func (c *cRole) List(ctx context.Context, req *role.ListReq) (res *role.ListRes, err error) {
	var in adminin.RoleListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.AdminRole().List(ctx, in)
	if err != nil {
		return
	}

	res = new(role.ListRes)
	res.RoleListModel = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Edit 修改角色
func (c *cRole) Edit(ctx context.Context, req *role.EditReq) (res *role.EditRes, err error) {
	var in adminin.RoleEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminRole().Edit(ctx, in)
	return
}

// Delete 删除
func (c *cRole) Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error) {
	var in adminin.RoleDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminRole().Delete(ctx, in)
	return
}

// Dynamic 动态路由
func (c *cRole) Dynamic(ctx context.Context, _ *role.DynamicReq) (res *role.DynamicRes, err error) {
	return service.AdminMenu().GetMenuList(ctx, contexts.GetUserId(ctx))
}

// GetPermissions 获取指定角色权限
func (c *cRole) GetPermissions(ctx context.Context, req *role.GetPermissionsReq) (res *role.GetPermissionsRes, err error) {
	var in adminin.GetPermissionsInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	data, err := service.AdminRole().GetPermissions(ctx, in)
	if err != nil {
		return
	}

	res = new(role.GetPermissionsRes)
	res.GetPermissionsModel = data
	return
}

// UpdatePermissions 修改角色菜单权限
func (c *cRole) UpdatePermissions(ctx context.Context, req *role.UpdatePermissionsReq) (res *role.UpdatePermissionsRes, err error) {
	var in adminin.UpdatePermissionsInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminRole().UpdatePermissions(ctx, in)
	return
}

// DataScopeSelect 获取数据权限选项
func (c *cRole) DataScopeSelect(_ context.Context, _ *role.DataScopeSelectReq) (res *role.DataScopeSelectRes, err error) {
	res = new(role.DataScopeSelectRes)
	res.List = service.AdminRole().DataScopeSelect()
	return
}

// DataScopeEdit 获取数据权限选项
func (c *cRole) DataScopeEdit(ctx context.Context, req *role.DataScopeEditReq) (res *role.DataScopeEditRes, err error) {
	var in adminin.DataScopeEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	in.CustomDept = req.CustomDept
	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminRole().DataScopeEdit(ctx, &in)
	return
}
