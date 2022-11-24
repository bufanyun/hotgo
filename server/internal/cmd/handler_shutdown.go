// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/grpool"
	"hotgo/internal/crons"
	"hotgo/internal/websocket"
	"os"
)

func signalHandlerForCron(sig os.Signal) {
	crons.StopALL()
}

func signalHandlerForWebSocket(sig os.Signal) {
	websocket.Stop()
}

func signalHandlerForOverall(sig os.Signal) {
	serverCloseSignal <- struct{}{}
}

func signalListen(ctx context.Context, handler ...gproc.SigHandler) {
	err := grpool.AddWithRecover(ctx, func(ctx context.Context) {
		gproc.AddSigHandlerShutdown(handler...)
		gproc.Listen()
	})
	if err != nil {
		g.Log().Fatal(ctx, "signalListen Fatal:", err)
		return
	}
}
