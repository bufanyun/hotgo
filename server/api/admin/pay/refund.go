// Package pay
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.5.3
// @AutoGenerate Date 2023-04-15 15:59:58
package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/payin"
)

// RefundListReq 查询交易退款列表
type RefundListReq struct {
	g.Meta `path:"/payRefund/list" method:"get" tags:"交易退款" summary:"获取交易退款列表"`
	payin.PayRefundListInp
}

type RefundListRes struct {
	form.PageRes
	List []*payin.PayRefundListModel `json:"list"   dc:"数据列表"`
}

// RefundExportReq 导出交易退款列表
type RefundExportReq struct {
	g.Meta `path:"/payRefund/export" method:"get" tags:"交易退款" summary:"导出交易退款列表"`
	payin.PayRefundListInp
}

type RefundExportRes struct{}
