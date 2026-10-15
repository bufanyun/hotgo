// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import "github.com/gogf/gf/v2/frame/g"

// HealthReq 健康检查。挂在根路径，不走 /api 前缀。
type HealthReq struct {
	g.Meta `path:"/health" method:"get" tags:"基础" summary:"健康检查"`
}

type HealthRes struct {
	Status string `json:"status" dc:"ok 表示数据库可连通"`
}
