// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/common"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/captcha"
	"hotgo/internal/library/jwt"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
)

var Site = cSite{}

type cSite struct{}

// Ping ping
func (c *cSite) Ping(ctx context.Context, req *common.SitePingReq) (res *common.SitePingRes, err error) {
	return
}

// Config 获取配置
func (c *cSite) Config(ctx context.Context, req *common.SiteConfigReq) (res *common.SiteConfigRes, err error) {
	res = &common.SiteConfigRes{
		Version: consts.VersionApp,
		WsAddr:  c.getWsAddr(ctx),
	}
	return
}

func (c *cSite) getWsAddr(ctx context.Context) string {
	ws := g.Cfg().MustGet(ctx, "hotgo.wsAddr", "ws://127.0.0.1:8000/socket")
	return ws.String()

	//// nginx负载均衡部署
	//// 如果是IP访问，则认为是调试模式，走配置中的ws地址，否则走实际请求中的域名+协议
	//if !validate.IsDNSName(ghttp.RequestFromCtx(ctx).Host) {
	//	ws := g.Cfg().MustGet(ctx, "hotgo.wsAddr", "ws://127.0.0.1:8000/socket")
	//	return ws.String()
	//}
	//
	//if !validate.IsHTTPS(ctx) {
	//	return fmt.Sprintf("ws://%s/socket", url.GetDomain(ctx))
	//}
	//
	//return fmt.Sprintf("wss://%s/socket", url.GetDomain(ctx))
}

// Captcha 登录验证码
func (c *cSite) Captcha(ctx context.Context, req *common.LoginCaptchaReq) (res *common.LoginCaptchaRes, err error) {

	// 获取生成的验证码图片
	Cid, Base64 := captcha.GetVerifyImgString(ctx)
	res = &common.LoginCaptchaRes{Cid: Cid, Base64: Base64}

	return
}

// Login 提交登录
func (c *cSite) Login(ctx context.Context, req *common.LoginReq) (res *common.LoginRes, err error) {

	//// 校验 验证码
	//if !captcha.VerifyString(req.Cid, req.Code) {
	//	err = gerror.New("验证码错误")
	//	return
	//}
	//
	var in adminin.MemberLoginInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	model, err := service.AdminMember().Login(ctx, in)
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(model, &res); err != nil {
		return nil, err
	}
	return
}

// Logout 注销登录
func (c *cSite) Logout(ctx context.Context, req *common.LoginLogoutReq) (res *common.LoginLogoutRes, err error) {

	var authorization = jwt.GetAuthorization(ghttp.RequestFromCtx(ctx))

	// 获取jwtToken
	jwtToken := consts.RedisJwtToken + gmd5.MustEncryptString(authorization)
	if len(jwtToken) == 0 {
		err = gerror.New("当前用户未登录！")
		return res, err
	}

	// 删除登录token
	ca := cache.New()
	_, err = ca.Remove(ctx, jwtToken)
	if err != nil {
		return res, err
	}

	return
}
