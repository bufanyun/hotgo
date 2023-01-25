// Package monitor
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package monitor

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
)

// OfflineReq 下线用户
type OfflineReq struct {
	g.Meta `path:"/monitor/offline" method:"post" tags:"在线用户" summary:"下线用户"`
	Id     string `json:"id" v:"required#SID不能为空" description:"SID"`
}
type OfflineRes struct{}

// OnlineListReq 获取在线用户列表
type OnlineListReq struct {
	g.Meta `path:"/monitor/onlineList" method:"get" tags:"在线用户" summary:"获取监控列表"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	UserId int64  `json:"userId"   description:"用户ID"`
	Addr   string `json:"addr" description:"登录地址"`
}

type OnlineListRes struct {
	List []*OnlineModel `json:"list"   description:"数据列表"`
	form.PageRes
}

// OnlineViewReq  获取指定信息
type OnlineViewReq struct {
	g.Meta `path:"/monitor/onlineView" method:"get" tags:"在线用户" summary:"获取指定用户信息"`
	Id     string `json:"id" v:"required#SID不能为空" description:"SID"`
}
type OnlineViewRes struct {
	*OnlineModel
}

type OnlineModel struct {
	ID            string `json:"id"`            // 连接唯一标识
	Addr          string `json:"addr"`          // 客户端地址
	Os            string `json:"os"`            // 客户端系统名称
	Browser       string `json:"browser"`       // 浏览器
	FirstTime     uint64 `json:"firstTime"`     // 首次连接时间
	HeartbeatTime uint64 `json:"heartbeatTime"` // 用户上次心跳时间
	App           string `json:"app"`           // 应用名称
	UserId        int64  `json:"userId"`        // 用户ID
	Username      string `json:"username"`      // 用户名
	Avatar        string `json:"avatar"`        // 头像
	ExpTime       int64  `json:"expTime"`       // 过期时间
}

type OnlineModels []*OnlineModel

func (p OnlineModels) Len() int {
	return len(p)
}
func (p OnlineModels) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p OnlineModels) Less(i, j int) bool {
	return p[j].FirstTime < p[i].FirstTime
}
