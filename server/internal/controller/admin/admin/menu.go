// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"hotgo/api/admin/menu"
	"hotgo/internal/service"
)

// Menu 菜单
var (
	Menu = cMenu{}
)

type cMenu struct{}

// Delete 删除
func (c *cMenu) Delete(ctx context.Context, req *menu.DeleteReq) (res *menu.DeleteRes, err error) {
	err = service.AdminMenu().Delete(ctx, &req.MenuDeleteInp)
	return
}

// Edit 更新
func (c *cMenu) Edit(ctx context.Context, req *menu.EditReq) (res *menu.EditRes, err error) {
	err = service.AdminMenu().Edit(ctx, &req.MenuEditInp)
	return
}

// List 获取列表
func (c *cMenu) List(ctx context.Context, req *menu.ListReq) (res menu.ListRes, err error) {
	res.MenuListModel, err = service.AdminMenu().List(ctx, &req.MenuListInp)
	return
}
