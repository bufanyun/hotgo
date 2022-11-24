// Package config
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/sysin"
)

// GetReq 获取指定分组的配置
type GetReq struct {
	Group  string `json:"group" dc:"分组名称" v:"required#分组名称不能为空" `
	g.Meta `path:"/config/get" method:"get" tags:"配置" summary:"获取指定分组的配置"`
}
type GetRes struct {
	*sysin.GetConfigModel
}

// UpdateReq 获取指定分组的配置
type UpdateReq struct {
	Group  string `json:"group" dc:"分组名称" v:"required#分组名称不能为空" `
	List   g.Map  `json:"list" dc:"更新配置列表" `
	g.Meta `path:"/config/update" method:"post" tags:"配置" summary:"获取指定分组的配置"`
}
type UpdateRes struct {
}
