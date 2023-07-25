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
	Id          interface{} // 变动ID
	MemberId    interface{} // 管理员ID
	AppId       interface{} // 应用id
	AddonsName  interface{} // 插件名称
	CreditType  interface{} // 变动类型
	CreditGroup interface{} // 变动组别
	BeforeNum   interface{} // 变动前
	Num         interface{} // 变动数据
	AfterNum    interface{} // 变动后
	Remark      interface{} // 备注
	Ip          interface{} // 操作人IP
	MapId       interface{} // 关联ID
	Status      interface{} // 状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
