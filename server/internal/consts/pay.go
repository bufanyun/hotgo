// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import (
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

// 支付方式

func init() {
	dict.RegisterEnums("payType", "支付方式", PayTypeOptions)
	dict.RegisterEnums("payStatus", "支付状态", PayStatusOptions)
}

const (
	PayTypeALL    = ""       // 全部
	PayTypeWxPay  = "wxpay"  // 微信支付
	PayTypeAliPay = "alipay" // 支付宝
	PayTypeQQPay  = "qqpay"  // QQ支付
)

var (
	PayTypeSlice = []string{
		PayTypeWxPay, PayTypeAliPay, PayTypeQQPay,
	}

	PayTypeNameMap = map[string]string{
		PayTypeALL:    "全部",
		PayTypeWxPay:  "微信支付",
		PayTypeAliPay: "支付宝",
		PayTypeQQPay:  "QQ支付",
	}
)

// 交易方式

const (
	// 微信
	TradeTypeWxMP   = "mp"   // 公众号
	TradeTypeWxMini = "mini" // 小程序
	TradeTypeWxApp  = "app"  // APP
	TradeTypeWxScan = "scan" // 二维码扫码
	TradeTypeWxPos  = "pos"  // 二维码收款
	TradeTypeWxH5   = "h5"   // H5

	// 支付宝
	TradeTypeAliWeb  = "web"  // 电脑网页
	TradeTypeAliApp  = "app"  // APP
	TradeTypeAliScan = "scan" // 二维码扫码
	TradeTypeAliWap  = "wap"  // 手机网页
	TradeTypeAliPos  = "pos"  // 二维码收款

	// QQ
	TradeTypeQQWeb = "qqweb" // PC网页
	TradeTypeQQWap = "qqwap" // 移动端
)

var (
	TradeTypeWxSlice  = []string{TradeTypeWxMP, TradeTypeWxMini, TradeTypeWxApp, TradeTypeWxScan, TradeTypeWxPos, TradeTypeWxH5}
	TradeTypeAliSlice = []string{TradeTypeAliWeb, TradeTypeAliApp, TradeTypeAliScan, TradeTypeAliWap, TradeTypeAliPos}
	TradeTypeQQSlice  = []string{TradeTypeQQWeb, TradeTypeQQWap}
)

// 支付状态

const (
	PayStatusWait = 1 // 待支付
	PayStatusOk   = 2 // 已支付
)

// PayStatusOptions 支付状态选项
var PayStatusOptions = []*model.Option{
	dict.GenDefaultOption(PayStatusWait, "待支付"),
	dict.GenSuccessOption(PayStatusOk, "已支付"),
}

// 退款状态

const (
	RefundStatusNo     = 1 // 未退款
	RefundStatusApply  = 2 // 申请退款
	RefundStatusReject = 3 // 拒绝退款
	RefundStatusAgree  = 4 // 同意退款，已退款
)

// PayTypeOptions 支付方式选项
var PayTypeOptions = []*model.Option{
	dict.GenSuccessOption(PayTypeWxPay, "微信支付"),
	dict.GenInfoOption(PayTypeAliPay, "支付宝"),
	dict.GenDefaultOption(PayTypeQQPay, "QQ支付"),
}
