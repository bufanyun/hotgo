// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCron is the golang structure of table hg_sys_cron for DAO operations like Where/Data.
type SysCron struct {
	g.Meta    `orm:"table:hg_sys_cron, do:true"`
	Id        any         // 任务ID
	GroupId   any         // 分组ID
	Title     any         // 任务标题
	Name      any         // 任务方法
	Params    any         // 函数参数
	Pattern   any         // 表达式
	Policy    any         // 策略
	Count     any         // 执行次数
	Sort      any         // 排序
	Remark    any         // 备注
	Status    any         // 任务状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
