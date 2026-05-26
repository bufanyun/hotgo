// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PayRefund is the golang structure for table pay_refund.
type PayRefund struct {
	Id            int64       `json:"id"            orm:"id"              description:"主键ID"`
	MemberId      int64       `json:"memberId"      orm:"member_id"       description:"会员ID"`
	AppId         string      `json:"appId"         orm:"app_id"          description:"应用ID"`
	OrderSn       string      `json:"orderSn"       orm:"order_sn"        description:"业务订单号"`
	RefundTradeNo string      `json:"refundTradeNo" orm:"refund_trade_no" description:"退款交易号"`
	RefundMoney   float64     `json:"refundMoney"   orm:"refund_money"    description:"退款金额"`
	RefundWay     int         `json:"refundWay"     orm:"refund_way"      description:"退款方式"`
	Ip            string      `json:"ip"            orm:"ip"              description:"申请者IP"`
	Reason        string      `json:"reason"        orm:"reason"          description:"申请退款原因"`
	Remark        string      `json:"remark"        orm:"remark"          description:"退款备注"`
	Status        int         `json:"status"        orm:"status"          description:"退款状态"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"申请时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`
}
