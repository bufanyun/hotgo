// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// LogListInp 获取菜单列表
type LogListInp struct {
	form.PageReq
	form.StatusReq
	ReqId           string   `json:"reqId"          dc:"对外ID"`
	Module          string   `json:"module"         dc:"应用端口"`
	MemberId        int      `json:"memberId"       dc:"用户ID"`
	TakeUpTime      string   `json:"takeUpTime"     dc:"请求耗时"`
	Method          string   `json:"method"         dc:"请求方式"`
	Url             string   `json:"url"            dc:"请求路径"`
	Ip              string   `json:"ip"             dc:"访问IP"`
	ErrorCode       string   `json:"errorCode"      dc:"状态码"`
	CreatedAt       []int64  `json:"createdAt"      dc:"创建时间"`
	Keyword         string   `json:"keyword"        dc:"关键词"`
	ComplexMemberId []string `json:"complexMemberId" dc:"操作人筛选"`
}

type LogListModel struct {
	entity.SysLog
	MemberName  string `json:"memberName"`
	Region      string `json:"region"`
	CityLabel   string `json:"cityLabel"    dc:"城市标签"`
	Tags        string `json:"tags"         dc:"接口所属的标签，用于接口分类"`
	Summary     string `json:"summary"      dc:"接口/参数概要描述"`
	Description string `json:"description"  dc:"接口/参数详细描述"`
}

// LogViewInp 获取信息
type LogViewInp struct {
	Id string `json:"id" v:"required#日志ID不能为空" description:"日志ID"`
}

type LogViewModel struct {
	entity.SysLog
	CityLabel   string `json:"cityLabel"    dc:"城市标签"`
	Tags        string `json:"tags"         dc:"接口所属的标签，用于接口分类"`
	Summary     string `json:"summary"      dc:"接口/参数概要描述"`
	Description string `json:"description"  dc:"接口/参数详细描述"`
}

// LogDeleteInp 删除
type LogDeleteInp struct {
	Id interface{} `json:"id" v:"required#日志ID不能为空" description:"日志ID"`
}

type LogDeleteModel struct{}
