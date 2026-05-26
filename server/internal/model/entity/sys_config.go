// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	Id           int64       `json:"id"           orm:"id"            description:"配置ID"`
	Group        string      `json:"group"        orm:"group"         description:"配置分组"`
	Name         string      `json:"name"         orm:"name"          description:"参数名称"`
	Type         string      `json:"type"         orm:"type"          description:"键值类型:string,int,uint,bool,TIMESTAMP,date"`
	Key          string      `json:"key"          orm:"key"           description:"参数键名"`
	Value        string      `json:"value"        orm:"value"         description:"参数键值"`
	DefaultValue string      `json:"defaultValue" orm:"default_value" description:"默认值"`
	Sort         int         `json:"sort"         orm:"sort"          description:"排序"`
	Tip          string      `json:"tip"          orm:"tip"           description:"变量描述"`
	IsDefault    int         `json:"isDefault"    orm:"is_default"    description:"是否为系统默认"`
	Status       int         `json:"status"       orm:"status"        description:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
}
