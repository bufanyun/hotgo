// Package flashbanner
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package flashbanner

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	_ "hotgo/addons/flashbanner/crons"
	"hotgo/addons/flashbanner/global"
	_ "hotgo/addons/flashbanner/logic"
	_ "hotgo/addons/flashbanner/queues"
	"hotgo/addons/flashbanner/router"
	"hotgo/addons/migrations"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"


)

type module struct {
	skeleton *addons.Skeleton
	ctx      context.Context
	sync.Mutex
}

func init() {
	newModule()
}

func newModule() {
	m := &module{
		skeleton: &addons.Skeleton{
			Label:       `轮播图管理`,
			Name:        `flashbanner`,
			Group:       6,
			Logo:        "",
			Brief:       ``,
			Description: ``,
			Author:      ``,
			Version:     `v1.0.0`, // 当该版本号高于已安装的版本号时，会提示可以更新
		},
		ctx: gctx.New(),
	}
	addons.RegisterModule(m)
}

// Start 启动模块
func (m *module) Start(option *addons.Option) (err error) {
	// 初始化模块
	global.Init(m.ctx, m.skeleton)

	// 注册插件路由
	option.Server.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Addon)
		router.Admin(m.ctx, group)
		router.Api(m.ctx, group)
		router.Home(m.ctx, group)
		router.WebSocket(m.ctx, group)
	})
	return
}

// Stop 停止模块
func (m *module) Stop() (err error) {
	return
}

// Ctx 上下文
func (m *module) Ctx() context.Context {
	return m.ctx
}

// GetSkeleton 获取模块
func (m *module) GetSkeleton() *addons.Skeleton {
	return m.skeleton
}

// Install 安装模块
func (m *module) Install(ctx context.Context) (err error) {
	// 执行数据库安装文件
	// 获取当前目录
	sqlPath := gfile.Pwd() + gfile.Separator + "addons/migrations/flashbanner/install.sql"
	g.Log().Debug(ctx, "安装模块", m.skeleton.Label, "路径", sqlPath)
	result, err := migrations.DoSqlContent(ctx, sqlPath)
	if err != nil {
		g.Log().Error(ctx, "安装模块", m.skeleton.Label, "失败", err)
		return
	}
	g.Log().Debug(ctx, "安装模块", m.skeleton.Label, "成功", result)
	// 复制web目录下的文件到管理后台对应位置
	// 插件的前端配置文件位于插件目录下的web子目录
	sourceWebPath := gfile.Pwd() + gfile.Separator + "addons/" + m.skeleton.Name + "/web/src/views/addons/" + m.skeleton.Name
	targetWebPath := "../web/src/views/addons/" + m.skeleton.Name
	g.Log().Debug(ctx, "复制前端配置文件", "源路径:", sourceWebPath, "目标路径:", targetWebPath)

	// 检查源路径是否存在
	if gfile.Exists(sourceWebPath) {
		err = gfile.CopyDir(sourceWebPath, targetWebPath)
		if err != nil {
			g.Log().Error(ctx, "复制前端配置文件失败:", err)
		} else {
			g.Log().Debug(ctx, "复制前端配置文件成功")
		}
	} else {
		g.Log().Warning(ctx, "前端配置文件源路径不存在:", sourceWebPath)
	}

	// 复制API文件
	sourceApiPath := gfile.Pwd() + gfile.Separator + "addons/" + m.skeleton.Name + "/web/src/api/addons/" + m.skeleton.Name
	targetApiPath := "../web/src/api/addons/" + m.skeleton.Name
	g.Log().Debug(ctx, "复制API文件", "源路径:", sourceApiPath, "目标路径:", targetApiPath)
	if gfile.Exists(sourceApiPath) {
		err = gfile.CopyDir(sourceApiPath, targetApiPath)
		if err != nil {
			g.Log().Error(ctx, "复制API文件失败:", err)
		} else {
			g.Log().Debug(ctx, "复制API文件成功")
		}
	} else {
		g.Log().Warning(ctx, "API文件源路径不存在:", sourceApiPath)
	}
	return
}

// Upgrade 更新模块
func (m *module) Upgrade(ctx context.Context) (err error) {
	// ...
	return
}

// UnInstall 卸载模块
func (m *module) UnInstall(ctx context.Context) (err error) {
	// ...
	// 移除数据库安装文件
	sqlPath := gfile.Pwd() + gfile.Separator + "addons/migrations/flashbanner/uninstall.sql"
	g.Log().Debug(ctx, "卸载模块", m.skeleton.Label, "路径", sqlPath)
	result, err := migrations.DoSqlContent(ctx, sqlPath)
	if err != nil {	
		g.Log().Error(ctx, "卸载模块", m.skeleton.Label, "失败", err)
		return
	}
	g.Log().Debug(ctx, "卸载模块", m.skeleton.Label, "成功", result)
	// 删除前端文件
	targetWebPath := "../web/src/views/addons/" + m.skeleton.Name
	targetApiPath := "../web/src/api/addons/" + m.skeleton.Name

	// 删除配置页面文件
	if gfile.Exists(targetWebPath) {
		err = gfile.Remove(targetWebPath)
		if err != nil {
			g.Log().Warning(ctx, "删除前端配置文件失败:", err)
		} else {
			g.Log().Debug(ctx, "删除前端配置文件成功")
		}
	}

	// 删除API文件
	if gfile.Exists(targetApiPath) {
		err = gfile.Remove(targetApiPath)
		if err != nil {
			g.Log().Warning(ctx, "删除API文件失败:", err)
		} else {
			g.Log().Debug(ctx, "删除API文件成功")
		}
	}

	return
}
