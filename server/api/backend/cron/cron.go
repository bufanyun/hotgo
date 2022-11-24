// Package cron
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package cron

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
	g.Meta  `path:"/cron/list" method:"get" tags:"定时任务" summary:"获取定时任务列表"`
}

type ListRes struct {
	List []*sysin.CronListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取信息
type ViewReq struct {
	Id     int64 `json:"id" v:"required#定时任务ID不能为空" dc:"定时任务ID"`
	g.Meta `path:"/cron/view" method:"get" tags:"定时任务" summary:"获取指定信息"`
}
type ViewRes struct {
	*sysin.CronViewModel
}

// EditReq 修改/新增
type EditReq struct {
	entity.SysCron
	g.Meta `path:"/cron/edit" method:"post" tags:"定时任务" summary:"修改/新增定时任务"`
}
type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#定时任务ID不能为空" dc:"定时任务ID"`
	g.Meta `path:"/cron/delete" method:"post" tags:"定时任务" summary:"删除定时任务"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	Id     int64 `json:"id" dc:"定时任务ID"`
	g.Meta `path:"/cron/max_sort" method:"get" tags:"定时任务" summary:"定时任务最大排序"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新状态
type StatusReq struct {
	entity.SysCron
	g.Meta `path:"/cron/status" method:"post" tags:"定时任务" summary:"更新定时任务状态"`
}
type StatusRes struct{}
