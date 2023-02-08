package hook

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberInfo 后台用户信息
var MemberInfo = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		if err != nil {
			return
		}
		for i, record := range result {
			// 部门
			if !record["dept_id"].IsEmpty() {
				deptName, err := g.Model("admin_dept").Ctx(ctx).
					Fields("name").
					Where("id", record["dept_id"]).
					Value()
				if err != nil {
					break
				}
				record["deptName"] = deptName
			}

			// 角色
			if !record["role_id"].IsEmpty() {
				roleName, err := g.Model("admin_role").Ctx(ctx).
					Fields("name").
					Where("id", record["role_id"]).
					Value()
				if err != nil {
					break
				}
				record["roleName"] = roleName
			}

			if !record["password_hash"].IsEmpty() {
				record["password_hash"] = gvar.New("")
			}

			if !record["salt"].IsEmpty() {
				record["salt"] = gvar.New("")
			}

			if !record["auth_key"].IsEmpty() {
				record["auth_key"] = gvar.New("")
			}

			result[i] = record
		}
		return
	},
}
