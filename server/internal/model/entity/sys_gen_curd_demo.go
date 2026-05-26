// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCurdDemo is the golang structure for table sys_gen_curd_demo.
type SysGenCurdDemo struct {
	Id          int64       `json:"id"          orm:"id"          description:"ID"`
	CategoryId  int64       `json:"categoryId"  orm:"category_id" description:"分类ID"`
	Title       string      `json:"title"       orm:"title"       description:"标题"`
	Description string      `json:"description" orm:"description" description:"描述"`
	Content     string      `json:"content"     orm:"content"     description:"内容"`
	Image       string      `json:"image"       orm:"image"       description:"单图"`
	Attachfile  string      `json:"attachfile"  orm:"attachfile"  description:"附件"`
	CityId      int64       `json:"cityId"      orm:"city_id"     description:"所在城市"`
	Switch      int         `json:"switch"      orm:"switch"      description:"显示开关"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	Status      int         `json:"status"      orm:"status"      description:"状态"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:"创建者"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:"更新者"`
	DeletedBy   int64       `json:"deletedBy"   orm:"deleted_by"  description:"删除者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`
}
