// Package dept
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package dept

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
)

// NameUniqueReq 名称是否唯一
type NameUniqueReq struct {
	Name   string `json:"name" v:"required#部门名称不能为空"  dc:"部门名称"`
	Id     int64  `json:"id" dc:"部门ID"`
	g.Meta `path:"/dept/name_unique" method:"get" tags:"部门" summary:"部门名称是否唯一"`
}
type NameUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

// ListReq 查询列表
type ListReq struct {
	Name   string `json:"name" dc:"部门名称"`
	Code   string `json:"code" dc:"部门编码"`
	g.Meta `path:"/dept/list" method:"get" tags:"部门" summary:"获取部门列表"`
}

type ListRes struct {
	adminin.DeptListModel
}

// ViewReq 获取指定信息
type ViewReq struct {
	Id     int64 `json:"id" v:"required#部门ID不能为空" dc:"部门ID"`
	g.Meta `path:"/dept/view" method:"get" tags:"部门" summary:"获取指定信息"`
}
type ViewRes struct {
	*adminin.DeptViewModel
}

// EditReq 修改/新增字典数据
type EditReq struct {
	entity.AdminDept
	g.Meta `path:"/dept/edit" method:"post" tags:"部门" summary:"修改/新增部门"`
}
type EditRes struct{}

// DeleteReq 删除字典类型
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#部门ID不能为空" dc:"部门ID"`
	g.Meta `path:"/dept/delete" method:"post" tags:"部门" summary:"删除部门"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	Id     int64 `json:"id" dc:"部门ID"`
	g.Meta `path:"/dept/maxSort" method:"get" tags:"部门" summary:"部门最大排序"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新部门状态
type StatusReq struct {
	entity.AdminDept
	g.Meta `path:"/dept/status" method:"post" tags:"部门" summary:"更新部门状态"`
}
type StatusRes struct{}
