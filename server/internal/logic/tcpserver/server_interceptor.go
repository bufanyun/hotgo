// Package tcpserver
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpserver

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/library/network/tcp"
	"hotgo/utility/validate"
)

// 免登录路由
var noLoginRouter = map[string]struct{}{
	"ServerLoginReq": {}, // 服务登录
}

// 免验证路由
var noVerifyRouter = map[string]struct{}{
	"ServerHeartbeatReq": {}, // 心跳
}

func (s *sTCPServer) isNoLoginRouter(router string) bool {
	_, ok := noLoginRouter[router]
	return ok
}

func (s *sTCPServer) isNoVerifyRouter(router string) bool {
	_, ok := noVerifyRouter[router]
	return ok
}

// DefaultInterceptor 默认拦截器
func (s *sTCPServer) DefaultInterceptor(ctx context.Context, msg *tcp.Message) (err error) {
	conn := tcp.ConnFromCtx(ctx)
	//g.Log().Debugf(ctx, "DefaultInterceptor msg:%+v, conn:%+v", msg, gjson.New(conn).String())

	// 免登录
	if s.isNoLoginRouter(msg.Router) {
		return
	}

	if conn.Auth == nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "未进行登录认证，请先登录")
		return
	}

	// 检查授权有效期
	if conn.Auth.EndAt.Before(gtime.Now()) {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "授权已过期")
		return
	}

	// 免验证
	if s.isNoVerifyRouter(msg.Router) {
		return
	}

	// 验证路由权限
	if len(conn.Auth.Routes) > 0 && !validate.InSlice(conn.Auth.Routes, msg.Router) {
		err = gerror.NewCodef(gcode.CodeNotAuthorized, "没有授权路由访问权限：%v", msg.Router)
		return
	}
	return
}

// PreFilterInterceptor 预处理
func (s *sTCPServer) PreFilterInterceptor(ctx context.Context, msg *tcp.Message) (err error) {
	//g.Log().Debug(ctx, "PreFilterInterceptor...")
	return
}
