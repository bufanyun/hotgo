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
	Module     string  `json:"module"         dc:"应用端口"`
	MemberId   int     `json:"member_id"      dc:"用户ID"`
	TakeUpTime int     `json:"take_up_time"   dc:"请求耗时"`
	Method     string  `json:"method"         dc:"请求方式"`
	Url        string  `json:"url"            dc:"请求路径"`
	Ip         string  `json:"ip"             dc:"访问IP"`
	ErrorCode  string  `json:"error_code"     dc:"状态码"`
	CreatedAt  []int64 `json:"created_at"     dc:"创建时间"`
}

type LogListModel struct {
	entity.SysLog
	MemberName string `json:"memberName"`
	Region     string `json:"region"`
}

// LogViewInp 获取信息
type LogViewInp struct {
	Id string `json:"id" v:"required#日志ID不能为空" description:"日志ID"`
}

type LogViewModel struct {
	entity.SysLog
	CityLabel string `json:"cityLabel"          description:"城市标签"`
}

// LogDeleteInp 删除
type LogDeleteInp struct {
	Id interface{} `json:"id" v:"required#日志ID不能为空" description:"日志ID"`
}

type LogDeleteModel struct{}
