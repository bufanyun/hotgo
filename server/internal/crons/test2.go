// Package crons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package crons

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"time"
)

func init() {
	cronList = append(cronList, Test2)
}

// Test2 测试2任务（带参数）
var Test2 = &cTest2{name: "test2"}

type cTest2 struct {
	name string
}

func (c *cTest2) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cTest2) Execute(ctx context.Context) {
	args, ok := ctx.Value(consts.ContextKeyCronArgs).([]string)
	if !ok {
		g.Log().Warning(ctx, "参数解析失败!")
		return
	}
	if len(args) != 3 {
		g.Log().Warning(ctx, "test2 传入参数不正确!")
		return
	}

	var (
		name = args[0]
		age  = args[1]
		msg  = args[2]
	)

	g.Log().Infof(ctx, "cron test2 Execute:%v, name:%v, age:%v, msg:%v", time.Now(), name, age, msg)
}
