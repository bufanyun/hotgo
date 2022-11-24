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
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Config = cConfig{}
)

type cConfig struct{}

// GetConfig 获取指定分组的配置
func (c *cConfig) GetConfig(ctx context.Context, req *config.GetReq) (*config.GetRes, error) {
	var (
		in  sysin.GetConfigInp
		res config.GetRes
		err error
	)
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	res.GetConfigModel, err = service.SysConfig().GetConfigByGroup(ctx, in)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UpdateConfig 更新指定分组的配置
func (c *cConfig) UpdateConfig(ctx context.Context, req *config.UpdateReq) (*config.UpdateRes, error) {
	var (
		in  sysin.UpdateConfigInp
		res config.UpdateRes
		err error
	)
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = service.SysConfig().UpdateConfigByGroup(ctx, in); err != nil {
		return nil, err
	}

	return &res, nil
}
