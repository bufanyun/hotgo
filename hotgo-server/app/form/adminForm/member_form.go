//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

//  更新会员资料
type MemberUpdateProfileReq struct {
	Mobile   int    `json:"mobile"   description:"手机号"`
	Email    string `json:"email"   description:"邮箱"`
	Realname string `json:"realname"   description:"真实姓名"`
	g.Meta   `path:"/member/update_profile" method:"post" tags:"会员" summary:"更新会员资料"`
}
type MemberUpdateProfileRes struct{}

//  修改登录密码
type MemberUpdatePwdReq struct {
	OldPassword string `json:"oldPassword" v:"required#原密码不能为空"  description:"原密码"`
	NewPassword string `json:"newPassword" v:"required|length:6,16#新密码不能为空#新密码需在6~16之间"  description:"新密码"`
	g.Meta      `path:"/member/update_pwd" method:"post" tags:"会员" summary:"重置密码"`
}
type MemberUpdatePwdRes struct{}

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

//  重置密码
type MemberResetPwdReq struct {
	Password string `json:"password" v:"required#密码不能为空"  description:"密码"`
	Id       int64  `json:"id" description:"会员ID"`
	g.Meta   `path:"/member/reset_pwd" method:"post" tags:"会员" summary:"重置密码"`
}
type MemberResetPwdRes struct{}

//  邮箱是否唯一
type MemberEmailUniqueReq struct {
	Email  string `json:"email" v:"required#邮箱不能为空"  description:"邮箱"`
	Id     int64  `json:"id" description:"会员ID"`
	g.Meta `path:"/member/email_unique" method:"get" tags:"会员" summary:"邮箱是否唯一"`
}
type MemberEmailUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  手机号是否唯一
type MemberMobileUniqueReq struct {
	Mobile string `json:"mobile" v:"required#手机号不能为空"  description:"手机号"`
	Id     int64  `json:"id" description:"会员ID"`
	g.Meta `path:"/member/mobile_unique" method:"get" tags:"会员" summary:"手机号是否唯一"`
}
type MemberMobileUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  名称是否唯一
type MemberNameUniqueReq struct {
	Username string `json:"username" v:"required#会员名称不能为空"  description:"会员名称"`
	Id       int64  `json:"id" description:"会员ID"`
	g.Meta   `path:"/member/name_unique" method:"get" tags:"会员" summary:"会员名称是否唯一"`
}
type MemberNameUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  查询列表
type MemberListReq struct {
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
	g.Meta    `path:"/member/list" method:"get" tags:"会员" summary:"获取会员列表"`
}

