// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.15.7
package sys

import (
	"context"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

type sSysTestCategory struct{}

func NewSysTestCategory() *sSysTestCategory {
	return &sSysTestCategory{}
}

func init() {
	service.RegisterSysTestCategory(NewSysTestCategory())
	dict.RegisterFunc("testCategoryOption", "测试分类选项", service.SysTestCategory().Option)
}

// Model 测试分类ORM模型
func (s *sSysTestCategory) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TestCategory.Ctx(ctx), option...)
}

// List 获取测试分类列表
func (s *sSysTestCategory) List(ctx context.Context, in *sysin.TestCategoryListInp) (list []*sysin.TestCategoryListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.TestCategoryListModel{})

	// 查询分类ID
	if in.Id > 0 {
		mod = mod.Where(dao.TestCategory.Columns().Id, in.Id)
	}

	// 查询分类名称
	if in.Name != "" {
		mod = mod.WhereLike(dao.TestCategory.Columns().Name, "%"+in.Name+"%")
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.TestCategory.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TestCategory.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.TestCategory.Columns().Sort).OrderDesc(dao.TestCategory.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取测试分类列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增测试分类
func (s *sSysTestCategory) Edit(ctx context.Context, in *sysin.TestCategoryEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.TestCategoryUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改测试分类失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.TestCategoryInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增测试分类失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除测试分类
func (s *sSysTestCategory) Delete(ctx context.Context, in *sysin.TestCategoryDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除测试分类失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取测试分类最大排序
func (s *sSysTestCategory) MaxSort(ctx context.Context, in *sysin.TestCategoryMaxSortInp) (res *sysin.TestCategoryMaxSortModel, err error) {
	if err = dao.TestCategory.Ctx(ctx).Fields(dao.TestCategory.Columns().Sort).OrderDesc(dao.TestCategory.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试分类最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.TestCategoryMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取测试分类指定信息
func (s *sSysTestCategory) View(ctx context.Context, in *sysin.TestCategoryViewInp) (res *sysin.TestCategoryViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试分类信息，请稍后重试！")
		return
	}
	return
}

// Status 更新测试分类状态
func (s *sSysTestCategory) Status(ctx context.Context, in *sysin.TestCategoryStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TestCategory.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新测试分类状态失败，请稍后重试！")
		return
	}
	return
}

// Option 获取测试分类选项
func (s *sSysTestCategory) Option(ctx context.Context) (opts []*model.Option, err error) {
	var models []*entity.TestCategory
	if err = s.Model(ctx).Fields(dao.TestCategory.Columns().Id, dao.TestCategory.Columns().Name).
		OrderAsc(dao.TestCategory.Columns().Sort).OrderDesc(dao.TestCategory.Columns().Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取测试分类选项失败！")
		return
	}

	opts = make([]*model.Option, len(models))
	for k, v := range models {
		opts[k] = dict.GenHashOption(v.Id, gconv.String(v.Name))
	}
	return
}
