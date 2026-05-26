// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPost is the golang structure for table admin_post.
type AdminPost struct {
	Id        int64       `json:"id"        orm:"id"         description:"岗位ID"`
	Code      string      `json:"code"      orm:"code"       description:"岗位编码"`
	Name      string      `json:"name"      orm:"name"       description:"岗位名称"`
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Status    int         `json:"status"    orm:"status"     description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
