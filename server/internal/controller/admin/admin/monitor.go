// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/api/admin/monitor"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
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

// Offline 下线用户
func (c *cMonitor) Offline(ctx context.Context, req *monitor.OfflineReq) (res *monitor.OfflineRes, err error) {
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

// View 获取指定信息
func (c *cMonitor) View(ctx context.Context, req *monitor.OnlineViewReq) (res *monitor.OnlineViewRes, err error) {
	return
}

// OnlineList 获取在线列表
func (c *cMonitor) OnlineList(ctx context.Context, req *monitor.OnlineListReq) (res *monitor.OnlineListRes, err error) {
	var (
		clients []*monitor.OnlineModel
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

		if req.Addr != "" && !gstr.Contains(conn.Addr, req.Addr) {
			continue
		}

		clients = append(clients, &monitor.OnlineModel{
			ID:            conn.ID,
			Addr:          conn.Addr,
			Os:            useragent.GetOs(conn.UserAgent),
			Browser:       useragent.GetBrowser(conn.UserAgent),
			FirstTime:     gtime.New(conn.User.LoginAt).Unix(),
			HeartbeatTime: conn.HeartbeatTime,
			App:           conn.User.App,
			UserId:        conn.User.Id,
			Username:      conn.User.Username,
			Avatar:        conn.User.Avatar,
		})
	}

	res = new(monitor.OnlineListRes)
	res.PageCount = form.CalPageCount(len(clients), req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	sort.Sort(monitor.OnlineModels(clients))
	isDemo := g.Cfg().MustGet(ctx, "hotgo.isDemo", false).Bool()
	_, perPage, offset := form.CalPage(ctx, req.Page, req.PerPage)

	for k, v := range clients {
		if k >= offset && i <= perPage {
			if isDemo {
				v.Addr = consts.DemoTips
			}
			res.List = append(res.List, v)
			i++
		}
	}
	return
}
