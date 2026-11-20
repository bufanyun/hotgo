// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package websocket

import (
	"context"
	"hotgo/internal/consts"
	"runtime/debug"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// handlerMsg 处理消息
func handlerMsg(client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(mctx, "handlerMsg recover, err:%+v, stack:%+v", r, string(debug.Stack()))
		}
	}()

	var request *WRequest
	if err := gconv.Struct(message, &request); err != nil {
		g.Log().Warningf(mctx, "handlerMsg 数据解析失败,err:%+v, message:%+v", err, string(message))
		return
	}

	if request.Event == "" {
		g.Log().Warning(mctx, "handlerMsg request.Event is null")
		return
	}

	fun, ok := routers[request.Event]
	if !ok {
		g.Log().Warningf(mctx, "handlerMsg function id %v: not registered", request.Event)
		return
	}

	// 前台(api) token 也能连主 WS，admin/ 事件只给后台 token
	if strings.HasPrefix(request.Event, "admin/") && (client.User == nil || client.User.App != consts.AppAdmin) {
		SendError(client, request.Event, gerror.New("没有访问权限"))
		return
	}

	err := msgGo.AddWithRecover(mctx,
		func(ctx context.Context) {
			fun(client, request)
		},
		func(ctx context.Context, err error) {
			g.Log().Warningf(mctx, "handlerMsg msgGo exec err:%+v", err)
		},
	)

	if err != nil {
		g.Log().Warningf(mctx, "handlerMsg msgGo Add err:%+v", err)
		return
	}
}

// RegisterMsg 注册消息
func RegisterMsg(handlers EventHandlers) {
	for id, f := range handlers {
		if _, ok := routers[id]; ok {
			g.Log().Fatalf(mctx, "RegisterMsg function id %v: already registered", id)
			return
		}
		routers[id] = f
	}
}
