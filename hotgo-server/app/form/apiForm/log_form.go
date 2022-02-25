package apiForm

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
	form.PageReq
	form.RangeDateReq
	Module     string `json:"module"   description:"应用端口"`
	MemberId   int    `json:"member_id"   description:"用户ID"`
	TakeUpTime int    `json:"take_up_time"   description:"请求耗时"`
	Method     string `json:"method"   description:"请求方式"`
	Url        string `json:"url"   description:"请求路径"`
	Ip         string `json:"ip"   description:"访问IP"`
	ErrorCode  string `json:"error_code"   description:"状态码"`
	g.Meta     `path:"/log/export" method:"get" tags:"日志" summary:"导出日志"`
}
type LogExportRes struct{}

//  获取菜单列表
type LogListReq struct {
	form.PageReq
	form.RangeDateReq
	Module     string `json:"module"   description:"应用端口"`
	MemberId   int    `json:"member_id"   description:"用户ID"`
	TakeUpTime int    `json:"take_up_time"   description:"请求耗时"`
	Method     string `json:"method"   description:"请求方式"`
	Url        string `json:"url"   description:"请求路径"`
	Ip         string `json:"ip"   description:"访问IP"`
	ErrorCode  string `json:"error_code"   description:"状态码"`
	g.Meta     `path:"/log/list" method:"get" tags:"日志" summary:"获取日志列表"`
}

type LogListRes struct {
	List []*input.LogListModel `json:"list"   description:"数据列表"`
	form.PageRes
}
