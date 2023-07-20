// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/payment"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/simple"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdminOrder struct{}

func NewAdminOrder() *sAdminOrder {
	return &sAdminOrder{}
}

func init() {
	service.RegisterAdminOrder(NewAdminOrder())
}

// Model 充值订单ORM模型
func (s *sAdminOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AdminOrder.Ctx(ctx), option...)
}

// AcceptRefund 受理申请退款
func (s *sAdminOrder) AcceptRefund(ctx context.Context, in *adminin.OrderAcceptRefundInp) (err error) {
	view, err := s.View(ctx, &adminin.OrderViewInp{Id: in.Id})
	if err != nil {
		return err
	}

	if view == nil {
		err = gerror.New("订单不存在")
		return
	}

	if view.Status != consts.OrderStatusReturnRequest {
		err = gerror.New("当前订单状态不是申请退款状态，无需受理")
		return
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 同意退款
		if in.Status == consts.OrderStatusReturned {
			// 更新余额
			_, err = service.AdminCreditsLog().SaveBalance(ctx, &adminin.CreditsLogSaveBalanceInp{
				MemberId:    view.MemberId,
				AppId:       contexts.GetModule(ctx),
				AddonsName:  contexts.GetAddonName(ctx),
				CreditGroup: consts.CreditGroupBalanceRefund,
				Num:         -view.Money,
				MapId:       view.Id,
				Remark:      fmt.Sprintf("余额退款:%v", in.Remark),
			})
			if err != nil {
				return err
			}

			_, err = service.PayRefund().Refund(ctx, &payin.PayRefundInp{
				OrderSn:     view.OrderSn,
				RefundMoney: view.Money,
				Reason:      view.RefundReason,
				Remark:      in.Remark,
			})
			if err != nil {
				return err
			}
		}

		update := g.Map{
			dao.AdminOrder.Columns().Status:             in.Status,
			dao.AdminOrder.Columns().RejectRefundReason: in.RejectRefundReason,
		}

		_, err = s.Model(ctx).Where(dao.AdminOrder.Columns().Id, in.Id).Data(update).Update()
		return
	})
	return
}

// ApplyRefund 申请退款
func (s *sAdminOrder) ApplyRefund(ctx context.Context, in *adminin.OrderApplyRefundInp) (err error) {
	view, err := s.View(ctx, &adminin.OrderViewInp{Id: in.Id})
	if err != nil {
		return err
	}

	if view == nil {
		err = gerror.New("订单不存在")
		return
	}

	if view.Status == consts.OrderStatusReturnRequest {
		err = gerror.New("当前订单退款正在申请处理，请勿重复提交！")
		return
	}

	if view.Status != consts.OrderStatusDone {
		err = gerror.New("当前订单状态不支持申请退款，如有疑问请联系管理员！")
		return
	}

	update := g.Map{
		dao.AdminOrder.Columns().Status:       consts.OrderStatusReturnRequest,
		dao.AdminOrder.Columns().RefundReason: in.RefundReason,
	}

	_, err = s.Model(ctx).Where(dao.AdminOrder.Columns().Id, in.Id).Data(update).Update()
	return
}

// PayNotify 支付成功通知
func (s *sAdminOrder) PayNotify(ctx context.Context, in *payin.NotifyCallFuncInp) (err error) {
	var models *entity.AdminOrder
	if err = s.Model(ctx).Where(dao.AdminOrder.Columns().OrderSn, in.Pay.OrderSn).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.New("订单不存在")
		return
	}

	if models.Status != consts.OrderStatusNotPay {
		err = gerror.New("订单已被处理，无需重复操作")
		return
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 更新订单状态
		_, err = s.Model(ctx).Where(dao.AdminOrder.Columns().Id, models.Id).Data(g.Map{
			dao.AdminOrder.Columns().Status: consts.OrderStatusDone,
		}).Update()
		if err != nil {
			return
		}

		// 更新余额
		_, err = service.AdminCreditsLog().SaveBalance(ctx, &adminin.CreditsLogSaveBalanceInp{
			MemberId:    models.MemberId,
			AppId:       in.Pay.AppId,
			AddonsName:  in.Pay.AddonsName,
			CreditGroup: consts.CreditGroupBalanceRecharge,
			Num:         models.Money,
			MapId:       models.Id,
			Remark:      in.Pay.Subject,
		})
		if err != nil {
			return err
		}
		return
	})

	if err != nil {
		return
	}

	// 推送通知
	response := &websocket.WResponse{
		Event: "admin/order/notify",
		Data:  in.Pay,
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		websocket.SendToUser(in.Pay.MemberId, response)
	})
	return
}

