// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Context 请求上下文结构
type Context struct {
	Module    string    // 应用模块 admin｜api｜home｜websocket
	AddonName string    // 插件名称 如果不是插件模块请求，可能为空
	User      *Identity // 上下文用户信息
	Response  *Response // 请求响应
	Data      g.Map     // 自定kv变量 业务模块根据需要设置，不固定
}

// Identity 通用身份模型
type Identity struct {
	Id       int64       `json:"id"              description:"用户ID"`
	Pid      int64       `json:"pid"             description:"上级ID"`
	DeptId   int64       `json:"deptId"          description:"部门ID"`
	DeptType string      `json:"deptType"        description:"部门类型"`
	RoleId   int64       `json:"roleId"          description:"角色ID"`
	RoleKey  string      `json:"roleKey"         description:"角色唯一标识符"`
	Username string      `json:"username"        description:"用户名"`
	RealName string      `json:"realName"        description:"姓名"`
	Avatar   string      `json:"avatar"          description:"头像"`
	Email    string      `json:"email"           description:"邮箱"`
	Mobile   string      `json:"mobile"          description:"手机号码"`
	App      string      `json:"app"             description:"登录应用"`
	LoginAt  *gtime.Time `json:"loginAt"         description:"登录时间"`
}
