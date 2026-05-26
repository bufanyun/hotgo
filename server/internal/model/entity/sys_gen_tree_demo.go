// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenTreeDemo is the golang structure for table sys_gen_tree_demo.
type SysGenTreeDemo struct {
	Id          int64       `json:"id"          orm:"id"          description:"ID"`
	Pid         int64       `json:"pid"         orm:"pid"         description:"上级ID"`
	Level       int         `json:"level"       orm:"level"       description:"关系树级别"`
	Tree        string      `json:"tree"        orm:"tree"        description:"关系树"`
	CategoryId  int64       `json:"categoryId"  orm:"category_id" description:"分类ID"`
	Title       string      `json:"title"       orm:"title"       description:"标题"`
	Description string      `json:"description" orm:"description" description:"描述"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	Status      int         `json:"status"      orm:"status"      description:"状态"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:"创建者"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:"更新者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`
}
