// Package handler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package handler

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/utility/convert"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

// FilterAuth 过滤数据权限
// 通过上下文中的用户角色权限和表中是否含有需要过滤的字段附加查询条件
func FilterAuth(m *gdb.Model) *gdb.Model {
	var (
		needAuth    bool
		filterField string
		fields      = convert.EscapeFieldsToSlice(m.GetFieldsStr())
	)

	// 优先级：created_by > member_id
	if gstr.InArray(fields, "created_by") {
		needAuth = true
		filterField = "created_by"
	}

	if !needAuth && gstr.InArray(fields, "member_id") {
		needAuth = true
		filterField = "member_id"
	}

	if !needAuth {
		return m
	}
	return m.Handler(FilterAuthWithField(filterField))
}

// FilterAuthWithField 过滤数据权限，设置指定字段
func FilterAuthWithField(filterField string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		var (
			role *entity.AdminRole
			ctx  = m.GetCtx()
			co   = contexts.Get(ctx)
		)

		if co == nil || co.User == nil {
			return m
		}

		err := dao.AdminRole.Ctx(ctx).Where(dao.AdminRole.Columns().Id, co.User.RoleId).Scan(&role)
		if err != nil {
			g.Log().Panicf(ctx, "failed to role information err:%+v", err)
		}

		if role == nil {
			g.Log().Panic(ctx, "failed to role information roleModel == nil")
		}

		// 超管拥有全部权限
		if role.Key == consts.SuperRoleKey {
			return m
		}

		getDeptIds := func(in interface{}) []gdb.Value {
			ds, err := dao.AdminMember.Ctx(ctx).Fields(dao.AdminMember.Columns().Id).Where(dao.AdminMember.Columns().DeptId, in).Array()
			if err != nil {
				g.Log().Panic(ctx, "failed to get member dept data")
			}
			return ds
		}

		switch role.DataScope {
		case consts.RoleDataAll: // 全部权限
			// ...
		case consts.RoleDataNowDept: // 当前部门
			m = m.WhereIn(filterField, getDeptIds(co.User.DeptId))
		case consts.RoleDataDeptAndSub: // 当前部门及以下部门ds
			m = m.WhereIn(filterField, getDeptIds(GetDeptAndSub(ctx, co.User.DeptId)))
		case consts.RoleDataDeptCustom: // 自定义部门
			m = m.WhereIn(filterField, getDeptIds(role.CustomDept.Var().Ints()))
		case consts.RoleDataSelf: // 仅自己
			m = m.Where(filterField, co.User.Id)
		case consts.RoleDataSelfAndSub: // 自己和直属下级
			m = m.WhereIn(filterField, GetSelfAndSub(ctx, co.User.Id))
		case consts.RoleDataSelfAndAllSub: // 自己和全部下级
			m = m.WhereIn(filterField, GetSelfAndAllSub(ctx, co.User.Id))
		default:
			g.Log().Panic(ctx, "dataScope is not registered")
		}
		return m
	}
}

// GetDeptAndSub 获取指定部门的所有下级，含本部门
func GetDeptAndSub(ctx context.Context, deptId int64) (ids []int64) {
	array, err := dao.AdminDept.Ctx(ctx).
		WhereLike(dao.AdminDept.Columns().Tree, "%"+tree.GetIdLabel(deptId)+"%").
		Fields(dao.AdminDept.Columns().Id).
		Array()
	if err != nil {
		g.Log().Panicf(ctx, "GetDeptAndSub err:%+v", err)
		return
	}

	for _, v := range array {
		ids = append(ids, v.Int64())
	}

	ids = append(ids, deptId)
	return
}

// GetSelfAndSub 获取直属下级，包含自己
func GetSelfAndSub(ctx context.Context, memberId int64) (ids []int64) {
	array, err := dao.AdminMember.Ctx(ctx).
		Where(dao.AdminMember.Columns().Pid, memberId).
		Fields(dao.AdminMember.Columns().Id).
		Array()
	if err != nil {
		g.Log().Panicf(ctx, "GetSelfAndSub err:%+v", err)
		return
	}

	for _, v := range array {
		ids = append(ids, v.Int64())
	}

	ids = append(ids, memberId)
	return
}

// GetSelfAndAllSub 获取全部下级，包含自己
func GetSelfAndAllSub(ctx context.Context, memberId int64) (ids []int64) {
	array, err := dao.AdminMember.Ctx(ctx).
		WhereLike(dao.AdminMember.Columns().Tree, "%"+tree.GetIdLabel(memberId)+"%").
		Fields(dao.AdminMember.Columns().Id).
		Array()
	if err != nil {
		g.Log().Panicf(ctx, "GetSelfAndAllSub err:%+v", err)
		return
	}

	for _, v := range array {
		ids = append(ids, v.Int64())
	}

	ids = append(ids, memberId)
	return
}
