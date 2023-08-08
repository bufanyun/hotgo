// Package wechat
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package wechat

import "hotgo/internal/model"

var config *model.WechatConfig

func SetConfig(c *model.WechatConfig) {
	config = c
}

func GetConfig() *model.WechatConfig {
	return config
}
