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

func GetTag(name string) string {
	return consts.AddonsTag + name
}

func Tpl(name, tpl string) string {
	return consts.AddonsDir + "/" + name + "/" + tpl
}

// RouterPrefix 路由前缀
// 最终效果：/应用名称/插件模块名称/xxx/xxx。如果你不喜欢现在的路由风格，可以自行调整
func RouterPrefix(ctx context.Context, app, name string) string {
	var prefix = "/"
	if app != "" {
		prefix = g.Cfg().MustGet(ctx, "router."+app+".prefix", "/"+app+"").String()
	}
	return prefix + "/" + name
}
