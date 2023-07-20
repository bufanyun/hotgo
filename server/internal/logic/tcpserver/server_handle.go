// Package tcpserver
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpserver

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/entity"
	"hotgo/utility/convert"
	"hotgo/utility/encrypt"
)

// onServerLogin 处理客户端登录
func (s *sTCPServer) onServerLogin(ctx context.Context, req *tcp.ServerLoginReq) {
	var (
		conn   = tcp.ConnFromCtx(ctx)
		models *entity.SysServeLicense
		res    = new(tcp.ServerLoginRes)
		cols   = dao.SysServeLicense.Columns()
	)

	if conn == nil {
		g.Log().Warningf(ctx, "conn is nil.")
		return
	}

	if err := dao.SysServeLicense.Ctx(ctx).Where(cols.Appid, req.AppId).Scan(&models); err != nil {
		return
	}
	if models == nil {
		res.SetError(gerror.New("授权信息不存在"))
		conn.Send(ctx, res)
		return
	}

	// 验证签名
	sign := encrypt.Md5ToString(fmt.Sprintf("%v%v%v", models.Appid, req.Timestamp, models.SecretKey))
	if sign != req.Sign {
		res.SetError(gerror.New("签名错误，请检查！"))
		conn.Send(ctx, res)
		return
	}

	if models.Status != consts.StatusEnabled {
		res.SetError(gerror.New("授权已禁用，请联系管理员"))
		conn.Send(ctx, res)
		return
	}

	if models.Group != req.Group {
		res.SetError(gerror.New("你登录的授权分组未得到授权，请联系管理员"))
		conn.Send(ctx, res)
		return
	}

	if models.EndAt.Before(gtime.Now()) {
		res.SetError(gerror.New("授权已过期，请联系管理员"))
		conn.Send(ctx, res)
		return
	}

	ip := gstr.StrTillEx(conn.RemoteAddr().String(), ":")
	if !convert.MatchIpStrategy(models.AllowedIps, ip) {
		res.SetError(gerror.New("IP(" + ip + ")未授权，请联系管理员"))
		conn.Send(ctx, res)
		return
	}

	var routes []string
	if err := models.Routes.Scan(&routes); err != nil {
		res.SetError(gerror.New("授权路由解析失败，请联系管理员"))
		conn.Send(ctx, res)
		return
	}

	// 拿出当前登录应用的所有客户端
	clients := s.serv.GetAppIdClients(models.Appid)

	// 检查多地登录，如果连接超过上限，则断开当前许可证下的所有连接
	if len(clients)+1 > models.OnlineLimit {
		for _, client := range clients {
			client.Close()
		}
		res.SetError(gerror.New("授权登录端超出上限，请勿多地登录"))
		conn.Send(ctx, res)
		return
	}

	for _, client := range clients {
		if client.Auth.Name == req.Name {
			res.SetError(gerror.Newf("应用名称[%v]已存在登录用户，当前连接已被拒绝。", req.Name))
			conn.Send(ctx, res)
			return
		}
	}

	auth := &tcp.AuthMeta{
		Name:      req.Name,
		Extra:     req.Extra,
		Group:     models.Group,
		AppId:     models.Appid,
		SecretKey: models.SecretKey,
		EndAt:     models.EndAt,
		Routes:    routes,
	}
	s.serv.AuthClient(conn, auth)

	update := g.Map{
		cols.LoginTimes:   models.LoginTimes + 1,
		cols.LastLoginAt:  gtime.Now(),
		cols.LastActiveAt: gtime.Now(),
		cols.RemoteAddr:   conn.RemoteAddr().String(),
	}
	if _, err := dao.SysServeLicense.Ctx(ctx).Where(cols.Id, models.Id).Data(update).Update(); err != nil {
		res.SetError(err)
		conn.Send(ctx, res)
		return
	}

	g.Log().Debugf(ctx, "onServerLogin succeed. appid:%v, group:%v, name:%v", auth.AppId, auth.Group, auth.Name)
	conn.Send(ctx, res)
}

// onServerHeartbeat 处理客户端心跳
func (s *sTCPServer) onServerHeartbeat(ctx context.Context, req *tcp.ServerHeartbeatReq) {
	var (
		conn = tcp.ConnFromCtx(ctx)
		res  = new(tcp.ServerHeartbeatRes)
	)

	if conn == nil {
		g.Log().Warningf(ctx, "conn is nil.")
		return
	}

	client := s.serv.GetClient(conn.Conn)
	if client == nil {
		res.SetError(gerror.New("登录异常，请重新登录"))
		conn.Send(ctx, res)
		return
	}

	// 更新心跳
	client.Heartbeat = gtime.Timestamp()

	// 更新活跃时间
	update := g.Map{
		dao.SysServeLicense.Columns().LastActiveAt: gtime.Now(),
	}
	if _, err := dao.SysServeLicense.Ctx(ctx).Where(dao.SysServeLicense.Columns().Appid, client.Auth.AppId).Data(update).Update(); err != nil {
		res.SetError(err)
		conn.Send(ctx, res)
		return
	}

	conn.Send(ctx, res)
}
