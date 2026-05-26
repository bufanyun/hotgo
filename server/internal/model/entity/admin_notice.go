// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNotice is the golang structure for table admin_notice.
type AdminNotice struct {
	Id        int64       `json:"id"        orm:"id"         description:"公告ID"`
	Title     string      `json:"title"     orm:"title"      description:"公告标题"`
	Type      int64       `json:"type"      orm:"type"       description:"公告类型"`
	Tag       int         `json:"tag"       orm:"tag"        description:"标签"`
	Content   string      `json:"content"   orm:"content"    description:"公告内容"`
	Receiver  *gjson.Json `json:"receiver"  orm:"receiver"   description:"接收者"`
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Status    int         `json:"status"    orm:"status"     description:"公告状态"`
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"发送人"`
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"修改人"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
