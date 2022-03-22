//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminController

import (
	"context"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/adminService"
	"github.com/gogf/gf/v2/util/gconv"
)

// 菜单
var Menu = menu{}

type menu struct{}

//
//  @Title  查询角色菜单列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) RoleList(ctx context.Context, req *adminForm.MenuRoleListReq) (*adminForm.MenuRoleListRes, error) {

	var in input.MenuRoleListInp
	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	data, err := adminService.Menu.RoleList(ctx, in)
	if err != nil {
		return nil, err
	}

	var res adminForm.MenuRoleListRes
	res.CheckedKeys = data.CheckedKeys
	res.Menus = data.Menus
	return &res, nil
}

//
//  @Title  查询菜单列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) SearchList(ctx context.Context, req *adminForm.MenuSearchListReq) (res *adminForm.MenuSearchListRes, err error) {

	res, err = adminService.Menu.SearchList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  最大排序
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) MaxSort(ctx context.Context, req *adminForm.MenuMaxSortReq) (res *adminForm.MenuMaxSortRes, err error) {

	res, err = adminService.Menu.MaxSort(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  菜单名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) NameUnique(ctx context.Context, req *adminForm.MenuNameUniqueReq) (res *adminForm.MenuNameUniqueRes, err error) {

	res, err = adminService.Menu.NameUnique(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  菜单编码是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) CodeUnique(ctx context.Context, req *adminForm.MenuCodeUniqueReq) (res *adminForm.MenuCodeUniqueRes, err error) {

	res, err = adminService.Menu.CodeUnique(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  删除
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) Delete(ctx context.Context, req *adminForm.MenuDeleteReq) (res *adminForm.MenuDeleteRes, err error) {

	if err = adminService.Menu.Delete(ctx, req); err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  修改/新增
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) Edit(ctx context.Context, req *adminForm.MenuEditReq) (res *adminForm.MenuEditRes, err error) {

	if err = adminService.Menu.Edit(ctx, req); err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  获取信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) View(ctx context.Context, req *adminForm.MenuViewReq) (res *adminForm.MenuViewRes, err error) {

	res, err = adminService.Menu.View(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  获取列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *menu) List(ctx context.Context, req *adminForm.MenuListReq) (res *adminForm.MenuListRes, err error) {

	res, err = adminService.Menu.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
