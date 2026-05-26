// Package addons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package addons

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"hotgo/internal/model/input/form"
	"sort"
	"sync"
	"time"
)

// Option 模块启动选项
type Option struct {
	Server *ghttp.Server // http服务器
	// 更多选项参数
	// ..
}

// Skeleton 模块骨架
type Skeleton struct {
	Label       string `json:"label"`       // 标识
	Name        string `json:"name"`        // 名称
	Group       int    `json:"group"`       // 分组
	Logo        string `json:"logo"`        // logo
	Brief       string `json:"brief"`       // 简介
	Description string `json:"description"` // 详细描述
	Author      string `json:"author"`      // 作者
	Version     string `json:"version"`     // 版本号
	RootPath    string `json:"rootPath"`    // 根路径
}

func (s *Skeleton) GetModule() Module {
	return GetModule(s.Name)
}

// Module 插件模块
type Module interface {
	Start(option *Option) (err error)          // 启动模块
	Stop() (err error)                         // 停止模块
	Ctx() context.Context                      // 上下文
	GetSkeleton() *Skeleton                    // 获取模块
	Install(ctx context.Context) (err error)   // 安装模块
	Upgrade(ctx context.Context) (err error)   // 更新模块
	UnInstall(ctx context.Context) (err error) // 卸载模块
}

var (
	modules = make(map[string]Module)
	mLock   sync.Mutex
)

// StartModules 启动所有已安装模块
func StartModules(ctx context.Context, option *Option) (err error) {
	for _, module := range filterInstalled() {
		if err = module.Start(option); err != nil {
			return
		}
	}

	// 为所有已安装模块设置静态资源路径
	AddStaticPath(ctx, option.Server)
	return
}

// StopModules 停止所有已安装模块
func StopModules(ctx context.Context) {
	for _, module := range filterInstalled() {
		if err := module.Stop(); err != nil {
			g.Log().Warningf(ctx, "StopModules err:%v, module:%v", err.Error(), module.GetSkeleton().Name)
			time.Sleep(time.Second)
		}
	}
	return
}

// RegisterModule 注册模块
func RegisterModule(m Module) Module {
	mLock.Lock()
	defer mLock.Unlock()
	name := m.GetSkeleton().Name
	_, ok := modules[name]
	if ok {
		panic("module repeat registration, name:" + name)
	}

	sk := m.GetSkeleton()
	if sk == nil {
		panic("module skeleton not initialized, name:" + name)
	}

	sk.RootPath = GetModulePath(name)
	modules[name] = m
	return m
}

// GetModule 获取指定名称模块
func GetModule(name string) Module {
	mLock.Lock()
	defer mLock.Unlock()
	m, ok := modules[name]
	if !ok {
		panic("implement not found for interface " + name + ", forgot register?")
	}
	return m
}

// GetSkeletons 获取所有模块骨架
func GetSkeletons() (list []*Skeleton) {
	var keys []string
	for k := range modules {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {
		list = append(list, GetModule(v).GetSkeleton())
	}
	return list
}

// GetModuleRealPath 获取指定模块绝对路径
func GetModuleRealPath(name string) string {
	path := gfile.RealPath(GetModulePath(name))
	if path == "" {
		panic("no path is found. please confirm that the path " + GetModulePath(name) + " exists?")
	}
	return path
}

// AddStaticPath 设置插件静态目录映射
func AddStaticPath(ctx context.Context, server *ghttp.Server) {
	basePath := GetResourcePath(ctx)
	if basePath == "" {
		return
	}

	for _, module := range filterInstalled() {
		name := module.GetSkeleton().Name
		prefix, path := StaticPath(name, basePath)
		if !gres.Contains(path) {
			if _, err := gfile.Search(path); err != nil {
				g.Log().Warningf(ctx, `addons AddStaticPath failed: %v`, err)
				continue
			}
		}
		server.AddStaticPath(prefix, path)
	}
}

// filterInstalled 过滤已安装模块
func filterInstalled() []Module {
	var ms []Module
	for _, module := range modules {
		if IsInstall(module) {
			ms = append(ms, module)
		}
	}
	return ms
}

// ModuleSelect 获取插件模块选项
func ModuleSelect() form.Selects {
	sks := GetSkeletons()
	lst := make(form.Selects, 0)
	if len(sks) == 0 {
		return lst
	}

	for _, skeleton := range sks {
		lst = append(lst, &form.Select{
			Value: skeleton.Name,
			Label: skeleton.Label,
			Name:  skeleton.Label,
		})
	}
	return lst
}
