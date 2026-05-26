// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginLog is the golang structure of table hg_sys_login_log for DAO operations like Where/Data.
type SysLoginLog struct {
	g.Meta     `orm:"table:hg_sys_login_log, do:true"`
	Id         any         // 日志ID
	ReqId      any         // 请求ID
	MemberId   any         // 用户ID
	Username   any         // 用户名
	Response   *gjson.Json // 响应数据
	LoginAt    *gtime.Time // 登录时间
	LoginIp    any         // 登录IP
	ProvinceId any         // 省编码
	CityId     any         // 市编码
	UserAgent  any         // UA信息
	ErrMsg     any         // 错误提示
	Status     any         // 状态
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
