// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/common"
	"hotgo/internal/consts"
	"hotgo/internal/library/captcha"
	"hotgo/internal/library/token"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var Site = cSite{}

type cSite struct{}

// Ping ping
func (c *cSite) Ping(_ context.Context, _ *common.SitePingReq) (res *common.SitePingRes, err error) {
	return
}

// Config 获取配置
func (c *cSite) Config(ctx context.Context, _ *common.SiteConfigReq) (res *common.SiteConfigRes, err error) {
	request := ghttp.RequestFromCtx(ctx)
	res = &common.SiteConfigRes{
		Version: consts.VersionApp,
		WsAddr:  c.getWsAddr(ctx, request),
		Domain:  c.getDomain(ctx, request),
	}
	return
}

func (c *cSite) getWsAddr(ctx context.Context, request *ghttp.Request) string {
	// 如果是本地IP访问，则认为是调试模式，走实际请求地址，否则走配置中的地址
	ip := ghttp.RequestFromCtx(ctx).GetHeader("hostname")
	if validate.IsLocalIPAddr(ip) {
		return "ws://" + ip + ":" + gstr.StrEx(request.Host, ":") + "/socket"
	}

	basic, err := service.SysConfig().GetBasic(ctx)
	if err != nil || basic == nil {
		return ""
	}
	return basic.WsAddr
}

func (c *cSite) getDomain(ctx context.Context, request *ghttp.Request) string {
	// 如果是本地IP访问，则认为是调试模式，走实际请求地址，否则走配置中的地址
	ip := request.GetHeader("hostname")
	if validate.IsLocalIPAddr(ip) {
		return "http://" + ip + ":" + gstr.StrEx(request.Host, ":")
	}

	basic, err := service.SysConfig().GetBasic(ctx)
	if err != nil || basic == nil {
		return ""
	}
	return basic.Domain
}

// LoginConfig 登录配置
func (c *cSite) LoginConfig(ctx context.Context, _ *common.SiteLoginConfigReq) (res *common.SiteLoginConfigRes, err error) {
	res = new(common.SiteLoginConfigRes)
	login, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	res.LoginConfig = login
	return
}

// Captcha 登录验证码
func (c *cSite) Captcha(ctx context.Context, _ *common.LoginCaptchaReq) (res *common.LoginCaptchaRes, err error) {
	cid, base64 := captcha.Generate(ctx)
	res = &common.LoginCaptchaRes{Cid: cid, Base64: base64}
	return
}

// Register 账号注册
func (c *cSite) Register(ctx context.Context, req *common.RegisterReq) (res *common.RegisterRes, err error) {
	err = service.AdminSite().Register(ctx, &req.RegisterInp)
	return
}

// AccountLogin 账号登录
func (c *cSite) AccountLogin(ctx context.Context, req *common.AccountLoginReq) (res *common.AccountLoginRes, err error) {
	login, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	if !req.IsLock && login.CaptchaSwitch == 1 {
		// 校验 验证码
		if !captcha.Verify(req.Cid, req.Code) {
			err = gerror.New("验证码错误")
			return
		}
	}

	model, err := service.AdminSite().AccountLogin(ctx, &req.AccountLoginInp)
	if err != nil {
		return
	}

	err = gconv.Scan(model, &res)
	return
}

// MobileLogin 手机号登录
func (c *cSite) MobileLogin(ctx context.Context, req *common.MobileLoginReq) (res *common.MobileLoginRes, err error) {
	model, err := service.AdminSite().MobileLogin(ctx, &req.MobileLoginInp)
	if err != nil {
		return
	}

	err = gconv.Scan(model, &res)
	return
}

// Logout 注销登录
func (c *cSite) Logout(ctx context.Context, _ *common.LoginLogoutReq) (res *common.LoginLogoutRes, err error) {
	err = token.Logout(ghttp.RequestFromCtx(ctx))
	return
}
