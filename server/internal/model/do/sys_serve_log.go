// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLog is the golang structure of table hg_sys_serve_log for DAO operations like Where/Data.
type SysServeLog struct {
	g.Meta      `orm:"table:hg_sys_serve_log, do:true"`
	Id          interface{} // 日志ID
	TraceId     interface{} // 链路ID
	LevelFormat interface{} // 日志级别
	Content     interface{} // 日志内容
	Stack       *gjson.Json // 打印堆栈
	Line        interface{} // 调用行
	TriggerNs   interface{} // 触发时间(ns)
	Status      interface{} // 状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
