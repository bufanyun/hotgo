// Package contexts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package contexts

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
	c := Get(ctx)
	if c == nil {
		g.Log().Warningf(ctx, "contexts.SetUser,  c == nil ")
		return
	}
	c.User = user
}

// SetResponse 设置组件响应 用于访问日志使用
func SetResponse(ctx context.Context, response *model.Response) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warningf(ctx, "contexts.SetResponse,  c == nil ")
		return
	}
	c.Response = response
}

// SetModule 设置应用模块
func SetModule(ctx context.Context, module string) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warningf(ctx, "contexts.SetModule,  c == nil ")
		return
	}
	c.Module = module
}

// SetTakeUpTime 设置请求耗时
func SetTakeUpTime(ctx context.Context, takeUpTime int64) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warningf(ctx, "contexts.SetTakeUpTime,  c == nil ")
		return
	}
	c.TakeUpTime = takeUpTime
}

// GetUser 获取用户信息
func GetUser(ctx context.Context) *model.Identity {
	c := Get(ctx)
	if c == nil {
		return nil
	}

	return c.User
}

// GetUserId 获取用户ID
func GetUserId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.Id
}

// GetRoleId 获取用户角色ID
func GetRoleId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.RoleId
}

// GetRoleKey 获取用户角色唯一编码
func GetRoleKey(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.RoleKey
}

// SetAddonName 设置插件信息
func SetAddonName(ctx context.Context, name string) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warningf(ctx, "contexts.SetAddonName,  c == nil ")
		return
	}
	Get(ctx).AddonName = name
}

// IsAddonRequest 是否为插件模块请求
func IsAddonRequest(ctx context.Context) bool {
	c := Get(ctx)
	if c == nil {
		return false
	}
	return GetAddonName(ctx) != ""
}

// GetAddonName 获取插件信息
func GetAddonName(ctx context.Context) string {
	c := Get(ctx)
	if c == nil {
		return ""
	}
	return Get(ctx).AddonName
}
