// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/library/hgrds/lock"
	"hotgo/internal/library/hgrds/pubsub"
	"hotgo/internal/service"
)

// SubscribeClusterSync 订阅集群同步，可以用来集中同步数据、状态等
func SubscribeClusterSync(ctx context.Context) {
	isCluster := g.Cfg().MustGet(ctx, "hotgo.isCluster").Bool()
	if !isCluster {
		return
	}

	err := pubsub.SubscribeMap(map[string]pubsub.SubHandler{
		consts.ClusterSyncSysconfig:     service.SysConfig().ClusterSync,             // 系统配置
		consts.ClusterSyncSysBlacklist:  service.SysBlacklist().ClusterSync,          // 系统黑名单
		consts.ClusterSyncSysSuperAdmin: service.AdminMember().ClusterSyncSuperAdmin, // 超管
	})

	if err != nil {
		g.Log().Fatal(ctx, err)
	}
}

// PublishClusterSync 推送集群同步消息，如果没有开启集群部署，则不进行推送
func PublishClusterSync(ctx context.Context, channel string, message interface{}) {
	isCluster := g.Cfg().MustGet(ctx, "hotgo.isCluster").Bool()
	if !isCluster {
		return
	}

	mutex := lock.Mutex(fmt.Sprintf("%s:%s", "lock", channel))
	err := mutex.LockFunc(ctx, func() {
		if _, err := pubsub.Publish(ctx, channel, message); err != nil {
			g.Log().Warningf(ctx, "PublishClusterSync %v err:%v", channel, err)
		}
	})
	if err != nil {
		g.Log().Warningf(ctx, "PublishClusterSync %v LockFunc err:%v", channel, err)
	}
}
