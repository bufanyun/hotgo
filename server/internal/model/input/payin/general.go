// Package payin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package payin

import (
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/os/gtime"
	officialJs "github.com/silenceper/wechat/v2/officialaccount/js"
	"hotgo/internal/model/entity"
)

// 和功能库[payment]包的通用输入/输出

// CreateOrderInp 统一创建订单入口
type CreateOrderInp struct {
	Pay *entity.PayLog
}

type CreateOrderModel struct {
	TradeType  string `json:"tradeType"     description:"交易类型"`
	PayURL     string `json:"payURL"        description:"支付地址"`
	OutTradeNo string `json:"outTradeNo"    description:"商户订单号"`
	JsApi      *JSAPI `json:"jsApi"         description:"jsapi支付参数"`
}

type JSAPI struct {
	Config *officialJs.Config     `json:"config" description:"js初始化配置"`
	Params *wechat.JSAPIPayParams `json:"params" description:"支付参数"`
}

// NotifyInp 统一异步通知处理入口
type NotifyInp struct {
}

type NotifyModel struct {
	OutTradeNo    string      `json:"outTradeNo"    description:"商户订单号"`
	TransactionId string      `json:"transactionId" description:"交易号"`
	PayAt         *gtime.Time `json:"payAt"         description:"支付时间"`
	ActualAmount  float64     `json:"actualAmount"  description:"实付金额"`
}

// NotifyCallFuncInp 异步通知回调，用于异步通知验签通过后回调到具体的业务中
type NotifyCallFuncInp struct {
	Pay *entity.PayLog
}

// RefundInp 统一退款处理入口
type RefundInp struct {
	Pay         *entity.PayLog
	RefundMoney float64 `json:"refundMoney"   dc:"退款金额"`
	RefundSn    string  `json:"refundSn"      dc:"退款单号"`
	Reason      string  `json:"reason"        dc:"申请退款原因"`
	Remark      string  `json:"remark"        dc:"退款备注"`
}

type RefundModel struct {
}
