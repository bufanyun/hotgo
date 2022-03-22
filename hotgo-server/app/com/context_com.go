//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package com

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 上下文
var Context = new(comContext)

type comContext struct{}

//
//  @Title  初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   r
//  @Param   customCtx
//
func (component *comContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

//
//  @Title  获得上下文变量，如果没有设置，那么返回nil
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Return  *model.Context
//
func (component *comContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

//
//  @Title  将上下文信息设置到上下文请求中，注意是完整覆盖
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   user
//
func (component *comContext) SetUser(ctx context.Context, user *model.Identity) {
	component.Get(ctx).User = user
}

//
//  @Title  设置组件响应 用于全局日志使用
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   response
//
func (component *comContext) SetResponse(ctx context.Context, response *model.Response) {
	component.Get(ctx).ComResponse = response
}

//
//  @Title  设置应用模块
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   module
//
func (component *comContext) SetModule(ctx context.Context, module string) {
	component.Get(ctx).Module = module
}

//
//  @Title  设置请求耗时
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   module
//
func (component *comContext) SetTakeUpTime(ctx context.Context, takeUpTime int64) {
	component.Get(ctx).TakeUpTime = takeUpTime
}

//
//  @Title  获取用户ID
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Return  int
//
func (component *comContext) GetUserId(ctx context.Context) int64 {
	user := component.Get(ctx).User
	if user == nil {
		return 0
	}

	return user.Id
}
