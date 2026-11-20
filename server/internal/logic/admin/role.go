// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
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
	"hotgo/utility/validate"
	"sort"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
		err  error
	)

	if user == nil {
		g.Log().Info(ctx, "admin Verify user = nil")
		return false
	}

	if service.AdminMember().VerifySuperId(ctx, user.Id) {
		return true
	}

	ok, err := casbin.Enforcer.Enforce(user.RoleKey, path, method)
	if err != nil {
		g.Log().Infof(ctx, "admin Verify Enforce err:%+v", err)
		return false
	}
	return ok
}

// List 获取列表
func (s *sAdminRole) List(ctx context.Context, in *adminin.RoleListInp) (res *adminin.RoleListModel, totalCount int, err error) {
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
	r, err := dao.AdminRole.Ctx(ctx).Fields(dao.AdminRole.Columns().Name).WherePri(id).OrderDesc(dao.AdminRole.Columns().Id).Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return r.String(), nil
}

// GetMemberList 获取指定用户的岗位列表
func (s *sAdminRole) GetMemberList(ctx context.Context, id int64) (list []*adminin.RoleListModel, err error) {
	if err = dao.AdminRole.Ctx(ctx).WherePri(id).OrderDesc(dao.AdminRole.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// GetPermissions 更改角色菜单权限
func (s *sAdminRole) GetPermissions(ctx context.Context, in *adminin.GetPermissionsInp) (res *adminin.GetPermissionsModel, err error) {
	if err = s.verifyRoleScope(ctx, in.RoleId); err != nil {
		return
	}

	values, err := dao.AdminRoleMenu.Ctx(ctx).Fields(dao.AdminRoleMenu.Columns().MenuId).Where(dao.AdminRoleMenu.Columns().RoleId, in.RoleId).Array()
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
func (s *sAdminRole) UpdatePermissions(ctx context.Context, in *adminin.UpdatePermissionsInp) (err error) {
	in.MenuIds = convert.UniqueSlice(in.MenuIds)
	isSuper := service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx))
	if !isSuper {
		if err = s.VerifyRoleId(ctx, in.RoleId); err != nil {
			return
		}
	}

	err = dao.AdminRoleMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if !isSuper && len(in.MenuIds) > 0 {
			// 加锁防止校验后、写入前目标菜单被并发改动
			existingQuery := dao.AdminRoleMenu.Ctx(ctx).
				Fields(dao.AdminRoleMenu.Columns().MenuId).
				Where(dao.AdminRoleMenu.Columns().RoleId, in.RoleId)
			// SQLite 不支持 FOR UPDATE；读后库被改的话写入会报 BUSY 失败，不会带着过期校验写入
			dbType := dao.AdminRoleMenu.DB().GetConfig().Type
			if dbType == consts.DBMysql || dbType == consts.DBPgsql {
				existingQuery = existingQuery.LockUpdate()
			}
			existingValues, err := existingQuery.Array()
			if err != nil {
				return gerror.Wrap(err, consts.ErrorORM)
			}

			values, err := dao.AdminRoleMenu.Ctx(ctx).
				Fields(dao.AdminRoleMenu.Columns().MenuId).
				Where(dao.AdminRoleMenu.Columns().RoleId, contexts.GetRoleId(ctx)).
				Array()
			if err != nil {
				return gerror.Wrap(err, consts.ErrorORM)
			}

			existing := make(map[int64]struct{}, len(existingValues))
			for _, v := range existingValues {
				existing[v.Int64()] = struct{}{}
			}
			owned := make(map[int64]struct{}, len(values))
			for _, v := range values {
				owned[v.Int64()] = struct{}{}
			}

			// 目标原有的菜单可能超出调用者权限，原样保留不算授予
			for _, menuId := range in.MenuIds {
				if _, ok := existing[menuId]; ok {
					continue
				}
				if _, ok := owned[menuId]; !ok {
					return gerror.New("不能授予自己没有的权限")
				}
			}
		}

		if _, err = dao.AdminRoleMenu.Ctx(ctx).Where("role_id", in.RoleId).Delete(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		if in.MenuIds = convert.UniqueSlice(in.MenuIds); len(in.MenuIds) == 0 {
			return
		}

		list := make(g.List, 0, len(in.MenuIds))
		for _, v := range in.MenuIds {
			list = append(list, g.Map{
				"role_id": in.RoleId,
				"menu_id": v,
			})
		}

		if _, err = dao.AdminRoleMenu.Ctx(ctx).Data(list).OmitEmptyData().Insert(); err != nil {
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

func (s *sAdminRole) Edit(ctx context.Context, in *adminin.RoleEditInp) (err error) {
	if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		if in.Id > 0 {
			if err = s.VerifyRoleId(ctx, in.Id); err != nil {
				return
			}
		}

		if in.Pid != contexts.GetRoleId(ctx) {
			if err = s.VerifyRoleId(ctx, in.Pid); err != nil {
				return
			}
		}

		if err = s.keepDataScope(ctx, in); err != nil {
			return
		}
	}

	if err = hgorm.IsUnique(ctx, &dao.AdminRole, g.Map{dao.AdminRole.Columns().Name: in.Name}, "名称已存在", in.Id); err != nil {
		return
	}

	if err = hgorm.IsUnique(ctx, &dao.AdminRole, g.Map{dao.AdminRole.Columns().Key: in.Key}, "编码已存在", in.Id); err != nil {
		return
	}

	if in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, &dao.AdminRole, in.Pid); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
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
		if err != nil {
			return
		}
		// 角色状态影响权限，事务提交后再刷新
		if err = casbin.Refresh(ctx); err != nil {
			return gerror.Wrap(err, "数据已保存，但权限刷新失败，请稍后重试保存或联系管理员")
		}
		return
	}

	// 新增
	if _, err = dao.AdminRole.Ctx(ctx).Fields(adminin.RoleInsertFields{}).Data(in).OmitEmptyData().Insert(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

func updateRoleChildrenTree(ctx context.Context, _id int64, _level int, _tree string) (err error) {
	var list []*entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Where("pid", _id).Scan(&list); err != nil {
		return
	}
	for _, child := range list {
		child.Level = _level + 1
		child.Tree = tree.GenLabel(_tree, child.Pid)

		if _, err = dao.AdminRole.Ctx(ctx).Where("id", child.Id).Data("level", child.Level, "tree", child.Tree).Update(); err != nil {
			return
		}

		if err = updateRoleChildrenTree(ctx, child.Id, child.Level, child.Tree); err != nil {
			return
		}
	}
	return
}

func (s *sAdminRole) Delete(ctx context.Context, in *adminin.RoleDeleteInp) (err error) {
	if err = s.verifyRoleScope(ctx, in.Id); err != nil {
		return
	}

	var models *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		return gerror.New("数据不存在或已删除！")
	}

	if models.Key == consts.SuperRoleKey {
		return gerror.New("超管角色禁止删除！")
	}

	exist, err := dao.AdminRole.Ctx(ctx).Where("pid", models.Id).Exist()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if exist {
		return gerror.New("请先删除该角色下得所有子级！")
	}

	exist, err = dao.AdminMember.Ctx(ctx).Where("role_id", models.Id).Exist()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if exist {
		return gerror.New("请先删除该角色下得所有用户账号！")
	}

	err = dao.AdminRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err = dao.AdminRoleMenu.Ctx(ctx).Where(dao.AdminRoleMenu.Columns().RoleId, in.Id).Delete(); err != nil {
			return gerror.Wrap(err, consts.ErrorORM)
		}
		if _, err = dao.AdminRole.Ctx(ctx).Where(dao.AdminRole.Columns().Id, in.Id).Delete(); err != nil {
			return gerror.Wrap(err, consts.ErrorORM)
		}
		return nil
	})
	if err != nil {
		return
	}
	// 角色删除后其权限不该继续生效，事务提交后再刷新
	if err = casbin.Refresh(ctx); err != nil {
		return gerror.Wrap(err, "数据已删除，但权限刷新失败，请联系管理员")
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
	if err = s.verifyRoleScope(ctx, in.Id); err != nil {
		return
	}

	if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		all, err := s.callerDataScopeAll(ctx)
		if err != nil {
			return err
		}
		if !all {
			return gerror.New("只有数据范围为全部权限的角色才能修改下级的数据范围")
		}
	}

	var models *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		return gerror.New("角色不存在")
	}

	if models.Key == consts.SuperRoleKey {
		return gerror.New("超管角色拥有全部权限，无需修改！")
	}

	if in.DataScope == consts.RoleDataDeptCustom && len(convert.UniqueSlice(in.CustomDept)) == 0 {
		return gerror.New("自定义权限必须配置自定义部门！")
	}

	models.DataScope = in.DataScope
	models.CustomDept = gjson.New(convert.UniqueSlice(in.CustomDept))

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

