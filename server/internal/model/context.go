// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Context 请求上下文结构
type Context struct {
	Module     string    // 应用模块
	TakeUpTime int64     // 请求耗时 ms
	User       *Identity // 上下文用户信息
	Response   *Response // 请求响应
	Data       g.Map     // 自定KV变量，业务模块根据需要设置，不固定
}

// Identity 通用身份模型
type Identity struct {
	Id         int64  `json:"id"    description:"会员ID"`
	Username   string `json:"username"    description:"用户名"`
	RealName   string `json:"realname"    description:"昵称"`
	Avatar     string `json:"avatar"       description:"头像"`
	Email      string `json:"email"              description:"邮箱"`
	Mobile     string `json:"mobile"             description:"手机号码"`
	VisitCount uint   `json:"visit_count"         description:"访问次数"`
	LastTime   int    `json:"last_time"           description:"最后一次登录时间"`
	LastIp     string `json:"last_ip"             description:"最后一次登录ip"`
	Role       int64  `json:"role"               description:"角色ID"`
	RoleKey    string `json:"role_key"               description:"角色唯一标识符"`
	Exp        int64  `json:"exp"               description:"登录有效期截止时间戳"`
	Expires    int64  `json:"expires"               description:"登录有效期"`
	App        string `json:"app"               description:"登录应用"`
}
