//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/gogf/gf/v2/frame/g"
)

//  清空日志
type LogClearReq struct {
	g.Meta `path:"/log/clear" method:"post" tags:"日志" summary:"清空日志"`
}
type LogClearRes struct{}

//  导出
type LogExportReq struct {
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
type LogExportRes struct{}

//  获取菜单列表
type LogListReq struct {
	g.Meta `path:"/log/list" method:"get" tags:"日志" summary:"获取日志列表"`
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

type LogListRes struct {
	List []*input.LogListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}
