// Package websocketin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package websocketin

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/websocket"
)

// SendToTagReq 发送标签消息
type SendToTagReq struct {
	g.Meta   `path:"/send/toTag" method:"post" tags:"WebSocket" summary:"发送标签消息"`
	Tag      string              `json:"tag" v:"required#tag不能为空" description:"标签"`
	Response websocket.WResponse `json:"response" v:"required#response不能为空"  description:"响应内容"`
}

type SendToTagRes struct {
}
