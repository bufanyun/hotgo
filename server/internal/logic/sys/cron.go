// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/cron"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sSysCron struct{}

func NewSysCron() *sSysCron {
	return &sSysCron{}
}

func init() {
	service.RegisterSysCron(NewSysCron())
}

func (s *sSysCron) StartCron(ctx context.Context) {
	var list []*entity.SysCron
	if err := dao.SysCron.Ctx(ctx).
		Where("status", consts.StatusEnabled).
		Order("sort asc,id desc").
		Scan(&list); err != nil {
		cron.Logger().Fatalf(ctx, "定时任务获取失败, err . %v", err)
		return
	}

	if err := cron.StartALL(list); err != nil {
		cron.Logger().Fatalf(ctx, "定时任务启动失败, err . %v", err)
		return
	}
}

// Delete 删除
func (s *sSysCron) Delete(ctx context.Context, in *sysin.CronDeleteInp) (err error) {
	var models *entity.SysCron
	if err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("定时任务不存在或已被删除")
		return
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Delete(); err != nil {
			return
		}
		return cron.Delete(models)
	})
	return
}

// Edit 修改/新增
func (s *sSysCron) Edit(ctx context.Context, in *sysin.CronEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("标题不能为空")
		return
	}

	// 修改
	if in.Id > 0 {
		if _, err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Data(in).Update(); err != nil {
			return
		}
		simple.SafeGo(ctx, func(ctx context.Context) {
			_ = cron.RefreshStatus(&in.SysCron)
		})
		return
	}

	// 新增
	in.SysCron.Id, err = dao.SysCron.Ctx(ctx).Data(in).OmitEmptyData().InsertAndGetId()
	if err != nil || in.SysCron.Id < 1 {
		return
	}

	if in.SysCron.Status == consts.StatusEnabled {
		simple.SafeGo(ctx, func(ctx context.Context) {
			_ = cron.Start(&in.SysCron)
		})
	}
	return
}

// Status 更新状态
func (s *sSysCron) Status(ctx context.Context, in *sysin.CronStatusInp) (err error) {
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

	var models *entity.SysCron
	if err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.New("定时任务不存在")
		return
	}

	_, err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		return
	}

	models.Status = in.Status
	simple.SafeGo(ctx, func(ctx context.Context) {
		_ = cron.RefreshStatus(models)
	})
	return
}

// MaxSort 最大排序
func (s *sSysCron) MaxSort(ctx context.Context, in *sysin.CronMaxSortInp) (res *sysin.CronMaxSortModel, err error) {
	if err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
		return
	}

	if res == nil {
		res = new(sysin.CronMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取指定信息
func (s *sSysCron) View(ctx context.Context, in *sysin.CronViewInp) (res *sysin.CronViewModel, err error) {
	if err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		return
	}
	return
}

// List 获取列表
func (s *sSysCron) List(ctx context.Context, in *sysin.CronListInp) (list []*sysin.CronListModel, totalCount int, err error) {
	mod := dao.SysCron.Ctx(ctx)

	if in.Name != "" {
		mod = mod.WhereLike("name", "%"+in.Name+"%")
	}

	if in.GroupId > 0 {
		mod = mod.Where("group_id", in.GroupId)
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	for _, v := range list {
		v.GroupName, _ = s.GetName(ctx, v.GroupId)
	}
	return
}

// GetName 获取分组名称
func (s *sSysCron) GetName(ctx context.Context, id int64) (name string, err error) {
	m := dao.SysCronGroup.Ctx(ctx).Fields(dao.SysCronGroup.Columns().Name).Where(dao.SysCronGroup.Columns().Id, id)
	list, err := m.Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return name, err
	}
	return list.String(), nil
}

// OnlineExec 在线执行
func (s *sSysCron) OnlineExec(ctx context.Context, in *sysin.OnlineExecInp) (err error) {
	var data *entity.SysCron
	if err = dao.SysCron.Ctx(ctx).Where(dao.SysCron.Columns().Id, in.Id).Scan(&data); err != nil {
		return
	}

	if data == nil {
		err = gerror.New("定时任务不存在")
		return
	}
	return cron.Once(gctx.New(), data)
}

// DispatchLog 查看指定任务的调度日志
func (s *sSysCron) DispatchLog(ctx context.Context, in *sysin.DispatchLogInp) (res *sysin.DispatchLogModel, err error) {
	var data *entity.SysCron
	if err = dao.SysCron.Ctx(ctx).Where(dao.SysCron.Columns().Id, in.Id).Scan(&data); err != nil {
		return
	}

	if data == nil {
		err = gerror.New("定时任务不存在")
		return
	}

	res = new(sysin.DispatchLogModel)
	res.Log, err = cron.DispatchLog(data)
	return
}
