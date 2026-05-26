// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOrder is the golang structure for table admin_order.
type AdminOrder struct {
	Id                 int64       `json:"id"                 orm:"id"                   description:"主键"`
	MemberId           int64       `json:"memberId"           orm:"member_id"            description:"管理员id"`
	OrderType          string      `json:"orderType"          orm:"order_type"           description:"订单类型"`
	ProductId          int64       `json:"productId"          orm:"product_id"           description:"产品id"`
	OrderSn            string      `json:"orderSn"            orm:"order_sn"             description:"关联订单号"`
	Money              float64     `json:"money"              orm:"money"                description:"充值金额"`
	Remark             string      `json:"remark"             orm:"remark"               description:"备注"`
	RefundReason       string      `json:"refundReason"       orm:"refund_reason"        description:"退款原因"`
	RejectRefundReason string      `json:"rejectRefundReason" orm:"reject_refund_reason" description:"拒绝退款原因"`
	Status             int         `json:"status"             orm:"status"               description:"状态"`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:"修改时间"`
}
