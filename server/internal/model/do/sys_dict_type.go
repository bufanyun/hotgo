// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure of table hg_sys_dict_type for DAO operations like Where/Data.
type SysDictType struct {
	g.Meta    `orm:"table:hg_sys_dict_type, do:true"`
	Id        any         // 字典类型ID
	Pid       any         // 父类字典类型ID
	Name      any         // 字典类型名称
	Type      any         // 字典类型
	Sort      any         // 排序
	Remark    any         // 备注
	Status    any         // 字典类型状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
