// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminDept is the golang structure of table hg_admin_dept for DAO operations like Where/Data.
type AdminDept struct {
	g.Meta    `orm:"table:hg_admin_dept, do:true"`
	Id        any         // 部门ID
	Pid       any         // 父部门ID
	Name      any         // 部门名称
	Code      any         // 部门编码
	Type      any         // 部门类型
	Leader    any         // 负责人
	Phone     any         // 联系电话
	Email     any         // 邮箱
	Level     any         // 关系树等级
	Tree      any         // 关系树
	Sort      any         // 排序
	Status    any         // 部门状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
