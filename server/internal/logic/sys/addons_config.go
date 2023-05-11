// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var AddonsMaskDemoField []string

type sSysAddonsConfig struct{}

func NewSysAddonsConfig() *sSysAddonsConfig {
	return &sSysAddonsConfig{}
}

func init() {
	service.RegisterSysAddonsConfig(NewSysAddonsConfig())
}

// GetConfigByGroup 获取指定分组的配置
func (s *sSysAddonsConfig) GetConfigByGroup(ctx context.Context, in sysin.GetAddonsConfigInp) (res *sysin.GetAddonsConfigModel, err error) {
	if in.AddonName == "" {
		err = gerror.New("插件名称不能为空")
		return
	}

	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}

	var (
		mod    = dao.SysAddonsConfig.Ctx(ctx)
		models []*entity.SysAddonsConfig
	)

	if err = mod.Fields("key", "value", "type").
		Where("addon_name", in.AddonName).
		Where("group", in.Group).
		Scan(&models); err != nil {
		return nil, err
	}
	isDemo := g.Cfg().MustGet(ctx, "hotgo.isDemo", false)

	if len(models) > 0 {
		res = new(sysin.GetAddonsConfigModel)
		res.List = make(g.Map, len(models))
		for _, v := range models {
			val, err := s.ConversionType(ctx, v)
			if err != nil {
				return nil, err
			}
			res.List[v.Key] = val
			if isDemo.Bool() && gstr.InArray(AddonsMaskDemoField, v.Key) {
				res.List[v.Key] = consts.DemoTips
				res.List[v.Key] = consts.DemoTips
			}
		}
	}

	return
}

// ConversionType 转换类型
func (s *sSysAddonsConfig) ConversionType(ctx context.Context, models *entity.SysAddonsConfig) (value interface{}, err error) {
	if models == nil {
		err = gerror.New("数据不存在")
		return
	}
	return consts.ConvType(models.Value, models.Type), nil
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysAddonsConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateAddonsConfigInp) (err error) {
	if in.AddonName == "" {
		err = gerror.New("插件名称不能为空")
		return
	}

	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}

	var (
		mod    = dao.SysAddonsConfig.Ctx(ctx)
		models []*entity.SysAddonsConfig
	)

	if err = mod.
		Where("addon_name", in.AddonName).
		Where("group", in.Group).
		Scan(&models); err != nil {
		return err
	}

	err = dao.SysAddonsConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for k, v := range in.List {
			row := s.getConfigByKey(k, models)
			// 新增
			if row == nil {
				//row.Id = 0
				//row.Key = k
				//row.Value = gconv.String(v)
				//row.Group = in.Group
				//row.Status = consts.StatusEnabled
				//row.CreatedAt = gtime.Now()
				//row.UpdatedAt = gtime.Now()
				//_, err := dao.SysAddonsConfig.Ctx(ctx).Data(row).Insert()
				//if err != nil {
				//	err = gerror.Wrap(err, consts.ErrorORM)
				//	return err
				//}
				//continue
				return gerror.Newf("暂不支持从前台添加变量，请先在数据库表[%v]中配置变量：%v", dao.SysAddonsConfig.Table(), k)
			}

			// 更新
			_, err = dao.SysAddonsConfig.Ctx(ctx).Where("id", row.Id).Data(g.Map{"value": v, "updated_at": gtime.Now()}).Update()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *sSysAddonsConfig) getConfigByKey(key string, models []*entity.SysAddonsConfig) *entity.SysAddonsConfig {
	if len(models) == 0 {
		return nil
	}

	for _, v := range models {
		if key == v.Key {
			return v
		}
	}

	return nil
}
