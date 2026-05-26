// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.13.1
package sys

import (
	"context"
	"fmt"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysTenantOrder struct{}

func NewSysTenantOrder() *sSysTenantOrder {
	return &sSysTenantOrder{}
}

func init() {
	service.RegisterSysTenantOrder(NewSysTenantOrder())
}

// Model 多租户功能演示ORM模型
func (s *sSysTenantOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		// 过滤多租户数据权限
		option = append(option, &handler.Option{
			FilterTenant: true,
			//FilterAuth:   true, // 如果还需要维护created_by、member_id等部门数据权限范围可放开注释
		})
	}
	return handler.Model(dao.AddonHgexampleTenantOrder.Ctx(ctx), option...)
}

// List 获取多租户功能演示列表
func (s *sSysTenantOrder) List(ctx context.Context, in *sysin.TenantOrderListInp) (list []*sysin.TenantOrderListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.TenantOrderListModel{})

	// 查询主键
	if in.Id > 0 {
		mod = mod.Where(dao.AddonHgexampleTenantOrder.Columns().Id, in.Id)
	}

	// 查询租户ID
	if in.TenantId > 0 {
		mod = mod.Where(dao.AddonHgexampleTenantOrder.Columns().TenantId, in.TenantId)
	}

	// 查询商户ID
	if in.MerchantId > 0 {
		mod = mod.Where(dao.AddonHgexampleTenantOrder.Columns().MerchantId, in.MerchantId)
	}

	// 查询用户ID
	if in.UserId > 0 {
		mod = mod.Where(dao.AddonHgexampleTenantOrder.Columns().UserId, in.UserId)
	}

	// 查询关联订单号
	if in.OrderSn != "" {
		mod = mod.Where(dao.AddonHgexampleTenantOrder.Columns().OrderSn, in.OrderSn)
	}

	// 查询支付状态
	if in.Status > 0 {
		mod = mod.Where(dao.AddonHgexampleTenantOrder.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AddonHgexampleTenantOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.AddonHgexampleTenantOrder.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取多租户功能演示列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出多租户功能演示
func (s *sSysTenantOrder) Export(ctx context.Context, in *sysin.TenantOrderListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.TenantOrderExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出多租户功能演示-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.TenantOrderExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增多租户功能演示
func (s *sSysTenantOrder) Edit(ctx context.Context, in *sysin.TenantOrderEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.TenantOrderUpdateFields{}).
				WherePri(in.Id).Data(in).
				Hook(hook.SaveTenant). // 自动维护租户关系更新
				Update(); err != nil {
			}
			return
		}

		// 新增
		if _, err = dao.AddonHgexampleTenantOrder.Ctx(ctx).
			Fields(sysin.TenantOrderInsertFields{}).
			Hook(hook.SaveTenant). // 自动维护租户关系更新
			Data(in).
			Insert(); err != nil {
		}
		return
	})
}

// Delete 删除多租户功能演示
func (s *sSysTenantOrder) Delete(ctx context.Context, in *sysin.TenantOrderDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除多租户功能演示失败，请稍后重试！")
		return
	}
	return
}

// View 获取多租户功能演示指定信息
func (s *sSysTenantOrder) View(ctx context.Context, in *sysin.TenantOrderViewInp) (res *sysin.TenantOrderViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取多租户功能演示信息，请稍后重试！")
		return
	}
	return
}
