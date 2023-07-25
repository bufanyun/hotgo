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
