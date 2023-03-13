package tcp

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

func (client *Client) getCronKey(s string) string {
	return fmt.Sprintf("tcp.client_%s_%s:%s", s, client.auth.Group, client.auth.Name)
}

func (client *Client) stopCron() {
	for _, v := range gcron.Entries() {
		gcron.Remove(v.Name)
	}
}

func (client *Client) startCron() {
	// 心跳超时检查
	if gcron.Search(client.getCronKey(cronHeartbeatVerify)) == nil {
		gcron.AddSingleton(client.Ctx, "@every 600s", func(ctx context.Context) {
			if client.heartbeat < gtime.Timestamp()-600 {
				client.Logger.Debugf(client.Ctx, "client  heartbeat timeout, about to reconnect..")
				client.Destroy()
			}
		}, client.getCronKey(cronHeartbeatVerify))
	}

	// 心跳
	if gcron.Search(client.getCronKey(cronHeartbeat)) == nil {
		gcron.AddSingleton(client.Ctx, "@every 120s", func(ctx context.Context) {
			client.serverHeartbeat()
		}, client.getCronKey(cronHeartbeat))
	}
}