func (s *sAdminRole) callerDataScopeAll(ctx context.Context) (all bool, err error) {
	var role *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).
		Fields(dao.AdminRole.Columns().DataScope).
		Where(dao.AdminRole.Columns().Id, contexts.GetRoleId(ctx)).
		Scan(&role); err != nil {
		return false, gerror.Wrap(err, consts.ErrorORM)
	}
	return role != nil && role.DataScope == consts.RoleDataAll, nil
}

// keepDataScope 防止没有全部数据权限的调用者借编辑角色放大数据范围
func (s *sAdminRole) keepDataScope(ctx context.Context, in *adminin.RoleEditInp) (err error) {
	all, err := s.callerDataScopeAll(ctx)
	if err != nil {
		return err
	}
	if all {
		return nil
	}

	roleId := in.Pid
	if in.Id > 0 {
		roleId = in.Id
	}

	var source *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).
		Fields(dao.AdminRole.Columns().DataScope, dao.AdminRole.Columns().CustomDept).
		Where(dao.AdminRole.Columns().Id, roleId).
		Scan(&source); err != nil {
		return gerror.Wrap(err, consts.ErrorORM)
	}
	if source == nil {
		return gerror.New("角色不存在")
	}

	in.DataScope = source.DataScope
	in.CustomDept = source.CustomDept
	return nil
}

