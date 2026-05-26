// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayLog is the golang structure of table hg_pay_log for DAO operations like Where/Data.
type PayLog struct {
	g.Meta        `orm:"table:hg_pay_log, do:true"`
	Id            any         // 主键
	MemberId      any         // 会员ID
	AppId         any         // 应用ID
	AddonsName    any         // 插件名称
	OrderSn       any         // 关联订单号
	OrderGroup    any         // 组别[默认统一支付类型]
	Openid        any         // openid
	MchId         any         // 商户支付账户
	Subject       any         // 订单标题
	Detail        *gjson.Json // 支付商品详情
	AuthCode      any         // 刷卡码
	OutTradeNo    any         // 商户订单号
	TransactionId any         // 交易号
	PayType       any         // 支付类型
	PayAmount     any         // 支付金额
	ActualAmount  any         // 实付金额
	PayStatus     any         // 支付状态
	PayAt         *gtime.Time // 支付时间
	TradeType     any         // 交易类型
	RefundSn      any         // 退款单号
	IsRefund      any         // 是否退款
	Custom        any         // 自定义参数
	CreateIp      any         // 创建者IP
	PayIp         any         // 支付者IP
	NotifyUrl     any         // 支付通知回调地址
	ReturnUrl     any         // 买家付款成功跳转地址
	TraceIds      *gjson.Json // 链路ID集合
	Status        any         // 状态
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 修改时间
}
