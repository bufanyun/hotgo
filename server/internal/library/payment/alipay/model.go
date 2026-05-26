// Package alipay
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package alipay

import "github.com/gogf/gf/v2/os/gtime"

// NotifyRequest 支付宝异步通知参数
// 文档：https://opendocs.alipay.com/open/203/105286
type NotifyRequest struct {
	NotifyTime        string              `json:"notify_time,omitempty"`
	NotifyType        string              `json:"notify_type,omitempty"`
	NotifyId          string              `json:"notify_id,omitempty"`
	AppId             string              `json:"app_id,omitempty"`
	Charset           string              `json:"charset,omitempty"`
	Version           string              `json:"version,omitempty"`
	SignType          string              `json:"sign_type,omitempty"`
	Sign              string              `json:"sign,omitempty"`
	AuthAppId         string              `json:"auth_app_id,omitempty"`
	TradeNo           string              `json:"trade_no,omitempty"`
	OutTradeNo        string              `json:"out_trade_no,omitempty"`
	OutBizNo          string              `json:"out_biz_no,omitempty"`
	BuyerId           string              `json:"buyer_id,omitempty"`
	BuyerLogonId      string              `json:"buyer_logon_id,omitempty"`
	SellerId          string              `json:"seller_id,omitempty"`
	SellerEmail       string              `json:"seller_email,omitempty"`
	TradeStatus       string              `json:"trade_status,omitempty"`
	TotalAmount       string              `json:"total_amount,omitempty"`
	ReceiptAmount     string              `json:"receipt_amount,omitempty"`
	InvoiceAmount     string              `json:"invoice_amount,omitempty"`
	BuyerPayAmount    string              `json:"buyer_pay_amount,omitempty"`
	PointAmount       string              `json:"point_amount,omitempty"`
	RefundFee         string              `json:"refund_fee,omitempty"`
	Subject           string              `json:"subject,omitempty"`
	Body              string              `json:"body,omitempty"`
	GmtCreate         string              `json:"gmt_create,omitempty"`
	GmtPayment        *gtime.Time         `json:"gmt_payment,omitempty"`
	GmtRefund         string              `json:"gmt_refund,omitempty"`
	GmtClose          string              `json:"gmt_close,omitempty"`
	FundBillList      []*FundBillListInfo `json:"fund_bill_list,omitempty"`
	PassbackParams    string              `json:"passback_params,omitempty"`
	VoucherDetailList []*VoucherDetail    `json:"voucher_detail_list,omitempty"`
	Method            string              `json:"method,omitempty"`    // 电脑网站支付 支付宝请求 return_url 同步返回参数
	Timestamp         string              `json:"timestamp,omitempty"` // 电脑网站支付 支付宝请求 return_url 同步返回参数
}

type FundBillListInfo struct {
	Amount      string `json:"amount,omitempty"`
	FundChannel string `json:"fundChannel,omitempty"` // 异步通知里是 fundChannel
}

type VoucherDetail struct {
	Id                         string `json:"id,omitempty"`
	Name                       string `json:"name,omitempty"`
	Type                       string `json:"type,omitempty"`
	Amount                     string `json:"amount,omitempty"`
	MerchantContribute         string `json:"merchant_contribute,omitempty"`
	OtherContribute            string `json:"other_contribute,omitempty"`
	Memo                       string `json:"memo,omitempty"`
	TemplateId                 string `json:"template_id,omitempty"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty"`
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute,omitempty"`
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty"`
}
