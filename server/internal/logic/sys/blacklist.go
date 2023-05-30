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
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/validate"
	"sync"
)

type sSysBlacklist struct {
	sync.RWMutex
}

func NewSysBlacklist() *sSysBlacklist {
	return &sSysBlacklist{}
}

func init() {
	service.RegisterSysBlacklist(NewSysBlacklist())
}

// Delete 删除
func (s *sSysBlacklist) Delete(ctx context.Context, in sysin.BlacklistDeleteInp) (err error) {
	defer s.VariableLoad(ctx, err)
	_, err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit 修改/新增
func (s *sSysBlacklist) Edit(ctx context.Context, in sysin.BlacklistEditInp) (err error) {
	defer s.VariableLoad(ctx, err)
	if in.Ip == "" {
		err = gerror.New("ip不能为空")
		return
	}

	// 修改
	if in.Id > 0 {
		_, err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		return
	}

	// 新增
	_, err = dao.SysBlacklist.Ctx(ctx).Data(in).Insert()
	return
}

// Status 更新部门状态
func (s *sSysBlacklist) Status(ctx context.Context, in sysin.BlacklistStatusInp) (err error) {
	defer s.VariableLoad(ctx, err)
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSliceInt(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	// 修改
	_, err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

// MaxSort 最大排序
func (s *sSysBlacklist) MaxSort(ctx context.Context, in sysin.BlacklistMaxSortInp) (res *sysin.BlacklistMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			return
		}
	}

	if res == nil {
		res = new(sysin.BlacklistMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return
}

// View 获取指定字典类型信息
func (s *sSysBlacklist) View(ctx context.Context, in sysin.BlacklistViewInp) (res *sysin.BlacklistViewModel, err error) {
	err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}

// List 获取列表
func (s *sSysBlacklist) List(ctx context.Context, in sysin.BlacklistListInp) (list []*sysin.BlacklistListModel, totalCount int, err error) {
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
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		return
	}
	return
}

// VariableLoad 变化加载
func (s *sSysBlacklist) VariableLoad(ctx context.Context, err error) {
	if err == nil {
		s.Load(ctx)
	}
}

// Load 加载黑名单
func (s *sSysBlacklist) Load(ctx context.Context) {
	s.RLock()
	defer s.RUnlock()

	global.Blacklists = make(map[string]struct{})

	array, err := dao.SysBlacklist.Ctx(ctx).
		Fields(dao.SysBlacklist.Columns().Ip).
		Where(dao.SysBlacklist.Columns().Status, consts.StatusEnabled).
		Array()
	if err != nil {
		g.Log().Fatalf(ctx, "load blacklist fail：%+v", err)
		return
	}

	for _, v := range array {
		list := convert.IpFilterStrategy(v.String())
		if len(list) > 0 {
			for k := range list {
				global.Blacklists[k] = struct{}{}
			}
		}
	}
}
