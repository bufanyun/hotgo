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
	form.PageReq
}

type RoleListModel struct {
	entity.AdminRole
	Label string `json:"label" dc:"标签"`
	Value int64  `json:"value" dc:"键值"`
}

// RoleMemberListInp 查询列表
type RoleMemberListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Role      int    `json:"role"   description:"角色ID"`
	DeptId    int    `json:"deptId"   description:"部门ID"`
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

type DataScopeEditInp struct {
	Id         int64   `json:"id" v:"required" dc:"角色ID"`
	DataScope  int     `json:"dataScope" v:"required" dc:"数据范围"`
	CustomDept []int64 `json:"customDept" dc:"自定义部门权限"`
}
