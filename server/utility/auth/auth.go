// Package auth
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/utility/validate"
)

// IsExceptAuth 是否是不需要验证权限的路由地址
func IsExceptAuth(ctx context.Context, path string) bool {
	var pathList []string

	except := g.Cfg().MustGet(ctx, "router.admin.exceptAuth")
	pathList = except.Strings()

	for i := 0; i < len(pathList); i++ {
		if validate.InSliceExistStr(pathList[i], path) {
			return true
		}
	}

	return false
}

// IsExceptLogin 是否是不需要登录的路由地址
func IsExceptLogin(ctx context.Context, path string) bool {
	var pathList []string

	except := g.Cfg().MustGet(ctx, "router.admin.exceptLogin")
	pathList = except.Strings()

	for i := 0; i < len(pathList); i++ {
		if validate.InSliceExistStr(pathList[i], path) {
			return true
		}
	}

	return false
}
