// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
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

	if err := r.GetError(); err != nil {
		g.Log().Print(ctx, err)
		// 记录到自定义错误日志文件
		//g.Log("exception").Error(err)
		////返回固定的友好信息
		//r.Response.ClearBuffer()
		//r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
	}

	// 已存在响应内容，且是comResponse返回的时，中断运行
	if r.Response.BufferLength() > 0 && comResponse != nil {
		return
	}

	if err = r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Print(r.Context(), "exception:", err)

		code = consts.CodeInternalError
		message = "服务器居然开小差了，请稍后再试吧！"

		// 是否输出错误到页面
		if debug, _ := g.Cfg().Get(ctx, "hotgo.debug", true); debug.Bool() {
			data = charset.GetStack(err)
			message = err.Error()
		}

		//} else if data, err = r.GetHandlerResponse(); err != nil {
		//	errCode := gerror.Code(err)
		//	if errCode == gcode.CodeNil {
		//		errCode = gcode.CodeInternalError
		//	}
		//	code = errCode.Code()
		//	message = err.Error()
		//}

	} else {
		data = r.GetHandlerResponse()
	}

	// 返回固定的友好信息
	response.RJson(r, code, message, data)
}
