// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/net/gtrace"
	"hotgo/internal/consts"
)

// initCtx 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改
func initCtx(ctx context.Context, model *Context) (newCtx context.Context, cancel context.CancelFunc) {
	if model.TraceID != "" {
		newCtx, _ = gtrace.WithTraceID(ctx, model.TraceID)
	} else {
		newCtx = ctx
	}
	newCtx = context.WithValue(newCtx, consts.ContextTCPKey, model)
	newCtx, cancel = context.WithCancel(newCtx)
	return
}

// SetCtx 设置上下文变量
func SetCtx(ctx context.Context, model *Context) {
	context.WithValue(ctx, consts.ContextTCPKey, model)
}

// GetCtx 获得上下文变量，如果没有设置，那么返回nil
func GetCtx(ctx context.Context) *Context {
	value := ctx.Value(consts.ContextTCPKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*Context); ok {
		return localCtx
	}
	return nil
}

// GetCtxConn .
func GetCtxConn(ctx context.Context) *gtcp.Conn {
	c := GetCtx(ctx)
	if c == nil {
		return nil
	}

	return c.Conn
}

// GetCtxAuth 认证元数据
func GetCtxAuth(ctx context.Context) *AuthMeta {
	c := GetCtx(ctx)
	if c == nil {
		return nil
	}

	return c.Auth
}
