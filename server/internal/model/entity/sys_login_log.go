// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginLog is the golang structure for table sys_login_log.
type SysLoginLog struct {
	Id         int64       `json:"id"         orm:"id"          description:"日志ID"`
	ReqId      string      `json:"reqId"      orm:"req_id"      description:"请求ID"`
	MemberId   int64       `json:"memberId"   orm:"member_id"   description:"用户ID"`
	Username   string      `json:"username"   orm:"username"    description:"用户名"`
	Response   *gjson.Json `json:"response"   orm:"response"    description:"响应数据"`
	LoginAt    *gtime.Time `json:"loginAt"    orm:"login_at"    description:"登录时间"`
	LoginIp    string      `json:"loginIp"    orm:"login_ip"    description:"登录IP"`
	ProvinceId int64       `json:"provinceId" orm:"province_id" description:"省编码"`
	CityId     int64       `json:"cityId"     orm:"city_id"     description:"市编码"`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"  description:"UA信息"`
	ErrMsg     string      `json:"errMsg"     orm:"err_msg"     description:"错误提示"`
	Status     int         `json:"status"     orm:"status"      description:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"修改时间"`
}
