// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/api/servmsg"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/servmsgin"
)

type (
	IAuthClient interface {
		// Instance 获取实例
		Instance() *tcp.Client
		// Start 启动服务
		Start(ctx context.Context)
		// Stop 停止服务
		Stop(ctx context.Context)
		// Summary 获取授权信息
		Summary() *servmsgin.AuthSummaryModel
		// OnResponseAuthSummary 响应授权信息
		OnResponseAuthSummary(ctx context.Context, req *servmsg.AuthSummaryRes)
		// OnResponseExampleHello 一个tcp请求例子
		OnResponseExampleHello(ctx context.Context, req *servmsg.ExampleHelloRes)
	}
	ICronClient interface {
		// Instance 获取实例
		Instance() *tcp.Client
		// Start 启动服务
		Start(ctx context.Context)
		// Stop 停止服务
		Stop(ctx context.Context)
		// OnCronDelete 删除任务
		OnCronDelete(ctx context.Context, req *servmsg.CronDeleteReq) (res *servmsg.CronDeleteRes, err error)
		// OnCronEdit 编辑任务
		OnCronEdit(ctx context.Context, req *servmsg.CronEditReq) (res *servmsg.CronEditRes, err error)
		// OnCronStatus 修改任务状态
		OnCronStatus(ctx context.Context, req *servmsg.CronStatusReq) (res *servmsg.CronStatusRes, err error)
		// OnCronOnlineExec 执行一次任务
		OnCronOnlineExec(ctx context.Context, req *servmsg.CronOnlineExecReq) (res *servmsg.CronOnlineExecRes, err error)
		// OnCronDispatchLog 查看调度日志
		OnCronDispatchLog(ctx context.Context, req *servmsg.CronDispatchLogReq) (res *servmsg.CronDispatchLogRes, err error)
		// DefaultInterceptor 默认拦截器
		DefaultInterceptor(ctx context.Context, msg *tcp.Message) (err error)
	}
)

var (
	localAuthClient IAuthClient
	localCronClient ICronClient
)

func AuthClient() IAuthClient {
	if localAuthClient == nil {
		panic("implement not found for interface IAuthClient, forgot register?")
	}
	return localAuthClient
}

func RegisterAuthClient(i IAuthClient) {
	localAuthClient = i
}

func CronClient() ICronClient {
	if localCronClient == nil {
		panic("implement not found for interface ICronClient, forgot register?")
	}
	return localCronClient
}

func RegisterCronClient(i ICronClient) {
	localCronClient = i
}
