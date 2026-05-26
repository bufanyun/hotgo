// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gerror"
	"sync"
	"time"
)

// RPC .
type RPC struct {
	mutex     sync.Mutex
	callbacks map[string]RPCResponseFunc
	task      RoutineTask
}

// RPCResponse 响应结构
type RPCResponse struct {
	res interface{}
	err error
}

type RPCResponseFunc func(resp interface{}, err error)

// NewRPC 初始化RPC
func NewRPC(task RoutineTask) *RPC {
	return &RPC{
		task:      task,
		callbacks: make(map[string]RPCResponseFunc),
	}
}

// Request 发起RPC请求
func (r *RPC) Request(ctx context.Context, msgId string, send func()) (res interface{}, err error) {
	resCh := make(chan RPCResponse, 1)
	isClose := gtype.NewBool(false)

	defer func() {
		isClose.Set(true)
		close(resCh)
		r.popCallback(msgId)
	}()

	r.mutex.Lock()
	r.callbacks[msgId] = func(res interface{}, err error) {
		if !isClose.Val() {
			resCh <- RPCResponse{res: res, err: err}
		}
	}
	r.mutex.Unlock()

	r.task(ctx, send)

	select {
	case <-time.After(time.Second * RPCTimeout):
		err = gerror.New("RPC response timeout")
		return
	case got := <-resCh:
		return got.res, got.err
	}
}

// Response RPC消息响应
func (r *RPC) Response(ctx context.Context, msg *Message) bool {
	if len(msg.MsgId) == 0 {
		return false
	}

	f, ok := r.popCallback(msg.MsgId)
	if !ok {
		return false
	}

	var msgError error
	if len(msg.Error) > 0 {
		msgError = gerror.New(msg.Error)
	}

	r.task(ctx, func() {
		f(msg.Data, msgError)
	})
	return true
}

// popCallback 弹出回调
func (r *RPC) popCallback(msgId string) (RPCResponseFunc, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	call, ok := r.callbacks[msgId]
	if ok {
		delete(r.callbacks, msgId)
	}
	return call, ok
}
