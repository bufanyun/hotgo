// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCronGroup is the golang structure for table sys_cron_group.
type SysCronGroup struct {
	Id        int64       `json:"id"        description:"任务分组ID"`
	Pid       int64       `json:"pid"       description:"父类任务分组ID"`
	Name      string      `json:"name"      description:"分组名称"`
	IsDefault int         `json:"isDefault" description:"是否默认"`
	Sort      int         `json:"sort"      description:"排序"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    int         `json:"status"    description:"分组状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
