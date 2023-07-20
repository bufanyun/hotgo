// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

// getCronKey 生成服务端定时任务名称
func (server *Server) getCronKey(s string) string {
	return fmt.Sprintf("tcp.server_%s_%s", s, server.name)
}

// stopCron 停止定时任务
func (server *Server) stopCron() {
	for _, v := range gcron.Entries() {
		gcron.Remove(v.Name)
	}
}

// startCron 启动定时任务
func (server *Server) startCron() {
	// 心跳超时检查
	if gcron.Search(server.getCronKey(CronHeartbeatVerify)) == nil {
		gcron.AddSingleton(server.ctx, "@every 300s", func(ctx context.Context) {
			if server == nil || server.clients == nil {
				return
			}
			for _, client := range server.clients {
				if client.Heartbeat < gtime.Timestamp()-HeartbeatTimeout {
					client.Conn.Close()
					server.logger.Debugf(server.ctx, "client heartbeat timeout, close conn. auth:%+v", client.Auth)
				}
			}
		}, server.getCronKey(CronHeartbeatVerify))
	}

	// 认证检查
	if gcron.Search(server.getCronKey(CronAuthVerify)) == nil {
		gcron.AddSingleton(server.ctx, "@every 300s", func(ctx context.Context) {
			if server == nil || server.clients == nil {
				return
			}
			for _, client := range server.clients {
				if client.Auth == nil {
					continue
				}
				if client.Auth.EndAt.Before(gtime.Now()) {
					client.Conn.Close()
					server.logger.Debugf(server.ctx, "client auth expired, close conn. auth:%+v", client.Auth)
				}
			}
		}, server.getCronKey(CronAuthVerify))
	}
}
