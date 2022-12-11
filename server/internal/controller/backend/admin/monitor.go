// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/backend/monitor"
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

	return res, nil
}

// View 获取指定信息
func (c *cMonitor) View(ctx context.Context, req *monitor.OnlineViewReq) (*monitor.OnlineViewRes, error) {
	var res monitor.OnlineViewRes
	// ...
	return &res, nil
}

// OnlineList 获取在线列表
func (c *cMonitor) OnlineList(ctx context.Context, req *monitor.OnlineListReq) (*monitor.OnlineListRes, error) {
	var (
		res     monitor.OnlineListRes
		clients []*monitor.OnlineModel
		i       int64
	)

	if c.wsManager.GetClientsLen() == 0 {
		return &res, nil
	}

	for c, _ := range c.wsManager.GetClients() {
		if c.SendClose || c.User == nil {
			continue
		}

		if req.UserId > 0 && req.UserId != c.User.Id {
			continue
		}
		clients = append(clients, &monitor.OnlineModel{
			ID:            c.ID,
			Addr:          c.Addr,
			Os:            useragent.GetOs(c.UserAgent),
			Browser:       useragent.GetBrowser(c.UserAgent),
			FirstTime:     c.FirstTime,
			HeartbeatTime: c.HeartbeatTime,
			App:           c.User.App,
			UserId:        c.User.Id,
			Username:      c.User.Username,
			Avatar:        c.User.Avatar,
			ExpTime:       c.User.Exp,
		})
	}

	res.PageCount = form.CalPageCount(int64(len(clients)), req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	sort.Sort(monitor.OnlineModels(clients))
	isDemo, _ := g.Cfg().Get(ctx, "hotgo.isDemo", false)
	_, perPage, offset := form.CalPage(ctx, req.Page, req.PerPage)

	for k, v := range clients {
		if int64(k) >= offset && i <= perPage {
			i++
			if isDemo.Bool() {
				v.Addr = consts.DemoTips
			}
			res.List = append(res.List, v)
		}
	}

	return &res, nil
}
