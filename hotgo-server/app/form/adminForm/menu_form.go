package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  菜单最大排序
type MenuMaxSortReq struct {
	Id     int64 `json:"id" description:"菜单ID"`
	g.Meta `path:"/menu/max_sort" method:"get" tags:"菜单" summary:"菜单最大排序"`
}
type MenuMaxSortRes struct {
	Sort int `json:"sort" description:"排序"`
}

//  菜单编码是否唯一
type MenuCodeUniqueReq struct {
	Code   string `json:"code" v:"required#菜单编码不能为空"  description:"菜单编码"`
	Id     int64  `json:"id" description:"菜单ID"`
	g.Meta `path:"/menu/code_unique" method:"get" tags:"菜单" summary:"菜单编码是否唯一"`
}
type MenuCodeUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  菜单名称是否唯一
type MenuNameUniqueReq struct {
	Name   string `json:"name" v:"required#菜单名称不能为空"  description:"菜单名称"`
	Id     int64  `json:"id" description:"菜单ID"`
	g.Meta `path:"/menu/name_unique" method:"get" tags:"菜单" summary:"菜单名称是否唯一"`
}
type MenuNameUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  修改/新增字典数据
type MenuEditReq struct {
	entity.AdminMenu
	g.Meta `path:"/menu/edit" method:"post" tags:"菜单" summary:"修改/新增菜单"`
}
type MenuEditRes struct{}

//  删除字典类型
type MenuDeleteReq struct {
	Id     interface{} `json:"id" v:"required#菜单ID不能为空" description:"菜单ID"`
	g.Meta `path:"/menu/delete" method:"post" tags:"菜单" summary:"删除菜单"`
}
type MenuDeleteRes struct{}

//  获取指定字典数据信息
type MenuViewReq struct {
	Id     string `json:"id" v:"required#菜单ID不能为空" description:"菜单ID"`
	g.Meta `path:"/menu/view" method:"get" tags:"菜单" summary:"获取指定菜单信息"`
}
type MenuViewRes struct {
	*entity.AdminMenu
}

//  获取菜单列表
type MenuListReq struct {
	form.PageReq
	Pid    int64 `json:"pid" description:"父ID"`
	g.Meta `path:"/menu/list" method:"get" tags:"菜单" summary:"获取菜单列表"`
}

type MenuListRes struct {
	List []*entity.AdminMenu `json:"list"   description:"数据列表"`
	form.PageRes
}

//  查询菜单列表
type MenuSearchListReq struct {
	Name string `json:"name" description:"菜单名称"`
	form.StatusReq
	g.Meta `path:"/menu/search_list" method:"get" tags:"菜单" summary:"获取菜单列表"`
}

type MenuSearchListRes []*model.TreeMenu

//  查询角色菜单列表
type MenuRoleListReq struct {
	RoleId string `json:"role_id" description:"角色ID"`
	g.Meta `path:"/menu/role_list" method:"get" tags:"菜单" summary:"查询角色菜单列表"`
}

type MenuRoleListRes struct {
	Menus       []*model.LabelTreeMenu `json:"menus"   description:"菜单列表"`
	CheckedKeys []int64                `json:"checkedKeys"   description:"选择的菜单ID"`
}
