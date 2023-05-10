package payment

import "hotgo/internal/model"

var config *model.PayConfig

func SetConfig(c *model.PayConfig) {
	config = c
}

func GetConfig() *model.PayConfig {
	return config
}
