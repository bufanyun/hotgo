// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmeta"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/response"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
)

// ResponseHandler HTTP响应预处理
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	contentType := getContentType(r)
	// 已存在响应
	if contentType != consts.HTTPContentTypeStream && r.Response.BufferLength() > 0 && contexts.Get(r.Context()).Response != nil {
		return
	}

	switch contentType {
	case consts.HTTPContentTypeHtml:
		s.responseHtml(r)
		return
	case consts.HTTPContentTypeXml:
		s.responseXml(r)
		return
	case consts.HTTPContentTypeStream:
	default:
		responseJson(r)
	}
}

// responseHtml html模板响应
func (s *sMiddleware) responseHtml(r *ghttp.Request) {
	code, message, resp := parseResponse(r)
	if code == gcode.CodeOK.Code() {
		return
	}

	r.Response.ClearBuffer()
	_ = r.Response.WriteTplContent(simple.DefaultErrorTplContent(r.Context()), g.Map{"code": code, "message": message, "stack": resp})
}

// responseXml xml响应
func (s *sMiddleware) responseXml(r *ghttp.Request) {
	code, message, data := parseResponse(r)
	response.RXml(r, code, message, data)
}

// responseJson json响应
func responseJson(r *ghttp.Request) {
	code, message, data := parseResponse(r)
	response.RJson(r, code, message, data)
}

// parseResponse 解析响应数据
func parseResponse(r *ghttp.Request) (code int, message string, resp interface{}) {
	ctx := r.Context()
	err := r.GetError()
	if err == nil {
		return gcode.CodeOK.Code(), "操作成功", r.GetHandlerResponse()
	}

	// 是否输出错误堆栈到页面
	if g.Cfg().MustGet(ctx, "hotgo.debug", true).Bool() {
		message = gerror.Current(err).Error()
		if getContentType(r) == consts.HTTPContentTypeHtml {
			resp = charset.SerializeStack(err)
		} else {
			resp = charset.ParseErrStack(err)
		}
	} else {
		message = consts.ErrorMessage(gerror.Current(err))
	}

	code = gerror.Code(err).Code()

	// 记录异常日志
	if code == gcode.CodeNil.Code() {
		g.Log().Stdout(false).Printf(ctx, "exception:%v", err)
	} else {
		g.Log().Errorf(ctx, "exception:%v", err)
	}
	return
}

func getContentType(r *ghttp.Request) (contentType string) {
	contentType = r.Response.Header().Get("Content-Type")
	if contentType != "" {
		return
	}

	mime := gmeta.Get(r.GetHandlerResponse(), "mime").String()
	if mime == "" {
		contentType = consts.HTTPContentTypeJson
	} else {
		contentType = mime
	}
	return
}
