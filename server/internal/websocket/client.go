// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package websocket

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/gorilla/websocket"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/location"
	"hotgo/internal/model"
	"runtime/debug"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 5 * 60
)

// 用户登录
type login struct {
	UserId int64
	Client *Client
}

// GetKey 读取客户端数据
func (l *login) GetKey() (key string) {
	key = GetUserKey(l.UserId)
	return
}

// Client 客户端连接
type Client struct {
	Addr          string          // 客户端地址
	ID            string          // 连接唯一标识
	Socket        *websocket.Conn // 用户连接
	Send          chan *WResponse // 待发送的数据
	SendClose     bool            // 发送是否关闭
	FirstTime     uint64          // 首次连接时间
	HeartbeatTime uint64          // 用户上次心跳时间
	Tags          garray.StrArray // 标签
	User          *model.Identity // 用户信息
	context       context.Context // Custom context for internal usage purpose.
	IP            string          // 客户端IP
	UserAgent     string          // 用户代理
}

// NewClient 初始化
func NewClient(r *ghttp.Request, socket *websocket.Conn, firstTime uint64) (client *Client) {
	client = &Client{
		Addr:          socket.RemoteAddr().String(),
		ID:            guid.S(),
		Socket:        socket,
		Send:          make(chan *WResponse, 100),
		SendClose:     false,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		User:          contexts.Get(r.Context()).User,
		IP:            location.GetClientIp(r),
		UserAgent:     r.UserAgent(),
	}
	return
}

// 读取客户端数据
func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(ctxManager, "client read err: %+v, stack:%+v, user:%+v", r, string(debug.Stack()), c.User)
		}
	}()

	defer func() {
		c.close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		// 处理消息
		handlerMsg(c, message)
	}
}

// 向客户端写数据
func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Warningf(ctxManager, "client write err: %+v, stack:%+v, user:%+v", r, string(debug.Stack()), c.User)
		}
	}()
	defer func() {
		clientManager.Unregister <- c
		c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 发送数据错误 关闭连接
				g.Log().Warningf(ctxManager, "client write message, user:%+v", c.User)
				return
			}
			c.Socket.WriteJSON(message)
		}
	}
}

// SendMsg 发送数据
func (c *Client) SendMsg(msg *WResponse) {
	if c == nil || c.SendClose {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			g.Log().Infof(ctxManager, "SendMsg err:%+v, stack:%+v", r, string(debug.Stack()))
		}
	}()
	c.Send <- msg
}

// Context is alias for function GetCtx.
func (c *Client) Context() context.Context {
	if c.context == nil {
		c.context = gctx.New()
	}
	return c.context
}

// Heartbeat 心跳更新
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime
	return
}

// IsHeartbeatTimeout 心跳是否超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

// 关闭客户端
func (c *Client) close() {
	if c.SendClose {
		return
	}
	c.SendClose = true
	if _, ok := <-c.Send; !ok {
		g.Log().Warningf(ctxManager, "close of closed channel, client.id:%v", c.ID)
	} else {
		// 关闭 chan
		close(c.Send)
	}
}

// Close 关闭指定客户端连接
func Close(client *Client) {
	client.close()
}

// SendSuccess 发送成功消息
func SendSuccess(client *Client, event string, data ...interface{}) {
	d := interface{}(nil)
	if len(data) > 0 {
		d = data[0]
	}
	client.SendMsg(&WResponse{
		Event:     event,
		Data:      d,
		Code:      consts.CodeOK,
		Timestamp: gtime.Now().Unix(),
	})
	before(client)
}

// SendError 发送错误消息
func SendError(client *Client, event string, err error) {
	client.SendMsg(&WResponse{
		Event:     event,
		Code:      consts.CodeNil,
		ErrorMsg:  err.Error(),
		Timestamp: gtime.Now().Unix(),
	})
	before(client)
}

// before
func before(client *Client) {
	client.Heartbeat(uint64(gtime.Now().Unix()))
}
