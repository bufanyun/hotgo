// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/utility/encrypt"
)

// serverLogin 心跳
func (client *Client) serverHeartbeat() {
	if !client.isLogin.Val() {
		return
	}

	ctx := gctx.New()
	if err := client.conn.Send(ctx, &ServerHeartbeatReq{}); err != nil {
		client.logger.Warningf(ctx, "client ServerHeartbeat Send err:%+v", err)
		return
	}
}

// serverLogin 服务登陆
func (client *Client) serverLogin() {
	auth := client.config.Auth

	// 无需登录
	if auth == nil {
		return
	}

	data := &ServerLoginReq{
		Name:      auth.Name,
		Extra:     auth.Extra,
		Group:     auth.Group,
		AppId:     auth.AppId,
		Timestamp: gtime.Timestamp(),
	}

	// 签名
	data.Sign = encrypt.Md5ToString(fmt.Sprintf("%v%v%v", data.AppId, data.Timestamp, auth.SecretKey))

	ctx := gctx.New()
	if err := client.conn.Send(ctx, data); err != nil {
		client.logger.Warningf(ctx, "client ServerLogin Send err:%+v", err)
		return
	}
}

// onResponseServerLogin 接收服务登陆响应结果
func (client *Client) onResponseServerLogin(ctx context.Context, req *ServerLoginRes) {
	if err := req.GetError(); err != nil {
		client.isLogin.Set(false)
		client.logger.Warningf(ctx, "onResponseServerLogin destroy, err:%v", err)
		client.Destroy()
		return
	}

	client.isLogin.Set(true)

	if client.config.LoginEvent != nil {
		client.config.LoginEvent()
	}
}

// onResponseServerHeartbeat 接收心跳响应结果
func (client *Client) onResponseServerHeartbeat(ctx context.Context, req *ServerHeartbeatRes) {
	if err := req.GetError(); err != nil {
		client.logger.Warningf(ctx, "onResponseServerHeartbeat err:%v", err)
		return
	}
	client.conn.Heartbeat = gtime.Timestamp()
}
