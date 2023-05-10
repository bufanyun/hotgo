package pay

// 异步通知

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/location"
	"hotgo/internal/library/payment"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/payin"
)

// Notify 异步通知
func (s *sPay) Notify(ctx context.Context, in payin.PayNotifyInp) (res *payin.PayNotifyModel, err error) {
	data, err := payment.New(in.PayType).Notify(ctx, payin.NotifyInp{})
	if err != nil {
		return
	}

	var models *entity.PayLog
	if err = s.Model(ctx).Where(dao.PayLog.Columns().OutTradeNo, data.OutTradeNo).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.Newf("商户订单号[%v]不存在支付记录，请检查", data.OutTradeNo)
		return
	}

	if models.PayStatus != consts.PayStatusWait {
		err = gerror.Newf("商户订单号[%v]已被处理，请勿重复操作", data.OutTradeNo)
		return
	}

	var traceIds []string
	if err = models.TraceIds.Scan(&traceIds); err != nil {
		return
	}
	traceIds = append(traceIds, gctx.CtxId(ctx))

	models.TransactionId = data.TransactionId
	models.PayStatus = consts.PayStatusOk
	models.PayAt = data.PayAt
	models.ActualAmount = data.ActualAmount
	models.PayIp = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	models.TraceIds = gjson.New(traceIds)

	result, err := s.Model(ctx).
		Fields(
			dao.PayLog.Columns().TransactionId,
			dao.PayLog.Columns().PayStatus,
			dao.PayLog.Columns().PayAt,
			dao.PayLog.Columns().PayIp,
			dao.PayLog.Columns().TraceIds,
			dao.PayLog.Columns().ActualAmount,
		).
		Where(dao.PayLog.Columns().Id, models.Id).
		Where(dao.PayLog.Columns().PayStatus, consts.PayStatusWait).
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
		g.Log().Warningf(ctx, "没有被更新的数据行")
		return
	}

	// 回调业务
	payment.NotifyCall(ctx, payin.NotifyCallFuncInp{Pay: models})
	return
}
