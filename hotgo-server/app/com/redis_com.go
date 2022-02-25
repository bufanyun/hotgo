//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package com

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var Redis = new(redis)

type redis struct{}

//
//  @Title  实例化redis
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   name
//  @Return  *gredis.Redis
//
func (component *redis) Instance(name ...string) *gredis.Redis {
	return g.Redis(name...)
}

//
//  @Title  获取
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   key
//  @Return  *gvar.Var
//  @Return  error
//
func (component *redis) Get(ctx context.Context, key string) (*gvar.Var, error) {
	data, err := Redis.Instance().Do(ctx, "GET", key)
	if err != nil {
		err := gerror.New(err.Error())
		return nil, err
	}

	return data, nil
}

//
//  @Title  设置
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   key
//  @Param   value
//  @Param   expire
//  @Return  *gvar.Var
//  @Return  error
//
func (component *redis) Set(ctx context.Context, key string, value string, expire interface{}) (*gvar.Var, error) {

	redisInstance := Redis.Instance()
	response, err := redisInstance.Do(ctx, "SET", key, value)
	if err != nil {
		err := gerror.New(err.Error())
		return nil, err
	}

	exp := gconv.Int(expire)
	// TODO  设置有效期
	if exp > 0 {
		_, err = redisInstance.Do(ctx, "EXPIRE", key, exp)
		if err != nil {
			err := gerror.New(err.Error())
			return nil, err
		}
	}

	return response, nil
}
