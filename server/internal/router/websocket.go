// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	controller "hotgo/internal/controller/websocket"
	"hotgo/internal/controller/websocket/handler/admin"
	"hotgo/internal/controller/websocket/handler/common"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
)

// WebSocket ws路由配置
func WebSocket(ctx context.Context, group *ghttp.RouterGroup) {
	prefix := g.Cfg().MustGet(ctx, "router.ws.prefix", "/socket")
	group.Group(prefix.String(), func(group *ghttp.RouterGroup) {
		group.Middleware()
		group.Bind(
			controller.Send, // 通过http发送ws消息。方便测试没有放权限中间件，实际使用时请自行调整
		)

		// ws连接中间件
		group.Middleware(service.Middleware().WebSocketAuth)

		// ws
		group.GET("/", websocket.WsPage)
	})

	// 启动websocket监听
	websocket.Start(ctx)

	// 注册消息路由
	websocket.RegisterMsg(websocket.EventHandlers{
		"ping":                  common.Site.Ping,      // 心跳
		"join":                  common.Site.Join,      // 加入组
		"quit":                  common.Site.Quit,      // 退出组
		"admin/monitor/trends":  admin.Monitor.Trends,  // 后台监控，动态数据
		"admin/monitor/runInfo": admin.Monitor.RunInfo, // 后台监控，运行信息
	})

}
