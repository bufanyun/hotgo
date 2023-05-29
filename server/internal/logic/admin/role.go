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
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
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
	var (
		user = contexts.Get(ctx).User
		sk   = g.Cfg().MustGet(ctx, "hotgo.admin.superRoleKey")
		err  error
	)

	if user == nil {
		g.Log().Info(ctx, "admin Verify user = nil")
		return false
	}

	if service.AdminMember().VerifySuperId(ctx, user.Id) || user.RoleKey == sk.String() {
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
func (s *sAdminRole) GetName(ctx context.Context, id int64) (name string, err error) {
	r, err := dao.AdminRole.Ctx(ctx).
		Fields("name").
		WherePri(id).
		Order("id desc").
		Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	return r.String(), nil
}

// GetMemberList 获取指定用户的岗位列表
func (s *sAdminRole) GetMemberList(ctx context.Context, id int64) (list []*adminin.RoleListModel, err error) {
	if err = dao.AdminRole.Ctx(ctx).WherePri(id).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// GetPermissions 更改角色菜单权限
func (s *sAdminRole) GetPermissions(ctx context.Context, in adminin.GetPermissionsInp) (res *adminin.GetPermissionsModel, err error) {
	values, err := dao.AdminRoleMenu.Ctx(ctx).
		Fields("menu_id").
		Where("role_id", in.RoleId).
		Array()
	if err != nil {
		return
	}

	if len(values) == 0 {
		return
	}

	res = new(adminin.GetPermissionsModel)
	for i := 0; i < len(values); i++ {
		res.MenuIds = append(res.MenuIds, values[i].Int64())
	}
	return
}

// UpdatePermissions 更改角色菜单权限
func (s *sAdminRole) UpdatePermissions(ctx context.Context, in adminin.UpdatePermissionsInp) (err error) {
	err = dao.AdminRoleMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.AdminRoleMenu.Ctx(ctx).Where("role_id", in.RoleId).Delete(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		if len(in.MenuIds) == 0 {
			return
		}

		// 去重
		in.MenuIds = convert.UniqueSliceInt64(in.MenuIds)

		list := make(g.List, 0, len(in.MenuIds))
		for _, v := range in.MenuIds {
			list = append(list, g.Map{
				"role_id": in.RoleId,
				"menu_id": v,
			})
		}

		if _, err = dao.AdminRoleMenu.Ctx(ctx).Data(list).Insert(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		return
	})

	if err != nil {
		return
	}

	return casbin.Refresh(ctx)
}

func (s *sAdminRole) Edit(ctx context.Context, in adminin.RoleEditInp) (err error) {
	if err = hgorm.IsUnique(ctx, dao.AdminRole, g.Map{dao.AdminRole.Columns().Name: in.Name}, "名称已存在", in.Id); err != nil {
		return
	}

	if err = hgorm.IsUnique(ctx, dao.AdminRole, g.Map{dao.AdminRole.Columns().Key: in.Key}, "编码已存在", in.Id); err != nil {
		return
	}

	if in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, dao.AdminRole, in.Pid); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		// 获取父级tree
		var pTree gdb.Value
		pTree, err = dao.AdminRole.Ctx(ctx).Where("id", in.Pid).Fields("tree").Value()
		if err != nil {
			return
		}
		in.Tree = tree.GenLabel(pTree.String(), in.Id)

		err = dao.AdminRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 更新数据
			_, err = dao.AdminRole.Ctx(ctx).Fields(adminin.RoleUpdateFields{}).WherePri(in.Id).Data(in).Update()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return err
			}

			// 如果当前角色有子级,更新子级tree关系树
			return updateRoleChildrenTree(ctx, in.Id, in.Level, in.Tree)
		})
		return
	}

	// 新增
	if _, err = dao.AdminRole.Ctx(ctx).Fields(adminin.RoleInsertFields{}).Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

func updateRoleChildrenTree(ctx context.Context, _id int64, _level int, _tree string) (err error) {
	var list []*entity.AdminDept
	err = dao.AdminRole.Ctx(ctx).Where("pid", _id).Scan(&list)
	if err != nil {
		return
	}
	for _, child := range list {
		child.Level = _level + 1
		child.Tree = tree.GenLabel(_tree, child.Id)

		_, err = dao.AdminRole.Ctx(ctx).Where("id", child.Id).Data("level", child.Level, "tree", child.Tree).Update()
		if err != nil {
			return err
		}
		err = updateRoleChildrenTree(ctx, child.Id, child.Level, child.Tree)
		if err != nil {
			return
		}
	}
	return
}

func (s *sAdminRole) Delete(ctx context.Context, in adminin.RoleDeleteInp) (err error) {
	var models *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		return gerror.New("数据不存在或已删除！")
	}

	has, err := dao.AdminRole.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if !has.IsEmpty() {
		return gerror.New("请先删除该角色下得所有子级！")
	}

	if _, err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

func (s *sAdminRole) DataScopeSelect() (res form.Selects) {
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
	var (
		models *entity.AdminRole
		sk     = g.Cfg().MustGet(ctx, "hotgo.admin.superRoleKey")
	)

	if err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		return gerror.New("角色不存在")
	}

	if models.Key == sk.String() {
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

// treeList 角色树列表
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
