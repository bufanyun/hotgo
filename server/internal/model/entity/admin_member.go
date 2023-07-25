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
	Id                 int64       `json:"id"                 description:"管理员ID"`
	DeptId             int64       `json:"deptId"             description:"部门ID"`
	RoleId             int64       `json:"roleId"             description:"角色ID"`
	RealName           string      `json:"realName"           description:"真实姓名"`
	Username           string      `json:"username"           description:"帐号"`
	PasswordHash       string      `json:"passwordHash"       description:"密码"`
	Salt               string      `json:"salt"               description:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" description:"密码重置令牌"`
	Integral           float64     `json:"integral"           description:"积分"`
	Balance            float64     `json:"balance"            description:"余额"`
	Avatar             string      `json:"avatar"             description:"头像"`
	Sex                int         `json:"sex"                description:"性别"`
	Qq                 string      `json:"qq"                 description:"qq"`
	Email              string      `json:"email"              description:"邮箱"`
	Mobile             string      `json:"mobile"             description:"手机号码"`
	Birthday           *gtime.Time `json:"birthday"           description:"生日"`
	CityId             int64       `json:"cityId"             description:"城市编码"`
	Address            string      `json:"address"            description:"联系地址"`
	Pid                int64       `json:"pid"                description:"上级管理员ID"`
	Level              int         `json:"level"              description:"关系树等级"`
	Tree               string      `json:"tree"               description:"关系树"`
	InviteCode         string      `json:"inviteCode"         description:"邀请码"`
	Cash               *gjson.Json `json:"cash"               description:"提现配置"`
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       description:"最后活跃时间"`
	Remark             string      `json:"remark"             description:"备注"`
	Status             int         `json:"status"             description:"状态"`
	CreatedAt          *gtime.Time `json:"createdAt"          description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          description:"修改时间"`
}
