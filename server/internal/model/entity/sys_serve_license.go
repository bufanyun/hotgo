// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLicense is the golang structure for table sys_serve_license.
type SysServeLicense struct {
	Id           int64       `json:"id"           description:"许可ID"`
	Group        string      `json:"group"        description:"分组"`
	Name         string      `json:"name"         description:"许可名称"`
	Appid        string      `json:"appid"        description:"应用ID"`
	SecretKey    string      `json:"secretKey"    description:"应用秘钥"`
	RemoteAddr   string      `json:"remoteAddr"   description:"最后连接地址"`
	OnlineLimit  int         `json:"onlineLimit"  description:"在线限制"`
	LoginTimes   int64       `json:"loginTimes"   description:"登录次数"`
	LastLoginAt  *gtime.Time `json:"lastLoginAt"  description:"最后登录时间"`
	LastActiveAt *gtime.Time `json:"lastActiveAt" description:"最后心跳"`
	Routes       *gjson.Json `json:"routes"       description:"路由表，空使用默认分组路由"`
	AllowedIps   string      `json:"allowedIps"   description:"IP白名单"`
	EndAt        *gtime.Time `json:"endAt"        description:"授权有效期"`
	Remark       string      `json:"remark"       description:"备注"`
	Status       int         `json:"status"       description:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:"修改时间"`
}
