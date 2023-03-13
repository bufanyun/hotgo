package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
	"hotgo/internal/service"
	"os"
)

var (
	Auth = &gcmd.Command{
		Name:        "auth",
		Brief:       "系统授权，当为第三方客户开发应用项目不想将源码和可执行文件让其随意使用时，可以通过授权的方式约束使用方。",
		Description: `目前已实现，一对一、一对多、有效期授权，具体使用可以参考现有逻辑结合实际场景进行改造`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			service.TCPAuth().Start(ctx)

			// 退出信号监听
			signalListen(ctx, func(sig os.Signal) {
				service.TCPAuth().Stop(ctx)
			})

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)
			select {
			case <-serverCloseSignal:
				// ...
			}

			return
		},
	}
)
