// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCronGroup is the golang structure of table hg_sys_cron_group for DAO operations like Where/Data.
type SysCronGroup struct {
	g.Meta    `orm:"table:hg_sys_cron_group, do:true"`
	Id        any         // 任务分组ID
	Pid       any         // 父类任务分组ID
	Name      any         // 分组名称
	IsDefault any         // 是否默认
	Sort      any         // 排序
	Remark    any         // 备注
	Status    any         // 分组状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
