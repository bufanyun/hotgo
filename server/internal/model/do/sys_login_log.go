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
	g.Meta    `orm:"table:hg_sys_login_log, do:true"`
	Id        interface{} // 日志ID
	ReqId     interface{} // 请求ID
	MemberId  interface{} // 用户ID
	Username  interface{} // 用户名
	Response  *gjson.Json // 响应数据
	LoginAt   *gtime.Time // 登录时间
	LoginIp   interface{} // 登录IP
	ErrMsg    interface{} // 错误提示
	Status    interface{} // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
