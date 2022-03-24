//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminController

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/adminService"
)

// 角色
var Role = role{}

type role struct{}

//
//  @Title  获取角色下的会员列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *role) RoleMemberList(ctx context.Context, req *adminForm.RoleMemberListReq) (*adminForm.RoleMemberListRes, error) {

	var in input.AdminRoleMemberListInp
	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	list, totalCount, err := adminService.Member.RoleMemberList(ctx, in)
	if err != nil {
		return nil, err
	}

	var res adminForm.RoleMemberListRes
	res.List = list
	res.TotalCount = totalCount
	res.Limit = req.Page
	res.Limit = req.Limit

	return &res, nil
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
func (controller *role) List(ctx context.Context, req *adminForm.RoleListReq) (*adminForm.RoleListRes, error) {

	list, totalCount, err := adminService.Role.List(ctx, input.AdminRoleListInp{
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}

	var res adminForm.RoleListRes
	res.List = list
	res.TotalCount = totalCount
	res.Limit = req.Page
	res.Limit = req.Limit

	return &res, nil
}

//
//  @Title  动态路由
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *role) Dynamic(ctx context.Context, req *adminForm.RoleDynamicReq) (res *adminForm.RoleDynamicRes, err error) {

	res, err = adminService.Menu.GetMenuList(ctx, com.Context.GetUserId(ctx))
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  修改角色菜单权限
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *role) Edit(ctx context.Context, req *adminForm.RoleMenuEditReq) (res *adminForm.RoleMenuEditRes, err error) {
	err = adminService.Role.EditRoleMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
