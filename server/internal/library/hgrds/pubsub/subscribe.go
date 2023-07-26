// Package pubsub
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package pubsub

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/utility/simple"
	"sync"
)

type SubHandler func(ctx context.Context, message *gredis.Message)

type subscribeManager struct {
	mutex sync.RWMutex
	List  map[string]SubHandler
}

var subscribes = &subscribeManager{
	List: make(map[string]SubHandler),
}

// Subscribe 订阅消息
func Subscribe(channel string, hr SubHandler) (err error) {
	subscribes.mutex.Lock()
	defer subscribes.mutex.Unlock()

	if _, ok := subscribes.List[channel]; ok {
		err = gerror.Newf("repeat the subscribe:%v register", channel)
		return
	}
	subscribes.List[channel] = hr
	go doSubscribe(channel, hr)
	return
}

func doSubscribe(channel string, hr SubHandler) {
	ctx := gctx.New()
	conn, err := g.Redis().Conn(ctx)
	if err != nil {
		return
	}
	defer conn.Close(ctx)

	_, err = conn.Subscribe(ctx, channel)
	for {
		msg, err := conn.ReceiveMessage(ctx)
		if err != nil {
			g.Log().Warningf(ctx, "subscribe quit, err:%v", err)
			return
		}
		handleMessage(hr, msg)
	}
}

func handleMessage(hr SubHandler, message *gredis.Message) {
	simple.SafeGo(gctx.New(), func(ctx context.Context) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		hr(ctx, message)
	})
}
