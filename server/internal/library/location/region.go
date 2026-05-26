// Package location
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package location

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

func GetModel(ctx context.Context) *gdb.Model {
	return dao.SysProvinces.Ctx(ctx)
}

// ParseSimpleRegion 通过地区ID解析地区名称，自动加入上级地区
func ParseSimpleRegion(ctx context.Context, id int64, spilt ...string) (string, error) {
	if id == 0 {
		return "", nil
	}
	var (
		models *entity.SysProvinces
		err    error
	)

	if err = GetModel(ctx).Fields(dao.SysProvinces.Columns().Title, dao.SysProvinces.Columns().Level, dao.SysProvinces.Columns().Tree).Where(dao.SysProvinces.Columns().Id, id).Scan(&models); err != nil {
		return "", err
	}

	if models == nil {
		return "", gerror.Newf("the area code :%v is not in the database", id)
	}

	if models.Level == 1 {
		return models.Title, nil
	}

	ids := tree.GetIds(models.Tree)

	if models.Level == 2 {
		if len(ids) != 1 {
			return "", gerror.Newf("the region code is incorrectly configured, models:%+v, ids:%v", models, ids)
		}
		return ParseRegion(ctx, ids[0], id, 0, spilt...)
	}

	if models.Level == 3 {
		if len(ids) != 2 {
			return "", gerror.Newf("the region code is incorrectly configured, models:%+v, ids:%v", models, ids)
		}
		return ParseRegion(ctx, ids[0], ids[1], id, spilt...)
	}
	return "", gerror.New("currently, it is only supported to regional areas")
}

// ParseRegion 解析省市编码对应的地区名称
func ParseRegion(ctx context.Context, province int64, city int64, county int64, spilt ...string) (string, error) {
	var (
		provinceName *gvar.Var
		cityName     *gvar.Var
		countyName   *gvar.Var
		err          error
	)

	// 分隔符
	sp := consts.RegionSpilt
	if len(spilt) > 0 {
		sp = spilt[0]
	}

	if province > 0 && province < 999999 {
		provinceName, err = GetModel(ctx).Where(dao.SysProvinces.Columns().Id, province).Fields(dao.SysProvinces.Columns().Title).Value()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return "", err
		}

		if city > 0 {
			cityName, err = GetModel(ctx).Where(dao.SysProvinces.Columns().Id, city).Fields(dao.SysProvinces.Columns().Title).Value()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return "", err
			}
		}

		if county > 0 {
			countyName, err = GetModel(ctx).Where(dao.SysProvinces.Columns().Id, county).Fields(dao.SysProvinces.Columns().Title).Value()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return "", err
			}
		}
	} else {
		return "保留地址", nil
	}

	if province > 0 && city > 0 && county > 0 {
		return provinceName.String() + sp + cityName.String() + sp + countyName.String(), nil
	}

	if province > 0 && city > 0 {
		return provinceName.String() + sp + cityName.String(), nil
	}
	return provinceName.String(), nil
}
