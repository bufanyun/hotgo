//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminService

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/bufanyun/hotgo/app/service/internal/dto"
	"github.com/bufanyun/hotgo/app/utils"
)

var Role = new(role)

type role struct{}

//
//  @Title  验证权限
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   member_id
//  @Param   path
//  @Return  bool
//
func (service *role) Verify(ctx context.Context, member_id int, path string) bool {
	var (
		err error
	)

	if utils.Auth.IsExceptAuth(ctx, path) {
		return true
	}

	menu := Menu.WhereScan(ctx, dto.AdminMenu{
		Path:   path,
		Status: consts.StatusEnabled,
	})

	if menu == nil {
		err = gerror.New(consts.ErrorNotData)
		return false
	}

	g.Log().Print(ctx, "menu:", menu)
	g.Log().Print(ctx, "err:", err)

	return true
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
func (service *role) List(ctx context.Context, in input.AdminRoleListInp) (list []*input.AdminRoleListModel, totalCount int, err error) {

	mod := dao.AdminRole.Ctx(ctx)

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	err = mod.Page(in.Page, in.Limit).Order("id asc").Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	return list, totalCount, err
}

//
//  @Title  获取指定角色的名称
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   RoleId
//  @Return  name
//  @Return  err
//
func (service *role) GetName(ctx context.Context, RoleId int64) (name string, err error) {
	roleName, err := dao.AdminRole.Ctx(ctx).
		Fields("name").
		Where("id", RoleId).
		Order("id desc").
		Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return name, err
	}

	return roleName.String(), nil
}

//
//  @Title  获取指定会员的岗位列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *role) GetMemberList(ctx context.Context, RoleId int64) (list []*input.AdminRoleListModel, err error) {

	err = dao.AdminRole.Ctx(ctx).
		Where("id", RoleId).
		Order("id desc").
		Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	return list, err
}

//
//  @Title  更改角色菜单权限
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  err
//
func (service *role) EditRoleMenu(ctx context.Context, reqInfo *adminForm.RoleMenuEditReq) error {
	return dao.AdminRoleMenu.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		_, err = dao.AdminRoleMenu.Ctx(ctx).Where("role_id", reqInfo.RoleId).Delete()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}
		addMap := make(g.List, 0, len(reqInfo.MenuIds))
		for _, v := range reqInfo.MenuIds {
			addMap = append(addMap, g.Map{
				"role_id": reqInfo.RoleId,
				"menu_id": v,
			})
		}
		_, err = dao.AdminRoleMenu.Ctx(ctx).Data(addMap).Insert()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}
		return nil
	})
}
