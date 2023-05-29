// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.7.3
package genrouter

import "hotgo/internal/controller/admin/sys"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, sys.CurdDemo) // 生成演示
}
