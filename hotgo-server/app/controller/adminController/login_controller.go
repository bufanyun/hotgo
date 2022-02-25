//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminController

import (
	"context"
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/adminService"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

var Login = login{}

type login struct{}

//
//  @Title  登录验证码
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *login) Captcha(ctx context.Context, req *adminForm.LoginCaptchaReq) (res *adminForm.LoginCaptchaRes, err error) {

	// TODO  获取生成的验证码图片
	Cid, Base64 := com.Captcha.GetVerifyImgString(ctx)
	res = &adminForm.LoginCaptchaRes{Cid: Cid, Base64: Base64}

	return
}

//
//  @Title  提交登录
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *login) Sign(ctx context.Context, req *adminForm.LoginReq) (res *adminForm.LoginRes, err error) {

	//// 校验 验证码
	//if !com.Captcha.VerifyString(req.Cid, req.Code) {
	//	err = gerror.New("验证码错误")
	//	return
	//}

	var in input.AdminMemberLoginSignInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	model, err := adminService.Member.Login(ctx, in)
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(model, &res); err != nil {
		return nil, err
	}
	return
}

//
//  @Title  注销登录
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *login) Logout(ctx context.Context, req *adminForm.LoginLogoutReq) (res *adminForm.LoginLogoutRes, err error) {

	var authorization = com.Jwt.GetAuthorization(com.Context.Get(ctx).Request)

	// TODO  获取jwtToken
	jwtToken := consts.RedisJwtToken + gmd5.MustEncryptString(authorization)
	if len(jwtToken) == 0 {
		err = gerror.New("当前用户未登录！")
		return res, err
	}

	// TODO  删除登录token
	cache := com.Cache.New()
	_, err = cache.Remove(ctx, jwtToken)
	if err != nil {
		return res, err
	}

	return
}
