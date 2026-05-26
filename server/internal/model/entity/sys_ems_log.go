// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysEmsLog is the golang structure for table sys_ems_log.
type SysEmsLog struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`
	Event     string      `json:"event"     orm:"event"      description:"事件"`
	Email     string      `json:"email"     orm:"email"      description:"邮箱地址，多个用;隔开"`
	Code      string      `json:"code"      orm:"code"       description:"验证码"`
	Times     int64       `json:"times"     orm:"times"      description:"验证次数"`
	Content   string      `json:"content"   orm:"content"    description:"邮件内容"`
	Ip        string      `json:"ip"        orm:"ip"         description:"ip地址"`
	Status    int         `json:"status"    orm:"status"     description:"状态(1未验证,2已验证)"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
