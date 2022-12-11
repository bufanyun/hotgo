// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/crons"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
)

type sSysCron struct{}

func NewSysCron() *sSysCron {
	return &sSysCron{}
}

func init() {
	service.RegisterSysCron(NewSysCron())
}

func (s *sSysCron) StartCron(ctx context.Context) {
	var (
		list []*entity.SysCron
	)

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
func (s *sSysCron) Delete(ctx context.Context, in sysin.CronDeleteInp) error {
	_, err := dao.SysCron.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysCron) Edit(ctx context.Context, in sysin.CronEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("标题不能为空")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.SysCron.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
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

	if !convert.InSliceInt(consts.StatusMap, in.Status) {
		err = gerror.New("状态不正确")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	_, err = dao.SysCron.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 最大排序
func (s *sSysCron) MaxSort(ctx context.Context, in sysin.CronMaxSortInp) (*sysin.CronMaxSortModel, error) {
	var res sysin.CronMaxSortModel

	if in.Id > 0 {
		if err := dao.SysCron.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
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
func (s *sSysCron) List(ctx context.Context, in sysin.CronListInp) (list []*sysin.CronListModel, totalCount int64, err error) {
	mod := dao.SysCron.Ctx(ctx)

	// 访问路径
	if in.Name != "" {
		mod = mod.WhereLike("name", "%"+in.Name+"%")
	}

	// 请求方式
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

	if err = mod.Page(int(in.Page), int(in.PerPage)).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for _, v := range list {
		v.GroupName, _ = dao.SysCronGroup.GetName(ctx, v.GroupId)
	}

	return list, totalCount, err
}
