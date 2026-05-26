// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hook_test

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model"
	"testing"
)

func init() {
	g.DB().SetDryRun(true)
	g.DB().SetDebug(true)
}

// TestSaveTenant_User 以用户身份增改购买订单数据
func TestSaveTenant_User(t *testing.T) {

	// 设置上下文用户身份为用户
	ctx := context.WithValue(gctx.New(), consts.ContextHTTPKey, &model.Context{
		// 为了测试只设置了hook中需要用到的数据
		User: &model.Identity{
			Id:       12,
			DeptType: consts.DeptTypeUser,
		},
	})

	cols := dao.AddonHgexampleTenantOrder.Columns()
	orderSn := grand.Letters(32)

	// 以用户身份插入购买订单数据，自动维护租户关系字段
	data := g.Map{
		//cols.TenantId:    8,
		//cols.MerchantId:  11,
		//cols.UserId:      12,
		cols.ProductName: "无线运动耳机",
		cols.OrderSn:     orderSn,
		cols.Money:       99,
		cols.Status:      consts.PayStatusWait,
	}

	_, err := dao.AddonHgexampleTenantOrder.Ctx(ctx).Data(data).Hook(hook.SaveTenant).OmitEmptyData().Insert()
	if err != nil {
		panic(err)
	}

	// 以用户身份插入修改订单数据，自动维护租户关系字段
	update := g.Map{
		cols.Status: consts.PayStatusOk,
	}

	_, err = dao.AddonHgexampleTenantOrder.Ctx(ctx).Where(cols.OrderSn, orderSn).Data(update).Hook(hook.SaveTenant).Update()
	if err != nil {
		panic(err)
	}
}
