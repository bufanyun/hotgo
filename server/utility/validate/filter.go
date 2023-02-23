// Package validate
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package validate

import (
	"context"
)

type Filter interface {
	// Filter gf效验规则 https://goframe.org/pages/viewpage.action?pageId=1114367
	Filter(ctx context.Context) error
}

// PreFilter 预过滤
func PreFilter(ctx context.Context, in interface{}) error {
	if c, ok := in.(Filter); ok {
		return c.Filter(ctx)
	}
	return nil
}
