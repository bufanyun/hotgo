// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

type sSysBlacklist struct{}

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
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysBlacklist) Edit(ctx context.Context, in sysin.BlacklistEditInp) (err error) {
	defer s.VariableLoad(ctx, err)
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
	defer s.VariableLoad(ctx, err)
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

// VariableLoad 变化加载
func (s *sSysBlacklist) VariableLoad(ctx context.Context, err error) {
	if err == nil {
		s.Load(ctx)
	}
}

// Load 加载黑名单
func (s *sSysBlacklist) Load(ctx context.Context) {
	global.Blacklists = make(map[string]struct{})

	array, err := dao.SysBlacklist.Ctx(ctx).
		Fields(dao.SysBlacklist.Columns().Ip).
		Where(dao.SysBlacklist.Columns().Status, consts.StatusEnabled).
		Array()
	if err != nil {
		g.Log().Fatalf(ctx, "load blacklist fail：%+v", err)
		return
	}

	matchStrategy := func(originIp string) {
		// 多个IP
		if gstr.Contains(originIp, ",") {
			ips := gstr.Explode(",", originIp)
			if len(ips) > 0 {
				for _, ip := range ips {
					if !validate.IsIp(ip) {
						continue
					}
					global.Blacklists[ip] = struct{}{}
				}
			}

			return
		}

		// IP段
		if gstr.Contains(originIp, "/24") {
			segment := gstr.Replace(originIp, "/24", "")
			if !validate.IsIp(segment) {
				return
			}

			var (
				start  = gstr.Explode(".", segment)
				prefix = gstr.Implode(".", start[:len(start)-1]) + "."
				index  = gconv.Int(start[len(start)-1])
			)

			if index < 1 {
				index = 1
			}

			for i := index; i <= 254; i++ {
				global.Blacklists[prefix+gconv.String(i)] = struct{}{}
			}

			return
		}

		// IP范围
		if gstr.Contains(originIp, "-") {
			originIps := gstr.Explode("-", originIp)
			if len(originIps) != 2 {
				return
			}

			if !validate.IsIp(originIps[0]) || !validate.IsIp(originIps[1]) {
				return
			}

			var (
				start      = gstr.Explode(".", originIps[0])
				prefix     = gstr.Implode(".", start[:len(start)-1]) + "."
				startIndex = gconv.Int(gstr.SubStrFromREx(originIps[0], "."))
				endIndex   = gconv.Int(gstr.SubStrFromREx(originIps[1], "."))
			)

			if startIndex >= endIndex {
				global.Blacklists[originIps[0]] = struct{}{}
				return
			}

			if startIndex < 1 {
				startIndex = 1
			}

			if endIndex > 254 {
				endIndex = 254
			}

			for i := startIndex; i <= endIndex; i++ {
				global.Blacklists[prefix+gconv.String(i)] = struct{}{}
			}
			return
		}

		// 指定IP
		if validate.IsIp(originIp) {
			global.Blacklists[originIp] = struct{}{}
			return
		}
	}

	for _, v := range array {
		matchStrategy(v.String())
	}
}
