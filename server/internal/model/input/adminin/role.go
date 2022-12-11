// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// RoleListInp 获取列表
type RoleListInp struct {
	Page    int64
	PerPage int64
}

type RoleListModel struct {
	entity.AdminRole
}

// RoleMemberListInp 查询列表
type RoleMemberListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Role      int    `json:"role"   description:"角色ID"`
	DeptId    int    `json:"dept_id"   description:"部门ID"`
	Mobile    int    `json:"mobile"   description:"手机号"`
	Username  string `json:"username"   description:"用户名"`
	Realname  string `json:"realname"   description:"真实姓名"`
	StartTime string `json:"start_time"   description:"开始时间"`
	EndTime   string `json:"end_time"   description:"结束时间"`
	Name      string `json:"name"   description:"岗位名称"`
	Code      string `json:"code"   description:"岗位编码"`
}

type RoleMemberListModel []*MemberListModel

// MenuRoleListInp 查询角色菜单列表
type MenuRoleListInp struct {
	RoleId int64
}
type MenuRoleListModel struct {
	Menus       []*model.LabelTreeMenu `json:"menus"   description:"菜单列表"`
	CheckedKeys []int64                `json:"checkedKeys"   description:"选择的菜单ID"`
}
