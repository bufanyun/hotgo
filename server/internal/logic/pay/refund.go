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
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/location"
	"hotgo/internal/library/payment"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

// 订单退款.

type sPayRefund struct{}

func NewPayRefund() *sPayRefund {
	return &sPayRefund{}
}

func init() {
	service.RegisterPayRefund(NewPayRefund())
}

// Model 交易退款ORM模型
func (s *sPayRefund) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PayRefund.Ctx(ctx), option...)
}

// Refund 订单退款
func (s *sPayRefund) Refund(ctx context.Context, in *payin.PayRefundInp) (res *payin.PayRefundModel, err error) {
	var models *entity.PayLog
	if err = service.Pay().Model(ctx).Where(dao.PayLog.Columns().OrderSn, in.OrderSn).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.Newf("业务订单号[%v]不存在支付记录，请检查", in.OrderSn)
		return
	}

	if models.PayStatus != consts.PayStatusOk {
		err = gerror.Newf("业务订单号[%v]未支付，无需退款", in.OrderSn)
		return
	}

	if models.IsRefund != consts.RefundStatusNo {
		err = gerror.Newf("业务订单号[%v]退款已被处理，请勿重复操作", in.OrderSn)
		return
	}

	var traceIds []string
	if err = models.TraceIds.Scan(&traceIds); err != nil {
		return
	}
	traceIds = append(traceIds, gctx.CtxId(ctx))

	refundSn := payment.GenRefundSn()

	// 创建第三方平台退款
	req := payin.RefundInp{
		Pay:         models,
		RefundMoney: in.RefundMoney,
		Reason:      in.Reason,
		Remark:      in.Remark,
		RefundSn:    refundSn,
	}

	if _, err = payment.New(models.PayType).Refund(ctx, req); err != nil {
		return
	}

	models.RefundSn = refundSn
	models.IsRefund = consts.RefundStatusAgree
	models.TraceIds = gjson.New(traceIds)

	result, err := s.Model(ctx).
		Fields(
			dao.PayLog.Columns().RefundSn,
			dao.PayLog.Columns().IsRefund,
			dao.PayLog.Columns().TraceIds,
		).
		Where(dao.PayLog.Columns().Id, models.Id).
		OmitEmpty().
		Data(models).Update()
	if err != nil {
		return
	}

	ret, err := result.RowsAffected()
	if err != nil {
		return
	}

	if ret == 0 {
		g.Log().Warningf(ctx, "Refund 没有被更新的数据行")
	}

	data := &entity.PayRefund{
		Id:            0,
		MemberId:      models.MemberId,
		AppId:         models.AppId,
		OrderSn:       models.OrderSn,
		RefundTradeNo: "",
		RefundMoney:   in.RefundMoney,
		RefundWay:     1,
		Ip:            location.GetClientIp(ghttp.RequestFromCtx(ctx)),
		Reason:        in.Reason,
		Remark:        in.Remark,
		Status:        consts.RefundStatusAgree,
	}

	// 创建退款记录
	if _, err = s.Model(ctx).Data(data).Insert(); err != nil {
		return
	}
	return
}

// List 获取交易退款列表
func (s *sPayRefund) List(ctx context.Context, in *payin.PayRefundListInp) (list []*payin.PayRefundListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询变动ID
	if in.Id > 0 {
		mod = mod.Where(dao.PayRefund.Columns().Id, in.Id)
	}

	// 查询管理员ID
	if in.MemberId > 0 {
		mod = mod.Where(dao.PayRefund.Columns().MemberId, in.MemberId)
	}

	// 查询应用id
	if in.AppId != "" {
		mod = mod.WhereLike(dao.PayRefund.Columns().AppId, in.AppId)
	}

	// 查询备注
	if in.Remark != "" {
		mod = mod.WhereLike(dao.PayRefund.Columns().Remark, in.Remark)
	}

	// 查询操作人IP
	if in.Ip != "" {
		mod = mod.WhereLike(dao.PayRefund.Columns().Ip, in.Ip)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.PayRefund.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PayRefund.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count(1)
	if err != nil {
		return
	}

	if totalCount == 0 {
		return
	}

	err = mod.Fields(payin.PayRefundListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.PayRefund.Columns().Id).Scan(&list)
	return
}

// Export 导出交易退款
func (s *sPayRefund) Export(ctx context.Context, in *payin.PayRefundListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(payin.PayRefundExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出交易退款-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []payin.PayRefundExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}
