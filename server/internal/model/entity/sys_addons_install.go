// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAddonsInstall is the golang structure for table sys_addons_install.
type SysAddonsInstall struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`
	Name      string      `json:"name"      orm:"name"       description:"插件名称"`
	Version   string      `json:"version"   orm:"version"    description:"版本号"`
	Status    int         `json:"status"    orm:"status"     description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
