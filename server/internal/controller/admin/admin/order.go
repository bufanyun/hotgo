// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"hotgo/api/admin/order"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Order = cOrder{}
)

type cOrder struct{}

// AcceptRefund 受理申请退款
func (c *cOrder) AcceptRefund(ctx context.Context, req *order.AcceptRefundReq) (res *order.AcceptRefundRes, err error) {
	var in adminin.OrderAcceptRefundInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminOrder().AcceptRefund(ctx, in)
	return
}

// ApplyRefund 申请退款
func (c *cOrder) ApplyRefund(ctx context.Context, req *order.ApplyRefundReq) (res *order.ApplyRefundRes, err error) {
	var in adminin.OrderApplyRefundInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminOrder().ApplyRefund(ctx, in)
	return
}

// Option 获取订单状态选项
func (c *cOrder) Option(ctx context.Context, req *order.OptionReq) (res *order.OptionRes, err error) {
	res = &order.OptionRes{
		Status:             consts.OrderStatusOptions,
		AcceptRefundStatus: consts.OrderAcceptRefundOptions,
		PayType:            consts.PayTypeOptions,
	}
	return
}

// Create 创建充值订单
func (c *cOrder) Create(ctx context.Context, req *order.CreateReq) (res *order.CreateRes, err error) {
	var in adminin.OrderCreateInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	data, err := service.AdminOrder().Create(ctx, in)
	if err != nil {
		return
	}

	res = new(order.CreateRes)
	res.OrderCreateModel = data
	return
}

// List 查看充值订单列表
func (c *cOrder) List(ctx context.Context, req *order.ListReq) (res *order.ListRes, err error) {
	var in adminin.OrderListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	list, totalCount, err := service.AdminOrder().List(ctx, in)
	if err != nil {
		return
	}

	res = new(order.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Export 导出充值订单列表
func (c *cOrder) Export(ctx context.Context, req *order.ExportReq) (res *order.ExportRes, err error) {
	var in adminin.OrderListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminOrder().Export(ctx, in)
	return
}

// Edit 更新充值订单
func (c *cOrder) Edit(ctx context.Context, req *order.EditReq) (res *order.EditRes, err error) {
	var in adminin.OrderEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminOrder().Edit(ctx, in)
	return
}

// View 获取指定充值订单信息
func (c *cOrder) View(ctx context.Context, req *order.ViewReq) (res *order.ViewRes, err error) {
	var in adminin.OrderViewInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	data, err := service.AdminOrder().View(ctx, in)
	if err != nil {
		return
	}

	res = new(order.ViewRes)
	res.OrderViewModel = data
	return
}

// Delete 删除充值订单
func (c *cOrder) Delete(ctx context.Context, req *order.DeleteReq) (res *order.DeleteRes, err error) {
	var in adminin.OrderDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminOrder().Delete(ctx, in)
	return
}

// Status 更新充值订单状态
func (c *cOrder) Status(ctx context.Context, req *order.StatusReq) (res *order.StatusRes, err error) {
	var in adminin.OrderStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminOrder().Status(ctx, in)
	return
}
