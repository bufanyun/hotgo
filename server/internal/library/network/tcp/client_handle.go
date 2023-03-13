package tcp

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model/input/msgin"
)

// serverLogin 心跳
func (client *Client) serverHeartbeat() {
	if err := client.Write(&msgin.ServerHeartbeat{}); err != nil {
		client.Logger.Debugf(client.Ctx, "client WriteMsg ServerHeartbeat err:%+v", err)
		return
	}
}

// serverLogin 服务登陆
func (client *Client) serverLogin() {
	data := &msgin.ServerLogin{
		Group: client.auth.Group,
		Name:  client.auth.Name,
	}

	if err := client.Write(data); err != nil {
		client.Logger.Debugf(client.Ctx, "client WriteMsg ServerLogin err:%+v", err)
		return
	}

	if client.loginEvent != nil {
		client.loginEvent()
	}
}

func (client *Client) onResponseServerLogin(args ...interface{}) {
	var in *msgin.ResponseServerLogin
	if err := gconv.Scan(args[0], &in); err != nil {
		client.Logger.Infof(client.Ctx, "onResponseServerLogin message Scan failed:%+v, args:%+v", err, args[0])
		return
	}
	client.Logger.Infof(client.Ctx, "onResponseServerLogin in:%+v", *in)

	if in.Code != gcode.CodeOK.Code() {
		client.IsLogin = false
		client.Logger.Warningf(client.Ctx, "onResponseServerLogin quit err:%v", in.Message)
		client.Destroy()
		return
	}
	client.IsLogin = true
}

func (client *Client) onResponseServerHeartbeat(args ...interface{}) {
	var in *msgin.ResponseServerHeartbeat
	if err := gconv.Scan(args[0], &in); err != nil {
		client.Logger.Infof(client.Ctx, "onResponseServerHeartbeat message Scan failed:%+v, args:%+v", err, args)
		return
	}

	client.heartbeat = gtime.Timestamp()
	client.Logger.Infof(client.Ctx, "onResponseServerHeartbeat in:%+v", *in)
}
