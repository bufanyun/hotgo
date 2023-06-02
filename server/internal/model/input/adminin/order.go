// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/payin"
	"hotgo/utility/validate"
)

// OrderAcceptRefundInp 受理申请退款
type OrderAcceptRefundInp struct {
	Id                 int64  `json:"id" v:"required#ID不能为空"  dc:"ID"`
	RejectRefundReason string `json:"rejectRefundReason"         dc:"拒绝退款原因"`
	Status             int64  `json:"status"                     dc:"状态"`
	Remark             string `json:"remark"                     dc:"退款备注"`
}

func (in *OrderAcceptRefundInp) Filter(ctx context.Context) (err error) {
	if !validate.InSlice(consts.OrderStatusSlice, in.Status) {
		err = gerror.Newf("订单状态不正确")
		return
	}

	if in.Status == consts.OrderStatusReturnReject && in.Remark == "" {
		in.Remark = "退款申请被拒绝"
	}
	return
}

type OrderAcceptRefundModel struct {
}

// OrderApplyRefundInp 申请退款
type OrderApplyRefundInp struct {
	Id           int64  `json:"id" v:"required#ID不能为空"                     dc:"ID"`
	RefundReason string `json:"refundReason"  v:"required#退款原因不能为空"      dc:"退款原因"`
}

func (in *OrderApplyRefundInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderApplyRefundModel struct {
}

// OrderCreateInp 创建充值订单
type OrderCreateInp struct {
	PayType   string  `json:"payType"        dc:"支付方式"`
	TradeType string  `json:"tradeType"      dc:"交易类型"`
	OrderType string  `json:"orderType"      dc:"订单类型"`
	ProductId int64   `json:"productId"      dc:"产品id"`
	Money     float64 `json:"money"          dc:"充值金额"`
	Remark    string  `json:"remark"         dc:"备注"`
	ReturnUrl string  `json:"returnUrl"      dc:"买家付款成功跳转地址"`
}

func (in *OrderCreateInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderCreateModel struct {
	Order *payin.CreateOrderModel `json:"order"`
}

// OrderEditInp 修改/新增充值订单
type OrderEditInp struct {
	entity.AdminOrder
}

func (in *OrderEditInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderEditModel struct{}

// OrderDeleteInp 删除充值订单
type OrderDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *OrderDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderDeleteModel struct{}

// OrderViewInp 获取指定充值订单信息
type OrderViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *OrderViewInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderViewModel struct {
	entity.AdminOrder
}

// OrderListInp 获取充值订单列表
type OrderListInp struct {
	form.PageReq
	MemberId         uint64        `json:"memberId"          dc:"用户id"`
	OrderSn          string        `json:"orderSn"           dc:"业务订单号"`
	Status           int           `json:"status"            dc:"状态"`
	CreatedAt        []*gtime.Time `json:"createdAt"         dc:"创建时间"`
	PayLogOutTradeNo string        `json:"payLogOutTradeNo"  dc:"商户订单号"`
}

func (in *OrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderListModel struct {
	entity.AdminOrder
	OutTradeNo string `json:"payLogOutTradeNo"  dc:"商户订单号"`
	PayType    string `json:"payLogPayType"  dc:"支付类型"`
}

// OrderExportModel 导出充值订单
type OrderExportModel struct {
	Id        int64       `json:"id"        dc:"主键"`
	MemberId  uint64      `json:"memberId"  dc:"用户id"`
	OrderType string      `json:"orderType" dc:"订单类型"`
	ProductId int64       `json:"productId" dc:"产品id"`
	OrderSn   string      `json:"orderSn"   dc:"关联订单号"`
	Money     float64     `json:"money"     dc:"充值金额"`
	Remark    string      `json:"remark"    dc:"备注"`
	Status    int         `json:"status"    dc:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// OrderStatusInp 更新充值订单状态
type OrderStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status"                   dc:"状态"`
}

func (in *OrderStatusInp) Filter(ctx context.Context) (err error) {
	return
}

type OrderStatusModel struct{}
