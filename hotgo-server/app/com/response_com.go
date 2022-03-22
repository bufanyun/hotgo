//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package com

import (
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"
)

// 统一响应
var Response = new(response)

type response struct{}

//
//  @Title  返回JSON数据并退出当前HTTP执行函数
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   r
//  @Param   code
//  @Param   message
//  @Param   data
//
func (component *response) JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	component.RJson(r, code, message, data...)
	r.Exit()
}

//
//  @Title  标准返回结果数据结构封装
//  @Description  返回固定数据结构的JSON
//  @Author  Ms <133814250@qq.com>
//  @Param   r
//  @Param   code 状态码(200:成功,302跳转，和http请求状态码一至)
//  @Param   message 请求结果信息
//  @Param   data 请求结果,根据不同接口返回结果的数据结构不同
//
func (component *response) RJson(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	Res := &model.Response{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().Unix(),
		ReqId:     Context.Get(r.Context()).ReqId,
	}

	// TODO  如果不是正常的返回，则将data转为error
	if consts.CodeOK == code {
		Res.Data = responseData
	} else {
		Res.Error = responseData
	}

	// TODO  清空响应
	r.Response.ClearBuffer()

	// TODO  写入响应
	if err := r.Response.WriteJson(Res); err != nil {
		g.Log().Error(r.Context(), "响应异常：", err)
	}

	// TODO  加入到上下文
	Context.SetResponse(r.Context(), Res)
}

//
//  @Title  返回成功JSON
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   isExit
//  @Param   r
//  @Param   message
//  @Param   data
//
func (component *response) SusJson(isExit bool, r *ghttp.Request, message string, data ...interface{}) {
	if isExit {
		component.JsonExit(r, consts.CodeOK, message, data...)
	}
	component.RJson(r, consts.CodeOK, message, data...)
}

//
//  @Title  返回失败JSON
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   isExit
//  @Param   r
//  @Param   message
//  @Param   data
//
func (component *response) FailJson(isExit bool, r *ghttp.Request, message string, data ...interface{}) {
	if isExit {
		component.JsonExit(r, consts.CodeNil, message, data...)
	}
	component.RJson(r, consts.CodeNil, message, data...)
}

//
//  @Title  重定向
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   r
//  @Param   location
//  @Param   code
//
func (component *response) Redirect(r *ghttp.Request, location string, code ...int) {
	r.Response.RedirectTo(location, code...)
}

func (component *response) Download(r *ghttp.Request, location string, code ...int) {
	r.Response.ServeFileDownload("test.txt")
}
