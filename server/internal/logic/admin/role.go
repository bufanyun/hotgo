// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/admin/role"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/auth"
	"hotgo/utility/convert"
	"hotgo/utility/tree"
	"sort"
)

type sAdminRole struct{}

func NewAdminRole() *sAdminRole {
	return &sAdminRole{}
}

func init() {
	service.RegisterAdminRole(NewAdminRole())
}

// Verify 验证权限
func (s *sAdminRole) Verify(ctx context.Context, path, method string) bool {
	if auth.IsExceptAuth(ctx, path) {
		return true
	}

	var (
		user         = contexts.Get(ctx).User
		superRoleKey = g.Cfg().MustGet(ctx, "hotgo.admin.superRoleKey")
		err          error
	)

	if user == nil {
		g.Log().Info(ctx, "admin Verify user = nil")
		return false
	}

	if service.AdminMember().VerifySuperId(ctx, user.Id) || user.RoleKey == superRoleKey.String() {
		return true
	}
	ok, err := casbin.Enforcer.Enforce(user.RoleKey, path, method)
	if err != nil {
		g.Log().Infof(ctx, "admin Verify Enforce  err:%+v", err)
		return false
	}

	return ok
}

// List 获取列表
func (s *sAdminRole) List(ctx context.Context, in adminin.RoleListInp) (res *adminin.RoleListModel, totalCount int, err error) {
	var (
		mod    = dao.AdminRole.Ctx(ctx)
		models []*entity.AdminRole
		pid    int64 = 0
	)

	// 非超管只获取下级角色
	if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		pid = contexts.GetRoleId(ctx)
		mod = mod.WhereLike(dao.AdminRole.Columns().Tree, "%"+tree.GetIdLabel(pid)+"%")
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	res = new(adminin.RoleListModel)
	res.List = s.treeList(pid, models)
	return
}

// GetName 获取指定角色的名称
func (s *sAdminRole) GetName(ctx context.Context, RoleId int64) (name string, err error) {
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

// GetMemberList 获取指定用户的岗位列表
func (s *sAdminRole) GetMemberList(ctx context.Context, RoleId int64) (list []*adminin.RoleListModel, err error) {
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

// GetPermissions 更改角色菜单权限
func (s *sAdminRole) GetPermissions(ctx context.Context, reqInfo *role.GetPermissionsReq) (MenuIds []int64, err error) {
	values, err := dao.AdminRoleMenu.Ctx(ctx).
		Fields("menu_id").
		Where("role_id", reqInfo.RoleId).
		Array()
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return
	}

	for i := 0; i < len(values); i++ {
		MenuIds = append(MenuIds, values[i].Int64())
	}
	return
}

// UpdatePermissions 更改角色菜单权限
func (s *sAdminRole) UpdatePermissions(ctx context.Context, reqInfo *role.UpdatePermissionsReq) (err error) {
	err = dao.AdminRoleMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.AdminRoleMenu.Ctx(ctx).Where("role_id", reqInfo.RoleId).Delete()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		if len(reqInfo.MenuIds) == 0 {
			return nil
		}

		// 去重
		reqInfo.MenuIds = convert.UniqueSliceInt64(reqInfo.MenuIds)

		addMap := make(g.List, 0, len(reqInfo.MenuIds))
		for _, v := range reqInfo.MenuIds {
			addMap = append(addMap, g.Map{
				"role_id": reqInfo.RoleId,
				"menu_id": v,
			})
		}

		if _, err = dao.AdminRoleMenu.Ctx(ctx).Data(addMap).Insert(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return
	})

	if err != nil {
		return
	}

	return casbin.Refresh(ctx)
}

func (s *sAdminRole) Edit(ctx context.Context, in *role.EditReq) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}

	if in.Key == "" {
		err = gerror.New("编码不能为空")
		return
	}

	uniqueName, err := dao.AdminRole.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	if !uniqueName {
		err = gerror.New("名称已存在")
		return
	}

	uniqueCode, err := dao.AdminRole.IsUniqueCode(ctx, in.Id, in.Key)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	if !uniqueCode {
		err = gerror.New("编码已存在")
		return
	}

	if in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, dao.AdminRole, in.Pid); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		_, err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		return
	}

	// 新增
	_, err = dao.AdminRole.Ctx(ctx).Data(in).Insert()
	return
}

func (s *sAdminRole) Delete(ctx context.Context, in *role.DeleteReq) (err error) {
	if in.Id <= 0 {
		return gerror.New("ID不正确！")
	}

	var models *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		return gerror.New("数据不存在或已删除！")
	}

	pidExist, err := dao.AdminRole.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !pidExist.IsEmpty() {
		return gerror.New("请先删除该角色下得所有子级！")
	}

	_, err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

func (s *sAdminRole) DataScopeSelect(ctx context.Context) (res form.Selects) {
	for k, v := range consts.RoleDataNameMap {
		res = append(res, &form.Select{
			Value: k,
			Name:  v,
			Label: v,
		})
	}
	sort.Sort(res)
	return res
}

func (s *sAdminRole) DataScopeEdit(ctx context.Context, in *adminin.DataScopeEditInp) (err error) {
	if in.Id <= 0 {
		return gerror.New("角色ID不正确！")
	}

	var (
		models       *entity.AdminRole
		superRoleKey = g.Cfg().MustGet(ctx, "hotgo.admin.superRoleKey")
	)

	if err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		return gerror.New("角色不存在")
	}

	if models.Key == superRoleKey.String() {
		return gerror.New("超管角色拥有全部权限，无需修改！")
	}

	if in.DataScope == consts.RoleDataDeptCustom && len(convert.UniqueSliceInt64(in.CustomDept)) == 0 {
		return gerror.New("自定义权限必须配置自定义部门！")
	}

	models.DataScope = in.DataScope
	models.CustomDept = gjson.New(convert.UniqueSliceInt64(in.CustomDept))

	_, err = dao.AdminRole.Ctx(ctx).
		Fields(dao.AdminRole.Columns().DataScope, dao.AdminRole.Columns().CustomDept).
		Where("id", in.Id).
		Data(models).
		Update()

	return
}

// treeList 树状列表
func (s *sAdminRole) treeList(pid int64, nodes []*entity.AdminRole) (list []*adminin.RoleTree) {
	list = make([]*adminin.RoleTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(adminin.RoleTree)
			item.AdminRole = *v
			item.Label = v.Name
			item.Value = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}
