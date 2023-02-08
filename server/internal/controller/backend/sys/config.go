// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/config"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Config = cConfig{}
)

type cConfig struct{}

// GetConfig 获取指定分组的配置
func (c *cConfig) GetConfig(ctx context.Context, req *config.GetReq) (res *config.GetRes, err error) {
	var in sysin.GetConfigInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	res = new(config.GetRes)
	res.GetConfigModel, err = service.SysConfig().GetConfigByGroup(ctx, in)
	return
}

// UpdateConfig 更新指定分组的配置
func (c *cConfig) UpdateConfig(ctx context.Context, req *config.UpdateReq) (res *config.UpdateRes, err error) {
	var in sysin.UpdateConfigInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysConfig().UpdateConfigByGroup(ctx, in)
	return
}

// TypeSelect 数据类型选项
func (c *cConfig) TypeSelect(ctx context.Context, req *config.TypeSelectReq) (res config.TypeSelectRes, err error) {
	for _, v := range consts.ConfigTypes {
		res = append(res, form.Select{
			Value: v,
			Name:  v,
			Label: v,
		})
	}
	return
}
