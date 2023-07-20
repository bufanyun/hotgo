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
	Id        interface{} // 任务分组ID
	Pid       interface{} // 父类任务分组ID
	Name      interface{} // 分组名称
	IsDefault interface{} // 是否默认
	Sort      interface{} // 排序
	Remark    interface{} // 备注
	Status    interface{} // 分组状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
