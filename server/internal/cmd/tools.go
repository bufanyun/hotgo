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
	"github.com/gogf/gf/v2/os/gcmd"
	"hotgo/internal/library/casbin"
)

var (
	Tools = &gcmd.Command{
		Name:        "tools",
		Brief:       "工具",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			flags := parser.GetOptAll()
			g.Log().Warningf(ctx, "flags:%+v", flags)
			if len(flags) == 0 {
				g.Log().Fatal(ctx, "工具参数不能为空")
				return
			}

			method, ok := flags["m"]
			if !ok {
				g.Log().Fatal(ctx, "工具方法不能为空")
				return
			}

			switch method {
			case "casbin":
				a1, ok := flags["a1"]
				if !ok {
					g.Log().Fatal(ctx, "casbin参数不能为空")
					return
				}
				if a1 == "clear" {
					if err := casbin.Clear(ctx); err != nil {
						return err
					}
				} else if a1 == "refresh" {
					casbin.InitEnforcer(ctx)
					if err := casbin.Refresh(ctx); err != nil {
						return err
					}
				} else {
					g.Log().Fatalf(ctx, "casbin参数无效,a1：%+v", a1)
					return
				}
			default:
				g.Log().Fatal(ctx, "工具方法不存在")
			}
			g.Log().Info(ctx, "执行完成！")
			return
		},
	}
)
