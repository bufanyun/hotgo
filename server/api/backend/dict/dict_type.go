// Package dict
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package dict

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
)

// TypeTreeReq 字典类型树
type TypeTreeReq struct {
	g.Meta `path:"/dict_type/tree" tags:"字典类型" method:"get" summary:"字典类型树列表"`
}
type TypeTreeRes struct {
	List []map[string]interface{} `json:"list"   dc:"数据列表"`
}

// TypeEditReq 修改/新增字典数据
type TypeEditReq struct {
	entity.AdminDept
	g.Meta `path:"/dict_type/edit" method:"post" tags:"字典类型" summary:"修改/新增字典类型"`
}

type TypeEditRes struct{}

// TypeDeleteReq 删除字典类型
type TypeDeleteReq struct {
	Id     interface{} `json:"id" v:"required#字典类型ID不能为空" dc:"字典类型ID"`
	g.Meta `path:"/dict_type/delete" method:"post" tags:"字典类型" summary:"删除字典类型"`
}
type TypeDeleteRes struct{}

// TypeSelectReq 修改/新增字典数据
type TypeSelectReq struct {
	g.Meta `path:"/dict_type/select" method:"get" tags:"字典类型" summary:"字典类型选项"`
}

type TypeSelectRes sysin.DictTypeSelectModel
