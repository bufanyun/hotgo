// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
)

// LoginLogoutReq 注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/site/logout" method:"post" tags:"后台基础" summary:"注销登录"`
}
type LoginLogoutRes struct{}

// LoginCaptchaReq 获取登录验证码
type LoginCaptchaReq struct {
	g.Meta `path:"/site/captcha" method:"get" tags:"后台基础" summary:"获取登录验证码"`
}
type LoginCaptchaRes struct {
	Cid    string `json:"cid" dc:"验证码ID"`
	Base64 string `json:"base64" dc:"验证码"`
}

// LoginReq 提交登录
type LoginReq struct {
	g.Meta   `path:"/site/login" method:"post" tags:"后台基础" summary:"账号登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Cid      string `json:"cid"  dc:"验证码ID"`
	Code     string `json:"code" dc:"验证码"`
	IsLock   bool   `json:"isLock"  dc:"是否为锁屏状态"`
}
type LoginRes struct {
	*adminin.MemberLoginModel
}

// SiteConfigReq 获取配置
type SiteConfigReq struct {
	g.Meta `path:"/site/config" method:"get" tags:"后台基础" summary:"获取配置"`
}
type SiteConfigRes struct {
	Version string `json:"version" dc:"系统版本"`
	WsAddr  string `json:"wsAddr" dc:"客户端websocket地址"`
	Domain  string `json:"domain" dc:"对外域名"`
}

// SitePingReq ping
type SitePingReq struct {
	g.Meta `path:"/site/ping" method:"get" tags:"后台基础" summary:"ping"`
}
type SitePingRes struct{}
