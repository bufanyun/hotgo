// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/sysin"
)

// UpdateConfigInp 更新指定配置
type UpdateConfigInp struct {
	sysin.UpdateAddonsConfigInp
}

type GetConfigInp struct {
	sysin.GetAddonsConfigInp
}

type GetConfigModel struct {
	List g.Map `json:"list"`
}
