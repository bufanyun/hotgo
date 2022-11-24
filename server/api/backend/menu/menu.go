// Package menu
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package menu

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// MaxSortReq 菜单最大排序
type MaxSortReq struct {
	g.Meta `path:"/menu/max_sort" method:"get" tags:"菜单" summary:"菜单最大排序"`
	Id     int64 `json:"id" dc:"菜单ID"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// CodeUniqueReq 菜单编码是否唯一
type CodeUniqueReq struct {
	g.Meta `path:"/menu/code_unique" method:"get" tags:"菜单" summary:"菜单编码是否唯一"`
	Code   string `json:"code" v:"required#菜单编码不能为空"  dc:"菜单编码"`
	Id     int64  `json:"id" dc:"菜单ID"`
}
type CodeUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

// NameUniqueReq 菜单名称是否唯一
type NameUniqueReq struct {
	g.Meta `path:"/menu/name_unique" method:"get" tags:"菜单" summary:"菜单名称是否唯一"`
	Name   string `json:"name" v:"required#菜单名称不能为空"  dc:"菜单名称"`
	Id     int64  `json:"id" dc:"菜单ID"`
}
type NameUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

// EditReq 修改/新增菜单
type EditReq struct {
	g.Meta `path:"/menu/edit" method:"post" tags:"菜单" summary:"修改/新增菜单"`
	entity.AdminMenu
}
type EditRes struct{}

// DeleteReq 删除菜单
type DeleteReq struct {
	g.Meta `path:"/menu/delete" method:"post" tags:"菜单" summary:"删除菜单"`
	Id     interface{} `json:"id" v:"required#菜单ID不能为空" dc:"菜单ID"`
}
type DeleteRes struct{}

// ViewReq 获取指定菜单信息
type ViewReq struct {
	g.Meta `path:"/menu/view" method:"get" tags:"菜单" summary:"获取指定菜单信息"`
	Id     string `json:"id" v:"required#菜单ID不能为空" dc:"菜单ID"`
}
type ViewRes struct {
	*entity.AdminMenu
}

// ListReq 获取菜单列表
type ListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"菜单" summary:"获取菜单列表"`
	form.PageReq
	Pid int64 `json:"pid" dc:"父ID"`
}

type ListRes struct {
	List []map[string]interface{} `json:"list"   dc:"数据列表"`
}

// SearchListReq 查询菜单列表
type SearchListReq struct {
	g.Meta `path:"/menu/search_list" method:"get" tags:"菜单" summary:"获取菜单列表"`
	Name   string `json:"name" dc:"菜单名称"`
	form.StatusReq
}

type SearchListRes []*model.TreeMenu

// RoleListReq 查询角色菜单列表
type RoleListReq struct {
	g.Meta `path:"/menu/role_list" method:"get" tags:"菜单" summary:"查询角色菜单列表"`
	RoleId string `json:"role_id" dc:"角色ID"`
}

type RoleListRes struct {
	Menus       []*model.LabelTreeMenu `json:"menus"   dc:"菜单列表"`
	CheckedKeys []int64                `json:"checkedKeys"   dc:"选择的菜单ID"`
}
