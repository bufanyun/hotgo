// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	Id        int64       `json:"id"        orm:"id"         description:"字典类型ID"`
	Pid       int64       `json:"pid"       orm:"pid"        description:"父类字典类型ID"`
	Name      string      `json:"name"      orm:"name"       description:"字典类型名称"`
	Type      string      `json:"type"      orm:"type"       description:"字典类型"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`
	Status    int         `json:"status"    orm:"status"     description:"字典类型状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
