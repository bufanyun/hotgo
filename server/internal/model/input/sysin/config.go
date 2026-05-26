// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateConfigInp 更新指定分组的配置
type UpdateConfigInp struct {
	Group string `json:"group"`
	List  g.Map  `json:"list"`
}

// GetConfigInp 获取指定分组的配置
type GetConfigInp struct {
	Group string `json:"group"`
}

type GetConfigModel struct {
	List g.Map `json:"list"`
}
