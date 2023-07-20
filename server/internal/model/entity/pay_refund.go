// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PayRefund is the golang structure for table pay_refund.
type PayRefund struct {
	Id            uint64      `json:"id"            description:"主键ID"`
	MemberId      int64       `json:"memberId"      description:"会员ID"`
	AppId         string      `json:"appId"         description:"应用ID"`
	OrderSn       string      `json:"orderSn"       description:"业务订单号"`
	RefundTradeNo string      `json:"refundTradeNo" description:"退款交易号"`
	RefundMoney   float64     `json:"refundMoney"   description:"退款金额"`
	RefundWay     int         `json:"refundWay"     description:"退款方式"`
	Ip            string      `json:"ip"            description:"申请者IP"`
	Reason        string      `json:"reason"        description:"申请退款原因"`
	Remark        string      `json:"remark"        description:"退款备注"`
	Status        int         `json:"status"        description:"退款状态"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"申请时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
}