type MemberListRes struct {
	List []*input.AdminMemberListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

//  获取指定信息
type MemberViewReq struct {
	Id     int64 `json:"id" description:"会员ID"` // v:"required#会员ID不能为空"
	g.Meta `path:"/member/view" method:"get" tags:"会员" summary:"获取指定信息"`
}
type MemberViewRes struct {
	*input.AdminMemberViewModel
	Posts    []*input.AdminPostListModel `json:"posts" description:"可选岗位"`
	PostIds  []int64                     `json:"postIds" description:"当前岗位"`
	Roles    []*input.AdminRoleListModel `json:"roles" description:"可选角色"`
	RoleIds  []int64                     `json:"roleIds" description:"当前角色"`
	DeptName string                      `json:"dept_name" description:"部门名称"`
}

//  修改/新增
type MemberEditReq struct {
	input.AdminMemberEditInp
	g.Meta `path:"/member/edit" method:"post" tags:"会员" summary:"修改/新增会员"`
}
type MemberEditRes struct{}

//  删除
type MemberDeleteReq struct {
	Id     interface{} `json:"id" v:"required#会员ID不能为空" description:"会员ID"`
	g.Meta `path:"/member/delete" method:"post" tags:"会员" summary:"删除会员"`
}
type MemberDeleteRes struct{}

//  最大排序
type MemberMaxSortReq struct {
	Id     int64 `json:"id" description:"会员ID"`
	g.Meta `path:"/member/max_sort" method:"get" tags:"会员" summary:"会员最大排序"`
}
type MemberMaxSortRes struct {
	Sort int `json:"sort" description:"排序"`
}

//  获取登录用户信息
type MemberInfoReq struct {
	g.Meta `path:"/member/info" method:"get" tags:"会员" summary:"获取登录用户信息" description:"获取管理后台的登录用户信息"`
}

type PortalConfigContentOptions struct {
	TitleRequired bool   `json:"titleRequired"  titleRequired:""`
	MoreUrl       string `json:"moreUrl"  description:"模块地址"`
	Refresh       int    `json:"refresh"  description:"刷新"`
}

type PortalConfigContent struct {
	Id          int                           `json:"id"  description:"内容ID"`
	X           int                           `json:"x"  description:""`
	Y           int                           `json:"y"  description:""`
	W           int                           `json:"w"  description:"宽"`
	H           int                           `json:"h"  description:"高"`
	I           int                           `json:"i"  description:""`
	Key         string                        `json:"key"  description:""`
	IsShowTitle string                        `json:"isShowTitle"  description:""`
	IsAllowDrag bool                          `json:"isAllowDrag"  description:""`
	Name        string                        `json:"name"  description:""`
	Type        string                        `json:"type"  description:""`
	Url         string                        `json:"url"  description:""`
	Options     []*PortalConfigContentOptions `json:"options"  description:""`
	Moved       bool                          `json:"moved"  description:""`
}

type PortalConfig struct {
	CreateByName        string      `json:"createByName"  description:"创建者名称"`
	CreateDeptName      string      `json:"createDeptName"  description:"创建部门名称"`
	ImportErrInfo       string      `json:"importErrInfo"  description:"导出错误信息"`
	Id                  string      `json:"id"  description:"用户ID"`
	SearchValue         string      `json:"searchValue"  description:"搜索内容"`
	CreateBy            string      `json:"createBy"  description:"创建者名称"`
	CreateDept          string      `json:"createDept"  description:"创建部门名称"`
	CreateTime          *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateBy            string      `json:"updateBy"  description:"更新者名称"`
	UpdateTime          *gtime.Time `json:"updateTime"  description:"更新时间"`
	UpdateIp            string      `json:"updateIp"  description:"更新iP"`
	Remark              string      `json:"remark"  description:"备注"`
	Version             string      `json:"version"  description:"版本号"`
	DelFlag             string      `json:"delFlag"  description:"删除标签"`
	HandleType          string      `json:"handleType"  description:""`
	Params              string      `json:"params"  description:""`
	Name                string      `json:"name"  description:"配置名称"`
	Code                string      `json:"code"  description:"配置代码"`
	ApplicationRange    string      `json:"applicationRange"  description:""`
	IsDefault           string      `json:"isDefault"  description:"是否默认"`
	ResourceId          string      `json:"resourceId"  description:""`
	ResourceName        string      `json:"resourceName"  description:""`
	SystemDefinedId     string      `json:"systemDefinedId"  description:""`
	Sort                string      `json:"sort"  description:"排序"`
	SaveType            string      `json:"saveType"  description:""`
	Status              string      `json:"status"  description:"状态"`
	RecordLog           string      `json:"recordLog"  description:""`
	PortalConfigContent string      `json:"content"  description:"配置内容"`
}

type MemberInfoRes struct {
	DefaultPortalConfig []*PortalConfig       `json:"defaultPortalConfig"  description:"默认用户配置"`
	LincenseInfo        string                `json:"lincenseInfo"  description:"应用版本号"`
	Permissions         []string              `json:"permissions"description:"权限"`
	Roles               []string              `json:"roles" description:"角色"`
	SysNoticeList       []*entity.AdminNotice `json:"sysNoticeList" description:"系统公告"`
	UserPortalConfig    []*PortalConfig       `json:"userPortalConfig" description:"用户配置"`
	User                model.Identity        `json:"user"  description:"用户信息"`
}
