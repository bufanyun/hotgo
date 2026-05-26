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
	Id           int64       `json:"id"           orm:"id"             description:"许可ID"`
	Group        string      `json:"group"        orm:"group"          description:"分组"`
	Name         string      `json:"name"         orm:"name"           description:"许可名称"`
	Appid        string      `json:"appid"        orm:"appid"          description:"应用ID"`
	SecretKey    string      `json:"secretKey"    orm:"secret_key"     description:"应用秘钥"`
	RemoteAddr   string      `json:"remoteAddr"   orm:"remote_addr"    description:"最后连接地址"`
	OnlineLimit  int         `json:"onlineLimit"  orm:"online_limit"   description:"在线限制"`
	LoginTimes   int64       `json:"loginTimes"   orm:"login_times"    description:"登录次数"`
	LastLoginAt  *gtime.Time `json:"lastLoginAt"  orm:"last_login_at"  description:"最后登录时间"`
	LastActiveAt *gtime.Time `json:"lastActiveAt" orm:"last_active_at" description:"最后心跳"`
	Routes       *gjson.Json `json:"routes"       orm:"routes"         description:"路由表，空使用默认分组路由"`
	AllowedIps   string      `json:"allowedIps"   orm:"allowed_ips"    description:"IP白名单"`
	EndAt        *gtime.Time `json:"endAt"        orm:"end_at"         description:"授权有效期"`
	Remark       string      `json:"remark"       orm:"remark"         description:"备注"`
	Status       int         `json:"status"       orm:"status"         description:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     description:"修改时间"`
}
