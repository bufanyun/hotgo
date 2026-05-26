// Package wechat
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package wechat

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	officialJs "github.com/silenceper/wechat/v2/officialaccount/js"
	"hotgo/internal/model"
)

var (
	config          *model.WechatConfig
	officialAccount *officialaccount.OfficialAccount // 微信公众号
)

func SetConfig(c *model.WechatConfig) {
	config = c
	ctx := gctx.GetInitCtx()
	GetOfficialAccount(ctx, true)
}

func GetConfig() *model.WechatConfig {
	return config
}

// NewOfficialAccount 微信公众号实例
func NewOfficialAccount(ctx context.Context) *officialaccount.OfficialAccount {
	cfg := &offConfig.Config{
		AppID:          config.OfficialAppID,
		AppSecret:      config.OfficialAppSecret,
		Token:          config.OfficialToken,
		EncodingAESKey: config.OfficialEncodingAESKey,
		Cache:          NewCache(ctx),
		UseStableAK:    true,
	}
	wc := wechat.NewWechat()
	return wc.GetOfficialAccount(cfg)
}

// GetOfficialAccount 微信公众号实例
func GetOfficialAccount(ctx context.Context, refresh ...bool) *officialaccount.OfficialAccount {
	isRefresh := false
	if len(refresh) > 0 {
		isRefresh = refresh[0]
	}
	if officialAccount != nil && !isRefresh {
		return officialAccount
	}

	officialAccount = NewOfficialAccount(ctx)
	return officialAccount
}

// GetJsConfig 获取js配置
func GetJsConfig(ctx context.Context, uri string) (config *officialJs.Config, err error) {
	return GetOfficialAccount(ctx).GetJs().GetConfig(uri)
}
