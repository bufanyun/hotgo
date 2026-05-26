// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TestCategory is the golang structure for table test_category.
type TestCategory struct {
	Id          int64       `json:"id"          orm:"id"          description:"分类ID"`
	Name        string      `json:"name"        orm:"name"        description:"分类名称"`
	ShortName   string      `json:"shortName"   orm:"short_name"  description:"简称"`
	Description string      `json:"description" orm:"description" description:"描述"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	Remark      string      `json:"remark"      orm:"remark"      description:"备注"`
	Status      int         `json:"status"      orm:"status"      description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`
}
