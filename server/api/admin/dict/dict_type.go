// Package dict
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package dict

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/sysin"
)

// TypeTreeReq 字典类型树
type TypeTreeReq struct {
	g.Meta `path:"/dictType/tree" tags:"字典类型" method:"get" summary:"字典类型树列表"`
}

type TypeTreeRes struct {
	List []*sysin.DictTypeTree `json:"list"   dc:"数据列表"`
}

// TypeEditReq 修改/新增字典数据
type TypeEditReq struct {
	g.Meta `path:"/dictType/edit" method:"post" tags:"字典类型" summary:"修改/新增字典类型"`
	sysin.DictTypeEditInp
}

type TypeEditRes struct{}

// TypeDeleteReq 删除字典类型
type TypeDeleteReq struct {
	g.Meta `path:"/dictType/delete" method:"post" tags:"字典类型" summary:"删除字典类型"`
	sysin.DictTypeDeleteInp
}

type TypeDeleteRes struct{}
