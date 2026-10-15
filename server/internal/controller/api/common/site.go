// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	apicommon "hotgo/api/api/common"
)

var Site = cSite{}

type cSite struct{}

// Health 检查默认数据库是否可连通
func (c *cSite) Health(ctx context.Context, _ *apicommon.HealthReq) (res *apicommon.HealthRes, err error) {
	if err = g.DB().PingMaster(); err != nil {
		ghttp.RequestFromCtx(ctx).Response.WriteHeader(503)
		return nil, gerror.Wrap(err, "数据库连接失败")
	}
	return &apicommon.HealthRes{Status: "ok"}, nil
}
