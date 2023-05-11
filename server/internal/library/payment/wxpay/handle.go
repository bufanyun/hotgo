package wxpay

import (
	"context"
	"crypto/rsa"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	weOpen "hotgo/internal/library/wechat"
	"hotgo/internal/model"
	"hotgo/internal/model/input/payin"
	"time"
)

func New(config *model.PayConfig) *wxPay {
	return &wxPay{
		config: config,
	}
}

type wxPay struct {
	config *model.PayConfig
}

// Refund 订单退款
func (h *wxPay) Refund(ctx context.Context, in payin.RefundInp) (res *payin.RefundModel, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", in.Pay.OutTradeNo).
		Set("out_refund_no", in.RefundSn).
		Set("reason", in.Remark).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", int64(in.Pay.PayAmount*100)).
				Set("currency", "CNY").
				Set("refund", int64(in.RefundMoney*100))
		})

	refund, err := client.V3Refund(ctx, bm)
	if err != nil {
		return
	}

	if refund.Error != "" {
		err = gerror.Newf("微信支付发起退款失败,原因：%v", refund.Response.Status)
		return
	}

	if refund.Response.Status != "SUCCESS" && refund.Response.Status != "PROCESSING" {
		err = gerror.Newf("微信支付发起退款失败,状态码：%v", refund.Response.Status)
		return
	}
	return
}

// Notify 异步通知
func (h *wxPay) Notify(ctx context.Context, in payin.NotifyInp) (res *payin.NotifyModel, err error) {
	notifyReq, err := wechat.V3ParseNotify(ghttp.RequestFromCtx(ctx).Request)
	if err != nil {
		return
	}

	client, err := GetClient(h.config)
	if err != nil {
		return
	}

	// 获取微信平台证书
	certMap, err := getPublicKeyMap(client)
	if err != nil {
		return
	}

	// 验证异步通知的签名
	if err = notifyReq.VerifySignByPKMap(certMap); err != nil {
		return
	}

	notify, err := notifyReq.DecryptCipherText(h.config.WxPayAPIv3Key)
	if err != nil {
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
	res.PayAt = gtime.New(notify.SuccessTime)
	res.ActualAmount = float64(notify.Amount.PayerTotal / 100) // 转为元，和系统内保持一至

	return
}

// CreateOrder 创建订单
func (h *wxPay) CreateOrder(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
	switch in.Pay.TradeType {
	case consts.TradeTypeWxScan:
		return h.scan(ctx, in)
	case consts.TradeTypeWxMP, consts.TradeTypeWxMini:
		return h.jsapi(ctx, in)
	case consts.TradeTypeWxH5:
		return h.h5(ctx, in)
	default:
		err = gerror.Newf("暂未支持的交易方式：%v", in.Pay.TradeType)
	}

	return
}

func GetClient(config *model.PayConfig) (client *wechat.ClientV3, err error) {
	client, err = wechat.NewClientV3(config.WxPayMchId, config.WxPaySerialNo, config.WxPayAPIv3Key, config.WxPayPrivateKey)
	if err != nil {
		return
	}

	if _, _, err = client.GetAndSelectNewestCertALL(); err != nil {
		return nil, err
	}

	serialNo, snCertMap, err := client.GetAndSelectNewestCert()
	if err != nil {
		return
	}
	snPkMap := make(map[string]*rsa.PublicKey)
	for sn, cert := range snCertMap {
		pubKey, err := xpem.DecodePublicKey([]byte(cert))
		if err != nil {
			return nil, err
		}
		snPkMap[sn] = pubKey
	}

	client.SnCertMap = snPkMap
	client.WxSerialNo = serialNo

	// 打开Debug开关，输出日志，默认关闭
	if config.Debug {
		client.DebugSwitch = gopay.DebugOn
	}
	return
}

func getPublicKeyMap(client *wechat.ClientV3) (wxPublicKeyMap map[string]*rsa.PublicKey, err error) {
	serialNo, snCertMap, err := client.GetAndSelectNewestCert()
	if err != nil {
		return
	}

	snPkMap := make(map[string]*rsa.PublicKey)
	for sn, cert := range snCertMap {
		pubKey, err := xpem.DecodePublicKey([]byte(cert))
		if err != nil {
			return nil, err
		}
		snPkMap[sn] = pubKey
	}
	client.SnCertMap = snPkMap
	client.WxSerialNo = serialNo

	wxPublicKeyMap = client.WxPublicKeyMap()
	return
}

// scan 创建扫码支付订单
func (h *wxPay) scan(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return
	}

	bm := make(gopay.BodyMap)
	bm.Set("appid", h.config.WxPayAppId).
		Set("mchid", h.config.WxPayMchId).
		Set("description", in.Pay.Subject).
		Set("out_trade_no", in.Pay.OutTradeNo).
		Set("time_expire", time.Now().Add(2*time.Hour).Format(time.RFC3339)).
		Set("notify_url", in.Pay.NotifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", int64(in.Pay.PayAmount*100)).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3TransactionNative(ctx, bm)
	if err != nil {
		return
	}

	if wxRsp.Code != 0 {
		err = gerror.New(wxRsp.Error)
		return
	}

	res = new(payin.CreateOrderModel)
	res.TradeType = in.Pay.TradeType
	res.PayURL = wxRsp.Response.CodeUrl
	res.OutTradeNo = in.Pay.OutTradeNo
	return
}

// h5 创建H5支付订单
func (h *wxPay) h5(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return
	}

	// 初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("appid", h.config.WxPayAppId).
		Set("mchid", h.config.WxPayMchId).
		Set("description", in.Pay.Subject).
		Set("out_trade_no", in.Pay.OutTradeNo).
		Set("time_expire", time.Now().Add(2*time.Hour).Format(time.RFC3339)).
		Set("notify_url", in.Pay.NotifyUrl).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", int64(in.Pay.PayAmount*100)).
				Set("currency", "CNY")
		}).
		SetBodyMap("scene_info", func(b gopay.BodyMap) {
			b.Set("payer_client_ip", in.Pay.CreateIp).
				SetBodyMap("h5_info", func(b gopay.BodyMap) {
					b.Set("type", "Wap")
				})
		})

	// 请求支付下单，成功后得到结果
	wxRsp, err := client.V3TransactionH5(ctx, bm)
	if err != nil {
		return
	}

	if wxRsp.Code != 0 {
		err = gerror.New(wxRsp.Error)
		return
	}

	res = new(payin.CreateOrderModel)
	res.TradeType = in.Pay.TradeType
	res.PayURL = wxRsp.Response.H5Url
	res.OutTradeNo = in.Pay.OutTradeNo
	return
}

