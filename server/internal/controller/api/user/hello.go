// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/api/user"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *user.HelloReq) (res *user.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World api member!")
	return
}
