// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"reflect"
	"runtime"
)

// RouteHandler 路由处理器
type RouteHandler struct {
	Id     string        // 路由ID
	IsRPC  bool          // 是否支持rpc协议
	Func   reflect.Value // 路由处理方法
	Input  reflect.Value // 输入参数
	Output reflect.Value // 输出参数
}

// ParseRouteHandler 解析路由
func ParseRouteHandler(router interface{}, isRPC bool) (info *RouteHandler, err error) {
	funcName := runtime.FuncForPC(reflect.ValueOf(router).Pointer()).Name()
	funcType := reflect.ValueOf(router).Type()

	if funcType.NumIn() != 2 {
		err = gerror.Newf(ParseRouterErrInvalidParams, funcName)
		return
	}

	if funcType.In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
		err = gerror.Newf(ParseRouterErrInvalidFirstParam, funcName)
		return
	}

	inputType := funcType.In(1)
	if !(inputType.Kind() == reflect.Ptr && inputType.Elem().Kind() == reflect.Struct) {
		err = gerror.Newf(ParseRouterErrInvalidSecondParam, funcName)
		return
	}

	// The request struct should be named as `xxxReq`.
	if !gstr.HasSuffix(inputType.String(), `Req`) && !gstr.HasSuffix(inputType.String(), `Res`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming of the request: defined as "%s", but should be named with the "Req" or "Res" suffix, such as "XxxReq" or "XxxRes"`,
			inputType.String(),
		)
		return
	}

	info = &RouteHandler{
		Id:    gstr.SubStrFromREx(inputType.String(), `.`),
		IsRPC: isRPC,
		Func:  reflect.ValueOf(router),
		Input: reflect.New(inputType.Elem()),
	}

	if !isRPC {
		return
	}

	if funcType.NumOut() != 2 {
		err = gerror.Newf(ParseRouterRPCErrInvalidParams, funcName)
		return
	}
	outputType := funcType.Out(0)
	// The response struct should be named as `xxxRes`.
	if !gstr.HasSuffix(outputType.String(), `Res`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming for response: defined as "%s", but it should be named with "Res" suffix like "XxxRes"`,
			outputType.String(),
		)
		return
	}

	if !funcType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid handler: defined as "%s", but the last output parameter should be type of "error"`,
			reflect.TypeOf(funcType).String(),
		)
		return
	}

	info.Output = reflect.New(outputType.Elem())
	return
}
