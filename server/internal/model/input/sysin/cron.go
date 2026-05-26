// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"hotgo/internal/library/cron"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// CronMaxSortInp 最大排序
type CronMaxSortInp struct {
	Id int64
}

type CronMaxSortModel struct {
	Sort int
}

// CronEditInp 修改/新增
type CronEditInp struct {
	entity.SysCron
}
type CronEditModel struct{}

// CronDeleteInp 删除
type CronDeleteInp struct {
	Id interface{}
}
type CronDeleteModel struct{}

// CronViewInp 获取信息
type CronViewInp struct {
	Id int64
}

type CronViewModel struct {
	entity.SysCron
}

// CronListInp 获取列表
type CronListInp struct {
	form.PageReq
	form.StatusReq
	GroupId int64  `json:"groupId"   description:"分组ID"`
	Name    string `json:"name"      description:"任务名称"`
}

type CronListModel struct {
	entity.SysCron
	GroupName string `json:"groupName"`
}

// CronStatusInp 更新状态
type CronStatusInp struct {
	entity.SysCron
}
type CronStatusModel struct{}

// OnlineExecInp 在线执行
type OnlineExecInp struct {
	entity.SysCron
}
type OnlineExecModel struct{}

// DispatchLogInp 查看指定任务的调度日志
type DispatchLogInp struct {
	entity.SysCron
}
type DispatchLogModel struct {
	*cron.Log
}
