// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthMeta 认证元数据
type AuthMeta struct {
	Group     string      `json:"group"`
	Name      string      `json:"name"`
	AppId     string      `json:"appId"`
	SecretKey string      `json:"secretKey"`
	EndAt     *gtime.Time `json:"-"`
}

type Context struct {
	Conn    *gtcp.Conn `json:"conn"`
	Auth    *AuthMeta  `json:"auth"`    // 认证元数据
	TraceID string     `json:"traceID"` // 链路ID
}

// CallbackEvent 回调事件
type CallbackEvent func()
