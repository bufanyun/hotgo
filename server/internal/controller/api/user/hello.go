// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package user

import (
	"context"
	"fmt"
	"hotgo/api/api/user"
	"hotgo/utility/simple"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *user.HelloReq) (res *user.HelloRes, err error) {
	res = &user.HelloRes{
		Tips: fmt.Sprintf("hello %v, this is the api for %v applications.", req.Name, simple.AppName(ctx)),
	}
	return
}
