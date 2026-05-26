// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCurdDemo is the golang structure of table hg_sys_gen_curd_demo for DAO operations like Where/Data.
type SysGenCurdDemo struct {
	g.Meta      `orm:"table:hg_sys_gen_curd_demo, do:true"`
	Id          any         // ID
	CategoryId  any         // 分类ID
	Title       any         // 标题
	Description any         // 描述
	Content     any         // 内容
	Image       any         // 单图
	Attachfile  any         // 附件
	CityId      any         // 所在城市
	Switch      any         // 显示开关
	Sort        any         // 排序
	Status      any         // 状态
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
	DeletedBy   any         // 删除者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
