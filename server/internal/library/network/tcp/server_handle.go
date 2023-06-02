// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/msgin"
	"hotgo/utility/convert"
)

// onServerLogin 处理客户端登录
func (server *Server) onServerLogin(ctx context.Context, args ...interface{}) {
	var (
		in     = new(msgin.ServerLogin)
		user   = GetCtx(ctx)
		res    = new(msgin.ResponseServerLogin)
		models *entity.SysServeLicense
	)

	if err := gconv.Scan(args[0], &in); err != nil {
		server.Logger.Warningf(ctx, "onServerLogin message Scan failed:%+v, args:%+v", err, args)
		return
	}

	err := g.Model("sys_serve_license").Ctx(ctx).
		Where("appid = ?", in.AppId).
		Scan(&models)

	if err != nil {
		res.Code = 1
		res.Message = err.Error()
		_ = server.Write(user.Conn, res)
		return
	}

	if models == nil {
		res.Code = 2
		res.Message = "授权信息不存在"
		_ = server.Write(user.Conn, res)
		return
	}

	// 验证签名
	if _, err = VerifySign(in, models.Appid, models.SecretKey); err != nil {
		res.Code = 3
		res.Message = "签名错误，请联系管理员"
		_ = server.Write(user.Conn, res)
		return
	}

	if models.Status != consts.StatusEnabled {
		res.Code = 4
		res.Message = "授权已禁用，请联系管理员"
		_ = server.Write(user.Conn, res)
		return
	}

	if models.Group != in.Group {
		res.Code = 5
		res.Message = "你登录的授权分组未得到授权，请联系管理员"
		_ = server.Write(user.Conn, res)
		return
	}

	if models.EndAt.Before(gtime.Now()) {
		res.Code = 6
		res.Message = "授权已过期，请联系管理员"
		_ = server.Write(user.Conn, res)
		return
	}

	allowedIps := convert.IpFilterStrategy(models.AllowedIps)
	if _, ok := allowedIps["*"]; !ok {
		ip := gstr.StrTillEx(user.Conn.RemoteAddr().String(), ":")
		if _, ok2 := allowedIps[ip]; !ok2 {
			res.Code = 7
			res.Message = "IP(" + ip + ")未授权，请联系管理员"
			_ = server.Write(user.Conn, res)
			return
		}
	}

	// 检查是否存在多地登录，如果连接超出上限，直接将所有已连接断开
	clients := server.getAppIdClients(models.Appid)
	online := len(clients) + 1
	if online > models.OnlineLimit {
		res2 := new(msgin.ResponseServerLogin)
		res2.Code = 8
		res2.Message = "授权登录端超出上限已进行记录。请立即终止操作。如有疑问请联系管理员"
		for _, client := range clients {
			_ = server.Write(client.Conn, res2)
			_ = client.Conn.Close()
		}

		// 当前连接也踢掉
		_ = server.Write(user.Conn, res2)
		_ = user.Conn.Close()
		return
	}

	server.mutexConns.Lock()
	server.clients[user.Conn.RemoteAddr().String()] = &ClientConn{
		Conn: user.Conn,
		Auth: &AuthMeta{
			Group:     in.Group,
			Name:      in.Name,
			AppId:     in.AppId,
			SecretKey: models.SecretKey,
			EndAt:     models.EndAt,
		},
		heartbeat: gtime.Timestamp(),
	}
	server.mutexConns.Unlock()

	_, err = g.Model("sys_serve_license").Ctx(ctx).
		Where("id = ?", models.Id).Data(g.Map{
		"online":         online,
		"login_times":    models.LoginTimes + 1,
		"last_login_at":  gtime.Now(),
		"last_active_at": gtime.Now(),
		"remote_addr":    user.Conn.RemoteAddr().String(),
	}).Update()
	if err != nil {
		server.Logger.Warningf(ctx, "onServerLogin Update err:%+v", err)
	}

	res.AppId = in.AppId
	res.Code = consts.TCPMsgCodeSuccess
	_ = server.Write(user.Conn, res)
}

// onServerHeartbeat 处理客户端心跳
func (server *Server) onServerHeartbeat(ctx context.Context, args ...interface{}) {
	var (
		in  *msgin.ServerHeartbeat
		res = new(msgin.ResponseServerHeartbeat)
	)

	if err := gconv.Scan(args[0], &in); err != nil {
		server.Logger.Warningf(ctx, "onServerHeartbeat message Scan failed:%+v, args:%+v", err, args)
		return
	}

	client := args[1].(*ClientConn)
	client.heartbeat = gtime.Timestamp()

	_, err := g.Model("sys_serve_license").Ctx(ctx).
		Where("appid = ?", client.Auth.AppId).Data(g.Map{
		"last_active_at": gtime.Now(),
	}).Update()
	if err != nil {
		server.Logger.Warningf(ctx, "onServerHeartbeat Update err:%+v", err)
	}

	res.Code = consts.TCPMsgCodeSuccess
	_ = server.Write(client.Conn, res)
}
