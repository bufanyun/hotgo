// Package hgorm
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package hgorm

// 预处理
import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
)

// HandlerFilterAuth 过滤数据权限
// 通过上下文中的用户角色权限和表中是否含有需要过滤的字段附加查询条件
func HandlerFilterAuth(m *gdb.Model) *gdb.Model {
	var (
		needAuth    bool
		filterField string
		roleModel   *entity.AdminRole
		ctx         = m.GetCtx()
		fields      = escapeFieldsToSlice(m.GetFieldsStr())
		co          = contexts.Get(ctx)
	)

	if co == nil || co.User == nil {
		return m
	}

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

	err := g.Model("admin_role").Where("id", co.User.RoleId).Scan(&roleModel)
	if err != nil {
		panic(fmt.Sprintf("HandlerFilterAuth Failed to role information err:%+v", err))
	}

	if roleModel == nil {
		panic(fmt.Sprintf("HandlerFilterAuth Failed to role information err2:%+v", err))
	}

	// TODO 当前不是完整功能，预计在下个版本中完善
	switch roleModel.DataScope {
	case consts.RoleDataAll: // 全部权限
		// ...
	case consts.RoleDataNowDept: // 当前部门
		m = m.Where(filterField, co.User.DeptId)
	case consts.RoleDataDeptAndSub: // 当前部门及以下部门
		//m = m.Where(filterField, 1)
	case consts.RoleDataDeptCustom: // 自定义部门
		m = m.WhereIn(filterField, roleModel.CustomDept.Var().Ints())
	case consts.RoleDataSelf: // 仅自己
		m = m.Where(filterField, co.User.Id)
	case consts.RoleDataSelfAndSub: // 自己和直属下级
		//m = m.Where(filterField, 1)
	case consts.RoleDataSelfAndAllSub: // 自己和全部下级
		//m = m.Where(filterField, 1)

	default:
		panic("HandlerFilterAuth dataScope is not registered")
	}

	return m
}

// HandlerForceCache 强制缓存
func HandlerForceCache(m *gdb.Model) *gdb.Model {
	return m.Cache(gdb.CacheOption{Duration: -1, Force: true})
}

// escapeFieldsToSlice 将转义过的字段转换为字段集切片
func escapeFieldsToSlice(s string) []string {
	return gstr.Explode(",", gstr.Replace(gstr.Replace(s, "`,`", ","), "`", ""))
}
