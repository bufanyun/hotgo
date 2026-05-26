// Package payin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package payin

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// PayNotifyInp 异步通知
type PayNotifyInp struct {
	PayType string `json:"payType"       description:"支付类型"`
}

type PayNotifyModel struct {
	PayType string `json:"payType"       description:"支付类型"`
	Code    string `json:"code"       description:"状态码"`
	Message string `json:"message"       description:"响应文本消息"`
}

type PayNotifyUpdate struct {
	TransactionId string      `json:"transactionId" description:"交易号"`
	PayStatus     int         `json:"payStatus"     description:"支付状态"`
	PayAt         *gtime.Time `json:"payAt"         description:"支付时间"`
	PayIp         string      `json:"payIp"         description:"支付者ip"`
	TraceIds      *gjson.Json `json:"traceIds"      description:"链路id集合"`
}

// PayCreateInp 创建支付订单和日志
type PayCreateInp struct {
	Subject    string      `json:"subject"    description:"订单标题"`
	Detail     *gjson.Json `json:"detail"        description:"支付商品详情"`
	OrderSn    string      `json:"orderSn"       description:"关联订单号"`
	OrderGroup string      `json:"orderGroup"    description:"组别[默认统一支付类型]"`
	Openid     string      `json:"openid"        description:"openid"`
	PayType    string      `json:"payType"       description:"支付类型"`
	TradeType  string      `json:"tradeType"     description:"交易类型"`
	PayAmount  float64     `json:"payAmount"     description:"支付金额"`
	ReturnUrl  string      `json:"returnUrl"     description:"买家付款成功跳转地址"`
}

type PayCreateModel struct {
	Order *CreateOrderModel
}

// PayEditInp 修改/新增支付日志
type PayEditInp struct {
	entity.PayLog
}

func (in *PayEditInp) Filter(ctx context.Context) (err error) {
	return
}

type PayEditModel struct{}

// PayDeleteInp 删除支付日志
type PayDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *PayDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PayDeleteModel struct{}

// PayViewInp 获取指定支付日志信息
type PayViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *PayViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PayViewModel struct {
	entity.PayLog
}

// PayListInp 获取支付日志列表
type PayListInp struct {
	form.PageReq
	Id        int64         `json:"id"               dc:"ID"`
	Status    int           `json:"status"           dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt"        dc:"创建时间"`
}

func (in *PayListInp) Filter(ctx context.Context) (err error) {
	return
}

type PayListModel struct {
	Id            int64       `json:"id"            description:"主键"`
	MemberId      uint64      `json:"memberId"      description:"用户id"`
	AppId         string      `json:"appId"         description:"应用id"`
	AddonsName    string      `json:"addonsName"    description:"插件名称"`
	OrderSn       string      `json:"orderSn"       description:"关联订单号"`
	OrderGroup    string      `json:"orderGroup"    description:"组别[默认统一支付类型]"`
	Openid        string      `json:"openid"        description:"openid"`
	MchId         string      `json:"mchId"         description:"商户支付账户"`
	Body          string      `json:"body"          description:"创建支付报文"`
	AuthCode      string      `json:"authCode"      description:"刷卡码"`
	OutTradeNo    string      `json:"outTradeNo"    description:"商户订单号"`
	TransactionId string      `json:"transactionId" description:"交易号"`
	PayType       int         `json:"payType"       description:"支付类型"`
	PayFee        float64     `json:"payFee"        description:"支付金额"`
	PayStatus     int         `json:"payStatus"     description:"支付状态"`
	PayAt         *gtime.Time `json:"payAt"         description:"支付时间"`
	TradeType     string      `json:"tradeType"     description:"交易类型"`
	RefundSn      string      `json:"refundSn"      description:"退款编号"`
	RefundFee     float64     `json:"refundFee"     description:"退款金额"`
	RefundRemark  string      `json:"refundRemark"  description:"退款备注"`
	IsRefund      int         `json:"isRefund"      description:"是否退款"`
	CreateIp      string      `json:"createIp"      description:"创建者ip"`
	PayIp         string      `json:"payIp"         description:"支付者ip"`
	NotifyUrl     string      `json:"notifyUrl"     description:"支付通知回调地址"`
	ReturnUrl     string      `json:"returnUrl"     description:"买家付款成功跳转地址"`
	TraceIds      *gjson.Json `json:"traceIds"      description:"链路id集合"`
	Status        int         `json:"status"        description:"状态"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"修改时间"`
}

// PayExportModel 导出支付日志
type PayExportModel struct {
	Id            int64       `json:"id"            description:"主键"`
	MemberId      uint64      `json:"memberId"      description:"用户id"`
	AppId         string      `json:"appId"         description:"应用id"`
	AddonsName    string      `json:"addonsName"    description:"插件名称"`
	OrderSn       string      `json:"orderSn"       description:"关联订单号"`
	OrderGroup    string      `json:"orderGroup"    description:"组别[默认统一支付类型]"`
	Openid        string      `json:"openid"        description:"openid"`
	MchId         string      `json:"mchId"         description:"商户支付账户"`
	Body          string      `json:"body"          description:"创建支付报文"`
	AuthCode      string      `json:"authCode"      description:"刷卡码"`
	OutTradeNo    string      `json:"outTradeNo"    description:"商户订单号"`
	TransactionId string      `json:"transactionId" description:"交易号"`
	PayType       int         `json:"payType"       description:"支付类型"`
	PayFee        float64     `json:"payFee"        description:"支付金额"`
	PayStatus     int         `json:"payStatus"     description:"支付状态"`
	PayAt         *gtime.Time `json:"payAt"         description:"支付时间"`
	TradeType     string      `json:"tradeType"     description:"交易类型"`
	RefundSn      string      `json:"refundSn"      description:"退款编号"`
	RefundFee     float64     `json:"refundFee"     description:"退款金额"`
	RefundRemark  string      `json:"refundRemark"  description:"退款备注"`
	IsRefund      int         `json:"isRefund"      description:"是否退款"`
	CreateIp      string      `json:"createIp"      description:"创建者ip"`
	PayIp         string      `json:"payIp"         description:"支付者ip"`
	NotifyUrl     string      `json:"notifyUrl"     description:"支付通知回调地址"`
	ReturnUrl     string      `json:"returnUrl"     description:"买家付款成功跳转地址"`
	TraceIds      *gjson.Json `json:"traceIds"      description:"链路id集合"`
	Status        int         `json:"status"        description:"状态"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"修改时间"`
}

// PayMaxSortInp 获取支付日志最大排序
type PayMaxSortInp struct{}

func (in *PayMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type PayMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// PayStatusInp 更新支付日志状态
type PayStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *PayStatusInp) Filter(ctx context.Context) (err error) {
	return
}

type PayStatusModel struct{}

// PaySwitchInp 更新支付日志开关状态
type PaySwitchInp struct {
	form.SwitchReq
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *PaySwitchInp) Filter(ctx context.Context) (err error) {
	return
}

type PaySwitchModel struct{}
