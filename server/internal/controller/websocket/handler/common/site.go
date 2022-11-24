// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/websocket"
)

var (
	Site = cSite{}
)

type cSite struct{}

func (c *cSite) Join(client *websocket.Client, req *websocket.WRequest) {
	name := gconv.String(req.Data["id"])

	if !client.Tags.Contains(name) {
		client.Tags.Append(name)
	}

	websocket.SendSuccess(client, req.Event, client.Tags.Slice())
}

func (c *cSite) Quit(client *websocket.Client, req *websocket.WRequest) {
	name := gconv.String(req.Data["id"])
	if client.Tags.Contains(name) {
		client.Tags.RemoveValue(name)
	}
	websocket.SendSuccess(client, req.Event, client.Tags.Slice())
}

func (c *cSite) Ping(client *websocket.Client, req *websocket.WRequest) {
	websocket.SendSuccess(client, req.Event)
}
