// Package monitor
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package monitor

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/form"
)

// UserOfflineReq 下线用户
type UserOfflineReq struct {
	g.Meta `path:"/monitor/userOffline" method:"post" tags:"在线用户" summary:"下线用户"`
	Id     string `json:"id" v:"required#SID不能为空" description:"SID"`
}

type UserOfflineRes struct{}

// UserOnlineListReq 获取在线用户列表
type UserOnlineListReq struct {
	g.Meta `path:"/monitor/userOnlineList" method:"get" tags:"在线用户" summary:"获取在线用户列表"`
	form.PageReq
	UserId    int64         `json:"userId"      description:"用户ID"`
	Username  string        `json:"username"    description:"用户名"`
	Addr      string        `json:"addr"        description:"登录地址"`
	FirstTime []*gtime.Time `json:"firstTime"   description:"登录时间"`
}

type UserOnlineListRes struct {
	List []*UserOnlineModel `json:"list"   description:"数据列表"`
	form.PageRes
}

type UserOnlineModel struct {
	ID            string `json:"id"`            // 连接唯一标识
	Addr          string `json:"addr"`          // 客户端地址
	Os            string `json:"os"`            // 客户端系统名称
	Browser       string `json:"browser"`       // 浏览器
	FirstTime     int64  `json:"firstTime"`     // 首次连接时间
	HeartbeatTime uint64 `json:"heartbeatTime"` // 用户上次心跳时间
	App           string `json:"app"`           // 应用名称
	UserId        int64  `json:"userId"`        // 用户ID
	Username      string `json:"username"`      // 用户名
	Avatar        string `json:"avatar"`        // 头像
}

// NetOnlineListReq 获取在线服务列表
type NetOnlineListReq struct {
	g.Meta `path:"/monitor/netOnlineList" method:"get" tags:"在线服务" summary:"获取在线服务列表"`
	form.PageReq
	Name      string        `json:"name"        description:"应用名称"`
	Group     string        `json:"group"       description:"分组"`
	AppId     string        `json:"appId"       description:"APPID"`
	Addr      string        `json:"addr"        description:"登录地址"`
	FirstTime []*gtime.Time `json:"firstTime"   description:"登录时间"`
}

type NetOnlineListRes struct {
	List []*NetOnlineModel `json:"list"   description:"数据列表"`
	form.PageRes
}

type NetOnlineModel struct {
	*tcp.AuthMeta
	Id            int64  `json:"id"               description:"连接ID"`
	IsAuth        bool   `json:"isAuth"           description:"是否认证"`
	Addr          string `json:"addr"             description:"登录地址"`
	Port          string `json:"port"             description:"连接端口"`
	FirstTime     int64  `json:"firstTime"        description:"首次连接时间"`
	HeartbeatTime int64  `json:"heartbeatTime"    description:"上次心跳时间"`
	LicenseId     int64  `json:"licenseId"        description:"许可ID"`
	LicenseName   string `json:"licenseName"      description:"许可名称"`
	LoginTimes    int64  `json:"loginTimes"       description:"许可累计登录次数"`
	Online        int    `json:"online"           description:"许可在线数量"`
	OnlineLimit   int    `json:"onlineLimit"      description:"许可在线数量限制"`
	Desc          string `json:"desc"             description:"许可说明"`
	Proto         string `json:"proto"            description:"网络协议"`
}

// NetOptionReq 获取服务选项
type NetOptionReq struct {
	g.Meta `path:"/monitor/netOption" method:"get" tags:"在线服务" summary:"获取服务选项"`
}

type NetOptionRes struct {
	LicenseGroup form.Selects   `json:"licenseGroup" dc:"授权分组"`
	Routes       []*RouteSelect `json:"routes" dc:"路由选项"`
}

type RouteSelect struct {
	Value    interface{} `json:"value"`
	Label    string      `json:"label"`
	Disabled bool        `json:"disabled"`
	IsRPC    bool        `json:"isRPC"`
}

// NetOfflineReq 下线服务
type NetOfflineReq struct {
	g.Meta `path:"/monitor/netOffline" method:"post" tags:"在线服务" summary:"下线服务"`
	Id     int64 `json:"id" v:"required#连接ID不能为空" description:"连接ID"`
}

type NetOfflineRes struct{}
