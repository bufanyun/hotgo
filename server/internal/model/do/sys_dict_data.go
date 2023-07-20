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
	Id        interface{} // 字典数据ID
	Label     interface{} // 字典标签
	Value     interface{} // 字典键值
	ValueType interface{} // 键值数据类型：string,int,uint,bool,datetime,date
	Type      interface{} // 字典类型
	ListClass interface{} // 表格回显样式
	IsDefault interface{} // 是否为系统默认
	Sort      interface{} // 字典排序
	Remark    interface{} // 备注
	Status    interface{} // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
