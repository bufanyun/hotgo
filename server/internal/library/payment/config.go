// Package payment
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package payment

import "hotgo/internal/model"

var config *model.PayConfig

func SetConfig(c *model.PayConfig) {
	config = c
}

func GetConfig() *model.PayConfig {
	return config
}
