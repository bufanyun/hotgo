// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/api/admin/monitor"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/simple"
	"hotgo/utility/useragent"
	"sort"
)

// Monitor 监控
var Monitor = cMonitor{
	wsManager: websocket.Manager(),
}

type cMonitor struct {
	wsManager *websocket.ClientManager
}

// UserOffline 下线用户
func (c *cMonitor) UserOffline(ctx context.Context, req *monitor.UserOfflineReq) (res *monitor.UserOfflineRes, err error) {
	client := c.wsManager.GetClient(req.Id)
	if client == nil {
		err = gerror.New("客户端已离线")
		return
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		websocket.SendSuccess(client, "kick")
		websocket.Close(client)
	})
	return
}

// UserOnlineList 获取用户在线列表
func (c *cMonitor) UserOnlineList(ctx context.Context, req *monitor.UserOnlineListReq) (res *monitor.UserOnlineListRes, err error) {
	var (
		clients []*monitor.UserOnlineModel
		i       int
	)

	if c.wsManager.GetClientsLen() == 0 {
		return
	}

	for conn := range c.wsManager.GetClients() {
		if conn.SendClose || conn.User == nil {
			continue
		}

		if req.UserId > 0 && req.UserId != conn.User.Id {
			continue
		}

		if req.Username != "" && req.Username != conn.User.Username {
			continue
		}

		if req.IP != "" && !gstr.Contains(conn.IP, req.IP) {
			continue
		}

		if len(req.FirstTime) == 2 && (conn.User.LoginAt.Before(req.FirstTime[0]) || conn.User.LoginAt.After(req.FirstTime[1])) {
			continue
		}

		clients = append(clients, &monitor.UserOnlineModel{
			ID:            conn.ID,
			IP:            conn.IP,
			Os:            useragent.GetOs(conn.UserAgent),
			Browser:       useragent.GetBrowser(conn.UserAgent),
			FirstTime:     conn.User.LoginAt.Unix(),
			HeartbeatTime: conn.HeartbeatTime,
			App:           conn.User.App,
			UserId:        conn.User.Id,
			Username:      conn.User.Username,
			Avatar:        conn.User.Avatar,
		})
	}

	res = new(monitor.UserOnlineListRes)
	res.PageRes.Pack(req, len(clients))

	sort.Slice(clients, func(i, j int) bool {
		if clients[i].FirstTime == clients[j].FirstTime {
			return clients[i].ID < clients[j].ID
		}
		return clients[i].FirstTime < clients[j].FirstTime
	})

	_, perPage, offset := form.CalPage(req.Page, req.PerPage)
	for k, v := range clients {
		if k >= offset && i <= perPage {
			if simple.IsDemo(ctx) {
				v.IP = consts.DemoTips
			}
			res.List = append(res.List, v)
			i++
		}
	}
	return
}

// NetOnlineList 获取服务在线列表
func (c *cMonitor) NetOnlineList(ctx context.Context, req *monitor.NetOnlineListReq) (res *monitor.NetOnlineListRes, err error) {
	var (
		clients []*monitor.NetOnlineModel
		i       int
		cols    = dao.SysServeLicense.Columns()
		serv    = service.TCPServer().Instance()
		models  *entity.SysServeLicense
	)

	conns := serv.GetClients()
	if len(conns) == 0 {
		res = new(monitor.NetOnlineListRes)
		res.PageRes.Pack(req, 0)
		return
	}

	for _, conn := range conns {
		v := &monitor.NetOnlineModel{
			AuthMeta:      conn.Auth,
			Id:            conn.CID,
			IsAuth:        conn.Auth != nil,
			Addr:          conn.RemoteAddr().String(),
			Port:          gstr.SubStrFromEx(conn.LocalAddr().String(), `:`),
			FirstTime:     conn.FirstTime,
			HeartbeatTime: conn.Heartbeat,
			Proto:         "TCP",
		}

		if v.IsAuth {
			if err = dao.SysServeLicense.Ctx(ctx).Where(cols.Appid, conn.Auth.AppId).Where(cols.Group, conn.Auth.Group).Scan(&models); err != nil {
				return
			}

			if models == nil {
				continue
			}

			v.LicenseId = models.Id
			v.LicenseName = models.Name
			v.LoginTimes = models.LoginTimes
			v.Online = serv.GetAppIdOnline(models.Appid)
			v.OnlineLimit = models.OnlineLimit
		}

		if req.Addr != "" && !gstr.Contains(v.Addr, req.Addr) {
			continue
		}

		if req.Name != "" && !gstr.Contains(v.Name, req.Name) {
			continue
		}

		if req.Group != "" && (!v.IsAuth || v.Group != req.Group) {
			continue
		}

		if req.AppId != "" && (!v.IsAuth || v.AppId != req.AppId) {
			continue
		}

		ft := gtime.New(conn.FirstTime)
		if len(req.FirstTime) == 2 && (ft.Before(req.FirstTime[0]) || ft.After(req.FirstTime[1])) {
			continue
		}
		clients = append(clients, v)
	}

	res = new(monitor.NetOnlineListRes)
	res.PageRes.Pack(req, len(clients))

	sort.Slice(clients, func(i, j int) bool {
		return clients[i].Id > clients[j].Id
	})

	_, perPage, offset := form.CalPage(req.Page, req.PerPage)
	for k, v := range clients {
		if k >= offset && i <= perPage {
			res.List = append(res.List, v)
			i++
		}
	}
	return
}

// NetOption 获取服务选项
func (c *cMonitor) NetOption(ctx context.Context, req *monitor.NetOptionReq) (res *monitor.NetOptionRes, err error) {
	res = new(monitor.NetOptionRes)

	for _, v := range service.TCPServer().Instance().GetRoutes() {
		// 无需勾选的路由
		disabled := false
		if v.Id == "ServerLoginReq" || v.Id == "ServerHeartbeatReq" {
			disabled = true
		}

		res.Routes = append(res.Routes, &monitor.RouteSelect{
			Value:    v.Id,
			Label:    v.Id,
			Disabled: disabled,
			IsRPC:    v.IsRPC,
		})
	}

	sort.Slice(res.Routes, func(i, j int) bool {
		if res.Routes[i].IsRPC && !res.Routes[j].IsRPC {
			return true
		} else if !res.Routes[i].IsRPC && res.Routes[j].IsRPC {
			return false
		}
		return res.Routes[i].Label < res.Routes[j].Label
	})
	return
}

// NetOffline 下线服务
func (c *cMonitor) NetOffline(ctx context.Context, req *monitor.NetOfflineReq) (res *monitor.NetOfflineRes, err error) {
	conn := service.TCPServer().Instance().GetClientById(req.Id)
	if conn == nil {
		err = gerror.New("客户端不在线")
		return
	}
	// 关闭连接
	conn.Close()
	return
}
