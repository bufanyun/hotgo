package tcpclient

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model/input/msgin"
	"hotgo/internal/service"
)

// OnCronDelete 删除任务
func (s *sCronClient) OnCronDelete(ctx context.Context, args ...interface{}) {
	var (
		in  *msgin.CronDelete
		res = new(msgin.ResponseCronDelete)
	)

	if err := gconv.Scan(args[0], &in); err != nil {
		res.Code = 1
		res.Message = err.Error()
		s.client.Reply(ctx, res)
		return
	}

	if err := service.SysCron().Delete(ctx, in.CronDeleteInp); err != nil {
		res.Code = 2
		res.Message = err.Error()
	}

	s.client.Reply(ctx, res)
}

// OnCronEdit 编辑任务
func (s *sCronClient) OnCronEdit(ctx context.Context, args ...interface{}) {
	var (
		in  *msgin.CronEdit
		res = new(msgin.ResponseCronEdit)
	)

	if err := gconv.Scan(args[0], &in); err != nil {
		res.Code = 1
		res.Message = err.Error()
		s.client.Reply(ctx, res)
		return
	}

	if err := service.SysCron().Edit(ctx, in.CronEditInp); err != nil {
		res.Code = 2
		res.Message = err.Error()
	}

	s.client.Reply(ctx, res)
}

// OnCronStatus 修改任务状态
func (s *sCronClient) OnCronStatus(ctx context.Context, args ...interface{}) {
	var (
		in  *msgin.CronStatus
		res = new(msgin.ResponseCronStatus)
	)

	if err := gconv.Scan(args[0], &in); err != nil {
		res.Code = 1
		res.Message = err.Error()
		s.client.Reply(ctx, res)
		return
	}

	if err := service.SysCron().Status(ctx, in.CronStatusInp); err != nil {
		res.Code = 2
		res.Message = err.Error()
	}

	s.client.Reply(ctx, res)
}

// OnCronOnlineExec 执行一次任务
func (s *sCronClient) OnCronOnlineExec(ctx context.Context, args ...interface{}) {
	var (
		in  *msgin.CronOnlineExec
		res = new(msgin.ResponseCronOnlineExec)
	)

	if err := gconv.Scan(args[0], &in); err != nil {
		res.Code = 1
		res.Message = err.Error()
		s.client.Reply(ctx, res)
		return
	}

	if err := service.SysCron().OnlineExec(ctx, in.OnlineExecInp); err != nil {
		res.Code = 1
		res.Message = err.Error()
	}

	s.client.Reply(ctx, res)
}
