// Package order
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package order

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// AcceptRefundReq 受理申请退款
type AcceptRefundReq struct {
	g.Meta `path:"/order/acceptRefund" method:"post" tags:"充值订单" summary:"受理申请退款"`
	adminin.OrderAcceptRefundInp
}

type AcceptRefundRes struct {
}

// ApplyRefundReq 申请退款
type ApplyRefundReq struct {
	g.Meta `path:"/order/applyRefund" method:"post" tags:"充值订单" summary:"申请退款"`
	adminin.OrderApplyRefundInp
}

type ApplyRefundRes struct {
}

// CreateReq 创建充值订单
type CreateReq struct {
	g.Meta `path:"/order/create" method:"post" tags:"充值订单" summary:"创建充值订单"`
	adminin.OrderCreateInp
}

type CreateRes struct {
	*adminin.OrderCreateModel
}

// ListReq 查询充值订单列表
type ListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"充值订单" summary:"获取充值订单列表"`
	adminin.OrderListInp
}

type ListRes struct {
	form.PageRes
	List []*adminin.OrderListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出充值订单列表
type ExportReq struct {
	g.Meta `path:"/order/export" method:"get" tags:"充值订单" summary:"导出充值订单列表"`
	adminin.OrderListInp
}

type ExportRes struct{}

// ViewReq 获取充值订单指定信息
type ViewReq struct {
	g.Meta `path:"/order/view" method:"get" tags:"充值订单" summary:"获取充值订单指定信息"`
	adminin.OrderViewInp
}

type ViewRes struct {
	*adminin.OrderViewModel
}

// EditReq 修改/新增充值订单
type EditReq struct {
	g.Meta `path:"/order/edit" method:"post" tags:"充值订单" summary:"修改/新增充值订单"`
	adminin.OrderEditInp
}

type EditRes struct{}

// DeleteReq 删除充值订单
type DeleteReq struct {
	g.Meta `path:"/order/delete" method:"post" tags:"充值订单" summary:"删除充值订单"`
	adminin.OrderDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新充值订单状态
type StatusReq struct {
	g.Meta `path:"/order/status" method:"post" tags:"充值订单" summary:"更新充值订单状态"`
	adminin.OrderStatusInp
}

type StatusRes struct{}
