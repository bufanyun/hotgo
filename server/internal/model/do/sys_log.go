// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure of table hg_sys_log for DAO operations like Where/Data.
type SysLog struct {
	g.Meta     `orm:"table:hg_sys_log, do:true"`
	Id         interface{} // 日志ID
	ReqId      interface{} // 对外ID
	AppId      interface{} // 应用ID
	MerchantId interface{} // 商户ID
	MemberId   interface{} // 用户ID
	Method     interface{} // 提交类型
	Module     interface{} // 访问模块
	Url        interface{} // 提交url
	GetData    *gjson.Json // get数据
	PostData   *gjson.Json // post数据
	HeaderData *gjson.Json // header数据
	Ip         interface{} // IP地址
	ProvinceId interface{} // 省编码
	CityId     interface{} // 市编码
	ErrorCode  interface{} // 报错code
	ErrorMsg   interface{} // 对外错误提示
	ErrorData  *gjson.Json // 报错日志
	UserAgent  interface{} // UA信息
	TakeUpTime interface{} // 请求耗时
	Timestamp  interface{} // 响应时间
	Status     interface{} // 状态
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
