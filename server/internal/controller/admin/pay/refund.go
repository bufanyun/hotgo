// Package pay
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.5.3
// @AutoGenerate Date 2023-04-15 15:59:58
package pay

import (
	"context"
	"hotgo/api/admin/pay"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Refund = cRefund{}
)

type cRefund struct{}

// List 查看交易退款列表
func (c *cRefund) List(ctx context.Context, req *pay.RefundListReq) (res *pay.RefundListRes, err error) {
	var in payin.PayRefundListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	list, totalCount, err := service.PayRefund().List(ctx, in)
	if err != nil {
		return
	}

	res = new(pay.RefundListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Export 导出交易退款列表
func (c *cRefund) Export(ctx context.Context, req *pay.RefundExportReq) (res *pay.RefundExportRes, err error) {
	var in payin.PayRefundListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.PayRefund().Export(ctx, in)
	return
}
