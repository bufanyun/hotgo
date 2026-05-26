// Package payment
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package payment

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/payment/alipay"
	"hotgo/internal/library/payment/qqpay"
	"hotgo/internal/library/payment/wxpay"
	"hotgo/internal/model/input/payin"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

// PayClient 支付客户端
type PayClient interface {
	// CreateOrder 创建订单
	CreateOrder(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error)
	// Notify 异步通知
	Notify(ctx context.Context, in payin.NotifyInp) (res *payin.NotifyModel, err error)
	// Refund 订单退款
	Refund(ctx context.Context, in payin.RefundInp) (res *payin.RefundModel, err error)
}

func New(name ...string) PayClient {
	var (
		payType = consts.PayTypeWxPay
		client  PayClient
	)

	if len(name) > 0 && name[0] != "" {
		payType = name[0]
	}

	switch payType {
	case consts.PayTypeAliPay:
		client = alipay.New(config)
	case consts.PayTypeWxPay:
		client = wxpay.New(config)
	case consts.PayTypeQQPay:
		client = qqpay.New(config)
	default:
		panic(fmt.Sprintf("暂不支持的支付方式:%v", payType))
	}
	return client
}

// GenOrderSn 生成业务订单号
func GenOrderSn(ctx context.Context) string {
	orderSn := fmt.Sprintf("HG@%v%v", gtime.Now().Format("YmdHis"), grand.S(4))
	count, err := dao.PayLog.Ctx(ctx).Where(dao.PayLog.Columns().OrderSn, orderSn).Count()
	if err != nil {
		panic(fmt.Sprintf("payment.GenOrderSn err:%+v", err))
	}
	if count > 0 {
		return GenOrderSn(ctx)
	}
	return orderSn
}

// GenOutTradeNo 生成商户订单号
func GenOutTradeNo(ctx context.Context) string {
	outTradeNo := fmt.Sprintf("%v%v", gtime.Now().Format("YmdHis"), grand.N(10000000, 99999999))
	count, err := dao.PayLog.Ctx(ctx).Where(dao.PayLog.Columns().OutTradeNo, outTradeNo).Count()
	if err != nil {
		panic(fmt.Sprintf("payment.GenOutTradeNo err:%+v", err))
	}
	if count > 0 {
		return GenOutTradeNo(ctx)
	}
	return outTradeNo
}

// GenRefundSn 生成退款订单号
func GenRefundSn(ctx context.Context) string {
	outTradeNo := fmt.Sprintf("%v%v", gtime.Now().Format("YmdHis"), grand.N(10000, 99999))
	count, err := dao.PayRefund.Ctx(ctx).Where(dao.PayRefund.Columns().RefundTradeNo, outTradeNo).Count()
	if err != nil {
		panic(fmt.Sprintf("payment.GenRefundSn err:%+v", err))
	}
	if count > 0 {
		return GenRefundSn(ctx)
	}
	return outTradeNo
}

// AutoTradeType 根据userAgent自动识别交易方式，在实际支付场景中你可以手动调整识别规则
func AutoTradeType(payType, userAgent string) (tradeType string) {
	isMobile := validate.IsMobileVisit(userAgent)
	switch payType {
	case consts.PayTypeAliPay:
		if isMobile {
			return consts.TradeTypeAliWap
		}
		return consts.TradeTypeAliWeb
	case consts.PayTypeWxPay:
		if isMobile {
			if validate.IsWxBrowserVisit(userAgent) {
				return consts.TradeTypeWxMP
			}

			if validate.IsWxMiniProgramVisit(userAgent) {
				return consts.TradeTypeWxMini
			}

			return consts.TradeTypeWxH5
		}
		return consts.TradeTypeWxScan
	case consts.PayTypeQQPay:
		if isMobile {
			return consts.TradeTypeQQWap
		}
		return consts.TradeTypeQQWeb
	default:
	}
	return
}
