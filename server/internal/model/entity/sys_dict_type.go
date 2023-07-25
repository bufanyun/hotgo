// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	Id        int64       `json:"id"        description:"字典类型ID"`
	Pid       int64       `json:"pid"       description:"父类字典类型ID"`
	Name      string      `json:"name"      description:"字典类型名称"`
	Type      string      `json:"type"      description:"字典类型"`
	Sort      int         `json:"sort"      description:"排序"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    int         `json:"status"    description:"字典类型状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
