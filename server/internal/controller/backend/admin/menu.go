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
	"hotgo/api/backend/menu"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
)

// Menu 菜单
var (
	Menu = cMenu{}
)

type cMenu struct{}

// RoleList 查询角色菜单列表
func (c *cMenu) RoleList(ctx context.Context, req *menu.RoleListReq) (res *menu.RoleListRes, err error) {
	var in adminin.MenuRoleListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	data, err := service.AdminMenu().RoleList(ctx, in)
	if err != nil {
		return
	}

	res = new(menu.RoleListRes)
	res.CheckedKeys = data.CheckedKeys
	res.Menus = data.Menus
	return
}

// SearchList 查询菜单列表
func (c *cMenu) SearchList(ctx context.Context, req *menu.SearchListReq) (res *menu.SearchListRes, err error) {
	res, err = service.AdminMenu().SearchList(ctx, req)
	return
}

// MaxSort 最大排序
func (c *cMenu) MaxSort(ctx context.Context, req *menu.MaxSortReq) (res *menu.MaxSortRes, err error) {
	res, err = service.AdminMenu().MaxSort(ctx, req)
	return
}

// NameUnique 菜单名称是否唯一
func (c *cMenu) NameUnique(ctx context.Context, req *menu.NameUniqueReq) (res *menu.NameUniqueRes, err error) {
	res, err = service.AdminMenu().NameUnique(ctx, req)
	return
}

// CodeUnique 菜单编码是否唯一
func (c *cMenu) CodeUnique(ctx context.Context, req *menu.CodeUniqueReq) (res *menu.CodeUniqueRes, err error) {
	res, err = service.AdminMenu().CodeUnique(ctx, req)
	return
}

// Delete 删除
func (c *cMenu) Delete(ctx context.Context, req *menu.DeleteReq) (res *menu.DeleteRes, err error) {
	err = service.AdminMenu().Delete(ctx, req)
	return
}

// Edit 更新
func (c *cMenu) Edit(ctx context.Context, req *menu.EditReq) (res *menu.EditRes, err error) {
	err = service.AdminMenu().Edit(ctx, req)
	return
}

// View 获取信息
func (c *cMenu) View(ctx context.Context, req *menu.ViewReq) (res *menu.ViewRes, err error) {
	res, err = service.AdminMenu().View(ctx, req)
	return
}

// List 获取列表
func (c *cMenu) List(ctx context.Context, req *menu.ListReq) (res menu.ListRes, err error) {
	res.List, err = service.AdminMenu().List(ctx, req)
	return
}
