// Package config
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// GetReq 获取指定分组的配置
type GetReq struct {
	g.Meta `path:"/config/get" method:"get" tags:"配置" summary:"获取指定分组的配置"`
	sysin.GetConfigInp
}

type GetRes struct {
	*sysin.GetConfigModel
}

// UpdateReq 获取指定分组的配置
type UpdateReq struct {
	g.Meta `path:"/config/update" method:"post" tags:"配置" summary:"获取指定分组的配置"`
	sysin.UpdateConfigInp
}

type UpdateRes struct {
}

// TypeSelectReq 数据类型选项
type TypeSelectReq struct {
	g.Meta `path:"/config/typeSelect" method:"get" tags:"配置" summary:"数据类型选项"`
}

type TypeSelectRes []form.Select

// GetCashReq 获取提现配置
type GetCashReq struct {
	g.Meta `path:"/config/getCash" method:"get" tags:"配置" summary:"获取提现配置"`
}

type GetCashRes struct {
	*sysin.GetConfigModel
}