// jsapi 创建jsapi支付订单
func (h *wxPay) jsapi(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
	jsApi := new(payin.JSAPI)
	jsApi.Config, err = weOpen.GetJsConfig(ctx, in.Pay.ReturnUrl)
	if err != nil {
		return
	}

	client, err := GetClient(h.config)
	if err != nil {
		return
	}

	bm := make(gopay.BodyMap)
	bm.Set("appid", h.config.WxPayAppId).
		Set("mchid", h.config.WxPayMchId).
		Set("description", in.Pay.Subject).
		Set("out_trade_no", in.Pay.OutTradeNo).
		Set("time_expire", time.Now().Add(2*time.Hour).Format(time.RFC3339)).
		Set("notify_url", in.Pay.NotifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", int64(in.Pay.PayAmount*100)).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", in.Pay.Openid)
		})

	wxRsp, err := client.V3TransactionJsapi(ctx, bm)
	if err != nil {
		return
	}

	if wxRsp.Code != 0 {
		err = gerror.New(wxRsp.Error)
		return
	}

	js, err := client.PaySignOfJSAPI(h.config.WxPayAppId, wxRsp.Response.PrepayId)
	if err != nil {
		return
	}
	jsApi.Params = js

	res = new(payin.CreateOrderModel)
	res.TradeType = in.Pay.TradeType
	res.OutTradeNo = in.Pay.OutTradeNo
	res.JsApi = jsApi
	return
}
