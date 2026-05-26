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
	Id         any         // 日志ID
	ReqId      any         // 对外ID
	AppId      any         // 应用ID
	MerchantId any         // 商户ID
	MemberId   any         // 用户ID
	Method     any         // 提交类型
	Module     any         // 访问模块
	Url        any         // 提交url
	GetData    *gjson.Json // get数据
	PostData   *gjson.Json // post数据
	HeaderData *gjson.Json // header数据
	Ip         any         // IP地址
	ProvinceId any         // 省编码
	CityId     any         // 市编码
	ErrorCode  any         // 报错code
	ErrorMsg   any         // 对外错误提示
	ErrorData  *gjson.Json // 报错日志
	UserAgent  any         // UA信息
	TakeUpTime any         // 请求耗时
	Timestamp  any         // 响应时间
	Status     any         // 状态
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
