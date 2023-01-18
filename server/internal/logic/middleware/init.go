// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package middleware

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/jwt"
	"hotgo/internal/library/response"
	"hotgo/internal/model"
	"hotgo/internal/service"
	"net/http"
	"strings"
)

type sMiddleware struct {
	LoginUrl      string // 登录路由地址
	DemoWhiteList g.Map  // 演示模式放行的路由白名單
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/common",
		DemoWhiteList: g.Map{
			"/admin/site/login": struct{}{}, // 后台登录
		},
	}
}

// Ctx 初始化请求上下文
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	contexts.Init(r, &model.Context{
		Data:   make(g.Map),
		Module: getModule(r.URL.Path),
	})

	r.Middleware.Next()
}

func getModule(path string) (module string) {
	slice := strings.Split(path, "/")
	if len(slice) < 2 {
		module = consts.AppDefault
		return
	}

	if slice[1] == "" {
		module = consts.AppDefault
		return
	}

	return slice[1]
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// DemoLimit 演示系統操作限制
func (s *sMiddleware) DemoLimit(r *ghttp.Request) {
	isDemo := g.Cfg().MustGet(r.Context(), "hotgo.isDemo", false)
	if !isDemo.Bool() {
		r.Middleware.Next()
		return
	}

	if r.Method == http.MethodPost {
		if _, ok := s.DemoWhiteList[r.URL.Path]; ok {
			r.Middleware.Next()
			return
		}
		response.JsonExit(r, gcode.CodeNotSupported.Code(), "演示系统禁止操作！")
		return
	}

	r.Middleware.Next()
}

// inspectAuth 检查并完成身份认证
func inspectAuth(r *ghttp.Request, appName string) error {
	var (
		ctx           = r.Context()
		user          = new(model.Identity)
		authorization = jwt.GetAuthorization(r)
		c             = cache.New()
		customCtx     = &model.Context{}
	)

	if authorization == "" {
		return gerror.New("请先登录！")
	}

	// 获取jwtToken
	jwtToken := consts.RedisJwtToken + gmd5.MustEncryptString(authorization)
	jwtSign := g.Cfg().MustGet(ctx, "jwt.sign", "hotgo")

	data, ParseErr := jwt.ParseToken(authorization, jwtSign.Bytes())
	if ParseErr != nil {
		return gerror.Newf("token不正确或已过期! err :%+v", ParseErr.Error())
	}

	parseErr := gconv.Struct(data, &user)
	if parseErr != nil {
		return gerror.Newf("登录信息解析异常，请重新登录！ err :%+v", ParseErr.Error())
	}

	// 判断token跟redis的缓存的token是否一样
	isContains, containsErr := c.Contains(ctx, jwtToken)
	if containsErr != nil {
		return gerror.Newf("token无效！ err :%+v", ParseErr.Error())
	}
	if !isContains {
		return gerror.New("token已过期")
	}

	// 是否开启多端登录
	if multiPort := g.Cfg().MustGet(ctx, "jwt.multiPort", true); !multiPort.Bool() {
		key := consts.RedisJwtUserBind + appName + ":" + gconv.String(user.Id)
		originJwtToken, originErr := c.Get(ctx, key)
		if originErr != nil {
			return gerror.Newf("信息异常，请重新登录！ err :%+v", originErr.Error())
		}

		if originJwtToken == nil || originJwtToken.IsEmpty() {
			return gerror.New("token已过期！")
		}

		if jwtToken != originJwtToken.String() {
			return gerror.New("账号已在其他地方登录！")
		}
	}

	// 保存到上下文
	if user != nil {
		customCtx.User = &model.Identity{
			Id:         user.Id,
			Pid:        user.Pid,
			DeptId:     user.DeptId,
			RoleId:     user.RoleId,
			RoleKey:    user.RoleKey,
			Username:   user.Username,
			RealName:   user.RealName,
			Avatar:     user.Avatar,
			Email:      user.Email,
			Mobile:     user.Mobile,
			VisitCount: user.VisitCount,
			LastTime:   user.LastTime,
			LastIp:     user.LastIp,
			Exp:        user.Exp,
			Expires:    user.Expires,
			App:        user.App,
		}
	}
	contexts.SetUser(ctx, customCtx.User)

	return nil
}
