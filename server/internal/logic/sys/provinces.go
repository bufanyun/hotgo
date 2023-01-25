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
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/tree"
	"hotgo/utility/validate"
)

type sSysProvinces struct{}

func NewSysProvinces() *sSysProvinces {
	return &sSysProvinces{}
}

func init() {
	service.RegisterSysProvinces(NewSysProvinces())
}

// Tree 关系树选项列表
func (s *sSysProvinces) Tree(ctx context.Context) (list []g.Map, err error) {
	var models []*entity.SysProvinces
	if err = dao.SysProvinces.Ctx(ctx).Order("pid asc,id asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	list = gconv.SliceMap(models)
	for k, v := range list {
		list[k]["key"] = v["id"]
		list[k]["label"] = v["title"]
	}

	return tree.GenTree(list), nil
}

// Delete 删除
func (s *sSysProvinces) Delete(ctx context.Context, in sysin.ProvincesDeleteInp) error {

	var (
		models *entity.SysProvinces
	)
	err := dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Scan(&models)
	if err != nil {
		return err
	}

	if models == nil {
		return gerror.New("数据不存在或已删除！")
	}

	pidExist, err := dao.SysProvinces.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !pidExist.IsEmpty() {
		return gerror.New("请先删除该地区下得所有子级！")
	}

	_, err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysProvinces) Edit(ctx context.Context, in sysin.ProvincesEditInp) (err error) {
	if in.Title == "" {
		err = gerror.New("标题不能为空")
		return
	}

	if in.Id <= 0 {
		err = gerror.New("地区Id必须大于0")
		return
	}

	// 关系树
	in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, dao.SysProvinces, in.Pid)
	if err != nil {
		return err
	}

	isUpdate := false
	models, err := s.View(ctx, sysin.ProvincesViewInp{Id: in.Id})
	if err != nil {
		return
	}

	if models != nil {
		isUpdate = true
	}

	// 修改
	if isUpdate {
		_, err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	_, err = dao.SysProvinces.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// Status 更新部门状态
func (s *sSysProvinces) Status(ctx context.Context, in sysin.ProvincesStatusInp) (err error) {
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

	// 修改
	in.UpdatedAt = gtime.Now()
	_, err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 最大排序
func (s *sSysProvinces) MaxSort(ctx context.Context, in sysin.ProvincesMaxSortInp) (res *sysin.ProvincesMaxSortModel, err error) {
	if err = dao.SysProvinces.Ctx(ctx).Fields(dao.SysProvinces.Columns().Sort).OrderDesc(dao.SysProvinces.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.Sort = res.Sort + g.Cfg().MustGet(ctx, "hotgo.admin.maxSortIncrement").Int()
	return res, nil
}

// View 获取指定字典类型信息
func (s *sSysProvinces) View(ctx context.Context, in sysin.ProvincesViewInp) (res *sysin.ProvincesViewModel, err error) {
	if err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
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

	return list, totalCount, err
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
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, nil
	}

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	return list, totalCount, err
}

// UniqueId 获取省市区下级列表
func (s *sSysProvinces) UniqueId(ctx context.Context, in sysin.ProvincesUniqueIdInp) (res *sysin.ProvincesUniqueIdModel, err error) {
	res = new(sysin.ProvincesUniqueIdModel)
	res.IsUnique = true
	if in.NewId == 0 {
		return
	}

	if err = hgorm.IsUnique(ctx, dao.SysProvinces, g.Map{dao.Test.Columns().Id: in.NewId}, "", in.OldId); err != nil {
		res.IsUnique = false
		return res, nil
	}

	return res, nil
}
