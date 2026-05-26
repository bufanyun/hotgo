// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenTreeDemo is the golang structure of table hg_sys_gen_tree_demo for DAO operations like Where/Data.
type SysGenTreeDemo struct {
	g.Meta      `orm:"table:hg_sys_gen_tree_demo, do:true"`
	Id          any         // ID
	Pid         any         // 上级ID
	Level       any         // 关系树级别
	Tree        any         // 关系树
	CategoryId  any         // 分类ID
	Title       any         // 标题
	Description any         // 描述
	Sort        any         // 排序
	Status      any         // 状态
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
