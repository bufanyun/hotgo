// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gres"
	"hotgo/internal/library/casbin"
)

var (
	Tools = &gcmd.Command{
		Name:        "tools",
		Brief:       "常用工具",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			args := parser.GetOptAll()
			g.Log().Debugf(ctx, "tools args:%v", args)
			if len(args) == 0 {
				err = gerror.New("tools args cannot be empty.")
				return
			}

			method, ok := args["m"]
			if !ok {
				err = gerror.New("tools method cannot be empty.")
				return
			}

			switch method {
			case "casbin":
				err = handleCasbin(ctx, args)
			case "gres":
				err = handleGRes(ctx, args)
			default:
				err = gerror.Newf("tools method[%v] does not exist", method)
			}

			if err == nil {
				g.Log().Info(ctx, "tools exec successful!")
			}
			return
		},
	}
)

// handleCasbin casbin.
func handleCasbin(ctx context.Context, args map[string]string) (err error) {
	a1, ok := args["a1"]
	if !ok {
		err = gerror.New("casbin args cannot be empty.")
		return
	}

	casbin.InitEnforcer(ctx)
	switch a1 {
	case "clear":
		err = casbin.Clear(ctx)
	case "refresh":
		err = casbin.Refresh(ctx)
	default:
		err = gerror.Newf("casbin a1 is invalid, a1:%v", a1)
	}
	return
}

func handleGRes(ctx context.Context, args map[string]string) (err error) {
	a1, ok := args["a1"]
	if !ok {
		err = gerror.New("gres args cannot be empty.")
		return
	}

	switch a1 {
	case "dump":
		gres.Dump()
	case "content":
		path, ok := args["a2"]
		if !ok {
			err = gerror.New("缺少查看文件路径参数：`a2`")
			return
		}

		if !gres.Contains(path) {
			err = gerror.Newf("没有找到资源文件:%v", path)
			return
		}
		content := string(gres.GetContent(path))

		if len(content) == 0 {
			err = gerror.Newf("没有找到资源文件内容，请确认传入`a2`参数是一个文件，a2:%v", path)
			return
		}
		fmt.Println("以下是资源文件内容:")
		fmt.Println(content)
	default:
		err = gerror.Newf("handleGRes a1 is invalid, a1:%v", a1)
	}
	return
}
