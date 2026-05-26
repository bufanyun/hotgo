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
	Id          any         // 日志ID
	TraceId     any         // 链路ID
	LevelFormat any         // 日志级别
	Content     any         // 日志内容
	Stack       *gjson.Json // 打印堆栈
	Line        any         // 调用行
	TriggerNs   any         // 触发时间(ns)
	Status      any         // 状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
