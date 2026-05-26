// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/payin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IPay interface {
		// Create 创建支付订单和日志
		Create(ctx context.Context, in payin.PayCreateInp) (res *payin.PayCreateModel, err error)
		// GenNotifyURL 生成支付通知地址
		GenNotifyURL(ctx context.Context, in payin.PayCreateInp) (notifyURL string, err error)
		// RegisterNotifyCall 注册支付成功回调方法
		RegisterNotifyCall()
		// Notify 异步通知
		Notify(ctx context.Context, in *payin.PayNotifyInp) (res *payin.PayNotifyModel, err error)
		// Model 支付日志ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取支付日志列表
		List(ctx context.Context, in payin.PayListInp) (list []*payin.PayListModel, totalCount int, err error)
		// Export 导出支付日志
		Export(ctx context.Context, in payin.PayListInp) (err error)
		// Edit 修改/新增支付日志
		Edit(ctx context.Context, in payin.PayEditInp) (err error)
		// Delete 删除支付日志
		Delete(ctx context.Context, in payin.PayDeleteInp) (err error)
		// View 获取支付日志指定信息
		View(ctx context.Context, in payin.PayViewInp) (res *payin.PayViewModel, err error)
		// Status 更新支付日志状态
		Status(ctx context.Context, in payin.PayStatusInp) (err error)
	}
	IPayRefund interface {
		// Model 交易退款ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Refund 订单退款
		Refund(ctx context.Context, in *payin.PayRefundInp) (res *payin.PayRefundModel, err error)
		// List 获取交易退款列表
		List(ctx context.Context, in *payin.PayRefundListInp) (list []*payin.PayRefundListModel, totalCount int, err error)
		// Export 导出交易退款
		Export(ctx context.Context, in *payin.PayRefundListInp) (err error)
	}
)

var (
	localPay       IPay
	localPayRefund IPayRefund
)

func Pay() IPay {
	if localPay == nil {
		panic("implement not found for interface IPay, forgot register?")
	}
	return localPay
}

func RegisterPay(i IPay) {
	localPay = i
}

func PayRefund() IPayRefund {
	if localPayRefund == nil {
		panic("implement not found for interface IPayRefund, forgot register?")
	}
	return localPayRefund
}

func RegisterPayRefund(i IPayRefund) {
	localPayRefund = i
}
