// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysEmsLog is the golang structure for table sys_ems_log.
type SysEmsLog struct {
	Id        int64       `json:"id"        description:"主键"`
	Event     string      `json:"event"     description:"事件"`
	Email     string      `json:"email"     description:"邮箱地址，多个用;隔开"`
	Code      string      `json:"code"      description:"验证码"`
	Times     int64       `json:"times"     description:"验证次数"`
	Content   string      `json:"content"   description:"邮件内容"`
	Ip        string      `json:"ip"        description:"ip地址"`
	Status    int         `json:"status"    description:"状态(1未验证,2已验证)"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
