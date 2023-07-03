// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

type sSysProvinces struct{}

func NewSysProvinces() *sSysProvinces {
	return &sSysProvinces{}
}

func init() {
	service.RegisterSysProvinces(NewSysProvinces())
}

// Tree 关系树选项列表
func (s *sSysProvinces) Tree(ctx context.Context) (list []*sysin.ProvincesTree, err error) {
	var models []*entity.SysProvinces
	if err = dao.SysProvinces.Ctx(ctx).Order("pid asc,id asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取省市区关系树选项列表失败！")
		return
	}

	list = s.treeList(0, models)
	return
}

// Delete 删除省市区数据
func (s *sSysProvinces) Delete(ctx context.Context, in sysin.ProvincesDeleteInp) (err error) {
	var models *entity.SysProvinces
	if err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取省市区数据失败！")
		return
	}

	if models == nil {
		err = gerror.New("数据不存在或已删除！")
		return
	}

	has, err := dao.SysProvinces.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, "删除省市区数据时获取上级数据失败！")
		return
	}

	if !has.IsEmpty() {
		err = gerror.New("请先删除该地区下得所有子级！")
		return
	}

	if _, err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除省市区数据失败！")
		return
	}
	return
}

// Edit 修改/新增省市区数据
func (s *sSysProvinces) Edit(ctx context.Context, in sysin.ProvincesEditInp) (err error) {
	// 关系树
	in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, &dao.SysProvinces, in.Pid)
	if err != nil {
		return
	}

	models, err := s.View(ctx, sysin.ProvincesViewInp{Id: in.Id})
	if err != nil {
		return
	}

	// 修改
	if models != nil {
		if _, err = dao.SysProvinces.Ctx(ctx).Fields(sysin.ProvincesUpdateFields{}).WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改省市区数据失败！")
		}
		return
	}

	// 新增
	if _, err = dao.SysProvinces.Ctx(ctx).Fields(sysin.ProvincesInsertFields{}).Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增省市区数据失败！")
	}
	return
}

// Status 更新省市区状态
func (s *sSysProvinces) Status(ctx context.Context, in sysin.ProvincesStatusInp) (err error) {
	if _, err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update(); err != nil {
		err = gerror.Wrap(err, "更新省市区状态失败！")
	}
	return
}

// MaxSort 最大排序
func (s *sSysProvinces) MaxSort(ctx context.Context, in sysin.ProvincesMaxSortInp) (res *sysin.ProvincesMaxSortModel, err error) {
	if err = dao.SysProvinces.Ctx(ctx).Fields(dao.SysProvinces.Columns().Sort).OrderDesc(dao.SysProvinces.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取省市区最大排序失败！")
		return
	}

	if res == nil {
		res = new(sysin.ProvincesMaxSortModel)
	}
	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return
}

// View 获取省市区信息
func (s *sSysProvinces) View(ctx context.Context, in sysin.ProvincesViewInp) (res *sysin.ProvincesViewModel, err error) {
	if err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取省市区信息失败！")
	}
	return
}

// List 获取列表
func (s *sSysProvinces) List(ctx context.Context, in sysin.ProvincesListInp) (list []*sysin.ProvincesListModel, totalCount int, err error) {
	mod := dao.SysProvinces.Ctx(ctx)

	if in.Title != "" {
		mod = mod.WhereLike("title", "%"+in.Title+"%")
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取省市区数据行失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取省市区列表失败！")
	}
	return
}

// ChildrenList 获取省市区下级列表
func (s *sSysProvinces) ChildrenList(ctx context.Context, in sysin.ProvincesChildrenListInp) (list []*sysin.ProvincesChildrenListModel, totalCount int, err error) {
	mod := dao.SysProvinces.Ctx(ctx)

	if in.Title != "" {
		mod = mod.WhereLike("title", "%"+in.Title+"%")
	}

	if in.Pid > 0 {
		mod = mod.Where("pid", in.Pid)
	}

	if in.Id > 0 {
		mod = mod.Where("id", in.Id)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取省市区下级数据行失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取省市区下级列表失败！")
	}
	return
}

// UniqueId 获取省市区下级列表
func (s *sSysProvinces) UniqueId(ctx context.Context, in sysin.ProvincesUniqueIdInp) (res *sysin.ProvincesUniqueIdModel, err error) {
	res = new(sysin.ProvincesUniqueIdModel)
	res.IsUnique = true
	if in.NewId == 0 {
		return
	}

	if err = hgorm.IsUnique(ctx, &dao.SysProvinces, g.Map{dao.SysProvinces.Columns().Id: in.NewId}, "", in.OldId); err != nil {
		res.IsUnique = false
		return
	}
	return
}

// Select 省市区选项
func (s *sSysProvinces) Select(ctx context.Context, in sysin.ProvincesSelectInp) (res *sysin.ProvincesSelectModel, err error) {
	res = new(sysin.ProvincesSelectModel)
	mod := dao.SysProvinces.Ctx(ctx).Fields("id as value, title as label, level").Where("pid", in.Value)

	if err = mod.Order("sort asc,id asc").Scan(&res.List); err != nil {
		err = gerror.Wrap(err, "获取省市区选项失败！")
		return
	}

	for _, v := range res.List {
		if in.DataType == "p" {
			v.IsLeaf = true
			continue
		}

		if in.DataType == "pc" && v.Level >= 2 {
			v.IsLeaf = true
			continue
		}

		if in.DataType == "pca" && v.Level >= 3 {
			v.IsLeaf = true
			continue
		}
	}
	return
}

// treeList 树状列表
func (s *sSysProvinces) treeList(pid int64, nodes []*entity.SysProvinces) (list []*sysin.ProvincesTree) {
	list = make([]*sysin.ProvincesTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(sysin.ProvincesTree)
			item.SysProvinces = *v
			item.Label = v.Title
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
