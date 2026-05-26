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
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sSysCronGroup struct{}

func NewSysCronGroup() *sSysCronGroup {
	return &sSysCronGroup{}
}

func init() {
	service.RegisterSysCronGroup(NewSysCronGroup())
}

// Delete 删除
func (s *sSysCronGroup) Delete(ctx context.Context, in *sysin.CronGroupDeleteInp) (err error) {
	_, err = dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit 修改/新增
func (s *sSysCronGroup) Edit(ctx context.Context, in *sysin.CronGroupEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = dao.SysCronGroup.Ctx(ctx).Fields(sysin.CronGroupUpdateFields{}).Where("id", in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
		}
		return
	}

	// 新增
	if _, err = dao.SysCronGroup.Ctx(ctx).Fields(sysin.CronGroupInsertFields{}).Data(in).OmitEmptyData().Insert(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// Status 更新状态
func (s *sSysCronGroup) Status(ctx context.Context, in *sysin.CronGroupStatusInp) (err error) {
	if _, err = dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// MaxSort 最大排序
func (s *sSysCronGroup) MaxSort(ctx context.Context, in *sysin.CronGroupMaxSortInp) (res *sysin.CronGroupMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	}

	if res == nil {
		res = new(sysin.CronGroupMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取指定信息
func (s *sSysCronGroup) View(ctx context.Context, in *sysin.CronGroupViewInp) (res *sysin.CronGroupViewModel, err error) {
	if err = dao.SysCronGroup.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// List 获取列表
func (s *sSysCronGroup) List(ctx context.Context, in *sysin.CronGroupListInp) (list []*sysin.CronGroupListModel, totalCount int, err error) {
	mod := dao.SysCronGroup.Ctx(ctx)

	if in.Name != "" {
		mod = mod.WhereLike("name", "%"+in.Name+"%")
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

	if err = mod.Page(in.Page, in.PerPage).OrderDesc(dao.SysCronGroup.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}

	for _, v := range list {
		if v.Pid < 1 {
			continue
		}
		name, err := dao.SysCronGroup.Ctx(ctx).Fields(dao.SysCronGroup.Columns().Name).WherePri(v.Pid).Value()
		if err != nil {
			return nil, 0, err
		}
		v.SupName = name.String()
	}
	return
}

// Select 选项
func (s *sSysCronGroup) Select(ctx context.Context, in *sysin.CronGroupSelectInp) (res *sysin.CronGroupSelectModel, err error) {
	var (
		mod    = dao.SysCronGroup.Ctx(ctx)
		models []*entity.SysCronGroup
	)

	if err = mod.Order("pid asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	res = new(sysin.CronGroupSelectModel)
	res.List = s.treeList(0, models)
	return
}

// treeList 树状列表
func (s *sSysCronGroup) treeList(pid int64, nodes []*entity.SysCronGroup) (list []*sysin.CronGroupTree) {
	list = make([]*sysin.CronGroupTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(sysin.CronGroupTree)
			item.SysCronGroup = *v
			item.Label = v.Name
			item.Value = v.Id
			item.Key = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}
