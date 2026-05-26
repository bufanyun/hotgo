// Package handler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package handler

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/library/contexts"
	"hotgo/utility/convert"
)

// FilterTenant 过滤多租户数据权限
// 根据部门类型识别当前租户、商户、用户身份，过滤只属于自己的数据
func FilterTenant(m *gdb.Model) *gdb.Model {
	var (
		needAuth    bool
		filterField string
		fields      = convert.EscapeFieldsToSlice(m.GetFieldsStr())
		ctx         = m.GetCtx()
	)

	// 租户
	if contexts.IsTenantDept(ctx) && gstr.InArray(fields, "tenant_id") {
		needAuth = true
		filterField = "tenant_id"
	}

	// 商户
	if contexts.IsMerchantDept(ctx) && gstr.InArray(fields, "merchant_id") {
		needAuth = true
		filterField = "merchant_id"
	}

	// 用户
	if contexts.IsUserDept(ctx) && gstr.InArray(fields, "user_id") {
		needAuth = true
		filterField = "user_id"
	}

	if !needAuth {
		return m
	}

	m = m.Where(filterField, contexts.GetUserId(ctx))
	return m
}
