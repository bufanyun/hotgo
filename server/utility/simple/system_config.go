// Package simple
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package simple

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// AppName 应用名称
func AppName(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.appName", "hotgo").String()
}

// Debug debug
func Debug(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "system.debug", true).Bool()
}

// IsDemo 是否为演示系统
func IsDemo(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "system.isDemo", true).Bool()
}

// IsCluster 是否为集群部署
func IsCluster(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "system.isCluster", true).Bool()
}
