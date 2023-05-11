package qqpay

// NotifyRequest QQ支付异步通知参数
// 文档：https://mp.qpay.tenpay.cn/buss/wiki/38/1204
type NotifyRequest struct {
	Appid         string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId         string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr      string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign          string `xml:"sign,omitempty" json:"sign,omitempty"`
	DeviceInfo    string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	TradeType     string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState    string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType      string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType       string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee      string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee       string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CouponFee     string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach        string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd       string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	Openid        string `xml:"openid,omitempty" json:"openid,omitempty"`
}
