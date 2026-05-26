// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCreditsLog is the golang structure of table hg_admin_credits_log for DAO operations like Where/Data.
type AdminCreditsLog struct {
	g.Meta      `orm:"table:hg_admin_credits_log, do:true"`
	Id          any         // 变动ID
	MemberId    any         // 管理员ID
	AppId       any         // 应用id
	AddonsName  any         // 插件名称
	CreditType  any         // 变动类型
	CreditGroup any         // 变动组别
	BeforeNum   any         // 变动前
	Num         any         // 变动数据
	AfterNum    any         // 变动后
	Remark      any         // 备注
	Ip          any         // 操作人IP
	MapId       any         // 关联ID
	Status      any         // 状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
