// Package addons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package addons

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
)

var cacheResourcePath string

// GetResourcePath 获取插件资源路径
func GetResourcePath(ctx context.Context) string {
	if len(cacheResourcePath) > 0 {
		return cacheResourcePath
	}
	basePath := g.Cfg().MustGet(ctx, "hotgo.addonsResourcePath").String()
	if basePath == "" {
		g.Log().Warning(ctx, "addons GetResourcePath not config found:'hotgo.addonsResourcePath', use default values:'resource'")
		basePath = "resource"
	}

	cacheResourcePath = basePath
	return basePath
}

// GetModulePath 获取指定模块相对路径
func GetModulePath(name string) string {
	return "./" + consts.AddonsDir + "/" + name
}

// ViewPath 默认的插件模板路径
// 模板路径：resource/addons/插件模块名称/template
// 例如：resource/addons/hgexample/template
// 如果你不喜欢现在的风格，可以自行调整
func ViewPath(name, basePath string) string {
	return basePath + "/" + consts.AddonsDir + "/" + name + "/template"
}

// StaticPath 默认的插件静态路映射关系
// 静态资源路径：resource/public/addons/插件模块名称/public
// 例如访问：http://127.0.0.1:8000/addons/hgexample/default 则指向文件-> resource/addons/hgexample/public/default
// 如果你不喜欢现在的风格，可以自行调整
func StaticPath(name, basePath string) (string, string) {
	return "/" + consts.AddonsDir + "/" + name, basePath + "/" + consts.AddonsDir + "/" + name + "/public"
}

// RouterPrefix 路由前缀
// 最终效果：/应用名称/插件模块名称/xxx/xxx。
// 如果你不喜欢现在的风格，可以自行调整
func RouterPrefix(ctx context.Context, app, name string) string {
	var prefix = "/"
	if app != "" {
		prefix = g.Cfg().MustGet(ctx, "router."+app+".prefix", "/"+app+"").String()
	}
	return prefix + "/" + name
}
