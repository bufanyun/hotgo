// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.13.1
package genrouter

import "hotgo/addons/hgexample/controller/admin/sys"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, sys.TenantOrder) // 多租户功能演示
}