// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// MenuMaxSortReq 菜单最大排序
type MenuMaxSortReq struct {
	g.Meta `path:"/menu/max_sort" method:"get" tags:"菜单" summary:"菜单最大排序"`
	Id     int64 `json:"id" dc:"菜单ID"`
}
type MenuMaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// MenuCodeUniqueReq 菜单编码是否唯一
type MenuCodeUniqueReq struct {
	g.Meta `path:"/menu/code_unique" method:"get" tags:"菜单" summary:"菜单编码是否唯一"`
	Code   string `json:"code" v:"required#菜单编码不能为空"  dc:"菜单编码"`
	Id     int64  `json:"id" dc:"菜单ID"`
}
type MenuCodeUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

// MenuNameUniqueReq 菜单名称是否唯一
type MenuNameUniqueReq struct {
	g.Meta `path:"/menu/name_unique" method:"get" tags:"菜单" summary:"菜单名称是否唯一"`
	Name   string `json:"name" v:"required#菜单名称不能为空"  dc:"菜单名称"`
	Id     int64  `json:"id" dc:"菜单ID"`
}
type MenuNameUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

// MenuEditReq 修改/新增菜单
type MenuEditReq struct {
	g.Meta `path:"/menu/edit" method:"post" tags:"菜单" summary:"修改/新增菜单"`
	entity.AdminMenu
}
type MenuEditRes struct{}

// MenuDeleteReq 删除菜单
type MenuDeleteReq struct {
	g.Meta `path:"/menu/delete" method:"post" tags:"菜单" summary:"删除菜单"`
	Id     interface{} `json:"id" v:"required#菜单ID不能为空" dc:"菜单ID"`
}
type MenuDeleteRes struct{}

// MenuViewReq 获取指定菜单信息
type MenuViewReq struct {
	g.Meta `path:"/menu/view" method:"get" tags:"菜单" summary:"获取指定菜单信息"`
	Id     string `json:"id" v:"required#菜单ID不能为空" dc:"菜单ID"`
}
type MenuViewRes struct {
	*entity.AdminMenu
}

// MenuListReq 获取菜单列表
type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"菜单" summary:"获取菜单列表"`
	form.PageReq
	Pid int64 `json:"pid" dc:"父ID"`
}

type MenuListRes struct {
	List []*entity.AdminMenu `json:"list"   dc:"数据列表"`
	form.PageRes
}

// MenuSearchListReq 查询菜单列表
type MenuSearchListReq struct {
	g.Meta `path:"/menu/search_list" method:"get" tags:"菜单" summary:"获取菜单列表"`
	Name   string `json:"name" dc:"菜单名称"`
	form.StatusReq
}

type MenuSearchListRes []*model.TreeMenu

// MenuRoleListReq 查询角色菜单列表
type MenuRoleListReq struct {
	g.Meta `path:"/menu/role_list" method:"get" tags:"菜单" summary:"查询角色菜单列表"`
	RoleId string `json:"role_id" dc:"角色ID"`
}

type MenuRoleListRes struct {
	Menus       []*model.LabelTreeMenu `json:"menus"   dc:"菜单列表"`
	CheckedKeys []int64                `json:"checkedKeys"   dc:"选择的菜单ID"`
}

// MenuTree 菜单树结构
type MenuTree struct {
	// 适配n-tree
	Id    int64  `json:"key" `
	Title string `json:"label"`
	entity.AdminMenu
}

// MenuRouteMeta 菜单路由
type MenuRouteMeta struct {
	// 解释参考：https://naive-ui-admin-docs.vercel.app/guide/router.html#%E5%A4%9A%E7%BA%A7%E8%B7%AF%E7%94%B1
	Title string `json:"title"` // 菜单名称 一般必填
	//Disabled   bool   `json:"disabled,omitempty"`   // 禁用菜单
	Icon       string `json:"icon,omitempty"`       // 菜单图标
	KeepAlive  bool   `json:"keepAlive,omitempty"`  // 缓存该路由
	Hidden     bool   `json:"hidden,omitempty"`     // 隐藏菜单
	Sort       int    `json:"sort,omitempty"`       // 排序越小越排前
	AlwaysShow bool   `json:"alwaysShow,omitempty"` // 取消自动计算根路由模式
	ActiveMenu string `json:"activeMenu,omitempty"` // 当路由设置了该属性，则会高亮相对应的侧边栏。
	// 这在某些场景非常有用，比如：一个列表页路由为：/list/basic-list
	// 点击进入详情页，这时候路由为/list/basic-info/1，但你想在侧边栏高亮列表的路由，就可以进行如下设置
	// 注意是配置高亮路由 `name`，不是path
	IsRoot      bool   `json:"isRoot,omitempty"`      // 是否跟路由 顶部混合菜单，必须传 true，否则左侧会显示异常（场景就是，分割菜单之后，当一级菜单没有子菜单）
	FrameSrc    string `json:"frameSrc,omitempty" `   // 内联外部地址
	Permissions string `json:"permissions,omitempty"` // 菜单包含权限集合，满足其中一个就会显示
	Affix       bool   `json:"affix,omitempty"`       // 是否固定 设置为 true 之后 多页签不可删除

	// 自定义
	Type int `json:"type"` // 菜单类型
}

type MenuRoute struct {
	Name      string        `json:"name"`
	Path      string        `json:"path"`
	Redirect  string        `json:"redirect"`
	Component string        `json:"component"`
	Meta      MenuRouteMeta `json:"meta"`
	Children  []MenuRoute   `json:"children,omitempty" dc:"子路由"`
}

// MenuRouteSummary 菜单树结构
type MenuRouteSummary struct {
	entity.AdminMenu
	Children []MenuRouteSummary
}
