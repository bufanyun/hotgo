// Package pay
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package pay

import (
	"context"
	"hotgo/api/api/pay"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
)

var (
	Notify = cNotify{}
)

type cNotify struct{}

// AliPay 支付宝回调
func (c *cNotify) AliPay(ctx context.Context, req *pay.NotifyAliPayReq) (res *pay.NotifyAliPayRes, err error) {
	_, err = service.Pay().Notify(ctx, payin.PayNotifyInp{PayType: consts.PayTypeAliPay})
	if err != nil {
		return nil, err
	}
	res = &pay.NotifyAliPayRes{PayType: consts.PayTypeAliPay, Message: "success"}
	return
}

// WxPay 微信支付回调
func (c *cNotify) WxPay(ctx context.Context, req *pay.NotifyWxPayReq) (res *pay.NotifyWxPayRes, err error) {
	_, err = service.Pay().Notify(ctx, payin.PayNotifyInp{PayType: consts.PayTypeWxPay})
	if err != nil {
		return nil, err
	}

	res = &pay.NotifyWxPayRes{PayType: consts.PayTypeWxPay, Code: "SUCCESS", Message: "收单成功"}

	return
}

// QQPay QQ支付回调
func (c *cNotify) QQPay(ctx context.Context, req *pay.NotifyQQPayReq) (res *pay.NotifyQQPayRes, err error) {
	_, err = service.Pay().Notify(ctx, payin.PayNotifyInp{PayType: consts.PayTypeQQPay})
	if err != nil {
		return nil, err
	}

	res = &pay.NotifyQQPayRes{PayType: consts.PayTypeQQPay, Message: "SUCCESS"}
	return
}
