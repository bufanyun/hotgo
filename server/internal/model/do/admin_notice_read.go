// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNoticeRead is the golang structure of table hg_admin_notice_read for DAO operations like Where/Data.
type AdminNoticeRead struct {
	g.Meta    `orm:"table:hg_admin_notice_read, do:true"`
	Id        any         // 记录ID
	NoticeId  any         // 公告ID
	MemberId  any         // 会员ID
	Clicks    any         // 已读次数
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 阅读时间
}
