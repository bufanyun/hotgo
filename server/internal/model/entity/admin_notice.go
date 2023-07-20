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
	Id        int64       `json:"id"        description:"公告ID"`
	Title     string      `json:"title"     description:"公告标题"`
	Type      int64       `json:"type"      description:"公告类型"`
	Tag       int         `json:"tag"       description:"标签"`
	Content   string      `json:"content"   description:"公告内容"`
	Receiver  *gjson.Json `json:"receiver"  description:"接收者"`
	Remark    string      `json:"remark"    description:"备注"`
	Sort      int         `json:"sort"      description:"排序"`
	Status    int         `json:"status"    description:"公告状态"`
	CreatedBy int64       `json:"createdBy" description:"发送人"`
	UpdatedBy int64       `json:"updatedBy" description:"修改人"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
