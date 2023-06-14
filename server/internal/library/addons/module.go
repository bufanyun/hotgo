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
	"github.com/gogf/gf/v2/os/gview"
	"hotgo/internal/model/input/form"
	"sort"
	"sync"
)

// Skeleton 模块骨架
type Skeleton struct {
	Label       string      `json:"label"`       // 标识
	Name        string      `json:"name"`        // 名称
	Group       int         `json:"group"`       // 分组
	Logo        string      `json:"logo"`        // logo
	Brief       string      `json:"brief"`       // 简介
	Description string      `json:"description"` // 详细描述
	Author      string      `json:"author"`      // 作者
	Version     string      `json:"version"`     // 版本号
	RootPath    string      `json:"rootPath"`    // 根路径
	View        *gview.View `json:"view"`        // 模板引擎
}

func (s *Skeleton) GetModule() Module {
	return GetModule(s.Name)
}

// Module 插件模块
type Module interface {
	Init(ctx context.Context)                                 // 初始化
	InitRouter(ctx context.Context, group *ghttp.RouterGroup) // 初始化并注册路由
	Ctx() context.Context                                     // 上下文
	GetSkeleton() *Skeleton                                   // 架子
	Install(ctx context.Context) error                        // 安装模块
	Upgrade(ctx context.Context) error                        // 更新模块
	UnInstall(ctx context.Context) error                      // 卸载模块
}

var (
	modules = make(map[string]Module, 0)
	mLock   sync.Mutex
)

// InitModules 初始化所有已注册模块
func InitModules(ctx context.Context) {
	for _, module := range modules {
		module.Init(ctx)
	}
}

// RegisterModulesRouter 注册所有已安装模块路由
func RegisterModulesRouter(ctx context.Context, group *ghttp.RouterGroup) {
	for _, module := range filterInstalled() {
		module.InitRouter(ctx, group)
	}
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
	sk.View = NewView(m.Ctx(), name)
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

// NewView 初始化一个插件的模板引擎
func NewView(ctx context.Context, name string) *gview.View {
	view := gview.New()

	if err := view.SetPath(ViewPath(name)); err != nil {
		g.Log().Warningf(ctx, "NewView SetPath err:%+v", err)
		return nil
	}

	// 默认和主模块使用一致的变量分隔符号
	delimiters := g.Cfg().MustGet(ctx, "viewer.delimiters", []string{"@{", "}"}).Strings()
	if len(delimiters) != 2 {
		g.Log().Warning(ctx, "NewView delimiters config error")
		return nil
	}
	view.SetDelimiters(delimiters[0], delimiters[1])

	//// 更多配置
	//view.SetI18n()
	//// ...
	return view
}

// AddStaticPath 设置插件静态目录映射
func AddStaticPath(ctx context.Context, server *ghttp.Server, p ...string) {
	basePath := g.Cfg().MustGet(ctx, "server.serverRoot").String()
	if len(p) > 0 {
		basePath = p[0]
	}

	if basePath == "" {
		return
	}

	for _, module := range filterInstalled() {
		name := module.GetSkeleton().Name
		prefix, path := StaticPath(name, basePath)
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
