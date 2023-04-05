package tcp

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/msgin"
	"hotgo/utility/convert"
)

func (server *Server) onServerLogin(args ...interface{}) {
	var (
		in     = new(msgin.ServerLogin)
		conn   = args[1].(*gtcp.Conn)
		res    = new(msgin.ResponseServerLogin)
		models *entity.SysServeLicense
	)

	if err := gconv.Scan(args[0], &in); err != nil {
		server.Logger.Infof(server.Ctx, "onServerLogin message Scan failed:%+v, args:%+v", err, args)
		return
	}
	server.Logger.Infof(server.Ctx, "onServerLogin in:%+v", *in)

	err := g.Model("sys_serve_license").
		Ctx(server.Ctx).
		Where("appid = ?", in.AppId).
		Scan(&models)

	if err != nil {
		res.Code = 1
		res.Message = err.Error()
		server.Write(conn, res)
		return
	}

	if models == nil {
		res.Code = 2
		res.Message = "授权信息不存在"
		server.Write(conn, res)
		return
	}

	// 验证签名
	if err = VerifySign(in, models.Appid, models.SecretKey); err != nil {
		res.Code = 3
		res.Message = "签名错误，请联系管理员"
		server.Write(conn, res)
		return
	}

	if models.Status != consts.StatusEnabled {
		res.Code = 4
		res.Message = "授权已禁用，请联系管理员"
		server.Write(conn, res)
		return
	}

	if models.Group != in.Group {
		res.Code = 5
		res.Message = "你登录的授权分组未得到授权，请联系管理员"
		server.Write(conn, res)
		return
	}

	if models.EndAt.Before(gtime.Now()) {
		res.Code = 6
		res.Message = "授权已过期，请联系管理员"
		server.Write(conn, res)
		return
	}

	allowedIps := convert.IpFilterStrategy(models.AllowedIps)
	if _, ok := allowedIps["*"]; !ok {
		ip := gstr.StrTillEx(conn.RemoteAddr().String(), ":")
		if _, ok2 := allowedIps[ip]; !ok2 {
			res.Code = 7
			res.Message = "IP(" + ip + ")未授权，请联系管理员"
			server.Write(conn, res)
			return
		}
	}

	// 检查是否存在多地登录，如果连接超出上限，直接将所有已连接断开
	clients := server.getAppIdClients(models.Appid)
	online := len(clients) + 1
	if online > models.OnlineLimit {
		online = 0
		res2 := new(msgin.ResponseServerLogin)
		res2.Code = 8
		res2.Message = "授权登录端超出上限已进行记录。请立即终止操作。如有疑问请联系管理员"
		for _, client := range clients {
			server.Write(client.Conn, res2)
			client.Conn.Close()
		}

		// 当前连接也踢掉
		server.Write(conn, res2)
		conn.Close()
		return
	}

	server.mutexConns.Lock()
	server.clients[conn.RemoteAddr().String()] = &ClientConn{
		Conn: conn,
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

	server.Write(conn, res)

	_, err = g.Model("sys_serve_license").
		Ctx(server.Ctx).
		Where("id = ?", models.Id).Data(g.Map{
		"online":         online,
		"login_times":    models.LoginTimes + 1,
		"last_login_at":  gtime.Now(),
		"last_active_at": gtime.Now(),
		"remote_addr":    conn.RemoteAddr().String(),
	}).Update()
	if err != nil {
		server.Logger.Warningf(server.Ctx, "onServerLogin Update err:%+v", err)
	}
}

func (server *Server) onServerHeartbeat(args ...interface{}) {
	var in *msgin.ServerHeartbeat
	if err := gconv.Scan(args, &in); err != nil {
		server.Logger.Infof(server.Ctx, "onServerHeartbeat message Scan failed:%+v, args:%+v", err, args)
		return
	}
	client := args[1].(*ClientConn)
	client.heartbeat = gtime.Timestamp()

	server.Write(client.Conn, &msgin.ResponseServerHeartbeat{})

	_, err := g.Model("sys_serve_license").
		Ctx(server.Ctx).
		Where("appid = ?", client.Auth.AppId).Data(g.Map{
		"last_active_at": gtime.Now(),
	}).Update()
	if err != nil {
		server.Logger.Warningf(server.Ctx, "onServerHeartbeat Update err:%+v", err)
	}
}
