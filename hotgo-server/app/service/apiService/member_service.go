//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package apiService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var Member = new(member)

type member struct {}

func (service *member) Test() (ctx context.Context) {
	g.Log().Print(ctx, "apiService--WithMember--test...")

	//g.Log().Print(ctx, "api调用：" ,  service.App.Admin.Member.Test())
	return
}