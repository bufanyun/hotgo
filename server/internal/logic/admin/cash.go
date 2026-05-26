// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/library/location"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

type sAdminCash struct{}

func NewAdminCash() *sAdminCash {
	return &sAdminCash{}
}

func init() {
	service.RegisterAdminCash(NewAdminCash())
}

// View 获取指定提现信息
func (s *sAdminCash) View(ctx context.Context, in *adminin.CashViewInp) (res *adminin.CashViewModel, err error) {
	// 这里做了强制限制非超管不允许访问，如果你想通过菜单权限控制，请注释掉以下验证
	if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		err = gerror.New("没有访问权限")
		return
	}

	if err = dao.AdminCash.Ctx(ctx).Where(dao.AdminCash.Columns().Id, in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if res == nil {
		err = gerror.New("提现信息获取失败")
		return
	}

	var mem *entity.AdminMember
	err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, res.MemberId).Scan(&mem)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if !mem.Cash.IsNil() {
		if err = mem.Cash.Scan(&res.MemberCash); err != nil {
			return
		}
	}
	return
}

// List 获取列表
func (s *sAdminCash) List(ctx context.Context, in *adminin.CashListInp) (list []*adminin.CashListModel, totalCount int, err error) {
	var (
		mod        = dao.AdminCash.Ctx(ctx)
		opMemberId = contexts.GetUserId(ctx)
		isSuper    = service.AdminMember().VerifySuperId(ctx, opMemberId)
	)

	if in.MemberId > 0 {
		mod = mod.Where(dao.AdminCash.Columns().MemberId, in.MemberId)
	}

	// 用户筛选
	if len(in.ComplexMemberId) == 2 && len(in.ComplexMemberId[0]) > 0 {
		memberIds, err := service.AdminMember().GetComplexMemberIds(ctx, in.ComplexMemberId[0], in.ComplexMemberId[1])
		if err != nil {
			return nil, 0, err
		}
		if len(memberIds) == 0 {
			return nil, 0, nil
		}
		mod = mod.WhereIn(dao.AdminCash.Columns().MemberId, memberIds)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AdminCash.Columns().CreatedAt, gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	// 请求方式
	if in.Status > 0 {
		mod = mod.Where(dao.AdminCash.Columns().Status, in.Status)
	}

	if !isSuper {
		mod = mod.Where(dao.AdminCash.Columns().MemberId, opMemberId)
	}

	// 申请人摘要信息
	mod = mod.Hook(hook.MemberSummary)

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order(dao.AdminCash.Columns().Id + " DESC").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

// Apply 申请提现
func (s *sAdminCash) Apply(ctx context.Context, in *adminin.CashApplyInp) (err error) {
	var (
		config *model.CashConfig
		member *entity.AdminMember
	)

	if in.Money <= 0 {
		err = gerror.New("请输入提现金额")
		return
	}

	if in.Money != float64(int(in.Money)) {
		err = gerror.New("提现金额必须是正整数")
		return
	}

	count, err := dao.AdminCash.Ctx(ctx).
		Where(dao.AdminCash.Columns().MemberId, in.MemberId).
		Where(dao.AdminCash.Columns().Status, consts.StatusEnabled).
		Count()
	if err != nil {
		return
	}

	if count > 0 {
		err = gerror.New("存在正在处理中的提现，请耐心等待处理后再试！")
		return
	}

	if err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, in.MemberId).Scan(&member); err != nil {
		err = gerror.Newf("获取用户信息失败:%+v", err.Error())
		return
	}

	if member == nil {
		err = gerror.Newf("获取用户信息失败")
		return
	}

	if member.Balance < in.Money {
		err = gerror.Newf("余额不足")
		return
	}

	// 提现信息
	var cash adminin.MemberCash
	if member.Cash.IsNil() {
		err = gerror.Newf("请先设置提现账户！")
		return
	}

	if err = gconv.Scan(member.Cash, &cash); err != nil {
		return
	}

	if cash.Name == "" || cash.Account == "" || cash.PayeeCode == "" {
		err = gerror.New("请设置完整的提现账户信息！")
		return
	}

	conf, err := service.SysConfig().GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "cash"})
	if err != nil {
		return
	}

	if err = gconv.Struct(conf.List, &config); err != nil {
		return err
	}

	if !config.Switch {
		err = gerror.New("提现通道正在升级维护中，请稍后再试！")
		return
	}

	if in.Money < config.MinMoney {
		err = gerror.Newf("单次提现金额不能低于 %v 元", config.MinMoney)
		return
	}

	fee := in.Money * config.MinFeeRatio
	if fee < config.MinFee {
		fee = config.MinFee
	}

	lastMoney := in.Money - fee
	if lastMoney <= 1 {
		err = gerror.Newf("提现金额过少，请增加提现金额！")
		return
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 添加提现记录
		lastInsertId, err := dao.AdminCash.Ctx(ctx).Data(g.Map{
			dao.AdminCash.Columns().MemberId:  in.MemberId,
			dao.AdminCash.Columns().Money:     in.Money,
			dao.AdminCash.Columns().Fee:       fee,
			dao.AdminCash.Columns().LastMoney: lastMoney,
			dao.AdminCash.Columns().CreatedAt: gtime.Now(),
			dao.AdminCash.Columns().Status:    consts.CashStatusWait,
			dao.AdminCash.Columns().Msg:       "",
			dao.AdminCash.Columns().Ip:        location.GetClientIp(ghttp.RequestFromCtx(ctx)),
		}).OmitEmptyData().InsertAndGetId()
		if err != nil {
			return
		}

		// 更新余额
		_, err = service.AdminCreditsLog().SaveBalance(ctx, &adminin.CreditsLogSaveBalanceInp{
			MemberId:    in.MemberId,
			AppId:       contexts.GetModule(ctx),
			AddonsName:  contexts.GetAddonName(ctx),
			CreditGroup: consts.CreditGroupApplyCash,
			Num:         -in.Money,
			MapId:       lastInsertId,
			Remark:      "后台申请提现",
		})

		return
	})

	if err != nil {
		err = gerror.Newf("申请提现失败, %+v", err)
		return
	}
	return
}

// Payment 提现打款处理
func (s *sAdminCash) Payment(ctx context.Context, in *adminin.CashPaymentInp) (err error) {
	if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		err = gerror.New("没有访问权限")
		return
	}

	var models *entity.AdminCash
	if err = dao.AdminCash.Ctx(ctx).Where(dao.AdminCash.Columns().Id, in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.New("未找到提现信息")
		return
	}

	if models.Status == consts.CashStatusOk {
		err = gerror.New("该提现已处理成功，不能再次操作！")
		return
	}

	_, err = dao.AdminCash.Ctx(ctx).Where(dao.AdminCash.Columns().Id, models.Id).Data(g.Map{
		dao.AdminCash.Columns().HandleAt: gtime.Now(),
		dao.AdminCash.Columns().Status:   in.Status,
		dao.AdminCash.Columns().Msg:      gstr.Trim(in.Msg, " "),
	}).Update()
	return
}
