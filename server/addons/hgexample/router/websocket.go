// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/addons/hgexample/controller/websocket"
	"hotgo/addons/hgexample/global"
	"hotgo/internal/consts"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"
	ws "hotgo/internal/websocket"
)

// WebSocket ws路由配置
func WebSocket(ctx context.Context, group *ghttp.RouterGroup) {
	prefix := addons.RouterPrefix(ctx, consts.AppWebSocket, global.GetSkeleton().Name)
	group.Group(prefix, func(group *ghttp.RouterGroup) {
		group.Bind(
			// 无需验证的路由
			websocket.Index,
		)
		// ws连接中间件
		group.Middleware(service.Middleware().WebSocketToken)
		group.Bind(
		// 需要验证的路由
		// ..
		)
	})

	// 注册消息路由
	ws.RegisterMsg(ws.EventHandlers{
		// ...
	})

}
