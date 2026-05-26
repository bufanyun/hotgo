// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayLog is the golang structure for table pay_log.
type PayLog struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键"`
	MemberId      int64       `json:"memberId"      orm:"member_id"      description:"会员ID"`
	AppId         string      `json:"appId"         orm:"app_id"         description:"应用ID"`
	AddonsName    string      `json:"addonsName"    orm:"addons_name"    description:"插件名称"`
	OrderSn       string      `json:"orderSn"       orm:"order_sn"       description:"关联订单号"`
	OrderGroup    string      `json:"orderGroup"    orm:"order_group"    description:"组别[默认统一支付类型]"`
	Openid        string      `json:"openid"        orm:"openid"         description:"openid"`
	MchId         string      `json:"mchId"         orm:"mch_id"         description:"商户支付账户"`
	Subject       string      `json:"subject"       orm:"subject"        description:"订单标题"`
	Detail        *gjson.Json `json:"detail"        orm:"detail"         description:"支付商品详情"`
	AuthCode      string      `json:"authCode"      orm:"auth_code"      description:"刷卡码"`
	OutTradeNo    string      `json:"outTradeNo"    orm:"out_trade_no"   description:"商户订单号"`
	TransactionId string      `json:"transactionId" orm:"transaction_id" description:"交易号"`
	PayType       string      `json:"payType"       orm:"pay_type"       description:"支付类型"`
	PayAmount     float64     `json:"payAmount"     orm:"pay_amount"     description:"支付金额"`
	ActualAmount  float64     `json:"actualAmount"  orm:"actual_amount"  description:"实付金额"`
	PayStatus     int         `json:"payStatus"     orm:"pay_status"     description:"支付状态"`
	PayAt         *gtime.Time `json:"payAt"         orm:"pay_at"         description:"支付时间"`
	TradeType     string      `json:"tradeType"     orm:"trade_type"     description:"交易类型"`
	RefundSn      string      `json:"refundSn"      orm:"refund_sn"      description:"退款单号"`
	IsRefund      int         `json:"isRefund"      orm:"is_refund"      description:"是否退款"`
	Custom        string      `json:"custom"        orm:"custom"         description:"自定义参数"`
	CreateIp      string      `json:"createIp"      orm:"create_ip"      description:"创建者IP"`
	PayIp         string      `json:"payIp"         orm:"pay_ip"         description:"支付者IP"`
	NotifyUrl     string      `json:"notifyUrl"     orm:"notify_url"     description:"支付通知回调地址"`
	ReturnUrl     string      `json:"returnUrl"     orm:"return_url"     description:"买家付款成功跳转地址"`
	TraceIds      *gjson.Json `json:"traceIds"      orm:"trace_ids"      description:"链路ID集合"`
	Status        int         `json:"status"        orm:"status"         description:"状态"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"修改时间"`
}
