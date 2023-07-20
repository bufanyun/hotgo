// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOrder is the golang structure for table admin_order.
type AdminOrder struct {
	Id                 int64       `json:"id"                 description:"主键"`
	MemberId           int64       `json:"memberId"           description:"管理员id"`
	OrderType          string      `json:"orderType"          description:"订单类型"`
	ProductId          int64       `json:"productId"          description:"产品id"`
	OrderSn            string      `json:"orderSn"            description:"关联订单号"`
	Money              float64     `json:"money"              description:"充值金额"`
	Remark             string      `json:"remark"             description:"备注"`
	RefundReason       string      `json:"refundReason"       description:"退款原因"`
	RejectRefundReason string      `json:"rejectRefundReason" description:"拒绝退款原因"`
	Status             int         `json:"status"             description:"状态"`
	CreatedAt          *gtime.Time `json:"createdAt"          description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          description:"修改时间"`
}
