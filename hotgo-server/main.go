package main

import (
	"github.com/bufanyun/hotgo/boot"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {

	var (
		ctx = gctx.New()
	)

	if err := boot.Main.RunWithError(ctx); err != nil {
		g.Log().Fatal(ctx, err)
	}
}
