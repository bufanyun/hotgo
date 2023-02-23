// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import "github.com/gogf/gf/v2/frame/g"

// SendTestSmsReq 发送测试短信
type SendTestSmsReq struct {
	Event  string `json:"event" v:"required#事件模板不能为空" dc:"事件模板"`
	Mobile string `json:"mobile" v:"required#接收手机号不能为空" dc:"接收手机号"`
	Code   string `json:"code" v:"required#接收验证码不能为空" dc:"接收验证码"`
	g.Meta `path:"/sms/sendTest" tags:"短信" method:"post" summary:"发送测试短信"`
}
type SendTestSmsRes struct {
}

// SendBindSmsReq 发送换绑短信
type SendBindSmsReq struct {
	g.Meta `path:"/sms/sendBind" tags:"短信" method:"post" summary:"发送换绑短信"`
}
type SendBindSmsRes struct {
}
