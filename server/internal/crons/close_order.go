// Package crons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package crons

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/cron"
	"hotgo/internal/service"
)

func init() {
	cron.Register(CloseOrder)
}

// CloseOrder 取消过期订单
var CloseOrder = &cCloseOrder{name: "close_order"}

type cCloseOrder struct {
	name string
}

func (c *cCloseOrder) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cCloseOrder) Execute(ctx context.Context) {
	_, err := service.AdminOrder().Model(ctx).
		Where(dao.AdminOrder.Columns().Status, consts.OrderStatusNotPay).
		WhereLTE(dao.AdminOrder.Columns().CreatedAt, gtime.Now().AddDate(0, 0, -1)).
		Data(g.Map{
			dao.AdminOrder.Columns().Status: consts.OrderStatusClose,
		}).Update()
	if err != nil {
		cron.Logger().Warning(ctx, "cron CloseOrder Execute err:%+v", err)
	}
}
