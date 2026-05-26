// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAddonsConfig is the golang structure of table hg_sys_addons_config for DAO operations like Where/Data.
type SysAddonsConfig struct {
	g.Meta       `orm:"table:hg_sys_addons_config, do:true"`
	Id           any         // 配置ID
	AddonName    any         // 插件名称
	Group        any         // 分组
	Name         any         // 参数名称
	Type         any         // 键值类型:string,int,uint,bool,TIMESTAMP,date
	Key          any         // 参数键名
	Value        any         // 参数键值
	DefaultValue any         // 默认值
	Sort         any         // 排序
	Tip          any         // 变量描述
	IsDefault    any         // 是否为系统默认
	Status       any         // 状态
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
