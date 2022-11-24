// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/role"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
)

var (
	Role = cRole{}
)

type cRole struct{}

// RoleMemberList 获取角色下的会员列表
func (c *cRole) RoleMemberList(ctx context.Context, req *role.MemberListReq) (*role.MemberListRes, error) {

	var in adminin.RoleMemberListInp
	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	list, totalCount, err := service.AdminMember().RoleMemberList(ctx, in)
	if err != nil {
		return nil, err
	}

	var res role.MemberListRes
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.PerPage = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// List 获取列表
func (c *cRole) List(ctx context.Context, req *role.ListReq) (*role.ListRes, error) {

	list, totalCount, err := service.AdminRole().List(ctx, adminin.RoleListInp{
		Page:    req.Page,
		PerPage: req.PerPage,
	})
	if err != nil {
		return nil, err
	}

	var res role.ListRes
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.PerPage = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// Edit 修改角色
func (c *cRole) Edit(ctx context.Context, req *role.EditReq) (res *role.EditRes, err error) {
	err = service.AdminRole().Edit(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete 删除
func (c *cRole) Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error) {

	if err = service.AdminRole().Delete(ctx, req); err != nil {
		return nil, err
	}
	return res, nil
}

// Dynamic 动态路由
func (c *cRole) Dynamic(ctx context.Context, req *role.DynamicReq) (res role.DynamicRes, err error) {

	res, err = service.AdminMenu().GetMenuList(ctx, contexts.GetUserId(ctx))
	if err != nil {
		return res, err
	}
	return res, nil
}

// GetPermissions 获取指定角色权限
func (c *cRole) GetPermissions(ctx context.Context, req *role.GetPermissionsReq) (res *role.GetPermissionsRes, err error) {
	MenuIds, err := service.AdminRole().GetPermissions(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &role.GetPermissionsRes{
		MenuIds: []int64{},
	}
	if MenuIds != nil {
		res.MenuIds = MenuIds
	}

	return res, nil
}

// UpdatePermissions 修改角色菜单权限
func (c *cRole) UpdatePermissions(ctx context.Context, req *role.UpdatePermissionsReq) (res *role.UpdatePermissionsRes, err error) {
	err = service.AdminRole().UpdatePermissions(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
