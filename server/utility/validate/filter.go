// Package validate
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package validate

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// Filter 通用过滤器

// Filter 预处理和数据过滤接口，目前主要用于input层的输入过滤和内部效验。
// 你可以在任意地方使用它，只要实现了Filter接口即可
type Filter interface {
	// Filter gf效验规则 https://goframe.org/pages/viewpage.action?pageId=1114367
	Filter(ctx context.Context) error
}

// PreFilter 预过滤
func PreFilter(ctx context.Context, in interface{}) error {
	return g.Try(ctx, func(ctx context.Context) {
		if c, ok := in.(Filter); ok {
			if err := c.Filter(ctx); err != nil {
				g.Throw(err)
			}
			return
		}
	})
}
