package wechat

import "hotgo/internal/model"

var config *model.WechatConfig

func SetConfig(c *model.WechatConfig) {
	config = c
}

func GetConfig() *model.WechatConfig {
	return config
}
