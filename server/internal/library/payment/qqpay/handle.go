package qqpay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/qq"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"hotgo/internal/model/input/payin"
)

func New(config *model.PayConfig) *qqPay {
	return &qqPay{
		config: config,
	}
}

type qqPay struct {
	config *model.PayConfig
}

// Refund 订单退款
func (h *qqPay) Refund(ctx context.Context, in payin.RefundInp) (res *payin.RefundModel, err error) {
	err = gerror.New("暂不支持QQ支付申请退款，如有疑问请联系管理员")
	return
}

// Notify 异步通知
func (h *qqPay) Notify(ctx context.Context, in payin.NotifyInp) (res *payin.NotifyModel, err error) {
	notifyReq, err := qq.ParseNotifyToBodyMap(ghttp.RequestFromCtx(ctx).Request)
	if err != nil {
		return
	}

	// 验签操作
	ok, err := qq.VerifySign(h.config.QQPayApiKey, qq.SignType_MD5, notifyReq)
	if err != nil {
		return
	}

	if !ok {
		err = gerror.New("QQ支付验签不通过！")
		return
	}

	var notify *NotifyRequest
	if err = gconv.Scan(notifyReq, &notify); err != nil {
		return
	}

	if notify == nil {
		err = gerror.New("解析订单参数失败！")
		return
	}

	if notify.TradeState != "SUCCESS" {
		err = gerror.New("非交易支付成功状态，无需处理！")
		// 这里如果相对非交易支付成功状态进行处理，可自行调整此处逻辑
		// ...
		return
	}

	if notify.OutTradeNo == "" {
		err = gerror.New("订单中没有找到商户单号！")
		return
	}

	res = new(payin.NotifyModel)
	res.TransactionId = notify.TransactionId
	res.OutTradeNo = notify.OutTradeNo
	res.PayAt = gtime.New(notify.TimeEnd)
	res.ActualAmount = gconv.Float64(notify.CouponFee) / 100 // 用户本次交易中，实际支付的金额 转为元，和系统内保持一至

	return
}

// CreateOrder 创建订单
func (h *qqPay) CreateOrder(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
	client := GetClient(h.config)

	switch in.Pay.TradeType {
	case consts.TradeTypeQQWeb, consts.TradeTypeQQWap:
		bm := make(gopay.BodyMap)
		bm.
			Set("mch_id", h.config.QQPayMchId).
			Set("body", in.Pay.Subject).
			Set("out_trade_no", in.Pay.OutTradeNo).
			Set("notify_url", in.Pay.NotifyUrl).
			Set("nonce_str", grand.Letters(32)).
			Set("spbill_create_ip", in.Pay.CreateIp).
			Set("trade_type", "NATIVE"). // MICROPAY、APP、JSAPI、NATIVE
			Set("total_fee", int64(in.Pay.PayAmount*100))

		qqRsp, err := client.UnifiedOrder(ctx, bm)
		if err != nil {
			return nil, err
		}

		if qqRsp.ReturnCode != "SUCCESS" {
			err = gerror.New(qqRsp.ReturnMsg)
			return nil, err
		}

		if qqRsp.ResultCode != "SUCCESS" {
			err = gerror.New(qqRsp.ErrCodeDes)
			return nil, err
		}

		res = new(payin.CreateOrderModel)
		res.TradeType = in.Pay.TradeType
		res.PayURL = qqRsp.CodeUrl
		res.OutTradeNo = in.Pay.OutTradeNo

	default:
		err = gerror.Newf("暂未支持的交易方式：%v", in.Pay.TradeType)
	}

	return
}

func GetClient(config *model.PayConfig) (client *qq.Client) {
	client = qq.NewClient(config.QQPayMchId, config.QQPayApiKey)

	// 打开Debug开关，输出日志，默认关闭
	if config.Debug {
		client.DebugSwitch = gopay.DebugOn
	}
	return
}
