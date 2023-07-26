// Package pubsub
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package pubsub

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// Publish 推送消息
func Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	return g.Redis().Publish(ctx, channel, message)
}
