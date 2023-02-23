// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateAddonsConfigInp 更新指定插件的配置
type UpdateAddonsConfigInp struct {
	AddonName string `json:"addonName"`
	Group     string `json:"group"`
	List      g.Map  `json:"list"`
}

// GetAddonsConfigInp 获取指定插件的配置
type GetAddonsConfigInp struct {
	AddonName string `json:"addonName"`
	Group     string `json:"group"`
}
type GetAddonsConfigModel struct {
	List g.Map `json:"list"`
}
