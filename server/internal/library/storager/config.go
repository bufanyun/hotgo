// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model"
)

var config *model.UploadConfig

func SetConfig(c *model.UploadConfig) {
	config = c
}

func GetConfig() *model.UploadConfig {
	return config
}

func GetModel(ctx context.Context) *gdb.Model {
	return g.Model("sys_attachment").Ctx(ctx)
}
