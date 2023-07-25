// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

// 定时任务
const (
	CronHeartbeatVerify = "tcpHeartbeatVerify"
	CronHeartbeat       = "tcpHeartbeat"
	CronAuthVerify      = "tcpAuthVerify"
)

const (
	HeartbeatTimeout = 300 // tcp心跳超时，默认300s
	RPCTimeout       = 10  // rpc通讯超时时间， 默认10s
)

const (
	ParseRouterErrInvalidParams      = "register router[%v] method must have two params"
	ParseRouterRPCErrInvalidParams   = "register RPC router [%v] method must have two response params"
	ParseRouterErrInvalidFirstParam  = "the first params of the processing method that registers the router[%v] must be of type context.Context"
	ParseRouterErrInvalidSecondParam = "the second params of the processing method that registers the router[%v] must be of type pointer to a struct"
)

type CtxKey string

// ContextKey 上下文
const (
	ContextKey CtxKey = "tcpContext" // tcp上下文变量名称
)
