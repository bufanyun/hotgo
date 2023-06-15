// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/response"
	"hotgo/internal/model/input/payin"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
	"net/http"
)

// ResponseHandler HTTP响应预处理
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 已存在响应
	if r.Response.BufferLength() > 0 && contexts.Get(r.Context()).Response != nil {
		return
	}

	// html模板响应
	if r.Response.Header().Get("Content-Type") == "text/html" {
		s.responseHtml(r)
		return
	}

	// 支付通知响应
	if _, ok := s.PayNotifyRoutes[r.Router.Uri]; ok {
		s.responsePayNotify(r)
		return
	}

	// 默认json响应
	responseJson(r)
}

// responseHtml html模板响应
func (s *sMiddleware) responseHtml(r *ghttp.Request) {
	code, message, resp := parseResponse(r)
	if code == gcode.CodeOK.Code() {
		return
	}

	r.Response.ClearBuffer()
	_ = r.Response.WriteTplContent(simple.DefaultErrorTplContent(r.Context()), g.Map{
		"code":    code,
		"message": message,
		"stack":   resp,
	})
	return
}

// responsePayNotify 支付通知响应
func (s *sMiddleware) responsePayNotify(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		err  error
		data *payin.PayNotifyModel
	)

	code, message, resp := parseResponse(r)
	if code != gcode.CodeOK.Code() {
		response.RJson(r, code, message, data)
		return
	}

	if err = gconv.Scan(resp, &data); err != nil || data == nil {
		g.Log("exception").Errorf(ctx, "middleware.responsePayNotify Scan err:%+v, data:%+v", err, data)
		r.Response.ClearBuffer()
		r.Response.WriteStatus(http.StatusInternalServerError, err.Error())
		return
	}

	switch data.PayType {
	case consts.PayTypeAliPay:
		response.RText(r, data.Message)

	case consts.PayTypeWxPay:
		r.Response.ClearBuffer()
		r.Response.WriteJson(fmt.Sprintf(`{"code": "%v","message": "%v"}`, data.Code, data.Message))

	case consts.PayTypeQQPay:
		r.Response.ClearBuffer()
		r.Response.Write(`<?xml version="1.0" encoding="UTF-8"?>`)
		r.Response.WriteXml(g.Map{
			"return_code": data.Message,
		})

	default:
		err = gerror.Newf("无效的支付方式，这可能是没有配置通知回调响应方式导致的：%+v", data)
		g.Log("exception").Error(ctx, err)
		r.Response.ClearBuffer()
		r.Response.WriteStatus(http.StatusInternalServerError, err.Error())
	}
}

// responseJson json响应
func responseJson(r *ghttp.Request) {
	code, message, data := parseResponse(r)
	response.RJson(r, code, message, data)
}

// parseResponse 解析响应数据
func parseResponse(r *ghttp.Request) (code int, message string, resp interface{}) {
	var (
		ctx = r.Context()
		err = r.GetError()
	)

	if err == nil {
		return gcode.CodeOK.Code(), "操作成功", r.GetHandlerResponse()
	}

	// 是否输出错误堆栈到页面
	if g.Cfg().MustGet(ctx, "hotgo.debug", true).Bool() {
		message = gerror.Current(err).Error()
		resp = charset.ParseErrStack(err)
	} else {
		message = consts.ErrorMessage(gerror.Current(err))
	}

	// 解析错误状态码
	code = gerror.Code(err).Code()

	// 记录异常日志
	if code == gcode.CodeNil.Code() {
		g.Log().Stdout(false).Printf(ctx, "exception:%v", err)
	} else {
		g.Log().Errorf(ctx, "exception:%v", err)
	}
	return
}
