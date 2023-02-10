package crons

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
)

func webRequest(ctx context.Context, url string) {
	args, _ := ctx.Value(consts.ContextKeyCronArgs).([]string)

	var (
		method = "GET"
	)
	
	for _, v := range args {
		if gstr.Contains(v, "method") {
			method_ := gstr.Split("method", "=")
			if len(method_) == 2 {
				method = gstr.ToUpper(method_[1])
			}
		}
	}
	client := g.Client()
	client.DoRequest(ctx, method, url, nil)
}
