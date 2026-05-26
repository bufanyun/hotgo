// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import (
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

func init() {
	dict.RegisterEnums("orderStatus", "订单状态", OrderStatusOptions)
	dict.RegisterEnums("acceptRefundStatus", "订单退款受理状态", OrderAcceptRefundOptions)
}

// 订单分组
// 为不同的业务订单设置不同的分组，分组可以设置不同的业务回调方法

const (
	OrderGroupDefault    = "order"       // 普通订单
	OrderGroupAdminOrder = "admin_order" // 后台充值订单
	// 还可以设置其他，方便后期扩展..
)

// 订单类型

const (
	OrderTypeBalance = "balance" // 余额充值
	OrderTypeProduct = "product" // 购买产品
)

const (
	OrderStatusALL           = -1 // 全部状态
	OrderStatusNotPay        = 1  // 待付款
	OrderStatusPay           = 2  // 已付款
	OrderStatusShipments     = 3  // 已发货
	OrderStatusDone          = 4  // 已完成
	OrderStatusClose         = 5  // 已关闭
	OrderStatusReturnRequest = 6  // 申请退款
	OrderStatusReturning     = 7  // 退款中
	OrderStatusReturned      = 8  // 已退款
	OrderStatusReturnReject  = 9  // 拒绝退款
)

// OrderStatusOptions 订单状态选项
var OrderStatusOptions = []*model.Option{
	dict.GenInfoOption(OrderStatusALL, "全部"),
	dict.GenInfoOption(OrderStatusNotPay, "待付款"),
	dict.GenInfoOption(OrderStatusPay, "已付款"),
	dict.GenInfoOption(OrderStatusShipments, "已发货"),
	dict.GenSuccessOption(OrderStatusDone, "已完成"),
	dict.GenDefaultOption(OrderStatusClose, "已关闭"),
	dict.GenWarningOption(OrderStatusReturnRequest, "申请退款"),
	dict.GenDefaultOption(OrderStatusReturning, "退款中"),
	dict.GenErrorOption(OrderStatusReturned, "已退款"),
	dict.GenWarningOption(OrderStatusReturnReject, "拒绝退款"),
}

// OrderAcceptRefundOptions 订单退款受理状态
var OrderAcceptRefundOptions = []*model.Option{
	dict.GenWarningOption(OrderStatusReturnRequest, "申请退款"),
	dict.GenSuccessOption(OrderStatusReturned, "已退款"),
	dict.GenErrorOption(OrderStatusReturnReject, "拒绝退款"),
}
