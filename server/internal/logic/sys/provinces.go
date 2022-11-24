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
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
)

type sSysProvinces struct{}

func NewSysProvinces() *sSysProvinces {
	return &sSysProvinces{}
}

func init() {
	service.RegisterSysProvinces(NewSysProvinces())
}

// Delete 删除
func (s *sSysProvinces) Delete(ctx context.Context, in sysin.ProvincesDeleteInp) error {

	_, err := dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Delete()
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
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
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

	if !convert.InSliceInt(consts.StatusMap, in.Status) {
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
func (s *sSysProvinces) MaxSort(ctx context.Context, in sysin.ProvincesMaxSortInp) (*sysin.ProvincesMaxSortModel, error) {
	var res sysin.ProvincesMaxSortModel

	if in.Id > 0 {
		if err := dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
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

	// 访问路径
	if in.Title != "" {
		mod = mod.WhereLike("title", "%"+in.Title+"%")
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

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	return list, totalCount, err
}
