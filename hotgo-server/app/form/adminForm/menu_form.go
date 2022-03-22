//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  菜单最大排序
type MenuMaxSortReq struct {
	g.Meta `path:"/menu/max_sort" method:"get" tags:"菜单" summary:"菜单最大排序"`
	Id     int64 `json:"id" dc:"菜单ID"`
}
type MenuMaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

//  菜单编码是否唯一
type MenuCodeUniqueReq struct {
	g.Meta `path:"/menu/code_unique" method:"get" tags:"菜单" summary:"菜单编码是否唯一"`
	Code   string `json:"code" v:"required#菜单编码不能为空"  dc:"菜单编码"`
	Id     int64  `json:"id" dc:"菜单ID"`
}
type MenuCodeUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  菜单名称是否唯一
type MenuNameUniqueReq struct {
	g.Meta `path:"/menu/name_unique" method:"get" tags:"菜单" summary:"菜单名称是否唯一"`
	Name   string `json:"name" v:"required#菜单名称不能为空"  dc:"菜单名称"`
	Id     int64  `json:"id" dc:"菜单ID"`
}
type MenuNameUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  修改/新增字典数据
type MenuEditReq struct {
	g.Meta `path:"/menu/edit" method:"post" tags:"菜单" summary:"修改/新增菜单"`
	entity.AdminMenu
}
type MenuEditRes struct{}

//  删除字典类型
type MenuDeleteReq struct {
	g.Meta `path:"/menu/delete" method:"post" tags:"菜单" summary:"删除菜单"`
	Id     interface{} `json:"id" v:"required#菜单ID不能为空" dc:"菜单ID"`
}
type MenuDeleteRes struct{}

//  获取指定字典数据信息
type MenuViewReq struct {
	g.Meta `path:"/menu/view" method:"get" tags:"菜单" summary:"获取指定菜单信息"`
	Id     string `json:"id" v:"required#菜单ID不能为空" dc:"菜单ID"`
}
type MenuViewRes struct {
	*entity.AdminMenu
}

//  获取菜单列表
type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"菜单" summary:"获取菜单列表"`
	form.PageReq
	Pid int64 `json:"pid" dc:"父ID"`
}

type MenuListRes struct {
	List []*entity.AdminMenu `json:"list"   dc:"数据列表"`
	form.PageRes
}

//  查询菜单列表
type MenuSearchListReq struct {
	g.Meta `path:"/menu/search_list" method:"get" tags:"菜单" summary:"获取菜单列表"`
	Name   string `json:"name" dc:"菜单名称"`
	form.StatusReq
}

type MenuSearchListRes []*model.TreeMenu

//  查询角色菜单列表
type MenuRoleListReq struct {
	g.Meta `path:"/menu/role_list" method:"get" tags:"菜单" summary:"查询角色菜单列表"`
	RoleId string `json:"role_id" dc:"角色ID"`
}

type MenuRoleListRes struct {
	Menus       []*model.LabelTreeMenu `json:"menus"   dc:"菜单列表"`
	CheckedKeys []int64                `json:"checkedKeys"   dc:"选择的菜单ID"`
}
