// Package blacklist
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package blacklist

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Title   string `json:"title"`
	Content string `json:"content"`
	g.Meta  `path:"/blacklist/list" method:"get" tags:"黑名单" summary:"获取黑名单列表"`
}

type ListRes struct {
	List []*sysin.BlacklistListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取信息
type ViewReq struct {
	Id     int64 `json:"id" v:"required#黑名单ID不能为空" dc:"黑名单ID"`
	g.Meta `path:"/blacklist/view" method:"get" tags:"黑名单" summary:"获取指定信息"`
}
type ViewRes struct {
	*sysin.BlacklistViewModel
}

// EditReq 修改/新增
type EditReq struct {
	entity.SysBlacklist
	g.Meta `path:"/blacklist/edit" method:"post" tags:"黑名单" summary:"修改/新增黑名单"`
}
type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#黑名单ID不能为空" dc:"黑名单ID"`
	g.Meta `path:"/blacklist/delete" method:"post" tags:"黑名单" summary:"删除黑名单"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	Id     int64 `json:"id" dc:"黑名单ID"`
	g.Meta `path:"/blacklist/maxSort" method:"get" tags:"黑名单" summary:"黑名单最大排序"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新状态
type StatusReq struct {
	entity.SysBlacklist
	g.Meta `path:"/blacklist/status" method:"post" tags:"黑名单" summary:"更新黑名单状态"`
}
type StatusRes struct{}
