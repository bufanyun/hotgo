//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package com

import (
	"context"
	"fmt"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

// jwt鉴权
type JWT struct{}

var Jwt = new(JWT)

//
//  @Title  为指定用户生成token
//  @Description  主要用于登录成功的jwt鉴权绑定
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   user 用户信息
//  @Param   isRefresh 是否是刷新token
//  @Return  interface{}
//  @Return  error
//
func (component *JWT) GenerateLoginToken(ctx context.Context, user *model.Identity, isRefresh bool) (interface{}, error) {

	jwtVersion, _ := g.Cfg().Get(ctx, "jwt.version", "1.0")
	jwtSign, _ := g.Cfg().Get(ctx, "jwt.sign", "hotGo")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          user.Id,
		"username":    user.Username,
		"realname":    user.Realname,
		"avatar":      user.Avatar,
		"email":       user.Email,
		"mobile":      user.Mobile,
		"last_time":   user.LastTime,
		"last_ip":     user.LastIp,
		"exp":         user.Exp,
		"expires":     user.Expires,
		"app":         user.App,
		"role":        user.Role,
		"visit_count": user.VisitCount,
		"is_refresh":  isRefresh,
		"jwt_version": jwtVersion.String(),
	})

	tokenString, err := token.SignedString(jwtSign.Bytes())
	if err != nil {
		err := gerror.New(err.Error())
		return nil, err
	}

	tokenStringMd5 := gmd5.MustEncryptString(tokenString)

	// TODO  绑定登录token
	cache := Cache.New()
	key := consts.RedisJwtToken + tokenStringMd5

	// TODO  将有效期转为持续时间，单位：秒
	expires, _ := time.ParseDuration(fmt.Sprintf("+%vs", user.Expires))

	err = cache.Set(ctx, key, tokenString, expires)
	if err != nil {
		err := gerror.New(err.Error())
		return nil, err
	}
	_ = cache.Set(ctx, consts.RedisJwtUserBind+user.App+":"+gconv.String(user.Id), key, expires)

	return tokenString, err
}

//
//  @Title  解析token
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   tokenString
//  @Param   secret
//  @Return  jwt.MapClaims
//  @Return  error
//
func (component *JWT) ParseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	if tokenString == "" {
		err := gerror.New("token 为空")
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if token == nil {
		err := gerror.New("token不存在")
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

/**
token有效正确返回用户id
*/
//func(component *JWT) VerifyLoginToken(tokenString string) (uint, err error) {
//	//if tokenString == "" {
//	//	err = gerror.New("token不能为空")
//	//	return 0, err
//	//}
//
//}

//
//  @Title  获取 authorization
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   r
//  @Return  string
//
func (component *JWT) GetAuthorization(r *ghttp.Request) string {

	// TODO  默认从请求头获取
	var authorization = r.Header.Get("Authorization")

	// TODO  如果请求头不存在则从get参数获取
	if authorization == "" {
		return r.Get("authorization").String()
	}

	return gstr.Replace(authorization, "Bearer ", "")
}

/**
清掉所以的相关的redis
*/
func (component *JWT) Layout(adminUserId int, tokenString string) {
	if tokenString == "" {
		return
	}
	//g.Redis().Do("HDEL", "VerifyLoginToken", gmd5.MustEncryptString(tokenString))
	//// 删除
	//g.Redis().Do("HDEL", "VerifyLoginTokenAdminUserId", adminUserId)
}
