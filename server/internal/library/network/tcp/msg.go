// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type ServerRes struct {
	Code    int    `json:"code" example:"2000"                     description:"状态码"`
	Message string `json:"message,omitempty" example:"操作成功"     description:"提示消息"`
}

// SetCode 设置状态码
func (i *ServerRes) SetCode(code ...int) {
	if len(code) > 0 {
		i.Code = code[0]
		return
	}

	// 默认值，转为成功的状态码
	if i.Code == 0 {
		i.Code = gcode.CodeOK.Code()
	}
}

// SetMessage 设置提示消息
func (i *ServerRes) SetMessage(msg ...string) {
	message := "操作成功"
	if len(msg) > 0 {
		message = msg[0]
	}
	i.Message = message
}

// SetError 设置响应中的错误
func (i *ServerRes) SetError(err error) {
	if err != nil {
		i.Code = gerror.Code(err).Code()
		i.Message = err.Error()
	}
}

// GetError 获取响应中的错误
func (i *ServerRes) GetError() (err error) {
	if i.Code != gcode.CodeOK.Code() {
		if i.Message == "" {
			i.Message = "操作失败"
		}
		err = gerror.NewCode(gcode.New(i.Code, i.Message, ""))
	}
	return
}

// ServerLoginReq 服务登录
type ServerLoginReq struct {
	Name      string `json:"name"             description:"客户端名称"` // 客户端名称，当同一个应用ID有多个客户端时请使用不同的名称区分。比如cron1,cron2
	Extra     g.Map  `json:"extra"            description:"自定义数据"` // 自定义数据，可以传递一些额外的自定义数据
	Group     string `json:"group"            description:"分组"`
	AppId     string `json:"appID"            description:"应用ID"`
	Timestamp int64  `json:"timestamp"        description:"服务器时间戳"`
	Sign      string `json:"sign"             description:"签名"`
}

// ServerLoginRes 响应服务登录
type ServerLoginRes struct {
	ServerRes
}

// ServerHeartbeatReq 心跳
type ServerHeartbeatReq struct {
}

// ServerHeartbeatRes 响应心跳
type ServerHeartbeatRes struct {
	ServerRes
}
