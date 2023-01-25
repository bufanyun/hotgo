// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/response"
	"hotgo/utility/charset"
)

// ResponseHandler HTTP响应预处理
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		ctx         = r.Context()
		comResponse = contexts.Get(ctx).Response
		code        = gcode.CodeOK.Code()
		message     = "操作成功"
		data        interface{}
		err         error
	)

	// 模板页面响应
	if "text/html" == r.Response.Header().Get("Content-Type") {
		r.Middleware.Next()
		return
	}

	// 已存在响应内容，且是comResponse返回的时，中断运行
	if r.Response.BufferLength() > 0 && comResponse != nil {
		return
	}

	if err = r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log().Warningf(ctx, "exception:%v", err)

		code = gerror.Code(err).Code()
		message = err.Error()

		// 是否输出错误到页面
		if g.Cfg().MustGet(ctx, "hotgo.debug", true).Bool() {
			data = charset.ParseErrStack(err)
		}
	} else {
		data = r.GetHandlerResponse()
	}

	// 返回固定的友好信息
	response.RJson(r, code, message, data)
}
