// Package log
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package log

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ClearReq 清空日志
type ClearReq struct {
	g.Meta `path:"/log/clear" method:"post" tags:"日志" summary:"清空日志"`
}
type ClearRes struct{}

// ExportReq 导出
type ExportReq struct {
	g.Meta `path:"/log/export" method:"get" tags:"日志" summary:"导出日志"`
	form.PageReq
	form.RangeDateReq
	Module     string `json:"module"   dc:"应用端口"`
	MemberId   int    `json:"member_id"   dc:"用户ID"`
	TakeUpTime int    `json:"take_up_time"   dc:"请求耗时"`
	Method     string `json:"method"   dc:"请求方式"`
	Url        string `json:"url"   dc:"请求路径"`
	Ip         string `json:"ip"   dc:"访问IP"`
	ErrorCode  string `json:"error_code"   dc:"状态码"`
}
type ExportRes struct{}

// ListReq 获取菜单列表
type ListReq struct {
	g.Meta `path:"/log/list" method:"get" tags:"日志" summary:"获取日志列表"`
	form.PageReq
	form.RangeDateReq
	Module     string  `json:"module"   dc:"应用端口"`
	MemberId   int     `json:"member_id"   dc:"用户ID"`
	TakeUpTime int     `json:"take_up_time"   dc:"请求耗时"`
	Method     string  `json:"method"   dc:"请求方式"`
	Url        string  `json:"url"   dc:"请求路径"`
	Ip         string  `json:"ip"   dc:"访问IP"`
	ErrorCode  string  `json:"error_code"   dc:"状态码"`
	CreatedAt  []int64 `json:"created_at "   dc:"访问时间区间"`
}

type ListRes struct {
	List []*sysin.LogListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/log/delete" method:"post" tags:"日志" summary:"删除日志"`
	Id     interface{} `json:"id" v:"required#日志ID不能为空" description:"日志ID"`
}
type DeleteRes struct{}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/log/view" method:"get" tags:"日志" summary:"获取指定信息"`
	Id     string `json:"id" v:"required#日志ID不能为空" description:"日志ID"`
}
type ViewRes struct {
	*sysin.LogViewModel
}
