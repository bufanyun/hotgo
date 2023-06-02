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
	"hotgo/internal/consts"
)

// getCronKey 生成客户端定时任务名称
func (client *Client) getCronKey(s string) string {
	return fmt.Sprintf("tcp.client_%s_%s:%s", s, client.auth.Group, client.auth.Name)
}

// stopCron 停止定时任务
func (client *Client) stopCron() {
	for _, v := range gcron.Entries() {
		gcron.Remove(v.Name)
	}
}

// startCron 启动定时任务
func (client *Client) startCron() {
	// 心跳超时检查
	if gcron.Search(client.getCronKey(consts.TCPCronHeartbeatVerify)) == nil {
		_, _ = gcron.AddSingleton(client.Ctx, "@every 600s", func(ctx context.Context) {
			if client.heartbeat < gtime.Timestamp()-consts.TCPHeartbeatTimeout {
				client.Logger.Debugf(client.Ctx, "client heartbeat timeout, about to reconnect..")
				client.Destroy()
			}
		}, client.getCronKey(consts.TCPCronHeartbeatVerify))
	}

	// 心跳
	if gcron.Search(client.getCronKey(consts.TCPCronHeartbeat)) == nil {
		_, _ = gcron.AddSingleton(client.Ctx, "@every 120s", func(ctx context.Context) {
			client.serverHeartbeat()
		}, client.getCronKey(consts.TCPCronHeartbeat))
	}
}
