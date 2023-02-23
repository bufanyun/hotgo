// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// LogListInp 获取菜单列表
type LogListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Module     string
	MemberId   int
	TakeUpTime int
	Method     string
	Url        string
	Ip         string
	ErrorCode  string
	StartTime  string
	EndTime    string
	CreatedAt  []int64
}

type LogListModel struct {
	entity.SysLog
	MemberName string `json:"memberName"`
	Region     string `json:"region"`
}

// LogViewInp 获取信息
type LogViewInp struct {
	Id string
}

type LogViewModel struct {
	entity.SysLog
	CityLabel string `json:"cityLabel"          description:"城市标签"`
}

// LogDeleteInp 删除
type LogDeleteInp struct {
	Id interface{}
}
type LogDeleteModel struct{}
