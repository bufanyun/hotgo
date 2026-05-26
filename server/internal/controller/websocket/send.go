// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package websocket

import (
	"context"
	"hotgo/api/websocket/base"
	"hotgo/internal/websocket"
	"hotgo/utility/simple"
)

// Send 通过http发送ws消息
var Send = send{}

type send struct{}

// SendToTag 发送标签消息
func (c *send) SendToTag(ctx context.Context, req *base.SendToTagReq) (res *base.SendToTagRes, err error) {
	simple.SafeGo(ctx, func(ctx context.Context) {
		websocket.SendToTag(req.Tag, &websocket.WResponse{
			Event: req.Response.Event,
			Data:  req.Response,
		})
	})
	return
}
