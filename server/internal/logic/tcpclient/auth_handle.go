// Package tcpclient
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpclient

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/servmsg"
)

// OnResponseAuthSummary 响应授权信息
func (s *sAuthClient) OnResponseAuthSummary(ctx context.Context, req *servmsg.AuthSummaryRes) {
	if err := req.GetError(); err != nil {
		g.Log().Warningf(ctx, "OnResponseAuthSummary GetError:%+v", err)
		return
	}

	// 拿到授权的数据，可以是一些动态的功能、路由、权限控制等
	s.summary = req.Data
}

// OnResponseExampleHello 一个tcp请求例子
func (s *sAuthClient) OnResponseExampleHello(ctx context.Context, req *servmsg.ExampleHelloRes) {
	if err := req.GetError(); err != nil {
		g.Log().Warningf(ctx, "OnResponseExampleHello GetError:%+v", err)
		return
	}

	g.Log().Infof(ctx, "OnResponseExampleHello data:%+v", req.Data)
}

// testExample 测试例子
func (s *sAuthClient) testExample(ctx context.Context) {
	// 发起tcp请求
	// 异步执行，服务端返回消息后会转到`OnResponseExampleHello`中
	_ = s.client.Send(ctx, &servmsg.ExampleHelloReq{
		Name: "Tom",
	})

	// 发起rpc请求
	// 同步执行，阻塞等待服务端返回消息
	var req = &servmsg.ExampleRPCHelloReq{
		Name: "Tony",
	}
	var res *servmsg.ExampleRPCHelloRes
	if err := s.client.RequestScan(ctx, req, &res); err != nil {
		g.Log().Warningf(ctx, "client.Request ExampleRPCHelloReq err:%+v", err)
		return
	}
	g.Log().Infof(ctx, "ExampleRPCHelloRes data:%+v", res.Data)
}
