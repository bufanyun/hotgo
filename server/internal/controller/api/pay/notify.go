// Package pay
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package pay

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/api/pay"
	"hotgo/internal/consts"
	"hotgo/internal/library/response"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
)

var (
	Notify = cNotify{}
)

type cNotify struct{}

// AliPay 支付宝回调
func (c *cNotify) AliPay(ctx context.Context, _ *pay.NotifyAliPayReq) (res *pay.NotifyAliPayRes, err error) {
	if _, err = service.Pay().Notify(ctx, payin.PayNotifyInp{PayType: consts.PayTypeAliPay}); err != nil {
		return nil, err
	}

	response.RText(g.RequestFromCtx(ctx), "success")
	return
}

// WxPay 微信支付回调
func (c *cNotify) WxPay(ctx context.Context, _ *pay.NotifyWxPayReq) (res *pay.NotifyWxPayRes, err error) {
	if _, err = service.Pay().Notify(ctx, payin.PayNotifyInp{PayType: consts.PayTypeWxPay}); err != nil {
		return
	}

	response.CustomJson(g.RequestFromCtx(ctx), `{"code": "SUCCESS","message": "收单成功"}`)
	return
}

// QQPay QQ支付回调
func (c *cNotify) QQPay(ctx context.Context, _ *pay.NotifyQQPayReq) (res *pay.NotifyQQPayRes, err error) {
	if _, err = service.Pay().Notify(ctx, payin.PayNotifyInp{PayType: consts.PayTypeQQPay}); err != nil {
		return
	}

	r := g.RequestFromCtx(ctx)
	r.Response.ClearBuffer()
	r.Response.Write(`<?xml version="1.0" encoding="UTF-8"?>`)
	r.Response.WriteXml(g.Map{
		"return_code": "SUCCESS",
	})
	return
}
