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

// GetModulePath 获取指定模块相对路径
func GetModulePath(name string) string {
	return "./" + consts.AddonsDir + "/" + name
}

// ViewPath 默认的插件模板路径
func ViewPath(name string) string {
	return consts.AddonsDir + "/" + name + "/" + "resource/template"
}

// StaticPath 默认的插件静态路映射关系
// 最终效果：对外访问地址：/addons/插件模块名称；静态资源路径：/addons/插件模块名称/设置的子路径。
// 如果你不喜欢现在的路由风格，可以自行调整
func StaticPath(name, path string) (string, string) {
	return "/" + consts.AddonsDir + "/" + name, consts.AddonsDir + "/" + name + "/" + path
}

// RouterPrefix 路由前缀
// 最终效果：/应用名称/插件模块名称/xxx/xxx。
// 如果你不喜欢现在的路由风格，可以自行调整
func RouterPrefix(ctx context.Context, app, name string) string {
	var prefix = "/"
	if app != "" {
		prefix = g.Cfg().MustGet(ctx, "router."+app+".prefix", "/"+app+"").String()
	}
	return prefix + "/" + name
}
