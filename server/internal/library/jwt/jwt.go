// Package jwt
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package jwt

import (
	"context"
	"fmt"
	j "github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/model"
	"time"
)

// GenerateLoginToken 为指定用户生成token
func GenerateLoginToken(ctx context.Context, user *model.Identity, isRefresh bool) (string, error) {
	var (
		jwtVersion = g.Cfg().MustGet(ctx, "jwt.version", "1.0")
		jwtSign    = g.Cfg().MustGet(ctx, "jwt.sign", "hotGo")
		token      = j.NewWithClaims(j.SigningMethodHS256, j.MapClaims{
			"id":         user.Id,
			"pid":        user.Pid,
			"deptId":     user.DeptId,
			"roleId":     user.RoleId,
			"roleKey":    user.RoleKey,
			"username":   user.Username,
			"realName":   user.RealName,
			"avatar":     user.Avatar,
			"email":      user.Email,
			"mobile":     user.Mobile,
			"lastTime":   user.LastTime,
			"lastIp":     user.LastIp,
			"exp":        user.Exp,
			"expires":    user.Expires,
			"app":        user.App,
			"visitCount": user.VisitCount,
			"isRefresh":  isRefresh,
			"jwtVersion": jwtVersion.String(),
		})
	)

	tokenString, err := token.SignedString(jwtSign.Bytes())
	if err != nil {
		return "", err
	}

	var (
		tokenStringMd5 = gmd5.MustEncryptString(tokenString)
		// 绑定登录token
		c   = cache.New()
		key = consts.RedisJwtToken + tokenStringMd5
		// 将有效期转为持续时间，单位：秒
		expires, _ = time.ParseDuration(fmt.Sprintf("+%vs", user.Expires))
	)

	err = c.Set(ctx, key, tokenString, expires)
	if err != nil {
		return "", err
	}

	err = c.Set(ctx, consts.RedisJwtUserBind+user.App+":"+gconv.String(user.Id), key, expires)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// ParseToken 解析token
func ParseToken(tokenString string, secret []byte) (j.MapClaims, error) {
	if tokenString == "" {
		err := gerror.New("token 为空")
		return nil, err
	}
	token, err := j.Parse(tokenString, func(token *j.Token) (interface{}, error) {
		if _, ok := token.Method.(*j.SigningMethodHMAC); !ok {
			return nil, gerror.Newf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if token == nil {
		err := gerror.New("token不存在")
		return nil, err
	}

	if claims, ok := token.Claims.(j.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
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
