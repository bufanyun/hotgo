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
	"hotgo/internal/service"
)

var (
	Refund = cRefund{}
)

type cRefund struct{}

// List 查看交易退款列表
func (c *cRefund) List(ctx context.Context, req *pay.RefundListReq) (res *pay.RefundListRes, err error) {
	list, totalCount, err := service.PayRefund().List(ctx, &req.PayRefundListInp)
	if err != nil {
		return
	}

	res = new(pay.RefundListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出交易退款列表
func (c *cRefund) Export(ctx context.Context, req *pay.RefundExportReq) (res *pay.RefundExportRes, err error) {
	err = service.PayRefund().Export(ctx, &req.PayRefundListInp)
	return
}
