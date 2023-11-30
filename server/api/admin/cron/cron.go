// Package cron
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cron

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/cron/list" method:"get" tags:"定时任务" summary:"获取定时任务列表"`
	sysin.CronListInp
}

type ListRes struct {
	List []*sysin.CronListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取信息
type ViewReq struct {
	g.Meta `path:"/cron/view" method:"get" tags:"定时任务" summary:"获取指定信息"`
	sysin.CronViewInp
}

type ViewRes struct {
	*sysin.CronViewModel
}

// EditReq 修改/新增
type EditReq struct {
	g.Meta `path:"/cron/edit" method:"post" tags:"定时任务" summary:"修改/新增定时任务"`
	sysin.CronEditInp
}

type EditRes struct {
	*sysin.CronEditModel
}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/cron/delete" method:"post" tags:"定时任务" summary:"删除定时任务"`
	sysin.CronDeleteInp
}

type DeleteRes struct {
	*sysin.CronDeleteModel
}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/cron/maxSort" method:"get" tags:"定时任务" summary:"定时任务最大排序"`
	sysin.CronMaxSortInp
}

type MaxSortRes struct {
	*sysin.CronMaxSortModel
}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/cron/status" method:"post" tags:"定时任务" summary:"更新定时任务状态"`
	sysin.CronStatusInp
}

type StatusRes struct {
	*sysin.CronStatusModel
}

// OnlineExecReq 在线执行
type OnlineExecReq struct {
	g.Meta `path:"/cron/onlineExec" method:"post" tags:"定时任务" summary:"在线执行"`
	sysin.OnlineExecInp
}

type OnlineExecRes struct {
	*sysin.OnlineExecModel
}

// DispatchLogReq 调度日志
type DispatchLogReq struct {
	g.Meta `path:"/cron/dispatchLog" method:"post" tags:"定时任务" summary:"调度日志"`
	sysin.DispatchLogInp
}

type DispatchLogRes struct {
	*sysin.DispatchLogModel
}
