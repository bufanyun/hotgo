package msgin

import (
	"hotgo/internal/model/input/sysin"
)

// CronDelete 删除任务
type CronDelete struct {
	RpcMsg
	sysin.CronDeleteInp
}

type ResponseCronDelete struct {
	Response
	sysin.CronDeleteModel
}

// CronEdit 编辑任务
type CronEdit struct {
	RpcMsg
	sysin.CronEditInp
}

type ResponseCronEdit struct {
	Response
	sysin.CronEditModel
}

// CronStatus 修改任务状态
type CronStatus struct {
	RpcMsg
	sysin.CronStatusInp
}

type ResponseCronStatus struct {
	Response
	sysin.CronStatusModel
}

// CronOnlineExec 在线执行
type CronOnlineExec struct {
	RpcMsg
	sysin.OnlineExecInp
}

type ResponseCronOnlineExec struct {
	Response
	sysin.OnlineExecModel
}
