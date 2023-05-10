// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/commonin"
)

// WechatAuthorizeReq 微信用户授权
type WechatAuthorizeReq struct {
	g.Meta `path:"/wechat/authorize" method:"get" tags:"微信" summary:"微信用户授权"`
	commonin.WechatAuthorizeInp
}
type WechatAuthorizeRes struct {
	*commonin.WechatAuthorizeModel
}

// WechatAuthorizeCallReq 微信用户授权回调
type WechatAuthorizeCallReq struct {
	g.Meta `path:"/wechat/authorizeCall" method:"get" tags:"微信" summary:"微信用户授权"`
	commonin.WechatAuthorizeCallInp
}
type WechatAuthorizeCallRes struct {
	*commonin.WechatAuthorizeCallModel
}
