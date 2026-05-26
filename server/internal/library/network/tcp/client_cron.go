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

// getCronKey 生成客户端定时任务名称
func (client *Client) getCronKey(s string) string {
	return fmt.Sprintf("tcp.client_%s:%d", s, client.conn.CID)
}

// stopCron 停止定时任务
func (client *Client) stopCron() {
	if client.conn == nil {
		return
	}
	gcron.Remove(client.getCronKey(CronHeartbeatVerify))
	gcron.Remove(client.getCronKey(CronHeartbeat))
}

// startCron 启动定时任务
func (client *Client) startCron() {
	// 心跳超时检查
	if gcron.Search(client.getCronKey(CronHeartbeatVerify)) == nil {
		_, _ = gcron.AddSingleton(client.ctx, "@every 600s", func(ctx context.Context) {
			if client == nil || client.conn == nil {
				return
			}
			if client.conn.Heartbeat < gtime.Timestamp()-HeartbeatTimeout {
				client.logger.Debugf(client.ctx, "client heartbeat timeout, about to reconnect..")
				client.Destroy()
			}
		}, client.getCronKey(CronHeartbeatVerify))
	}

	// 心跳
	if gcron.Search(client.getCronKey(CronHeartbeat)) == nil {
		_, _ = gcron.AddSingleton(client.ctx, "@every 300s", func(ctx context.Context) {
			client.serverHeartbeat()
		}, client.getCronKey(CronHeartbeat))
	}
}
