// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPost is the golang structure for table admin_post.
type AdminPost struct {
	Id        int64       `json:"id"        description:"岗位ID"`
	Code      string      `json:"code"      description:"岗位编码"`
	Name      string      `json:"name"      description:"岗位名称"`
	Remark    string      `json:"remark"    description:"备注"`
	Sort      int         `json:"sort"      description:"排序"`
	Status    int         `json:"status"    description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
