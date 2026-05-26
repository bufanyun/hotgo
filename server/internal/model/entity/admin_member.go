// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMember is the golang structure for table admin_member.
type AdminMember struct {
	Id                 int64       `json:"id"                 orm:"id"                   description:"管理员ID"`
	DeptId             int64       `json:"deptId"             orm:"dept_id"              description:"部门ID"`
	RoleId             int64       `json:"roleId"             orm:"role_id"              description:"角色ID"`
	RealName           string      `json:"realName"           orm:"real_name"            description:"真实姓名"`
	Username           string      `json:"username"           orm:"username"             description:"帐号"`
	PasswordHash       string      `json:"passwordHash"       orm:"password_hash"        description:"密码"`
	Salt               string      `json:"salt"               orm:"salt"                 description:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" orm:"password_reset_token" description:"密码重置令牌"`
	Integral           float64     `json:"integral"           orm:"integral"             description:"积分"`
	Balance            float64     `json:"balance"            orm:"balance"              description:"余额"`
	Avatar             string      `json:"avatar"             orm:"avatar"               description:"头像"`
	Sex                int         `json:"sex"                orm:"sex"                  description:"性别"`
	Qq                 string      `json:"qq"                 orm:"qq"                   description:"qq"`
	Email              string      `json:"email"              orm:"email"                description:"邮箱"`
	Mobile             string      `json:"mobile"             orm:"mobile"               description:"手机号码"`
	Birthday           *gtime.Time `json:"birthday"           orm:"birthday"             description:"生日"`
	CityId             int64       `json:"cityId"             orm:"city_id"              description:"城市编码"`
	Address            string      `json:"address"            orm:"address"              description:"联系地址"`
	Pid                int64       `json:"pid"                orm:"pid"                  description:"上级管理员ID"`
	Level              int         `json:"level"              orm:"level"                description:"关系树等级"`
	Tree               string      `json:"tree"               orm:"tree"                 description:"关系树"`
	InviteCode         string      `json:"inviteCode"         orm:"invite_code"          description:"邀请码"`
	Cash               *gjson.Json `json:"cash"               orm:"cash"                 description:"提现配置"`
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       orm:"last_active_at"       description:"最后活跃时间"`
	Remark             string      `json:"remark"             orm:"remark"               description:"备注"`
	Status             int         `json:"status"             orm:"status"               description:"状态"`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:"修改时间"`
}
