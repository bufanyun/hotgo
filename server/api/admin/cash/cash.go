// Package cash
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cash

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/cash/list" method:"get" tags:"提现" summary:"获取提现列表"`
	adminin.CashListInp
}

type ListRes struct {
	List []*adminin.CashListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/cash/view" method:"get" tags:"提现" summary:"获取指定信息"`
	adminin.CashViewInp
}

type ViewRes struct {
	*adminin.CashViewModel
}

// ApplyReq 申请提现
type ApplyReq struct {
	g.Meta `path:"/cash/apply" method:"post" tags:"提现" summary:"申请提现"`
	Money  float64 `json:"money"     description:"提现金额"`
}

type ApplyRes struct{}

// PaymentReq 提现打款处理
type PaymentReq struct {
	g.Meta `path:"/cash/payment" method:"post" tags:"提现" summary:"提现打款处理"`
	adminin.CashPaymentInp
}

type PaymentRes struct{}
