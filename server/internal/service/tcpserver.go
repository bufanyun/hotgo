// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/api/servmsg"
	"hotgo/internal/library/cron"
	"hotgo/internal/library/network/tcp"
)

type (
	ITCPServer interface {
		// OnAuthSummary 获取授权信息
		OnAuthSummary(ctx context.Context, req *servmsg.AuthSummaryReq)
		// CronDelete 删除任务
		CronDelete(ctx context.Context, in *servmsg.CronDeleteReq) (err error)
		// CronEdit 编辑任务
		CronEdit(ctx context.Context, in *servmsg.CronEditReq) (err error)
		// CronStatus 修改任务状态
		CronStatus(ctx context.Context, in *servmsg.CronStatusReq) (err error)
		// CronOnlineExec 执行一次任务
		CronOnlineExec(ctx context.Context, in *servmsg.CronOnlineExecReq) (err error)
		// DispatchLog 执行一次任务
		DispatchLog(ctx context.Context, in *servmsg.CronDispatchLogReq) (log *cron.Log, err error)
		// OnExampleHello 一个tcp请求例子
		OnExampleHello(ctx context.Context, req *servmsg.ExampleHelloReq)
		// OnExampleRPCHello 一个rpc请求例子
		OnExampleRPCHello(ctx context.Context, req *servmsg.ExampleRPCHelloReq) (res *servmsg.ExampleRPCHelloRes, err error)
		// Instance 获取实例
		Instance() *tcp.Server
		// Start 启动服务
		Start(ctx context.Context)
		// Stop 关闭服务
		Stop(ctx context.Context)
		// DefaultInterceptor 默认拦截器
		DefaultInterceptor(ctx context.Context, msg *tcp.Message) (err error)
		// PreFilterInterceptor 预处理
		PreFilterInterceptor(ctx context.Context, msg *tcp.Message) (err error)
	}
)

var (
	localTCPServer ITCPServer
)

func TCPServer() ITCPServer {
	if localTCPServer == nil {
		panic("implement not found for interface ITCPServer, forgot register?")
	}
	return localTCPServer
}

func RegisterTCPServer(i ITCPServer) {
	localTCPServer = i
}
