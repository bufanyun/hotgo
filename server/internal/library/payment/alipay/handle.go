package alipay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"hotgo/internal/model/input/payin"
)

func New(config *model.PayConfig) *aliPay {
	return &aliPay{
		config: config,
	}
}

type aliPay struct {
	config *model.PayConfig
}

// Refund 订单退款
func (h *aliPay) Refund(ctx context.Context, in payin.RefundInp) (res *payin.RefundModel, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return
	}

	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", in.Pay.OutTradeNo).
		Set("refund_amount", in.RefundMoney).
		Set("out_request_no", in.RefundSn).
		Set("refund_reason", in.Remark)

	refund, err := client.TradeRefund(ctx, bm)
	if err != nil {
		return
	}

	if refund.Response.FundChange != "Y" {
		err = gerror.New("支付宝本次退款未发生资金变化！")
		return
	}
	return
}

// Notify 异步通知
func (h *aliPay) Notify(ctx context.Context, in payin.NotifyInp) (res *payin.NotifyModel, err error) {
	notifyReq, err := alipay.ParseNotifyToBodyMap(ghttp.RequestFromCtx(ctx).Request)
	if err != nil {
		return
	}

	// 支付宝异步通知验签（公钥证书模式）
	ok, err := alipay.VerifySignWithCert(h.config.AliPayCertPublicKeyRSA2, notifyReq)
	if err != nil {
		return
	}

	if !ok {
		err = gerror.New("支付宝验签不通过！")
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

	if notify.TradeStatus != "TRADE_SUCCESS" {
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
	res.TransactionId = notify.TradeNo
	res.OutTradeNo = notify.OutTradeNo
	res.PayAt = notify.GmtPayment
	res.ActualAmount = gconv.Float64(notify.ReceiptAmount)

	return
}

// CreateOrder 创建订单
func (h *aliPay) CreateOrder(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return nil, err
	}

	// 设置回调地址
	client.SetReturnUrl(in.Pay.ReturnUrl).SetNotifyUrl(in.Pay.NotifyUrl)

	switch in.Pay.TradeType {
	case consts.TradeTypeAliScan, consts.TradeTypeAliWeb:
		return h.scan(ctx, in)
	case consts.TradeTypeAliWap:
		return h.wap(ctx, in)
	default:
		err = gerror.Newf("暂未支持的交易方式：%v", in.Pay.TradeType)
	}

	return
}

func GetClient(config *model.PayConfig) (client *alipay.Client, err error) {
	client, err = alipay.NewClient(config.AliPayAppId, gfile.GetContents(config.AliPayPrivateKey), true)
	if err != nil {
		err = gerror.Newf("创建支付宝客户端失败：%+v", err.Error())
		return
	}

	// 打开Debug开关，输出日志，默认关闭
	if config.Debug {
		client.DebugSwitch = gopay.DebugOn
	}

	client.SetLocation(alipay.LocationShanghai) // 设置时区，不设置或出错均为默认服务器时间

	// 证书路径
	err = client.SetCertSnByPath(config.AliPayAppCertPublicKey, config.AliPayRootCert, config.AliPayCertPublicKeyRSA2)
	return
}

// scan 扫码支付
func (h *aliPay) scan(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return nil, err
	}

	// 设置回调地址
	client.SetReturnUrl(in.Pay.ReturnUrl).SetNotifyUrl(in.Pay.NotifyUrl)

	bm := make(gopay.BodyMap)
	bm.Set("subject", in.Pay.Subject).
		Set("out_trade_no", in.Pay.OutTradeNo).
		Set("total_amount", in.Pay.PayAmount)

	payUrl, err := client.TradePagePay(ctx, bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			return nil, bizErr
		}
		return nil, err
	}

	res = new(payin.CreateOrderModel)
	res.TradeType = in.Pay.TradeType
	res.PayURL = payUrl
	res.OutTradeNo = in.Pay.OutTradeNo
	return
}

func (h *aliPay) wap(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return nil, err
	}

	// 设置回调地址
	client.SetReturnUrl(in.Pay.ReturnUrl).SetNotifyUrl(in.Pay.NotifyUrl)

	bm := make(gopay.BodyMap)
	bm.Set("subject", in.Pay.Subject).
		Set("out_trade_no", in.Pay.OutTradeNo).
		Set("total_amount", in.Pay.PayAmount).
		Set("product_code", "QUICK_WAP_WAY")

	// 手机网站支付请求
	payUrl, err := client.TradeWapPay(ctx, bm)
	if err != nil {
		return
	}

	res = new(payin.CreateOrderModel)
	res.TradeType = in.Pay.TradeType
	res.PayURL = payUrl
	res.OutTradeNo = in.Pay.OutTradeNo
	return
}
