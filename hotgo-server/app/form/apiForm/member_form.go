package apiForm

import (
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/gogf/gf/v2/frame/g"
)

//  获取登录用户的基本信息
type MemberProfileReq struct {
	g.Meta `path:"/member/profile" method:"get" tags:"会员" summary:"获取登录用户的基本信息"`
}
type MemberProfileRes struct {
	PostGroup string                      `json:"postGroup" description:"岗位名称"`
	RoleGroup string                      `json:"roleGroup" description:"角色名称"`
	User      *input.AdminMemberViewModel `json:"user" description:"用户基本信息"`
	SysDept   *input.AdminDeptViewModel   `json:"sysDept" description:"部门信息"`
	SysRoles  []*input.AdminRoleListModel `json:"sysRoles" description:"角色列表"`
	PostIds   int64                       `json:"postIds" description:"当前岗位"`
	RoleIds   int64                       `json:"roleIds" description:"当前角色"`
}
