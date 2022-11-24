// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package hook

import (
	"hotgo/internal/service"
)

type sHook struct {
}

func init() {
	service.RegisterHook(New())
}

func New() *sHook {
	return &sHook{}
}
