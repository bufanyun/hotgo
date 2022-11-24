// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package websocket

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"runtime/debug"
)

// handlerMsg 处理消息
func handlerMsg(client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(ctxManager, "handlerMsg recover, err:%+v, stack:%+v", r, string(debug.Stack()))
		}
	}()
	request := &WRequest{}
	err := gconv.Struct(message, request)
	if err != nil {
		g.Log().Warningf(ctxManager, "handlerMsg 数据解析失败,err:%+v, message:%+v", err, string(message))
		return
	}

	if request.Event == "" {
		g.Log().Warning(ctxManager, "handlerMsg request.Event is null")
		return
	}

	//g.Log().Infof(ctxManager, "websocket handlerMsg:%+v", request)

	fun, ok := routers[request.Event]
	if !ok {
		g.Log().Warningf(ctxManager, "handlerMsg function id %v: not registered", request.Event)
		return
	}
	fun(client, request)
}

// RegisterMsg 注册消息
func RegisterMsg(handlers EventHandlers) {
	for id, f := range handlers {
		if _, ok := routers[id]; ok {
			g.Log().Fatalf(ctxManager, "RegisterMsg function id %v: already registered", id)
			return
		}
		routers[id] = f
	}
}
