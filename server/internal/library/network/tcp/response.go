// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

type Response interface {
	PkgResponse()
	GetError() (err error)
}

// PkgResponse 打包响应消息
func PkgResponse(data interface{}) {
	if c, ok := data.(Response); ok {
		c.PkgResponse()
		return
	}
}

// GetResponseError 解析响应消息中的错误
func GetResponseError(data interface{}) (err error) {
	if c, ok := data.(Response); ok {
		return c.GetError()
	}
	return
}
