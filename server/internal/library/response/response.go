// Package response
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package response

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model"
	"time"
)

// JsonExit 返回JSON数据并退出当前HTTP执行函数
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	RJson(r, code, message, data...)
	r.Exit()
}

// RJson 标准返回结果数据结构封装
// @Description: 返回固定数据结构的JSON
// @param r
// @param code 状态码(200:成功,302跳转，和http请求状态码一至)
// @param message 请求结果信息
// @param data 请求结果,根据不同接口返回结果的数据结构不同
func RJson(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	res := &model.Response{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().Unix(),
		TraceID:   gctx.CtxId(r.Context()),
	}

	// 如果不是正常的返回，则将data转为error
	if gcode.CodeOK.Code() == code {
		res.Data = responseData
	} else {
		res.Error = responseData
	}

	// 清空响应
	r.Response.ClearBuffer()

	// 写入响应
	r.Response.WriteJson(res)

	// 加入到上下文
	contexts.SetResponse(r.Context(), res)
}

// SusJson 返回成功JSON
func SusJson(isExit bool, r *ghttp.Request, message string, data ...interface{}) {
	if isExit {
		JsonExit(r, gcode.CodeOK.Code(), message, data...)
		return
	}
	RJson(r, gcode.CodeOK.Code(), message, data...)
}

// FailJson 返回失败JSON
func FailJson(isExit bool, r *ghttp.Request, message string, data ...interface{}) {
	if isExit {
		JsonExit(r, gcode.CodeNil.Code(), message, data...)
		return
	}
	RJson(r, gcode.CodeNil.Code(), message, data...)
}

// Redirect 重定向
func Redirect(r *ghttp.Request, location string, code ...int) {
	r.Response.RedirectTo(location, code...)
}

// Download 下载文件
func Download(r *ghttp.Request, location string, code ...int) {
	r.Response.ServeFileDownload("test.txt")
}

// RText 返回成功文本
func RText(r *ghttp.Request, message string) {
	// 清空响应
	r.Response.ClearBuffer()

	// 写入响应
	r.Response.Write(message)
}
