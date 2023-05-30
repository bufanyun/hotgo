package tcpclient

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model/input/msgin"
)

// OnResponseAuthSummary 获取授权信息
func (s *sAuthClient) OnResponseAuthSummary(ctx context.Context, args ...interface{}) {
	var in *msgin.ResponseAuthSummary
	if err := gconv.Scan(args[0], &in); err != nil {
		s.client.Logger.Warningf(ctx, "OnResponseAuthSummary Scan err:%+v", err)
		return
	}

	if err := in.GetError(); err != nil {
		s.client.Logger.Warningf(ctx, "OnResponseAuthSummary GetError:%+v", err)
		return
	}

	// 拿到授权的数据，可以是一些动态的功能、路由、权限控制等
	s.summary = in.Data
}
