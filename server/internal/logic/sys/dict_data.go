// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

type sSysDictData struct{}

func NewSysDictData() *sSysDictData {
	return &sSysDictData{}
}

func init() {
	service.RegisterSysDictData(NewSysDictData())
}

// Delete 删除
func (s *sSysDictData) Delete(ctx context.Context, in sysin.DictDataDeleteInp) error {
	_, err := dao.SysDictData.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysDictData) Edit(ctx context.Context, in sysin.DictDataEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		_, err = dao.SysDictData.Ctx(ctx).Fields(sysin.DictDataUpdateFields{}).WherePri(in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.Type, err = dao.SysDictType.GetType(ctx, in.TypeID)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if in.Type == "" {
		return gerror.Wrap(err, "类型选择无效，请检查")
	}

	_, err = dao.SysDictData.Ctx(ctx).Fields(sysin.DictDataInsertFields{}).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// List 获取列表
func (s *sSysDictData) List(ctx context.Context, in sysin.DictDataListInp) (list []*sysin.DictDataListModel, totalCount int, err error) {
	mod := dao.SysDictData.Ctx(ctx)
	// 类型ID
	if in.TypeID > 0 {
		types, err := dao.SysDictType.GetTypes(ctx, in.TypeID)
		if err != nil {
			return list, totalCount, err
		}
		mod = mod.WhereIn("type", types)
	}

	if in.Type != "" {
		mod = mod.Where("type", in.Type)
	}

	// 访问路径
	if in.Label != "" {
		mod = mod.WhereLike("label", "%"+in.Label+"%")
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

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for _, v := range list {
		v.TypeID, _ = dao.SysDictType.GetId(ctx, v.Type)
	}

	return list, totalCount, err
}

// Select 获取列表
func (s *sSysDictData) Select(ctx context.Context, in sysin.DataSelectInp) (list sysin.DataSelectModel, err error) {
	mod := dao.SysDictData.Ctx(ctx).Where("type", in.Type)
	if in.Type != "" {
		mod = mod.Where("type", in.Type)
	}

	if err = mod.Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	for k, v := range list {
		list[k].Value = consts.ConvType(v.Value, v.ValueType)
		list[k].Key = list[k].Value
	}
	return list, err
}
