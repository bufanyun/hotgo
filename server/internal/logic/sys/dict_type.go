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
	"hotgo/utility/tree"
)

type sSysDictType struct{}

func NewSysDictType() *sSysDictType {
	return &sSysDictType{}
}

func init() {
	service.RegisterSysDictType(NewSysDictType())
}

// Tree 树
func (s *sSysDictType) Tree(ctx context.Context) (list []g.Map, err error) {
	var (
		mod    = dao.SysDictType.Ctx(ctx)
		models []*entity.SysDictType
	)

	if err = mod.Order("pid asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	for i := 0; i < len(models); i++ {
		list = append(list, g.Map{
			"index":      models[i].Id,
			"key":        models[i].Id,
			"label":      models[i].Name,
			"id":         models[i].Id,
			"pid":        models[i].Pid,
			"name":       models[i].Name,
			"type":       models[i].Type,
			"sort":       models[i].Sort,
			"remark":     models[i].Remark,
			"status":     models[i].Status,
			"updated_at": models[i].UpdatedAt,
			"created_at": models[i].CreatedAt,
		})
	}

	return tree.GenTree(list), nil
}

// Delete 删除
func (s *sSysDictType) Delete(ctx context.Context, in sysin.DictTypeDeleteInp) error {

	var (
		models *entity.SysDictType
	)
	err := dao.SysDictType.Ctx(ctx).Where("id", in.Id).Scan(&models)
	if err != nil {
		return err
	}

	if models.Id < 1 {
		return gerror.New("数据不存在或已删除！")
	}

	exist, err := dao.SysDictData.Ctx(ctx).Where("type", models.Type).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("请先删除该字典类型下得所有字典数据！")
	}

	pidExist, err := dao.SysDictType.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !pidExist.IsEmpty() {
		return gerror.New("请先删除该字典类型下得所有子级类型！")
	}

	_, err = dao.SysDictType.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysDictType) Edit(ctx context.Context, in sysin.DictTypeEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return err
	}

	uniqueName, err := dao.SysDictType.IsUniqueType(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("名称已存在")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysDictType.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.SysDictType.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// Select 选项
func (s *sSysDictType) Select(ctx context.Context, in sysin.DictTypeSelectInp) (list sysin.DictTypeSelectModel, err error) {
	var (
		mod      = dao.SysDictType.Ctx(ctx)
		models   []*entity.SysDictType
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

// TreeSelect 获取类型关系树选项
func (s *sSysDictType) TreeSelect(ctx context.Context, in sysin.DictTreeSelectInp) (list sysin.DictTreeSelectModel, err error) {
	var (
		mod      = dao.SysDictType.Ctx(ctx)
		models   []*entity.SysDictType
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

	maps := tree.GenTree(typeList)
	for _, v := range maps {
		// 父类一律禁止选中
		if _, ok := v["children"]; ok {
			v["disabled"] = true
		}
	}
	return tree.GenTree(typeList), nil
}
