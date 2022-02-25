//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package middleware

import (
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/utils"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

//
//  @Title  接口中间件
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   r
//
func (s *sMiddleware) ApiAuth(r *ghttp.Request) {

	var (
		ctx           = r.Context()
		user          = new(model.Identity)
		authorization = com.Jwt.GetAuthorization(r)
	)

	// TODO  替换掉模块前缀
	routerPrefix, _ := g.Cfg().Get(ctx, "router.api.prefix", "/api")
	path := gstr.Replace(r.URL.Path, routerPrefix.String(), "", 1)

	/// TODO  不需要验证登录的路由地址
	if utils.Auth.IsExceptLogin(ctx, path) {
		r.Middleware.Next()
		return
	}

	if authorization == "" {
		com.Response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "请先登录！")
		return
	}

	// TODO  获取jwtToken
	jwtToken := consts.RedisJwtToken + gmd5.MustEncryptString(authorization)
	jwtSign, _ := g.Cfg().Get(ctx, "jwt.sign", "hotgo")

	data, ParseErr := com.Jwt.ParseToken(authorization, jwtSign.Bytes())
	if ParseErr != nil {
		com.Response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "token不正确或已过期！", ParseErr.Error())
	}

	parseErr := gconv.Struct(data, &user)
	if parseErr != nil {
		com.Response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "登录信息解析异常，请重新登录！", parseErr.Error())
	}

	// TODO  判断token跟redis的缓存的token是否一样
	cache := com.Cache.New()
	isContains, containsErr := cache.Contains(ctx, jwtToken)
	if containsErr != nil {
		com.Response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "token无效！", containsErr.Error())
		return
	}
	if !isContains {
		com.Response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "token已过期！")
		return
	}

	// TODO  是否开启多端登录
	if multiPort, _ := g.Cfg().Get(ctx, "jwt.multiPort", true); !multiPort.Bool() {
		key := consts.RedisJwtUserBind + consts.AppApi + ":" + gconv.String(user.Id)
		originJwtToken, originErr := cache.Get(ctx, key)
		if originErr != nil {
			com.Response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "信息异常，请重新登录！", originErr.Error())
			return
		}

		if originJwtToken == nil || originJwtToken.IsEmpty() {
			com.Response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "token已过期！")
			return
		}

		if jwtToken != originJwtToken.String() {
			com.Response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "账号已在其他地方登录！")
			return
		}
	}

	// TODO  保存到上下文
	customCtx := &model.Context{}
	if user != nil {
		customCtx.User = &model.Identity{
			Id:         user.Id,
			Username:   user.Username,
			Realname:   user.Realname,
			Avatar:     user.Avatar,
			Email:      user.Email,
			Mobile:     user.Mobile,
			VisitCount: user.VisitCount,
			LastTime:   user.LastTime,
			LastIp:     user.LastIp,
			Role:       user.Role,
			Exp:        user.Exp,
			Expires:    user.Expires,
			App:        user.App,
		}
	}
	com.Context.SetUser(ctx, customCtx.User)
	com.Context.SetModule(ctx, consts.AppApi)

	//// TODO  验证路由访问权限
	//verify := adminService.Role.Verify(ctx, customCtx.User.Id, path)
	//if !verify {
	//	com.Response.JsonExit(r, gcode.CodeSecurityReason.Code(), "你没有访问权限！")
	//	return
	//}

	r.Middleware.Next()
}
