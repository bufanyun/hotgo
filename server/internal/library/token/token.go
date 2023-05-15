package token

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model"
	"hotgo/utility/simple"
	"time"
)

type Claims struct {
	*model.Identity
	jwt.RegisteredClaims
}

type Token struct {
	ExpireAt     int64 `json:"exp"` // token过期时间
	RefreshAt    int64 `json:"ra"`  // 刷新时间
	RefreshCount int64 `json:"rc"`  // 刷新次数
}

var (
	config          *model.TokenConfig
	errorLogin      = gerror.New("登录身份已失效，请重新登录！")
	errorMultiLogin = gerror.New("账号存在异地登录，如非本人操作请及时修改登录密码！")
)

func SetConfig(c *model.TokenConfig) {
	config = c
}

func GetConfig() *model.TokenConfig {
	return config
}

// Login 登录
func Login(ctx context.Context, user *model.Identity) (string, int64, error) {
	claims := Claims{
		user,
		jwt.RegisteredClaims{},
	}

	header, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", 0, err
	}

	var (
		now = gtime.Now()
		// 认证key
		authKey = GetAuthKey(header)
		// 登录token
		tokenKey = GetTokenKey(user.App, authKey)
		// 身份绑定
		bindKey = GetBindKey(user.App, user.Id)
		// 有效时长
		duration = time.Second * gconv.Duration(config.Expires)
	)

	token := &Token{
		ExpireAt:     now.Unix() + config.Expires,
		RefreshAt:    now.Unix(),
		RefreshCount: 0,
	}

	if err = cache.Instance().Set(ctx, tokenKey, token, duration); err != nil {
		return "", 0, err
	}

	if err = cache.Instance().Set(ctx, bindKey, tokenKey, duration); err != nil {
		return "", 0, err
	}

	return header, config.Expires, nil
}

// Logout 注销登录
func Logout(r *ghttp.Request) (err error) {
	var (
		ctx    = r.Context()
		header = GetAuthorization(r)
	)

	if header == "" {
		err = errorLogin
		return
	}

	claims, err := parseToken(ctx, header)
	if err != nil {
		return
	}

	var (
		// 认证key
		authKey = GetAuthKey(header)
		// 登录token
		tokenKey = GetTokenKey(contexts.GetModule(ctx), authKey)
		// 身份绑定
		bindKey = GetBindKey(contexts.GetModule(ctx), claims.Id)
	)

	// 删除token
	if _, err = cache.Instance().Remove(ctx, tokenKey); err != nil {
		return
	}

	if !config.MultiLogin {
		if _, err = cache.Instance().Remove(ctx, bindKey); err != nil {
			return
		}
	}
	return
}

// ParseLoginUser 解析登录用户信息
func ParseLoginUser(r *ghttp.Request) (user *model.Identity, err error) {
	var (
		ctx    = r.Context()
		header = GetAuthorization(r)
	)

	if header == "" {
		err = errorLogin
		return
	}

	claims, err := parseToken(ctx, header)
	if err != nil {
		return
	}

	var (
		authKey = GetAuthKey(header)
		// 登录token
		tokenKey = GetTokenKey(claims.App, authKey)
		// 身份绑定
		bindKey = GetBindKey(claims.App, claims.Id)
	)

	// 检查token是否存在
	tk, err := cache.Instance().Get(ctx, tokenKey)
	if err != nil {
		g.Log().Debugf(ctx, "get tokenKey err:%+v", err)
		err = errorLogin
		return
	}

	if tk.IsEmpty() {
		g.Log().Debug(ctx, "token isEmpty")
		err = errorLogin
		return
	}

	var token *Token
	if err = tk.Scan(&token); err != nil {
		g.Log().Debugf(ctx, "token scan err:%+v", err)
		err = errorLogin
		return
	}

	if token == nil {
		g.Log().Debugf(ctx, "token = nil")
		err = errorLogin
		return
	}

	now := gtime.Now()
	if token.ExpireAt < now.Unix() {
		g.Log().Debugf(ctx, "token expired.")
		err = errorLogin
		return
	}

	// 是否允许多端登录
	if !config.MultiLogin {
		origin, err := cache.Instance().Get(ctx, bindKey)
		if err != nil {
			g.Log().Debugf(ctx, "bindKey get err:%+v", err)
			err = errorLogin
			return nil, err
		}

		if origin == nil || origin.IsEmpty() {
			g.Log().Debug(ctx, "bindKey isEmpty")
			err = errorLogin
			return nil, err
		}

		if tokenKey != origin.String() {
			g.Log().Debugf(ctx, "bindKey offsite login tokenKey:%v, origin:%v", tokenKey, origin.String())
			err = errorMultiLogin
			return nil, err
		}
	}

	// 自动刷新token有效期
	refreshToken := func() {
		// 未开启自动刷新
		if !config.AutoRefresh {
			return
		}

		// 刷新次数已达上限
		if config.MaxRefreshTimes != -1 && token.RefreshCount >= config.MaxRefreshTimes {
			return
		}

		// 未达到刷新间隔
		if gtime.New(token.RefreshAt).Unix()+config.RefreshInterval > now.Unix() {
			return
		}

		// 刷新有效期
		token.ExpireAt = now.Unix() + config.Expires
		token.RefreshAt = now.Unix()
		token.RefreshCount += 1

		duration := time.Second * gconv.Duration(config.Expires)

		if err = cache.Instance().Set(ctx, tokenKey, token, duration); err != nil {
			return
		}

		if err = cache.Instance().Set(ctx, bindKey, tokenKey, duration); err != nil {
			return
		}
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		refreshToken()
	})

	user = new(model.Identity)
	user = claims.Identity
	return
}

// parseToken 解析jwt令牌
func parseToken(ctx context.Context, header string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(header, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		g.Log().Debugf(ctx, "parseToken err:%+v", err)
		return nil, err
	}

	if !token.Valid {
		return nil, errorLogin
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errorLogin
	}

	return claims, nil
}

// GetAuthorization 获取authorization
func GetAuthorization(r *ghttp.Request) string {
	// 默认从请求头获取
	var authorization = r.Header.Get("Authorization")

	// 如果请求头不存在则从get参数获取
	if authorization == "" {
		return r.Get("authorization").String()
	}

	return gstr.Replace(authorization, "Bearer ", "")
}

// GetAuthKey 认证key
func GetAuthKey(token string) string {
	return gmd5.MustEncryptString(token)
}

// GetTokenKey 令牌缓存key
func GetTokenKey(appName, authKey string) string {
	return fmt.Sprintf("%v:%v:%v", consts.CacheToken, appName, authKey)
}

// GetBindKey 令牌身份绑定key
func GetBindKey(appName string, userId int64) string {
	return fmt.Sprintf("%v:%v:%v", consts.CacheTokenBind, appName, userId)
}
