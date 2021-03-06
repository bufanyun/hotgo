// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNotice is the golang structure for table admin_notice.
type AdminNotice struct {
	Id        int64       `json:"id"         description:"公告ID"`
	Title     string      `json:"title"      description:"公告标题"`
	Type      string      `json:"type"       description:"公告类型（1通知 2公告）"`
	Content   string      `json:"content"    description:"公告内容"`
	Remark    string      `json:"remark"     description:"备注"`
	Status    string      `json:"status"     description:"公告状态"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"更新时间"`
}
