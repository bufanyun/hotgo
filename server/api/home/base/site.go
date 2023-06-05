// Package base
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package base

import "github.com/gogf/gf/v2/frame/g"

type SiteIndexReq struct {
	g.Meta `path:"/index" method:"get" summary:"首页" tags:"首页"`
}

type SiteIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
