// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// MemberUpdateCashInp 更新会员提现信息
type MemberUpdateCashInp struct {
	Name      string `json:"name" v:"required#支付宝姓名不能为空"  dc:"支付宝姓名"`
	PayeeCode string `json:"payeeCode" v:"required#支付宝收款码不能为空"  dc:"支付宝收款码"`
	Account   string `json:"account" v:"required#支付宝账号不能为空"  dc:"支付宝账号"`
	Password  string `json:"password" v:"required#密码不能为空"  dc:"密码"`
}

type MemberUpdateEmailInp struct {
	Email string `json:"email"  v:"required#换绑邮箱不能为空"       dc:"换绑邮箱"`
	Code  string `json:"code" dc:"原邮箱验证码"`
}

// MemberUpdateMobileInp 换绑手机号
type MemberUpdateMobileInp struct {
	Mobile string `json:"mobile"  v:"required#换绑手机号不能为空"       dc:"换绑手机号"`
	Code   string `json:"code" dc:"原号码短信验证码"`
}

// GetIdByCodeInp 通过邀请码获取用户ID
type GetIdByCodeInp struct {
	Code string `json:"code"`
}
type GetIdByCodeModel struct {
	Id int64
}

// MemberProfileInp 获取指定用户资料
type MemberProfileInp struct {
	Id int64
}
type MemberProfileModel struct {
	PostGroup string           `json:"postGroup" dc:"岗位名称"`
	RoleGroup string           `json:"roleGroup" dc:"角色名称"`
	User      *MemberViewModel `json:"member"    dc:"用户基本信息"`
	SysDept   *DeptViewModel   `json:"sysDept"   dc:"部门信息"`
	SysRoles  []*RoleListModel `json:"sysRoles"  dc:"角色列表"`
	PostIds   int64            `json:"postIds"   dc:"当前岗位"`
	RoleIds   int64            `json:"roleIds"   dc:"当前角色"`
}

// MemberUpdateProfileInp 更新用户资料
type MemberUpdateProfileInp struct {
	Avatar   string      `json:"avatar"   v:"required#头像不能为空"     dc:"头像"`
	RealName string      `json:"realName"  v:"required#真实姓名不能为空"       dc:"真实姓名"`
	Qq       string      `json:"qq"          dc:"QQ"`
	Birthday *gtime.Time `json:"birthday"    dc:"生日"`
	Sex      int         `json:"sex"         dc:"性别"`
	Address  string      `json:"address"     dc:"联系地址"`
	CityId   int64       `json:"cityId"      dc:"城市编码"`
}

// MemberUpdatePwdInp 修改登录密码
type MemberUpdatePwdInp struct {
	Id          int64
	OldPassword string `json:"oldPassword" v:"required#原密码不能为空"  dc:"原密码"`
	NewPassword string `json:"newPassword" v:"required|length:6,16#新密码不能为空#新密码需在6~16之间"  dc:"新密码"`
}

// MemberResetPwdInp 重置密码
type MemberResetPwdInp struct {
	Password string `json:"password" v:"required#密码不能为空"  dc:"密码"`
	Id       int64  `json:"id" dc:"用户ID"`
}

// MemberEmailUniqueInp 邮箱是否唯一
type MemberEmailUniqueInp struct {
	Email string `json:"email" v:"required#邮箱不能为空"  dc:"邮箱"`
	Id    int64  `json:"id" dc:"用户ID"`
}

type MemberEmailUniqueModel struct {
	IsUnique bool
}

// MemberMobileUniqueInp 手机号是否唯一
type MemberMobileUniqueInp struct {
	Mobile string `json:"mobile" v:"required#手机号不能为空"  dc:"手机号"`
	Id     int64  `json:"id" dc:"用户ID"`
}

type MemberMobileUniqueModel struct {
	IsUnique bool
}

// MemberNameUniqueInp 名称是否唯一
type MemberNameUniqueInp struct {
	Username string `json:"username" v:"required#用户名称不能为空"  dc:"用户名称"`
	Id       int64  `json:"id" dc:"用户ID"`
}

type MemberNameUniqueModel struct {
	IsUnique bool
}

// MemberMaxSortInp 最大排序
type MemberMaxSortInp struct {
	Id int64 `json:"id" dc:"用户ID"`
}

type MemberMaxSortModel struct {
	Sort int
}

// MemberEditInp 修改/新增管理员
type MemberEditInp struct {
	Id           int64       `json:"id"                                            dc:""`
	RoleId       int         `json:"roleId"    v:"required#角色不能为空"            dc:"角色ID"`
	PostIds      []int64     `json:"postIds"   v:"required#岗位不能为空"            dc:"岗位ID"`
	DeptId       int64       `json:"deptId"    v:"required#部门不能为空"            dc:"部门ID"`
	Username     string      `json:"username"   v:"required#账号不能为空"           dc:"帐号"`
	PasswordHash string      `json:"passwordHash"                                  dc:"密码hash"`
	Password     string      `json:"password"                                      dc:"密码"`
	RealName     string      `json:"realName"                                      dc:"真实姓名"`
	Avatar       string      `json:"avatar"                                        dc:"头像"`
	Sex          string      `json:"sex"                                           dc:"性别"`
	Qq           string      `json:"qq"                                            dc:"qq"`
	Email        string      `json:"email"                                         dc:"邮箱"`
	Birthday     *gtime.Time `json:"birthday"                                      dc:"生日"`
	ProvinceId   int         `json:"provinceId"                                    dc:"省"`
	CityId       int         `json:"cityId"                                        dc:"城市"`
	AreaId       int         `json:"areaId"                                        dc:"地区"`
	Address      string      `json:"address"                                       dc:"默认地址"`
	Mobile       string      `json:"mobile"                                        dc:"手机号码"`
	Remark       string      `json:"remark"                                        dc:"备注"`
	Status       string      `json:"status"                                        dc:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"                                     dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"                                     dc:"修改时间"`
}

