// Package cron
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
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

// GroupListReq 查询列表
type GroupListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Title   string `json:"title"`
	Content string `json:"content"`
	g.Meta  `path:"/cronGroup/list" method:"get" tags:"定时任务分组" summary:"获取定时任务分组列表"`
}

type GroupListRes struct {
	List []*sysin.CronGroupListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// GroupViewReq 获取信息
type GroupViewReq struct {
	Id     int64 `json:"id" v:"required#定时任务分组ID不能为空" dc:"定时任务分组ID"`
	g.Meta `path:"/cronGroup/view" method:"get" tags:"定时任务分组" summary:"获取指定信息"`
}
type GroupViewRes struct {
	*sysin.CronGroupViewModel
}

// GroupEditReq 修改/新增
type GroupEditReq struct {
	entity.SysCronGroup
	g.Meta `path:"/cronGroup/edit" method:"post" tags:"定时任务分组" summary:"修改/新增定时任务分组"`
}
type GroupEditRes struct{}

// GroupDeleteReq 删除
type GroupDeleteReq struct {
	Id     interface{} `json:"id" v:"required#定时任务分组ID不能为空" dc:"定时任务分组ID"`
	g.Meta `path:"/cronGroup/delete" method:"post" tags:"定时任务分组" summary:"删除定时任务分组"`
}
type GroupDeleteRes struct{}

// GroupMaxSortReq 最大排序
type GroupMaxSortReq struct {
	Id     int64 `json:"id" dc:"定时任务分组ID"`
	g.Meta `path:"/cronGroup/maxSort" method:"get" tags:"定时任务分组" summary:"定时任务分组最大排序"`
}
type GroupMaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// GroupStatusReq 更新状态
type GroupStatusReq struct {
	entity.SysCronGroup
	g.Meta `path:"/cronGroup/status" method:"post" tags:"定时任务分组" summary:"更新定时任务分组状态"`
}
type GroupStatusRes struct{}

// GroupSelectReq 定时任务分组选项
type GroupSelectReq struct {
	g.Meta `path:"/cronGroup/select" method:"get" tags:"定时任务分组" summary:"定时任务分组选项"`
}

type GroupSelectRes sysin.DictTypeSelectModel
