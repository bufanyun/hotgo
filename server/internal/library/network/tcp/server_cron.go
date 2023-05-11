package tcp

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
)

func (server *Server) getCronKey(s string) string {
	return fmt.Sprintf("tcp.server_%s_%s", s, server.name)
}

func (server *Server) stopCron() {
	for _, v := range gcron.Entries() {
		gcron.Remove(v.Name)
	}
}

func (server *Server) startCron() {
	// 心跳超时检查
	if gcron.Search(server.getCronKey(consts.TCPCronHeartbeatVerify)) == nil {
		gcron.AddSingleton(server.Ctx, "@every 300s", func(ctx context.Context) {
			if server.clients == nil {
				return
			}
			for _, client := range server.clients {
				if client.heartbeat < gtime.Timestamp()-300 {
					client.Conn.Close()
					server.Logger.Debugf(server.Ctx, "client heartbeat timeout, close conn. auth:%+v", client.Auth)
				}
			}
		}, server.getCronKey(consts.TCPCronHeartbeatVerify))
	}

	// 认证检查
	if gcron.Search(server.getCronKey(consts.TCPCronAuthVerify)) == nil {
		gcron.AddSingleton(server.Ctx, "@every 300s", func(ctx context.Context) {
			if server.clients == nil {
				return
			}
			for _, client := range server.clients {
				if client.Auth.EndAt.Before(gtime.Now()) {
					client.Conn.Close()
					server.Logger.Debugf(server.Ctx, "client auth expired, close conn. auth:%+v", client.Auth)
				}
			}
		}, server.getCronKey(consts.TCPCronAuthVerify))
	}
}
