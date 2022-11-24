// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package websocket

import (
	"context"
	"hotgo/internal/model/input/websocketin"
	"hotgo/internal/websocket"
)

// Send 通过http发送ws消息
var Send = send{}

type send struct{}

// ToTag 发送标签消息
func (c *send) ToTag(ctx context.Context, req *websocketin.SendToTagReq) (res *websocketin.SendToTagRes, err error) {

	go websocket.SendToTag(req.Tag, &websocket.WResponse{
		Event: req.Response.Event,
		Data:  req.Response,
	})
	return
}
