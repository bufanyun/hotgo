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
	Id         int64       `json:"id"         description:"日志ID"`
	ReqId      string      `json:"reqId"      description:"对外ID"`
	AppId      string      `json:"appId"      description:"应用ID"`
	MerchantId uint64      `json:"merchantId" description:"商户ID"`
	MemberId   int64       `json:"memberId"   description:"用户ID"`
	Method     string      `json:"method"     description:"提交类型"`
	Module     string      `json:"module"     description:"访问模块"`
	Url        string      `json:"url"        description:"提交url"`
	GetData    *gjson.Json `json:"getData"    description:"get数据"`
	PostData   *gjson.Json `json:"postData"   description:"post数据"`
	HeaderData *gjson.Json `json:"headerData" description:"header数据"`
	Ip         string      `json:"ip"         description:"IP地址"`
	ProvinceId int64       `json:"provinceId" description:"省编码"`
	CityId     int64       `json:"cityId"     description:"市编码"`
	ErrorCode  int         `json:"errorCode"  description:"报错code"`
	ErrorMsg   string      `json:"errorMsg"   description:"对外错误提示"`
	ErrorData  *gjson.Json `json:"errorData"  description:"报错日志"`
	UserAgent  string      `json:"userAgent"  description:"UA信息"`
	TakeUpTime int64       `json:"takeUpTime" description:"请求耗时"`
	Timestamp  int64       `json:"timestamp"  description:"响应时间"`
	Status     int         `json:"status"     description:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"修改时间"`
}
