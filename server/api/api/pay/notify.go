package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/payin"
)

// NotifyAliPayReq 支付宝回调
type NotifyAliPayReq struct {
	g.Meta `path:"/pay/notify/alipay" method:"post" tags:"支付异步通知" summary:"支付宝回调"`
}

type NotifyAliPayRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	*payin.PayNotifyModel
}

// NotifyWxPayReq 微信支付回调
type NotifyWxPayReq struct {
	g.Meta `path:"/pay/notify/wxpay" method:"post" tags:"支付异步通知" summary:"微信支付回调"`
}

type NotifyWxPayRes struct {
	*payin.PayNotifyModel
}

// NotifyQQPayReq QQ支付回调
type NotifyQQPayReq struct {
	g.Meta `path:"/pay/notify/qqpay" method:"post" tags:"支付异步通知" summary:"QQ支付回调"`
}

type NotifyQQPayRes struct {
	g.Meta `mime:"text/xml" type:"string"`
	*payin.PayNotifyModel
}
