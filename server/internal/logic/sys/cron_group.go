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
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/tree"
)

type sSysCronGroup struct{}

func NewSysCronGroup() *sSysCronGroup {
	return &sSysCronGroup{}
}

func init() {
	service.RegisterSysCronGroup(NewSysCronGroup())
}

// Delete 删除
func (s *sSysCronGroup) Delete(ctx context.Context, in sysin.CronGroupDeleteInp) error {
	_, err := dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysCronGroup) Edit(ctx context.Context, in sysin.CronGroupEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("分组名称不能为空")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.SysCronGroup.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// Status 更新部门状态
func (s *sSysCronGroup) Status(ctx context.Context, in sysin.CronGroupStatusInp) (err error) {
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
	_, err = dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 最大排序
func (s *sSysCronGroup) MaxSort(ctx context.Context, in sysin.CronGroupMaxSortInp) (*sysin.CronGroupMaxSortModel, error) {
	var res sysin.CronGroupMaxSortModel

	if in.Id > 0 {
		if err := dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

// View 获取指定字典类型信息
func (s *sSysCronGroup) View(ctx context.Context, in sysin.CronGroupViewInp) (res *sysin.CronGroupViewModel, err error) {
	if err = dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// List 获取列表
func (s *sSysCronGroup) List(ctx context.Context, in sysin.CronGroupListInp) (list []*sysin.CronGroupListModel, totalCount int64, err error) {
	mod := dao.SysCronGroup.Ctx(ctx)

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

	return list, totalCount, err
}

// Select 选项
func (s *sSysCronGroup) Select(ctx context.Context, in sysin.CronGroupSelectInp) (list sysin.CronGroupSelectModel, err error) {
	var (
		mod      = dao.SysCronGroup.Ctx(ctx)
		models   []*entity.SysCronGroup
		typeList []g.Map
	)

	if err = mod.Order("pid asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	for i := 0; i < len(models); i++ {
		typeList = append(typeList, g.Map{
			"index":      models[i].Id,
			"key":        models[i].Id,
			"label":      models[i].Name,
			"id":         models[i].Id,
			"pid":        models[i].Pid,
			"name":       models[i].Name,
			"sort":       models[i].Sort,
			"created_at": models[i].CreatedAt,
			"status":     models[i].Status,
		})
	}

	return tree.GenTree(typeList), nil
}
