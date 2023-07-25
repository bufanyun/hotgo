// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"github.com/gogf/gf/v2/net/gtrace"
)

// Context tcp上下文
type Context struct {
	Conn *Conn
}

// initCtx 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改
func initCtx(ctx context.Context, model *Context) context.Context {
	return context.WithValue(ctx, ContextKey, model)
}

// GetCtx 获得上下文变量，如果没有设置，那么返回nil
func GetCtx(ctx context.Context) *Context {
	value := ctx.Value(ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*Context); ok {
		return localCtx
	}
	return nil
}

// ConnFromCtx retrieves and returns the Conn object from context.
func ConnFromCtx(ctx context.Context) *Conn {
	user := GetCtx(ctx)
	if user == nil {
		return nil
	}
	return user.Conn
}

// SetCtxTraceID 将自定义跟踪ID注入上下文以进行传播
func SetCtxTraceID(ctx context.Context, traceID string) (context.Context, error) {
	if len(traceID) > 0 {
		return gtrace.WithTraceID(ctx, traceID)
	}
	return ctx, nil
}
