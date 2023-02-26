// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"sort"
)

// RoleListInp 获取列表
type RoleListInp struct {
	form.PageReq
}

type RoleTree struct {
	entity.AdminRole
	Label    string      `json:"label"     dc:"标签"`
	Value    int64       `json:"value"     dc:"键值"`
	Children []*RoleTree `json:"children"  dc:"子级"`
}

type RoleListModel struct {
	List []*RoleTree `json:"list"`
}

func Sort(v []*RoleTree) {
	sort.SliceStable(v, func(i, j int) bool {
		if v[i].Sort < v[j].Sort {
			return true
		}
		if v[i].Sort > v[j].Sort {
			return false
		}
		return v[i].Id < v[j].Id
	})
}

// RoleMemberListInp 查询列表
type RoleMemberListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Role      int    `json:"role"        dc:"角色ID"`
	DeptId    int    `json:"deptId"      dc:"部门ID"`
	Mobile    int    `json:"mobile"      dc:"手机号"`
	Username  string `json:"username"    dc:"用户名"`
	RealName  string `json:"realName"    dc:"真实姓名"`
	StartTime string `json:"start_time"  dc:"开始时间"`
	EndTime   string `json:"end_time"    dc:"结束时间"`
	Name      string `json:"name"        dc:"岗位名称"`
	Code      string `json:"code"        dc:"岗位编码"`
}

type RoleMemberListModel []*MemberListModel

// MenuRoleListInp 查询角色菜单列表
type MenuRoleListInp struct {
	RoleId int64
}
type MenuRoleListModel struct {
	Menus       []*model.LabelTreeMenu `json:"menus"         dc:"菜单列表"`
	CheckedKeys []int64                `json:"checkedKeys"   dc:"选择的菜单ID"`
}

type DataScopeEditInp struct {
	Id         int64   `json:"id" v:"required"        dc:"角色ID"`
	DataScope  int     `json:"dataScope" v:"required" dc:"数据范围"`
	CustomDept []int64 `json:"customDept"             dc:"自定义部门权限"`
}
