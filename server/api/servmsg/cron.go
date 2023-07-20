package servmsg

import (
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/sysin"
)

// CronDeleteReq 删除任务
type CronDeleteReq struct {
	*sysin.CronDeleteInp
}

type CronDeleteRes struct {
	tcp.ServerRes
	sysin.CronDeleteModel
}

// CronEditReq 编辑任务
type CronEditReq struct {
	*sysin.CronEditInp
}

type CronEditRes struct {
	tcp.ServerRes
	*sysin.CronEditModel
}

// CronStatusReq 修改任务状态
type CronStatusReq struct {
	*sysin.CronStatusInp
}

type CronStatusRes struct {
	tcp.ServerRes
	*sysin.CronStatusModel
}

// CronOnlineExecReq 在线执行
type CronOnlineExecReq struct {
	*sysin.OnlineExecInp
}

type CronOnlineExecRes struct {
	tcp.ServerRes
	*sysin.OnlineExecModel
}
