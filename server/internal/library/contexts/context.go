// Package contexts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package contexts

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/model"
)

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改
func Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func SetUser(ctx context.Context, user *model.Identity) {
	Get(ctx).User = user
}

// SetResponse 设置组件响应 用于全局日志使用
func SetResponse(ctx context.Context, response *model.Response) {
	Get(ctx).Response = response
}

// SetModule 设置应用模块
func SetModule(ctx context.Context, module string) {
	Get(ctx).Module = module
}

// SetTakeUpTime 设置请求耗时
func SetTakeUpTime(ctx context.Context, takeUpTime int64) {
	Get(ctx).TakeUpTime = takeUpTime
}

// GetUserId 获取用户ID
func GetUserId(ctx context.Context) int64 {
	user := Get(ctx).User
	if user == nil {
		return 0
	}

	return user.Id
}

// GetRoleId 获取用户角色ID
func GetRoleId(ctx context.Context) int64 {
	user := Get(ctx).User
	if user == nil {
		return 0
	}

	return user.Role
}

// GetRoleKey 获取用户角色唯一编码
func GetRoleKey(ctx context.Context) string {
	user := Get(ctx).User
	if user == nil {
		return ""
	}

	return user.RoleKey
}