// verifyRoleScope 防止非超管改自己或上级角色给自己提权
func (s *sAdminRole) verifyRoleScope(ctx context.Context, roleId int64) (err error) {
	if service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		return nil
	}
	return s.VerifyRoleId(ctx, roleId)
}

// VerifyRoleId 验证角色ID
func (s *sAdminRole) VerifyRoleId(ctx context.Context, id int64) (err error) {
	mb := contexts.GetUser(ctx)
	if mb == nil {
		err = gerror.New("用户信息获取失败！")
		return
	}

	ids, err := s.GetSubRoleIds(ctx, mb.RoleId, service.AdminMember().VerifySuperId(ctx, mb.Id))
	if err != nil {
		err = gerror.New("验证角色信息失败！")
		return
	}

	if !validate.InSlice(ids, id) {
		err = gerror.New("角色ID是无效的")
		return
	}
	return
}

// GetSubRoleIds 获取所有下级角色ID
func (s *sAdminRole) GetSubRoleIds(ctx context.Context, roleId int64, isSuper bool) (ids []int64, err error) {
	mod := dao.AdminRole.Ctx(ctx).Fields(dao.AdminRole.Columns().Id)
	if !isSuper {
		mod = mod.WhereNot(dao.AdminRole.Columns().Id, roleId).WhereLike(dao.AdminRole.Columns().Tree, "%"+tree.GetIdLabel(roleId)+"%")
	}

	columns, err := mod.Array()
	if err != nil {
		return nil, err
	}

	ids = g.NewVar(columns).Int64s()
	return
}
