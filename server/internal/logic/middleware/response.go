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
)

// ResponseHandler HTTP响应预处理
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 模板页面响应
	if "text/html" == r.Response.Header().Get("Content-Type") {
		r.Middleware.Next()
		return
	}

	// 支付通知响应
	if _, ok := s.PayNotifyRoutes[r.Router.Uri]; ok {
		s.responsePayNotify(r)
		return
	}

	responseJson(r)
}

// rTemplate 支付通知响应
func (s *sMiddleware) responsePayNotify(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		err  error
		data *payin.PayNotifyModel
	)

	// 异常
	if err = r.GetError(); err != nil {
		g.Log("exception").Error(ctx, err)
		r.Response.ClearBuffer()
		r.Response.WriteStatus(500, err.Error())
		return
	}

	if err = gconv.Scan(r.GetHandlerResponse(), &data); err != nil || data == nil {
		g.Log("exception").Errorf(ctx, "middleware.responsePayNotify Scan err:%+v, data:%+v", err, data)
		r.Response.ClearBuffer()
		r.Response.WriteStatus(500, err.Error())
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
		r.Response.WriteStatus(500, err.Error())
	}
}

func responseJson(r *ghttp.Request) {
	var (
		ctx         = r.Context()
		comResponse = contexts.Get(ctx).Response
		code        = gcode.CodeOK.Code()
		message     = "操作成功"
		data        interface{}
		err         error
	)

	// 已存在响应内容，且是comResponse返回的时，中断运行
	if r.Response.BufferLength() > 0 && comResponse != nil {
		return
	}

	if err = r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log().Warningf(ctx, "exception:%v", err)

		code = gerror.Code(err).Code()

		// 是否输出错误到页面
		if g.Cfg().MustGet(ctx, "hotgo.debug", true).Bool() {
			message = err.Error()
			data = charset.ParseErrStack(err)
		} else {
			message = consts.ErrorMessage(err)
		}
	} else {
		data = r.GetHandlerResponse()
	}

	// 返回固定的友好信息
	response.RJson(r, code, message, data)
}
