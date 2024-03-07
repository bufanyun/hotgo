// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/dict"
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
func (s *sSysDictData) Delete(ctx context.Context, in *sysin.DictDataDeleteInp) error {
	_, err := dao.SysDictData.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// Edit 修改/新增
func (s *sSysDictData) Edit(ctx context.Context, in *sysin.DictDataEditInp) (err error) {
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
	in.Type, err = s.GetType(ctx, in.TypeID)
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
func (s *sSysDictData) List(ctx context.Context, in *sysin.DictDataListInp) (list []*sysin.DictDataListModel, totalCount int, err error) {
	mod := dao.SysDictData.Ctx(ctx)
	// 类型ID
	if in.TypeID > 0 {
		types, err := s.GetTypes(ctx, in.TypeID)
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
		v.TypeID, _ = s.GetId(ctx, v.Type)
	}
	return list, totalCount, err
}

// GetId 获取指定类型的ID
func (s *sSysDictData) GetId(ctx context.Context, t string) (id int64, err error) {
	m := dao.SysDictType.Ctx(ctx).Fields("id").Where("type", t).Where("status", consts.StatusEnabled)
	val, err := m.Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return 0, err
	}
	return val.Int64(), nil
}

// GetType 获取指定ID的类型标识
func (s *sSysDictData) GetType(ctx context.Context, id int64) (types string, err error) {
	m := dao.SysDictType.Ctx(ctx).Fields("type").Where("id", id).Where("status", consts.StatusEnabled)
	val, err := m.Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return types, err
	}
	return val.String(), nil
}

// GetTypes 获取指定ID的所有类型标识，包含下级
func (s *sSysDictData) GetTypes(ctx context.Context, id int64) (types []string, err error) {
	columns, err := dao.SysDictType.Ctx(ctx).Fields("type").
		Where("id", id).WhereOr("pid", id).Where("status", consts.StatusEnabled).
		Array()
	types = g.NewVar(columns).Strings()
	return
}

// Select 获取列表
func (s *sSysDictData) Select(ctx context.Context, in *sysin.DataSelectInp) (list sysin.DataSelectModel, err error) {
	// 优先从内置字典中获取
	options, err := dict.GetOptions(ctx, in.Type)
	if err == nil {
		return options, nil
	}
	if !errors.Is(err, dict.NotExistKeyError) {
		return nil, err
	}

	mod := dao.SysDictData.Ctx(ctx).Where("type", in.Type)
	if err = mod.Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	for k, v := range list {
		list[k].Value = consts.ConvType(v.Value, v.ValueType)
		list[k].Key = list[k].Value
	}

	if len(list) == 0 {
		list = make(sysin.DataSelectModel, 0)
	}
	return
}
