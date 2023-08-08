// Package servmsg
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
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
}

// CronEditReq 编辑任务
type CronEditReq struct {
	*sysin.CronEditInp
}

type CronEditRes struct {
	tcp.ServerRes
}

// CronStatusReq 修改任务状态
type CronStatusReq struct {
	*sysin.CronStatusInp
}

type CronStatusRes struct {
	tcp.ServerRes
}

// CronOnlineExecReq 在线执行
type CronOnlineExecReq struct {
	*sysin.OnlineExecInp
}

type CronOnlineExecRes struct {
	tcp.ServerRes
}