type MemberAddInp struct {
	MemberEditInp
	Salt  string `json:"salt"               dc:"密码盐"`
	Pid   int64  `json:"pid"                dc:"上级ID"`
	Level int    `json:"level"              dc:"等级"`
	Tree  string `json:"tree"               dc:"关系树"`
}

func (in *MemberEditInp) Filter(ctx context.Context) (err error) {
	if in.Password != "" {
		if err := g.Validator().
			Rules("length:6,16").
			Messages("#新密码不能为空#新密码需在6~16之间").
			Data(in.Password).Run(ctx); err != nil {
			return err.Current()
		}
	}
	return
}

type MemberEditModel struct{}

// MemberDeleteInp 删除字典类型
type MemberDeleteInp struct {
	Id interface{} `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}
type MemberDeleteModel struct{}

// MemberViewInp 获取信息
type MemberViewInp struct {
	Id int64 `json:"id" dc:"用户ID"`
}

type MemberViewModel struct {
	entity.AdminMember
	DeptName string `json:"deptName"    dc:"所属部门"`
	RoleName string `json:"roleName"    dc:"所属角色"`
}

// MemberListInp 获取列表
type MemberListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	DeptId    int     `json:"deptId"     dc:"部门ID"`
	Mobile    int     `json:"mobile"     dc:"手机号"`
	Username  string  `json:"username"   dc:"用户名"`
	RealName  string  `json:"realName"   dc:"真实姓名"`
	Name      string  `json:"name"       dc:"岗位名称"`
	Code      string  `json:"code"       dc:"岗位编码"`
	CreatedAt []int64 `json:"createdAt"  dc:"创建时间"`
}

type MemberListModel struct {
	entity.AdminMember
	DeptName string  `json:"deptName"    dc:"所属部门"`
	RoleName string  `json:"roleName"    dc:"所属角色"`
	PostIds  []int64 `json:"postIds"     dc:"岗位"`
	DeptId   int64   `json:"deptId"      dc:"部门ID"`
}

// MemberLoginInp 登录
type MemberLoginInp struct {
	Username string
	Password string
}
type MemberLoginModel struct {
	Id      int64  `json:"id"              dc:"用户ID"`
	Token   string `json:"token"           dc:"登录token"`
	Expires int64  `json:"expires"         dc:"登录有效期"`
}

type LoginMemberInfoModel struct {
	Id          int64       `json:"id"                 dc:"用户ID"`
	DeptName    string      `json:"deptName"           dc:"所属部门"`
	RoleName    string      `json:"roleName"           dc:"所属角色"`
	Permissions []string    `json:"permissions"        dc:"角色信息"`
	DeptId      int64       `json:"-"                  dc:"部门ID"`
	RoleId      int64       `json:"-"                  dc:"角色ID"`
	Username    string      `json:"username"           dc:"用户名"`
	RealName    string      `json:"realName"           dc:"姓名"`
	Avatar      string      `json:"avatar"             dc:"头像"`
	Balance     float64     `json:"balance"            dc:"余额"`
	Sex         int         `json:"sex"                dc:"性别"`
	Qq          string      `json:"qq"                 dc:"qq"`
	Email       string      `json:"email"              dc:"邮箱"`
	Mobile      string      `json:"mobile"             dc:"手机号码"`
	Birthday    *gtime.Time `json:"birthday"           dc:"生日"`
	CityId      int64       `json:"cityId"             dc:"城市编码"`
	Address     string      `json:"address"            dc:"联系地址"`
	Cash        *MemberCash `json:"cash"               dc:"收款信息"`
	CreatedAt   *gtime.Time `json:"createdAt"          dc:"创建时间"`
	*MemberLoginStatModel
}

// MemberLoginPermissions 登录用户角色信息
type MemberLoginPermissions []string

// MemberCash 用户提现配置
type MemberCash struct {
	Name      string `json:"name"       dc:"收款人姓名"`
	Account   string `json:"account"    dc:"收款账户"`
	PayeeCode string `json:"payeeCode"  dc:"收款码"`
}

// MemberStatusInp  更新状态
type MemberStatusInp struct {
	entity.AdminPost
}
type MemberStatusModel struct{}

// MemberSelectInp 获取可选的后台用户选项
type MemberSelectInp struct {
}

type MemberSelectModel struct {
	Value    int64  `json:"value"    dc:"用户ID"`
	Label    string `json:"label"    dc:"真实姓名"`
	Username string `json:"username" dc:"用户名"`
	Avatar   string `json:"avatar"   dc:"头像"`
}

// MemberLoginStatInp 用户登录统计
type MemberLoginStatInp struct {
	MemberId int64
}

type MemberLoginStatModel struct {
	LoginCount  int         `json:"loginCount"  dc:"登录次数"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" dc:"最后登录时间"`
	LastLoginIp string      `json:"lastLoginIp" dc:"最后登录IP"`
}
