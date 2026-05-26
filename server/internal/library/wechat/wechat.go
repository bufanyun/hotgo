// Package wechat
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package wechat

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	officialOauth "github.com/silenceper/wechat/v2/officialaccount/oauth"
)

// GetUserInfo 获取用户信息
func GetUserInfo(ctx context.Context, token officialOauth.ResAccessToken) (info officialOauth.UserInfo, err error) {
	oauth := GetOfficialAccount(ctx).GetOauth()
	info, err = oauth.GetUserInfo(token.AccessToken, token.OpenID, "")
	return
}

// GetOauthURL 获取网页授权地址
func GetOauthURL(ctx context.Context, redirectURI, scope, state string) (location string, err error) {
	oauth := GetOfficialAccount(ctx).GetOauth()
	location, err = oauth.GetRedirectURL(redirectURI, scope, state)
	return
}

// GetUserAccessToken 通过网页授权的code 换取access_token
func GetUserAccessToken(ctx context.Context, code string) (accessToken officialOauth.ResAccessToken, err error) {
	oauth := GetOfficialAccount(ctx).GetOauth()
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
