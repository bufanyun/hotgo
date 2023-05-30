// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import "github.com/gogf/gf/v2/frame/g"

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

var OrderStatusSlice = []int64{
	OrderStatusALL,
	OrderStatusNotPay, OrderStatusPay, OrderStatusShipments, OrderStatusDone, OrderStatusClose,
	OrderStatusReturnRequest, OrderStatusReturning, OrderStatusReturned, OrderStatusReturnReject,
}

// OrderStatusOptions 订单状态选项
var OrderStatusOptions = []g.Map{
	{
		"key":       OrderStatusALL,
		"value":     OrderStatusALL,
		"label":     "全部",
		"listClass": "info",
	},
	{
		"key":       OrderStatusNotPay,
		"value":     OrderStatusNotPay,
		"label":     "待付款",
		"listClass": "info",
	},
	{
		"key":       OrderStatusPay,
		"value":     OrderStatusPay,
		"label":     "已付款",
		"listClass": "info",
	},
	{
		"key":       OrderStatusShipments,
		"value":     OrderStatusShipments,
		"label":     "已发货",
		"listClass": "info",
	},
	{
		"key":       OrderStatusDone,
		"value":     OrderStatusDone,
		"label":     "已完成",
		"listClass": "success",
	},
	{
		"key":       OrderStatusClose,
		"value":     OrderStatusClose,
		"label":     "已关闭",
		"listClass": "default",
	},
	{
		"key":       OrderStatusReturnRequest,
		"value":     OrderStatusReturnRequest,
		"label":     "申请退款",
		"listClass": "warning",
	},
	{
		"key":       OrderStatusReturning,
		"value":     OrderStatusReturning,
		"label":     "退款中",
		"listClass": "default",
	},
	{
		"key":       OrderStatusReturned,
		"value":     OrderStatusReturned,
		"label":     "已退款",
		"listClass": "success",
	},
	{
		"key":       OrderStatusReturnReject,
		"value":     OrderStatusReturnReject,
		"label":     "拒绝退款",
		"listClass": "error",
	},
}

// OrderAcceptRefundOptions 订单退款受理状态
var OrderAcceptRefundOptions = []g.Map{
	{
		"key":       OrderStatusReturnRequest,
		"value":     OrderStatusReturnRequest,
		"label":     "申请退款",
		"listClass": "warning",
	},
	{
		"key":       OrderStatusReturned,
		"value":     OrderStatusReturned,
		"label":     "已退款",
		"listClass": "success",
	},
	{
		"key":       OrderStatusReturnReject,
		"value":     OrderStatusReturnReject,
		"label":     "拒绝退款",
		"listClass": "error",
	},
}
