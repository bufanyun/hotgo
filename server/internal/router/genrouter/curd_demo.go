// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.1.0
// @AutoGenerate Date 2023-01-18 15:19:42
//
package genrouter

import "hotgo/internal/controller/backend/sys"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, sys.CurdDemo) // 生成演示
}
