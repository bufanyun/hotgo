// Package crons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package crons

import (
	"context"
	"hotgo/internal/library/cron"
	"time"
)

func init() {
	cron.Register(Test)
}

// Test 测试任务（无参数）
var Test = &cTest{name: "test"}

type cTest struct {
	name string
}

func (c *cTest) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cTest) Execute(ctx context.Context) {
	cron.Logger().Infof(ctx, "cron test Execute:%v", time.Now())
}
