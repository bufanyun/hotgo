// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCodes is the golang structure for table sys_gen_codes.
type SysGenCodes struct {
	Id            int64       `json:"id"            orm:"id"             description:"生成ID"`
	GenType       int         `json:"genType"       orm:"gen_type"       description:"生成类型"`
	GenTemplate   int         `json:"genTemplate"   orm:"gen_template"   description:"生成模板"`
	VarName       string      `json:"varName"       orm:"var_name"       description:"实体命名"`
	Options       *gjson.Json `json:"options"       orm:"options"        description:"配置选项"`
	DbName        string      `json:"dbName"        orm:"db_name"        description:"数据库名称"`
	TableName     string      `json:"tableName"     orm:"table_name"     description:"主表名称"`
	TableComment  string      `json:"tableComment"  orm:"table_comment"  description:"主表注释"`
	DaoName       string      `json:"daoName"       orm:"dao_name"       description:"主表dao模型"`
	MasterColumns *gjson.Json `json:"masterColumns" orm:"master_columns" description:"主表字段"`
	AddonName     string      `json:"addonName"     orm:"addon_name"     description:"插件名称"`
	Status        int         `json:"status"        orm:"status"         description:"生成状态"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
}
