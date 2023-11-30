// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCron is the golang structure for table sys_cron.
type SysCron struct {
	Id        int64       `json:"id"        description:"任务ID"`
	GroupId   int64       `json:"groupId"   description:"分组ID"`
	Title     string      `json:"title"     description:"任务标题"`
	Name      string      `json:"name"      description:"任务方法"`
	Params    string      `json:"params"    description:"函数参数"`
	Pattern   string      `json:"pattern"   description:"表达式"`
	Policy    int64       `json:"policy"    description:"策略"`
	Count     int64       `json:"count"     description:"执行次数"`
	Sort      int         `json:"sort"      description:"排序"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    int         `json:"status"    description:"任务状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
