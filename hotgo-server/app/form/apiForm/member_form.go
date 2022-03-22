//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package apiForm

import (
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/gogf/gf/v2/frame/g"
)

//  获取登录用户的基本信息
type MemberProfileReq struct {
	g.Meta `path:"/member/profile" method:"get" tags:"会员接口" summary:"获取登录用户的基本信息"`
}
type MemberProfileRes struct {
	PostGroup string                      `json:"postGroup" dc:"岗位名称"`
	RoleGroup string                      `json:"roleGroup" dc:"角色名称"`
	User      *input.AdminMemberViewModel `json:"user" dc:"用户基本信息"`
	SysDept   *input.AdminDeptViewModel   `json:"sysDept" dc:"部门信息"`
	SysRoles  []*input.AdminRoleListModel `json:"sysRoles" dc:"角色列表"`
	PostIds   int64                       `json:"postIds" dc:"当前岗位"`
	RoleIds   int64                       `json:"roleIds" dc:"当前角色"`
}
