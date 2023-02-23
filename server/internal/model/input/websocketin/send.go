// Package websocketin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package websocketin

import (
	"hotgo/internal/websocket"
)

// SendToTagInput 发送标签消息
type SendToTagInput struct {
	Tag      string              `json:"tag" v:"required#tag不能为空" description:"标签"`
	Response websocket.WResponse `json:"response" v:"required#response不能为空"  description:"响应内容"`
}
