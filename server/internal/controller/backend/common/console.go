// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"hotgo/api/backend/common"
)

var Console = cConsole{}

type cConsole struct{}

// Stat 综合数据统计
func (c *cConsole) Stat(ctx context.Context, req *common.ConsoleStatReq) (res *common.ConsoleStatRes, err error) {
	res = new(common.ConsoleStatRes)

	res.Visits.DayVisits = 12010
	res.Visits.Rise = 13501
	res.Visits.Decline = 10502
	res.Visits.Amount = 10403

	res.Saleroom.WeekSaleroom = 20501
	res.Saleroom.Amount = 21002
	res.Saleroom.Degree = 83.66

	res.OrderLarge.WeekLarge = 39901
	res.OrderLarge.Rise = 31012
	res.OrderLarge.Decline = 30603
	res.OrderLarge.Amount = 36084

	res.Volume.WeekLarge = 40021
	res.Volume.Rise = 40202
	res.Volume.Decline = 45003
	res.Volume.Amount = 49004
	return
}
