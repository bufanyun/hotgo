// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import "github.com/gogf/gf/v2/text/gstr"

// 错误解释
const (
	ErrorORM         = "sql执行异常"
	ErrorNotData     = "数据不存在"
	ErrorRotaPointer = "指针转换异常"
)

// 需要隐藏真实错误的Wrap，开启访问日志后仍然会将真实错误记录
var concealErrorSlice = []string{ErrorORM, ErrorRotaPointer}

// ErrorMessage 错误显示信息，非debug模式有效
func ErrorMessage(err error) (message string) {
	if err == nil {
		return "操作失败！~"
	}

	message = err.Error()
	for _, e := range concealErrorSlice {
		if gstr.Contains(message, e) {
			return "操作失败，请稍后重试！~"
		}
	}
	return
}
