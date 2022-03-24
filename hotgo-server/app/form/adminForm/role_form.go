//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
)

//  查询列表
type RoleMemberListReq struct {
	g.Meta `path:"/role/member_list" method:"get" tags:"角色" summary:"获取角色下的会员列表"`
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

type RoleMemberListRes struct {
	List []*input.AdminMemberListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

//  查询列表
type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色" summary:"获取角色列表"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	DeptId    int    `json:"dept_id"   description:"部门ID"`
	Mobile    int    `json:"mobile"   description:"手机号"`
	Username  string `json:"username"   description:"用户名"`
	Realname  string `json:"realname"   description:"真实姓名"`
	StartTime string `json:"start_time"   description:"开始时间"`
	EndTime   string `json:"end_time"   description:"结束时间"`
	Name      string `json:"name"   description:"岗位名称"`
	Code      string `json:"code"   description:"岗位编码"`
}

type RoleListRes struct {
	List []*input.AdminRoleListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

//  动态路由
type RoleDynamicReq struct {
	g.Meta `path:"/role/dynamic" method:"get" tags:"路由" summary:"获取动态路由" description:"获取登录用户动态路由"`
}

type RoleDynamicMeta struct {
	Title   string `json:"title"      description:"菜单标题"`
	Icon    string `json:"icon"      description:"菜单图标"`
	NoCache bool   `json:"noCache"      description:"是否缓存"`
	Remark  string `json:"remark"      description:"备注"`
}

type RoleDynamicBase struct {
	Id         int64            `json:"id"   description:"菜单ID"`
	Pid        int64            `json:"pid"   description:"父ID"`
	Name       string           `json:"name"      description:"菜单名称"`
	Code       string           `json:"code"      description:"菜单编码"`
	Path       string           `json:"path"      description:"路由地址"`
	Hidden     bool             `json:"hidden"      description:"是否隐藏"`
	Redirect   string           `json:"redirect"     description:"重定向"`
	Component  string           `json:"component" description:"组件路径"`
	AlwaysShow bool             `json:"alwaysShow"     description:"暂时不知道干啥"`
	IsFrame    string           `json:"isFrame"   description:"是否为外链（0是 1否）"`
	Meta       *RoleDynamicMeta `json:"meta"   description:"配置数据集"`
}

type RoleDynamicMenu struct {
	RoleDynamicBase
	Children []*RoleDynamicBase `json:"children"   description:"子菜单"`
}

type RoleDynamicRes []*RoleDynamicMenu

type RoleMenuEditReq struct {
	g.Meta  `path:"/role/edit" method:"post" tags:"角色" summary:"编辑角色菜单权限"`
	RoleId  int64   `json:"id"`
	MenuIds []int64 `json:"menuIds"`
}

type RoleMenuEditRes struct{}
