// Package tcpserver
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpserver

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/api/servmsg"
	"hotgo/internal/consts"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/servmsgin"
)

// OnExampleHello 一个tcp请求例子
func (s *sTCPServer) OnExampleHello(ctx context.Context, req *servmsg.ExampleHelloReq) {
	var (
		conn = tcp.ConnFromCtx(ctx)
		res  = new(servmsg.ExampleHelloRes)
		data = new(servmsgin.ExampleHelloModel)
	)

	if conn == nil {
		g.Log().Warningf(ctx, "conn is nil.")
		return
	}

	if conn.Auth == nil {
		res.SetError(gerror.New("连接未认证，请确认已登录成功！"))
		_ = conn.Send(ctx, res)
		return
	}

	data.Desc = fmt.Sprintf("Hello %v, 你的APPID：%v，当前HotGo版本：%v，你成功请求了`servmsg.ExampleHelloReq`接口！", req.Name, conn.Auth.AppId, consts.VersionApp)
	data.Timestamp = gtime.Now()
	res.Data = data
	_ = conn.Send(ctx, res)
}

// OnExampleRPCHello 一个rpc请求例子
func (s *sTCPServer) OnExampleRPCHello(ctx context.Context, req *servmsg.ExampleRPCHelloReq) (res *servmsg.ExampleRPCHelloRes, err error) {
	var data = new(servmsgin.ExampleHelloModel)
	data.Desc = fmt.Sprintf("Hello %v, 当前HotGo版本：%v，你成功请求了`servmsg.ExampleRPCHelloReq`接口！", req.Name, consts.VersionApp)
	data.Timestamp = gtime.Now()

	res = new(servmsg.ExampleRPCHelloRes)
	res.Data = data
	return
}
