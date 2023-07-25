// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.5.3
// @AutoGenerate Date 2023-04-15 15:59:58
package admin

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdminCreditsLog struct{}

func NewAdminCreditsLog() *sAdminCreditsLog {
	return &sAdminCreditsLog{}
}

func init() {
	service.RegisterAdminCreditsLog(NewAdminCreditsLog())
}

// Model 资产变动ORM模型
func (s *sAdminCreditsLog) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AdminCreditsLog.Ctx(ctx), option...)
}

// SaveBalance 更新余额
func (s *sAdminCreditsLog) SaveBalance(ctx context.Context, in *adminin.CreditsLogSaveBalanceInp) (res *adminin.CreditsLogSaveBalanceModel, err error) {
	if err = validate.PreFilter(ctx, in); err != nil {
		return
	}

	var (
		mb        *entity.AdminMember
		daoMember = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, in.MemberId)
		data      = new(entity.AdminCreditsLog)
	)

	err = dao.AdminMember.Ctx(ctx).
		Fields(dao.AdminMember.Columns().Balance).
		Where(dao.AdminMember.Columns().Id, in.MemberId).
		Scan(&mb)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New("用户不存在！")
		return
	}

	if in.Num > 0 {
		if _, err = daoMember.Increment(dao.AdminMember.Columns().Balance, in.Num); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	} else {
		num := in.Num * -1
		if mb.Balance < num {
			err = gerror.Newf("余额不足，当前余额为：%v，需要扣除的余额为：%v", mb.Balance, num)
			return
		}
		if _, err = daoMember.Decrement(dao.AdminMember.Columns().Balance, num); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	}

	data.MemberId = in.MemberId
	data.AppId = in.AppId
	data.AddonsName = in.AddonsName
	data.CreditType = consts.CreditTypeBalance
	data.CreditGroup = in.CreditGroup
	data.BeforeNum = mb.Balance
	data.Num = in.Num
	data.AfterNum = mb.Balance + in.Num
	data.Remark = in.Remark
	data.Ip = in.Ip
	data.MapId = in.MapId

	_, err = dao.AdminCreditsLog.Ctx(ctx).Data(data).Insert()
	return
}

// SaveIntegral 更新积分
func (s *sAdminCreditsLog) SaveIntegral(ctx context.Context, in *adminin.CreditsLogSaveIntegralInp) (res *adminin.CreditsLogSaveIntegralModel, err error) {
	if err = validate.PreFilter(ctx, in); err != nil {
		return
	}

	var (
		mb        *entity.AdminMember
		daoMember = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, in.MemberId)
		data      = new(entity.AdminCreditsLog)
	)

	err = dao.AdminMember.Ctx(ctx).
		Fields(dao.AdminMember.Columns().Integral).
		Where(dao.AdminMember.Columns().Id, in.MemberId).
		Scan(&mb)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New("用户不存在！")
		return
	}

	if in.Num > 0 {
		if _, err = daoMember.Increment(dao.AdminMember.Columns().Integral, in.Num); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	} else {
		num := in.Num * -1
		if mb.Integral < num {
			err = gerror.Newf("积分不足，当前积分为：%v，需要扣除的积分为：%v", mb.Integral, num)
			return
		}
		if _, err = daoMember.Decrement(dao.AdminMember.Columns().Integral, num); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	}

	data.MemberId = in.MemberId
	data.AppId = in.AppId
	data.AddonsName = in.AddonsName
	data.CreditType = consts.CreditTypeIntegral
	data.CreditGroup = in.CreditGroup
	data.BeforeNum = mb.Integral
	data.Num = in.Num
	data.AfterNum = mb.Integral + in.Num
	data.Remark = in.Remark
	data.Ip = in.Ip
	data.MapId = in.MapId

	_, err = dao.AdminCreditsLog.Ctx(ctx).Data(data).Insert()
	return
}

// List 获取资产变动列表
func (s *sAdminCreditsLog) List(ctx context.Context, in *adminin.CreditsLogListInp) (list []*adminin.CreditsLogListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询变动ID
	if in.Id > 0 {
		mod = mod.Where(dao.AdminCreditsLog.Columns().Id, in.Id)
	}

	// 查询管理员ID
	if in.MemberId > 0 {
		mod = mod.Where(dao.AdminCreditsLog.Columns().MemberId, in.MemberId)
	}

	// 查询应用id
	if in.AppId != "" {
		mod = mod.WhereLike(dao.AdminCreditsLog.Columns().AppId, in.AppId)
	}

	// 查询变动类型
	if in.CreditType != "" {
		mod = mod.WhereLike(dao.AdminCreditsLog.Columns().CreditType, in.CreditType)
	}

	// 查询变动的组别
	if in.CreditGroup != "" {
		mod = mod.WhereLike(dao.AdminCreditsLog.Columns().CreditGroup, in.CreditGroup)
	}

	// 查询备注
	if in.Remark != "" {
		mod = mod.WhereLike(dao.AdminCreditsLog.Columns().Remark, "%"+in.Remark+"%")
	}

	// 查询操作人IP
	if in.Ip != "" {
		mod = mod.WhereLike(dao.AdminCreditsLog.Columns().Ip, in.Ip)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.AdminCreditsLog.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AdminCreditsLog.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count(1)
	if err != nil {
		return
	}

	if totalCount == 0 {
		return
	}

	err = mod.Fields(adminin.CreditsLogListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.AdminCreditsLog.Columns().Id).Scan(&list)
	return
}

// Export 导出资产变动
func (s *sAdminCreditsLog) Export(ctx context.Context, in *adminin.CreditsLogListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(adminin.CreditsLogExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出资产变动-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []adminin.CreditsLogExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}
