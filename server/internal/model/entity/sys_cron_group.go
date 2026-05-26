// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCronGroup is the golang structure for table sys_cron_group.
type SysCronGroup struct {
	Id        int64       `json:"id"        orm:"id"         description:"任务分组ID"`
	Pid       int64       `json:"pid"       orm:"pid"        description:"父类任务分组ID"`
	Name      string      `json:"name"      orm:"name"       description:"分组名称"`
	IsDefault int         `json:"isDefault" orm:"is_default" description:"是否默认"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`
	Status    int         `json:"status"    orm:"status"     description:"分组状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