// Create 创建充值订单
func (s *sAdminOrder) Create(ctx context.Context, in *adminin.OrderCreateInp) (res *adminin.OrderCreateModel, err error) {
	var (
		subject = "支付订单"
		orderSn = payment.GenOrderSn()
	)

	switch in.OrderType {
	case consts.OrderTypeBalance:
		subject = fmt.Sprintf("余额充值:￥%v", in.Money)
	case consts.OrderTypeProduct:
		// 读取商品信息，读取商品最终支付价格
		// ...

		//in.Money = 999
		//subject = fmt.Sprintf("购买商品:%v", "测试商品名称")

	default:
		err = gerror.New("不支持的订单类型")
		return
	}

	res = new(adminin.OrderCreateModel)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = s.Model(ctx).Data(entity.AdminOrder{
			MemberId:  contexts.GetUserId(ctx),
			OrderType: in.OrderType,
			ProductId: in.ProductId,
			OrderSn:   orderSn,
			Money:     in.Money,
			Remark:    in.Remark,
			Status:    consts.OrderStatusNotPay,
		}).Insert()
		if err != nil {
			return
		}

		create, err := service.Pay().Create(ctx, payin.PayCreateInp{
			Subject:    subject,
			OrderSn:    orderSn,
			OrderGroup: consts.OrderGroupAdminOrder,
			PayType:    in.PayType,
			TradeType:  in.TradeType,
			PayAmount:  in.Money,
			ReturnUrl:  in.ReturnUrl,
		})
		if err != nil {
			return err
		}

		res.Order = create.Order
		return
	})
	return
}

// List 获取充值订单列表
func (s *sAdminOrder) List(ctx context.Context, in *adminin.OrderListInp) (list []*adminin.OrderListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询业务订单号
	if in.OrderSn != "" {
		mod = mod.WhereLike(dao.AdminOrder.Columns().OrderSn, in.OrderSn)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.AdminOrder.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AdminOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 查询商户订单号
	if in.PayLogOutTradeNo != "" {
		mod = mod.WhereLike(dao.PayLog.Columns().OutTradeNo, in.PayLogOutTradeNo)
	}

	if in.MemberId > 0 {
		mod = mod.Where(dao.AdminOrder.Columns().MemberId, in.MemberId)
	}

	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.AdminOrder.Table(), dao.AdminOrder.Columns().OrderSn, // 主表表名,关联条件
		dao.PayLog.Table(), "payLog", dao.PayLog.Columns().OrderSn, // 关联表表名,别名,关联条件
	)...)

	totalCount, err = mod.Clone().Count(1)
	if err != nil {
		return
	}

	if totalCount == 0 {
		return
	}

	//关联表select
	fields, err := hgorm.GenJoinSelect(ctx, adminin.OrderListModel{}, &dao.AdminOrder, []*hgorm.Join{
		{Dao: &dao.PayLog, Alias: "payLog"},
	})

	if err != nil {
		return
	}

	err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderDesc(dao.AdminOrder.Columns().Id).Scan(&list)
	return
}

// Export 导出充值订单
func (s *sAdminOrder) Export(ctx context.Context, in *adminin.OrderListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(adminin.OrderExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出充值订单-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []adminin.OrderExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增充值订单
func (s *sAdminOrder) Edit(ctx context.Context, in *adminin.OrderEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		_, err = s.Model(ctx).
			FieldsEx(
				dao.AdminOrder.Columns().Id,
				dao.AdminOrder.Columns().CreatedAt,
			).
			Where(dao.AdminOrder.Columns().Id, in.Id).Data(in).Update()
		return
	}

	// 新增
	_, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		FieldsEx(
			dao.AdminOrder.Columns().Id,
		).
		Data(in).Insert()
	return
}

// Delete 删除充值订单
func (s *sAdminOrder) Delete(ctx context.Context, in *adminin.OrderDeleteInp) (err error) {
	_, err = s.Model(ctx).
		Where(dao.AdminOrder.Columns().Id, in.Id).
		Where(dao.AdminOrder.Columns().Status, consts.OrderStatusClose).
		Delete()
	return
}

// View 获取充值订单指定信息
func (s *sAdminOrder) View(ctx context.Context, in *adminin.OrderViewInp) (res *adminin.OrderViewModel, err error) {
	err = s.Model(ctx).Where(dao.AdminOrder.Columns().Id, in.Id).Scan(&res)
	return
}

// Status 更新充值订单状态
func (s *sAdminOrder) Status(ctx context.Context, in *adminin.OrderStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	_, err = s.Model(ctx).Where(dao.AdminOrder.Columns().Id, in.Id).Data(g.Map{
		dao.AdminOrder.Columns().Status: in.Status,
	}).Update()
	return
}
