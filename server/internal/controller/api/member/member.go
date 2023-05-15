// Package member
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package member

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/api/member"
)

var (
	Member = cMember{}
)

type cMember struct{}

func (c *cMember) GetIdByCode(ctx context.Context, _ *member.GetIdByCodeReq) (res *member.GetIdByCodeRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World api member!")
	return
}
