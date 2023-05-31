// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

type CtxKey string

// ContextKey 上下文
const (
	ContextHTTPKey     CtxKey = "httpContext" // http上下文变量名称
	ContextKeyCronArgs CtxKey = "cronArgs"    // 定时任务参数上下文变量名称
	ContextTCPKey      CtxKey = "tcpContext"  // tcp上下文变量名称
)
