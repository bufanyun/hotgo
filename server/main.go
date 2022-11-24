// Package main
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package main

import (
	_ "hotgo/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/internal/cmd"
	"hotgo/internal/global"
	_ "hotgo/internal/logic"
)

func main() {
	var ctx = gctx.New()
	global.Init(ctx)
	cmd.Main.Run(ctx)
}
