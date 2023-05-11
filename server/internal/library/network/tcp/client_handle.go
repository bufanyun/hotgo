package tcp

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/msgin"
)

// serverLogin 心跳
func (client *Client) serverHeartbeat() {
	ctx := gctx.New()
	if err := client.Send(ctx, &msgin.ServerHeartbeat{}); err != nil {
		client.Logger.Debugf(ctx, "client WriteMsg ServerHeartbeat err:%+v", err)
		return
	}
}

// serverLogin 服务登陆
func (client *Client) serverLogin() {
	data := &msgin.ServerLogin{
		Group: client.auth.Group,
		Name:  client.auth.Name,
	}

	ctx := gctx.New()
	if err := client.Send(ctx, data); err != nil {
		client.Logger.Debugf(ctx, "client WriteMsg ServerLogin err:%+v", err)
		return
	}
}

func (client *Client) onResponseServerLogin(ctx context.Context, args ...interface{}) {
	var in *msgin.ResponseServerLogin
	if err := gconv.Scan(args[0], &in); err != nil {
		client.Logger.Infof(ctx, "onResponseServerLogin message Scan failed:%+v, args:%+v", err, args[0])
		return
	}

	if in.Code != consts.TCPMsgCodeSuccess {
		client.IsLogin = false
		client.Logger.Warningf(ctx, "onResponseServerLogin quit err:%v", in.Message)
		client.Destroy()
		return
	}

	client.IsLogin = true

	if client.loginEvent != nil {
		client.loginEvent()
	}
}

func (client *Client) onResponseServerHeartbeat(ctx context.Context, args ...interface{}) {
	var in *msgin.ResponseServerHeartbeat
	if err := gconv.Scan(args[0], &in); err != nil {
		client.Logger.Infof(ctx, "onResponseServerHeartbeat message Scan failed:%+v, args:%+v", err, args)
		return
	}

	if in.Code != consts.TCPMsgCodeSuccess {
		client.Logger.Warningf(ctx, "onResponseServerHeartbeat err:%v", in.Message)
		return
	}

	client.heartbeat = gtime.Timestamp()
}
