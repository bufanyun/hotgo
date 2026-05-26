// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import (
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

func init() {
	dict.RegisterEnums("deptType", "部门类型选项", DeptTypeOptions)
}

const (
	DeptTypeCompany  = "company"  // 公司
	DeptTypeTenant   = "tenant"   // 租户
	DeptTypeMerchant = "merchant" // 商户
	DeptTypeUser     = "user"     // 用户
)

// DeptTypeOptions 部门类型选项
var DeptTypeOptions = []*model.Option{
	dict.GenSuccessOption(DeptTypeCompany, "公司"),
	dict.GenErrorOption(DeptTypeTenant, "租户"),
	dict.GenInfoOption(DeptTypeMerchant, "商户"),
	dict.GenWarningOption(DeptTypeUser, "用户"),
}
