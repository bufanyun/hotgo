// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
)

// DictTypeEditInp 修改/新增字典数据
type DictTypeEditInp struct {
	entity.SysDictType
}
type DictTypeEditModel struct{}

// DictTypeDeleteInp 删除字典类型
type DictTypeDeleteInp struct {
	Id interface{}
}
type DictTypeDeleteModel struct{}

// DictTypeSelectInp 获取类型选项
type DictTypeSelectInp struct {
}

type DictTypeSelectModel []g.Map
