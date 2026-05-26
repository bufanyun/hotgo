// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMember is the golang structure of table hg_admin_member for DAO operations like Where/Data.
type AdminMember struct {
	g.Meta             `orm:"table:hg_admin_member, do:true"`
	Id                 any         // 管理员ID
	DeptId             any         // 部门ID
	RoleId             any         // 角色ID
	RealName           any         // 真实姓名
	Username           any         // 帐号
	PasswordHash       any         // 密码
	Salt               any         // 密码盐
	PasswordResetToken any         // 密码重置令牌
	Integral           any         // 积分
	Balance            any         // 余额
	Avatar             any         // 头像
	Sex                any         // 性别
	Qq                 any         // qq
	Email              any         // 邮箱
	Mobile             any         // 手机号码
	Birthday           *gtime.Time // 生日
	CityId             any         // 城市编码
	Address            any         // 联系地址
	Pid                any         // 上级管理员ID
	Level              any         // 关系树等级
	Tree               any         // 关系树
	InviteCode         any         // 邀请码
	Cash               *gjson.Json // 提现配置
	LastActiveAt       *gtime.Time // 最后活跃时间
	Remark             any         // 备注
	Status             any         // 状态
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 修改时间
}
