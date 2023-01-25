// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// ServeLogDeleteInp 删除服务日志
type ServeLogDeleteInp struct {
	Id interface{} `json:"id" v:"required#日志ID不能为空" dc:"日志ID"`
}

type ServeLogDeleteModel struct{}

// ServeLogViewInp 获取指定服务日志信息
type ServeLogViewInp struct {
	Id int64 `json:"id" v:"required#日志ID不能为空" dc:"日志ID"`
}

type ServeLogViewModel struct {
	entity.Test
}

// ServeLogListInp 获取服务日志列表
type ServeLogListInp struct {
	form.PageReq
	TraceId     string        `json:"traceId"          dc:"链路ID"`
	LevelFormat string        `json:"levelFormat"      dc:"日志级别"`
	TriggerNs   []int64       `json:"triggerNs"        dc:"触发时间(ns)"`
	CreatedAt   []*gtime.Time `json:"createdAt"        dc:"创建时间"`
}

type ServeLogListModel struct {
	Id          int64       `json:"id"               dc:"日志ID"`
	TraceId     string      `json:"traceId"          dc:"链路ID"`
	LevelFormat string      `json:"levelFormat"      dc:"日志级别"`
	Content     string      `json:"content"          dc:"日志内容"`
	Stack       string      `json:"stack"            dc:"堆栈"`
	Line        string      `json:"line"             dc:"调用行"`
	TriggerNs   int64       `json:"triggerNs"        dc:"触发时间(ns)"`
	CreatedAt   *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"        dc:"修改时间"`
	SysLogId    int64       `json:"sysLogId"         dc:"访问日志ID"`
}

func (in *ServeLogListInp) Filter(ctx context.Context) (err error) {
	return
}

// ServeLogExportModel 导出服务日志
type ServeLogExportModel struct {
	Id          int64       `json:"id"               dc:"日志ID"`
	Env         string      `json:"env"              dc:"环境"`
	TraceId     string      `json:"traceId"          dc:"链路ID"`
	LevelFormat string      `json:"levelFormat"      dc:"日志级别"`
	Content     string      `json:"content"          dc:"日志内容"`
	Line        string      `json:"line"             dc:"调用行"`
	TriggerNs   int64       `json:"triggerNs"        dc:"触发时间(ns)"`
	CreatedAt   *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"        dc:"修改时间"`
}
