// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gproc"
	"hotgo/utility/simple"
	"os"
)

func signalHandlerForOverall(sig os.Signal) {
	serverCloseSignal <- struct{}{}
}

func signalListen(ctx context.Context, handler ...gproc.SigHandler) {
	simple.SafeGo(ctx, func(ctx context.Context) {
		gproc.AddSigHandlerShutdown(handler...)
		gproc.Listen()
	})
}
