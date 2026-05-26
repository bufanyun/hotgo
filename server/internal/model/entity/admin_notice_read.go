// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNoticeRead is the golang structure for table admin_notice_read.
type AdminNoticeRead struct {
	Id        int64       `json:"id"        orm:"id"         description:"记录ID"`
	NoticeId  int64       `json:"noticeId"  orm:"notice_id"  description:"公告ID"`
	MemberId  int64       `json:"memberId"  orm:"member_id"  description:"会员ID"`
	Clicks    int         `json:"clicks"    orm:"clicks"     description:"已读次数"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"阅读时间"`
}
