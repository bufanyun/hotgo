// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import (
	"fmt"
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
)

func init() {
	dict.RegisterEnums("HTTPMethod", "HTTP请求方式选项", HTTPMethodOptions)
	dict.RegisterEnums("HTTPHandlerTime", "HTTP处理耗时选项", HTTPHandlerTimeOptions)
	dict.RegisterEnums("HTTPApiCode", "HTTP接口状态码选项", HTTPApiCodeOptions)
}

const (
	HTTPContentTypeXml         = "text/xml"
	HTTPContentTypeHtml        = "text/html"
	HTTPContentTypeStream      = "text/event-stream"
	HTTPContentTypeJson        = "application/json"
	HTTPContentTypeOctetStream = "application/octet-stream"
)

// HTTPMethodOptions HTTP请求方式选项
var HTTPMethodOptions = []*model.Option{
	dict.GenSuccessOption(http.MethodGet, "GET"),
	dict.GenInfoOption(http.MethodPost, "POST"),
	dict.GenSuccessOption(http.MethodPut, "PUT"),
	dict.GenInfoOption(http.MethodDelete, "DELETE"),
}

const (
	HTTPHandlerTime50          = "< 50"
	HTTPHandlerTime200         = "< 200"
	HTTPHandlerTime200To500    = "BETWEEN 200 AND 500"
	HTTPHandlerTime500To1000   = "BETWEEN 500 AND 1000"
	HTTPHandlerTime1000To10000 = "BETWEEN 1000 AND 10000"
	HTTPHandlerTime10000UP     = "> 10000"
)

// HTTPHandlerTimeOptions HTTP处理耗时选项
var HTTPHandlerTimeOptions = []*model.Option{
	dict.GenSuccessOption(HTTPHandlerTime50, "50ms以内"),
	dict.GenInfoOption(HTTPHandlerTime200, "200ms以内"),
	dict.GenSuccessOption(HTTPHandlerTime200To500, "200~500ms"),
	dict.GenSuccessOption(HTTPHandlerTime500To1000, "500~1000ms"),
	dict.GenInfoOption(HTTPHandlerTime1000To10000, "1000~10000ms"),
	dict.GenInfoOption(HTTPHandlerTime10000UP, "10000ms以上"),
}

// HTTPApiCodeOptions HTTP接口状态码选项
var HTTPApiCodeOptions = []*model.Option{
	dict.GenSuccessOption(gcode.CodeOK.Code(), fmt.Sprintf("%v %v", gcode.CodeOK.Code(), "成功")),
	dict.GenWarningOption(gcode.CodeNil.Code(), fmt.Sprintf("%v %v", gcode.CodeNil.Code(), "失败")),
	dict.GenWarningOption(gcode.CodeSecurityReason.Code(), fmt.Sprintf("%v %v", gcode.CodeSecurityReason.Code(), "无访问权限")),
}
