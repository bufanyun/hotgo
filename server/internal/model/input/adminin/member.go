// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// GetIdByCodeInp 通过邀请码获取会员ID
type GetIdByCodeInp struct {
	Code string `json:"code""`
}
type GetIdByCodeModel struct {
	Id int64
}

// MemberProfileInp 获取指定会员资料
type MemberProfileInp struct {
	Id int64
}
type MemberProfileModel struct {
	PostGroup string           `json:"postGroup" description:"岗位名称"`
	RoleGroup string           `json:"roleGroup" description:"角色名称"`
	User      *MemberViewModel `json:"member" description:"用户基本信息"`
	SysDept   *DeptViewModel   `json:"sysDept" description:"部门信息"`
	SysRoles  []*RoleListModel `json:"sysRoles" description:"角色列表"`
	PostIds   int64            `json:"postIds" description:"当前岗位"`
	RoleIds   int64            `json:"roleIds" description:"当前角色"`
}

// MemberUpdateProfileInp 更新会员资料
type MemberUpdateProfileInp struct {
	Mobile   int
	Email    string
	Realname string
}

// MemberUpdatePwdInp 修改登录密码
type MemberUpdatePwdInp struct {
	Id          int64
	OldPassword string
	NewPassword string
}

// MemberResetPwdInp 重置密码
type MemberResetPwdInp struct {
	Password string
	Id       int64
}

// MemberEmailUniqueInp 邮箱是否唯一
type MemberEmailUniqueInp struct {
	Email string
	Id    int64
}

type MemberEmailUniqueModel struct {
	IsUnique bool
}

// MemberMobileUniqueInp 手机号是否唯一
type MemberMobileUniqueInp struct {
	Mobile string
	Id     int64
}

type MemberMobileUniqueModel struct {
	IsUnique bool
}

// MemberNameUniqueInp 名称是否唯一
type MemberNameUniqueInp struct {
	Username string
	Id       int64
}

type MemberNameUniqueModel struct {
	IsUnique bool
}

// MemberMaxSortInp 最大排序
type MemberMaxSortInp struct {
	Id int64
}

type MemberMaxSortModel struct {
	Sort int
}

// MemberEditInp 修改/新增管理员
type MemberEditInp struct {
	Id         int64       `json:"id"                   description:""`
	RoleId     int         `json:"roleId"            v:"required#角色不能为空"         description:"角色ID"`
	PostIds    []int64     `json:"postIds"       v:"required#岗位不能为空"           description:"岗位ID"`
	DeptId     int64       `json:"deptId"       v:"required#部门不能为空"           description:"部门ID"`
	Username   string      `json:"username"   v:"required#账号不能为空"           description:"帐号"`
	Password   string      `json:"password"            description:"密码"`
	Realname   string      `json:"realName"             description:"真实姓名"`
	Avatar     string      `json:"avatar"               description:"头像"`
	Sex        string      `json:"sex"                  description:"性别"`
	Qq         string      `json:"qq"                   description:"qq"`
	Email      string      `json:"email"                description:"邮箱"`
	Birthday   *gtime.Time `json:"birthday"             description:"生日"`
	ProvinceId int         `json:"provinceId"          description:"省"`
	CityId     int         `json:"cityId"              description:"城市"`
	AreaId     int         `json:"areaId"              description:"地区"`
	Address    string      `json:"address"              description:"默认地址"`
	Mobile     string      `json:"mobile"               description:"手机号码"`
	Remark     string      `json:"remark"               description:"备注"`
	Status     string      `json:"status"               description:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"           description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"           description:"修改时间"`
}

type MemberAddInp struct {
	MemberEditInp
	PasswordHash string `json:"passwordHash"            description:"密码hash"`
	Salt         string `json:"salt"            description:"密码盐"`
	Pid          int64  `json:"pid"                description:"上级ID"`
	Level        int    `json:"level"              description:"等级"`
	Tree         string `json:"tree"               description:"关系树"`
}

type MemberEditModel struct{}

// MemberDeleteInp 删除字典类型
type MemberDeleteInp struct {
	Id interface{}
}
type MemberDeleteModel struct{}

// MemberViewInp 获取信息
type MemberViewInp struct {
	Id int64
}

type MemberViewModel struct {
	entity.AdminMember
}

// MemberListInp 获取列表
type MemberListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name      string
	Code      string
	DeptId    int
	Mobile    int
	Username  string
	Realname  string
	CreatedAt []int64
}

type MemberListModel struct {
	entity.AdminMember
	DeptName string  `json:"deptName"`
	RoleName string  `json:"roleName"`
	PostIds  []int64 `json:"postIds"`
	DeptId   int64   `json:"deptId"             description:"部门ID"`
}

// MemberLoginInp 登录
type MemberLoginInp struct {
	Username string
	Password string
}
type MemberLoginModel struct {
	UserId      int64                     `json:"userId"    description:"会员ID"`
	Username    string                    `json:"username"    description:"用户名"`
	RealName    string                    `json:"realName"    description:"昵称"`
	Avatar      string                    `json:"avatar"       description:"头像"`
	Token       string                    `json:"token" v:""  description:"登录token"`
	Permissions []*MemberLoginPermissions `json:"permissions"  description:"角色信息"`
}

// MemberLoginPermissions 登录用户角色信息
type MemberLoginPermissions struct {
	Label string `json:"label"    description:"标签"`
	Value string `json:"value"    description:"值"`
}

// MemberStatusInp  更新状态
type MemberStatusInp struct {
	entity.AdminPost
}
type MemberStatusModel struct{}
