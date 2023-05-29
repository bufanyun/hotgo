// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/cash"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
)

var (
	Cash = cCash{}
)

type cCash struct{}

// View 获取指定信息
func (c *cCash) View(ctx context.Context, req *cash.ViewReq) (res *cash.ViewRes, err error) {
	data, err := service.AdminCash().View(ctx, adminin.CashViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(cash.ViewRes)
	res.CashViewModel = data
	return
}

// List 查看列表
func (c *cCash) List(ctx context.Context, req *cash.ListReq) (res *cash.ListRes, err error) {
	var in adminin.CashListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.AdminCash().List(ctx, in)
	if err != nil {
		return
	}

	res = new(cash.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return
}

// Apply 申请提现
func (c *cCash) Apply(ctx context.Context, req *cash.ApplyReq) (res *cash.ApplyRes, err error) {
	err = service.AdminCash().Apply(ctx, adminin.CashApplyInp{
		Money:    req.Money,
		MemberId: contexts.GetUserId(ctx),
	})
	return
}

// Payment 提现打款处理
func (c *cCash) Payment(ctx context.Context, req *cash.PaymentReq) (res *cash.PaymentRes, err error) {
	err = service.AdminCash().Payment(ctx, adminin.CashPaymentInp{
		Id:     req.Id,
		Status: req.Status,
		Msg:    req.Msg,
	})
	return
}
