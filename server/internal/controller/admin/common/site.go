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
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var Site = cSite{}

type cSite struct{}

// Ping ping
func (c *cSite) Ping(ctx context.Context, req *common.SitePingReq) (res *common.SitePingRes, err error) {
	return
}

// Config 获取配置
func (c *cSite) Config(ctx context.Context, req *common.SiteConfigReq) (res *common.SiteConfigRes, err error) {
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

// Captcha 登录验证码
func (c *cSite) Captcha(ctx context.Context, req *common.LoginCaptchaReq) (res *common.LoginCaptchaRes, err error) {
	cid, base64 := captcha.Generate(ctx)
	res = &common.LoginCaptchaRes{Cid: cid, Base64: base64}
	return
}

// Login 提交登录
func (c *cSite) Login(ctx context.Context, req *common.LoginReq) (res *common.LoginRes, err error) {
	var in adminin.MemberLoginInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	defer func() {
		var response = new(adminin.MemberLoginModel)
		if res != nil && res.MemberLoginModel != nil {
			response = res.MemberLoginModel
		}
		service.SysLoginLog().Push(ctx, sysin.LoginLogPushInp{Input: in, Response: response, Err: err})
	}()

	// 校验 验证码
	if !req.IsLock && !captcha.Verify(req.Cid, req.Code) {
		err = gerror.New("验证码错误")
		return
	}

	model, err := service.AdminMember().Login(ctx, in)
	if err != nil {
		return
	}

	err = gconv.Scan(model, &res)
	return
}

// Logout 注销登录
func (c *cSite) Logout(ctx context.Context, req *common.LoginLogoutReq) (res *common.LoginLogoutRes, err error) {
	err = token.Logout(ghttp.RequestFromCtx(ctx))
	return
}
