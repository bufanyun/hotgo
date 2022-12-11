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

type sSysBlacklist struct{}

func NewSysBlacklist() *sSysBlacklist {
	return &sSysBlacklist{}
}

func init() {
	service.RegisterSysBlacklist(NewSysBlacklist())
}

// Delete 删除
func (s *sSysBlacklist) Delete(ctx context.Context, in sysin.BlacklistDeleteInp) error {
	_, err := dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysBlacklist) Edit(ctx context.Context, in sysin.BlacklistEditInp) (err error) {
	if in.Ip == "" {
		err = gerror.New("ip不能为空")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.SysBlacklist.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// Status 更新部门状态
func (s *sSysBlacklist) Status(ctx context.Context, in sysin.BlacklistStatusInp) (err error) {
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
	_, err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 最大排序
func (s *sSysBlacklist) MaxSort(ctx context.Context, in sysin.BlacklistMaxSortInp) (*sysin.BlacklistMaxSortModel, error) {
	var res sysin.BlacklistMaxSortModel
	if in.Id > 0 {
		if err := dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}
	res.Sort = res.Sort + 10
	return &res, nil
}

// View 获取指定字典类型信息
func (s *sSysBlacklist) View(ctx context.Context, in sysin.BlacklistViewInp) (res *sysin.BlacklistViewModel, err error) {
	if err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}
	return res, nil
}

// List 获取列表
func (s *sSysBlacklist) List(ctx context.Context, in sysin.BlacklistListInp) (list []*sysin.BlacklistListModel, totalCount int64, err error) {
	mod := dao.SysBlacklist.Ctx(ctx)

	// 访问路径
	if in.Ip != "" {
		mod = mod.Where("ip", in.Ip)
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
