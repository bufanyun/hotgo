package tcpserver

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/msgin"
)

// CronDelete 删除任务
func (s *sTCPServer) CronDelete(ctx context.Context, in *msgin.CronDelete) (err error) {
	clients := s.serv.GetGroupClients(consts.TCPClientGroupCron)
	if len(clients) == 0 {
		err = gerror.New("没有在线的定时任务服务")
		return
	}

	for _, client := range clients {
		res, err := s.serv.RpcRequest(ctx, client, in)
		if err != nil {
			return err
		}

		var resp = new(msgin.ResponseCronDelete)
		if err = gconv.Scan(res, &resp); err != nil {
			return err
		}

		if err = resp.GetError(); err != nil {
			return err
		}
	}

	return
}

// CronEdit 编辑任务
func (s *sTCPServer) CronEdit(ctx context.Context, in *msgin.CronEdit) (err error) {
	clients := s.serv.GetGroupClients(consts.TCPClientGroupCron)
	if len(clients) == 0 {
		err = gerror.New("没有在线的定时任务服务")
		return
	}

	for _, client := range clients {
		res, err := s.serv.RpcRequest(ctx, client, in)
		if err != nil {
			return err
		}

		var resp = new(msgin.ResponseCronEdit)
		if err = gconv.Scan(res, &resp); err != nil {
			return err
		}

		if err = resp.GetError(); err != nil {
			return err
		}
	}

	return
}

// CronStatus 修改任务状态
func (s *sTCPServer) CronStatus(ctx context.Context, in *msgin.CronStatus) (err error) {
	clients := s.serv.GetGroupClients(consts.TCPClientGroupCron)
	if len(clients) == 0 {
		err = gerror.New("没有在线的定时任务服务")
		return
	}

	for _, client := range clients {
		res, err := s.serv.RpcRequest(ctx, client, in)
		if err != nil {
			return err
		}

		var resp = new(msgin.ResponseCronStatus)
		if err = gconv.Scan(res, &resp); err != nil {
			return err
		}

		if err = resp.GetError(); err != nil {
			return err
		}
	}

	return
}

// CronOnlineExec 执行一次任务
func (s *sTCPServer) CronOnlineExec(ctx context.Context, in *msgin.CronOnlineExec) (err error) {
	clients := s.serv.GetGroupClients(consts.TCPClientGroupCron)
	if len(clients) == 0 {
		err = gerror.New("没有在线的定时任务服务")
		return
	}

	for _, client := range clients {
		res, err := s.serv.RpcRequest(ctx, client, in)
		if err != nil {
			return err
		}

		var resp = new(msgin.ResponseCronOnlineExec)
		if err = gconv.Scan(res, &resp); err != nil {
			return err
		}

		if err = resp.GetError(); err != nil {
			return err
		}
	}

	return
}
