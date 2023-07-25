// Package tcpclient
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpclient

import (
	"context"
	"hotgo/internal/library/network/tcp"
)

// DefaultInterceptor 默认拦截器
func (s *sCronClient) DefaultInterceptor(ctx context.Context, msg *tcp.Message) (err error) {
	// conn := tcp.ConnFromCtx(ctx)
	// g.Log().Debugf(ctx, "DefaultInterceptor msg:%+v, conn:%+v", msg, gjson.New(conn).String())
	return
}
