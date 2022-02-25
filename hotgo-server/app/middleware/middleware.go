//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package middleware

import (
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
	"go.opentelemetry.io/otel/trace"
)

type (
	// sMiddleware is service struct of module Middleware.
	sMiddleware struct{}
)

var (
	// insMiddleware is the instance of service Middleware.
	insMiddleware = sMiddleware{}
)

// Middleware returns the interface of Middleware service.
func Instance() *sMiddleware {
	return &insMiddleware
}

//
//  @Title  初始化请求上下文
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   r
//
func (s *sMiddleware) Ctx(r *ghttp.Request) {

	spanCtx := trace.SpanContextFromContext(r.Context())

	reqId := guid.S(grand.B(64))
	if traceId := spanCtx.TraceID(); traceId.IsValid() {
		reqId = traceId.String()
	}

	customCtx := &model.Context{
		Data:    make(g.Map),
		Request: r,
		ReqId:   reqId,
	}

	com.Context.Init(r, customCtx)

	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
