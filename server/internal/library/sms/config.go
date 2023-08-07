// Package sms
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sms

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model"
)

var config *model.SmsConfig

func SetConfig(c *model.SmsConfig) {
	config = c
}

func GetConfig() *model.SmsConfig {
	return config
}

func GetModel(ctx context.Context) *gdb.Model {
	return g.Model("sys_sms_log").Ctx(ctx)
}
