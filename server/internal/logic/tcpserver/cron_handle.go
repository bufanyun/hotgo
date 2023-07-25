// Package tcpserver
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpserver

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/api/servmsg"
	"hotgo/internal/consts"
)

// CronDelete 删除任务
func (s *sTCPServer) CronDelete(ctx context.Context, in *servmsg.CronDeleteReq) (err error) {
	clients := s.serv.GetGroupClients(consts.LicenseGroupCron)
	if len(clients) == 0 {
		err = gerror.New("没有在线的定时任务服务")
		return
	}

	for _, client := range clients {
		var res servmsg.CronDeleteRes
		if err = s.serv.RequestScan(ctx, client, in, &res); err != nil {
			return
		}

		if err = res.GetError(); err != nil {
			return
		}
	}
	return
}

// CronEdit 编辑任务
func (s *sTCPServer) CronEdit(ctx context.Context, in *servmsg.CronEditReq) (err error) {
	clients := s.serv.GetGroupClients(consts.LicenseGroupCron)
	if len(clients) == 0 {
		err = gerror.New("没有在线的定时任务服务")
		return
	}

	for _, client := range clients {
		var res servmsg.CronEditRes
		if err = s.serv.RequestScan(ctx, client, in, &res); err != nil {
			return
		}

		if err = res.GetError(); err != nil {
			return
		}
	}
	return
}

// CronStatus 修改任务状态
func (s *sTCPServer) CronStatus(ctx context.Context, in *servmsg.CronStatusReq) (err error) {
	clients := s.serv.GetGroupClients(consts.LicenseGroupCron)
	if len(clients) == 0 {
		err = gerror.New("没有在线的定时任务服务")
		return
	}

	for _, client := range clients {
		var res servmsg.CronStatusRes
		if err = s.serv.RequestScan(ctx, client, in, &res); err != nil {
			return
		}

		if err = res.GetError(); err != nil {
			return
		}
	}
	return
}

// CronOnlineExec 执行一次任务
func (s *sTCPServer) CronOnlineExec(ctx context.Context, in *servmsg.CronOnlineExecReq) (err error) {
	clients := s.serv.GetGroupClients(consts.LicenseGroupCron)
	if len(clients) == 0 {
		err = gerror.New("没有在线的定时任务服务")
		return
	}

	for _, client := range clients {
		var res servmsg.CronOnlineExecRes
		if err = s.serv.RequestScan(ctx, client, in, &res); err != nil {
			return
		}

		if err = res.GetError(); err != nil {
			return
		}
	}
	return
}
