// Package tcpclient
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpclient

import (
	"context"
	"hotgo/api/servmsg"
	"hotgo/internal/service"
)

// OnCronDelete 删除任务
func (s *sCronClient) OnCronDelete(ctx context.Context, req *servmsg.CronDeleteReq) (res *servmsg.CronDeleteRes, err error) {
	err = service.SysCron().Delete(ctx, req.CronDeleteInp)
	return
}

// OnCronEdit 编辑任务
func (s *sCronClient) OnCronEdit(ctx context.Context, req *servmsg.CronEditReq) (res *servmsg.CronEditRes, err error) {
	err = service.SysCron().Edit(ctx, req.CronEditInp)
	return
}

// OnCronStatus 修改任务状态
func (s *sCronClient) OnCronStatus(ctx context.Context, req *servmsg.CronStatusReq) (res *servmsg.CronStatusRes, err error) {
	err = service.SysCron().Status(ctx, req.CronStatusInp)
	return
}

// OnCronOnlineExec 执行一次任务
func (s *sCronClient) OnCronOnlineExec(ctx context.Context, req *servmsg.CronOnlineExecReq) (res *servmsg.CronOnlineExecRes, err error) {
	err = service.SysCron().OnlineExec(ctx, req.OnlineExecInp)
	return
}
