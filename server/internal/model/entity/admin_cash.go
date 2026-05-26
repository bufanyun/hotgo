// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCash is the golang structure for table admin_cash.
type AdminCash struct {
	Id        int64       `json:"id"        orm:"id"         description:"ID"`
	MemberId  int64       `json:"memberId"  orm:"member_id"  description:"管理员ID"`
	Money     float64     `json:"money"     orm:"money"      description:"提现金额"`
	Fee       float64     `json:"fee"       orm:"fee"        description:"手续费"`
	LastMoney float64     `json:"lastMoney" orm:"last_money" description:"最终到账金额"`
	Ip        string      `json:"ip"        orm:"ip"         description:"申请人IP"`
	Status    int64       `json:"status"    orm:"status"     description:"状态码"`
	Msg       string      `json:"msg"       orm:"msg"        description:"处理结果"`
	HandleAt  *gtime.Time `json:"handleAt"  orm:"handle_at"  description:"处理时间"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"申请时间"`
}
