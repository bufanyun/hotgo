// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	Id           int64       `json:"id"           description:"配置ID"`
	Group        string      `json:"group"        description:"配置分组"`
	Name         string      `json:"name"         description:"参数名称"`
	Type         string      `json:"type"         description:"键值类型:string,int,uint,bool,datetime,date"`
	Key          string      `json:"key"          description:"参数键名"`
	Value        string      `json:"value"        description:"参数键值"`
	DefaultValue string      `json:"defaultValue" description:"默认值"`
	Sort         int         `json:"sort"         description:"排序"`
	Tip          string      `json:"tip"          description:"变量描述"`
	IsDefault    int         `json:"isDefault"    description:"是否为系统默认"`
	Status       int         `json:"status"       description:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:"更新时间"`
}
