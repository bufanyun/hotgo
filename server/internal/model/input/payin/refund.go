// Package payin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.5.3
// @AutoGenerate Date 2023-04-15 15:59:58
package payin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// PayRefundInp 订单退款
type PayRefundInp struct {
	OrderSn     string  `json:"orderSn"       dc:"业务订单号"`
	RefundMoney float64 `json:"refundMoney"   dc:"退款金额"`
	Reason      string  `json:"reason"        dc:"申请退款原因"`
	Remark      string  `json:"remark"        dc:"退款备注"`
}

func (in *PayRefundInp) Filter(ctx context.Context) (err error) {
	return
}

type PayRefundModel struct {
	entity.PayRefund
}

// PayRefundListInp 获取资产变动列表
type PayRefundListInp struct {
	form.PageReq
	Id          int64         `json:"id"          dc:"变动ID"`
	MemberId    int64         `json:"memberId"    dc:"管理员ID"`
	AppId       string        `json:"appId"       dc:"应用id"`
	CreditType  string        `json:"creditType"  dc:"变动类型"`
	CreditGroup string        `json:"creditGroup" dc:"变动的组别"`
	Remark      string        `json:"remark"      dc:"备注"`
	Ip          string        `json:"ip"          dc:"操作人IP"`
	Status      int           `json:"status"      dc:"状态"`
	CreatedAt   []*gtime.Time `json:"createdAt"   dc:"创建时间"`
}

func (in *PayRefundListInp) Filter(ctx context.Context) (err error) {
	return
}

type PayRefundListModel struct {
	entity.PayRefund
}

// PayRefundExportModel 导出资产变动
type PayRefundExportModel struct {
	entity.PayRefund
}
