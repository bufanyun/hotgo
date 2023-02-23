// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/internal/consts"
	"hotgo/internal/crons"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
	"strings"
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
		g.Log().Fatalf(ctx, "定时任务获取失败, err . %v", err)
		return
	}

	if err := crons.StartALL(list); err != nil {
		g.Log().Fatalf(ctx, "定时任务启动失败, err . %v", err)
		return
	}
}

// Delete 删除
func (s *sSysCron) Delete(ctx context.Context, in sysin.CronDeleteInp) (err error) {
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
		_, err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Delete()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return crons.Delete(models)
	})
	return nil
}

// Edit 修改/新增
func (s *sSysCron) Edit(ctx context.Context, in sysin.CronEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("标题不能为空")
		return
	}

	// 修改
	if in.Id > 0 {
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Data(in).Update()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return err
			}

			simple.SafeGo(ctx, func(ctx context.Context) {
				crons.RefreshStatus(&in.SysCron)
			})
			return
		})
		return
	}

	// 新增
	_, err = dao.SysCron.Ctx(ctx).Data(in).Insert()
	return
}

// Status 更新部门状态
func (s *sSysCron) Status(ctx context.Context, in sysin.CronStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return err
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return err
	}

	if !validate.InSliceInt(consts.StatusMap, in.Status) {
		err = gerror.New("状态不正确")
		return err
	}

	var models *entity.SysCron
	if err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("定时任务不存在")
		return
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		models.Status = in.Status
		simple.SafeGo(ctx, func(ctx context.Context) {
			crons.RefreshStatus(models)
		})
		return
	})
	return
}

// MaxSort 最大排序
func (s *sSysCron) MaxSort(ctx context.Context, in sysin.CronMaxSortInp) (res *sysin.CronMaxSortModel, err error) {
	if err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}
	if res == nil {
		res = new(sysin.CronMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return res, nil
}

// View 获取指定字典类型信息
func (s *sSysCron) View(ctx context.Context, in sysin.CronViewInp) (res *sysin.CronViewModel, err error) {
	if err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// List 获取列表
func (s *sSysCron) List(ctx context.Context, in sysin.CronListInp) (list []*sysin.CronListModel, totalCount int, err error) {
	mod := dao.SysCron.Ctx(ctx)

	if in.Name != "" {
		mod = mod.WhereLike("name", "%"+in.Name+"%")
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, nil
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for _, v := range list {
		v.GroupName, _ = dao.SysCronGroup.GetName(ctx, v.GroupId)
	}
	return list, totalCount, err
}

// OnlineExec 在线执行
func (s *sSysCron) OnlineExec(ctx context.Context, in sysin.OnlineExecInp) (err error) {
	var data *entity.SysCron
	err = dao.SysCron.Ctx(ctx).Where(dao.SysCron.Columns().Id, in.Id).Scan(&data)
	if err != nil {
		return
	}

	if data == nil {
		return gerror.New("定时任务不存在")
	}

	newCtx := context.WithValue(gctx.New(), consts.ContextKeyCronArgs, strings.Split(data.Params, consts.CronSplitStr))
	return crons.Once(newCtx, data)
}
