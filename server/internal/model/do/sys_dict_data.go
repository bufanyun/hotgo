// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure of table hg_sys_dict_data for DAO operations like Where/Data.
type SysDictData struct {
	g.Meta    `orm:"table:hg_sys_dict_data, do:true"`
	Id        any         // 字典数据ID
	Label     any         // 字典标签
	Value     any         // 字典键值
	ValueType any         // 键值数据类型：string,int,uint,bool,TIMESTAMP,date
	Type      any         // 字典类型
	ListClass any         // 表格回显样式
	IsDefault any         // 是否为系统默认
	Sort      any         // 字典排序
	Remark    any         // 备注
	Status    any         // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
