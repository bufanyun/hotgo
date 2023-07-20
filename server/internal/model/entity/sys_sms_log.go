// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsLog is the golang structure for table sys_sms_log.
type SysSmsLog struct {
	Id        int64       `json:"id"        description:"主键"`
	Event     string      `json:"event"     description:"事件"`
	Mobile    string      `json:"mobile"    description:"手机号"`
	Code      string      `json:"code"      description:"验证码或短信内容"`
	Times     int64       `json:"times"     description:"验证次数"`
	Ip        string      `json:"ip"        description:"ip地址"`
	Status    int         `json:"status"    description:"状态(1未验证,2已验证)"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
