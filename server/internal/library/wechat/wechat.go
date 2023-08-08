// Package wechat
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package wechat

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	officialJs "github.com/silenceper/wechat/v2/officialaccount/js"
	officialOauth "github.com/silenceper/wechat/v2/officialaccount/oauth"
	"github.com/silenceper/wechat/v2/openplatform"
	openConfig "github.com/silenceper/wechat/v2/openplatform/config"
	"hotgo/internal/consts"
)

// NewOfficialAccount 微信公众号实例
func NewOfficialAccount(ctx context.Context) *officialaccount.OfficialAccount {
	cfg := &offConfig.Config{
		AppID:          config.OfficialAppID,
		AppSecret:      config.OfficialAppSecret,
		Token:          config.OfficialToken,
		EncodingAESKey: config.OfficialEncodingAESKey,
		Cache:          NewCache(ctx),
	}
	return wechat.NewWechat().GetOfficialAccount(cfg)
}

// NewOpenPlatform 开放平台实例
func NewOpenPlatform(ctx context.Context) *openplatform.OpenPlatform {
	cfg := &openConfig.Config{
		AppID:          config.OpenAppID,
		AppSecret:      config.OpenAppSecret,
		Token:          config.OpenToken,
		EncodingAESKey: config.OpenEncodingAESKey,
		Cache:          NewCache(ctx),
	}
	return wechat.NewWechat().GetOpenPlatform(cfg)
}

// GetOpenOauthURL 代第三方公众号 - 获取网页授权地址
func GetOpenOauthURL(ctx context.Context, redirectURI, scope, state string) (location string, err error) {
	op := NewOpenPlatform(ctx)
	appid := config.OfficialAppID // 公众号appid
	oauth := op.GetOfficialAccount(appid).PlatformOauth()
	if scope == "" {
		scope = consts.WechatScopeBase
	}
	location, err = oauth.GetRedirectURL(redirectURI, scope, state, appid)
	return
}

// GetOpenUserAccessToken 代第三方公众号 - 通过网页授权的code 换取access_token
func GetOpenUserAccessToken(ctx context.Context, code string) (accessToken officialOauth.ResAccessToken, err error) {
	op := NewOpenPlatform(ctx)
	appid := config.OfficialAppID // 公众号appid
	officialAccount := op.GetOfficialAccount(appid)
	componentAccessToken, err := op.GetComponentAccessToken()
	if err != nil {
		return
	}

	accessToken, err = officialAccount.PlatformOauth().GetUserAccessToken(code, appid, componentAccessToken)
	if err != nil {
		return
	}

	if accessToken.ErrCode > 0 {
		err = gerror.Newf("GetOpenUserAccessToken err:%+v", accessToken.ErrMsg)
		return
	}
	return
}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx context.Context, token officialOauth.ResAccessToken) (info officialOauth.UserInfo, err error) {
	oauth := NewOfficialAccount(ctx).GetOauth()
	info, err = oauth.GetUserInfo(token.AccessToken, token.OpenID, "")
	return
}

// GetOauthURL 获取网页授权地址
func GetOauthURL(ctx context.Context, redirectURI, scope, state string) (location string, err error) {
	oauth := NewOfficialAccount(ctx).GetOauth()
	location, err = oauth.GetRedirectURL(redirectURI, scope, state)
	return
}

// GetUserAccessToken 通过网页授权的code 换取access_token
func GetUserAccessToken(ctx context.Context, code string) (accessToken officialOauth.ResAccessToken, err error) {
	oauth := NewOfficialAccount(ctx).GetOauth()
	accessToken, err = oauth.GetUserAccessToken(code)
	if err != nil {
		return
	}

	if accessToken.ErrCode > 0 {
		err = gerror.Newf("GetUserAccessToken err:%+v", accessToken.ErrMsg)
		return
	}
	return
}

// GetJsConfig 获取js配置
func GetJsConfig(ctx context.Context, uri string) (config *officialJs.Config, err error) {
	return NewOfficialAccount(ctx).GetJs().GetConfig(uri)
}
