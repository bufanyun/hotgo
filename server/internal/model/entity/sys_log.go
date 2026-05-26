// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure for table sys_log.
type SysLog struct {
	Id         int64       `json:"id"         orm:"id"           description:"日志ID"`
	ReqId      string      `json:"reqId"      orm:"req_id"       description:"对外ID"`
	AppId      string      `json:"appId"      orm:"app_id"       description:"应用ID"`
	MerchantId int64       `json:"merchantId" orm:"merchant_id"  description:"商户ID"`
	MemberId   int64       `json:"memberId"   orm:"member_id"    description:"用户ID"`
	Method     string      `json:"method"     orm:"method"       description:"提交类型"`
	Module     string      `json:"module"     orm:"module"       description:"访问模块"`
	Url        string      `json:"url"        orm:"url"          description:"提交url"`
	GetData    *gjson.Json `json:"getData"    orm:"get_data"     description:"get数据"`
	PostData   *gjson.Json `json:"postData"   orm:"post_data"    description:"post数据"`
	HeaderData *gjson.Json `json:"headerData" orm:"header_data"  description:"header数据"`
	Ip         string      `json:"ip"         orm:"ip"           description:"IP地址"`
	ProvinceId int64       `json:"provinceId" orm:"province_id"  description:"省编码"`
	CityId     int64       `json:"cityId"     orm:"city_id"      description:"市编码"`
	ErrorCode  int         `json:"errorCode"  orm:"error_code"   description:"报错code"`
	ErrorMsg   string      `json:"errorMsg"   orm:"error_msg"    description:"对外错误提示"`
	ErrorData  *gjson.Json `json:"errorData"  orm:"error_data"   description:"报错日志"`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"   description:"UA信息"`
	TakeUpTime int64       `json:"takeUpTime" orm:"take_up_time" description:"请求耗时"`
	Timestamp  int64       `json:"timestamp"  orm:"timestamp"    description:"响应时间"`
	Status     int         `json:"status"     orm:"status"       description:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"   description:"修改时间"`
}
