// Package role
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package role

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色" summary:"获取角色列表"`
	adminin.RoleListInp
}

type ListRes struct {
	*adminin.RoleListModel
	form.PageRes
}

// DynamicReq 动态路由
type DynamicReq struct {
	g.Meta `path:"/role/dynamic" method:"get" tags:"路由" summary:"获取动态路由" description:"获取登录用户动态路由"`
}

type DynamicMeta struct {
	Title   string `json:"title"      description:"菜单标题"`
	Icon    string `json:"icon"      description:"菜单图标"`
	NoCache bool   `json:"noCache"      description:"是否缓存"`
	Remark  string `json:"remark"      description:"备注"`
}

type DynamicBase struct {
	Id         int64        `json:"id"   description:"菜单ID"`
	Pid        int64        `json:"pid"   description:"父ID"`
	Name       string       `json:"name"      description:"菜单名称"`
	Code       string       `json:"code"      description:"菜单编码"`
	Path       string       `json:"path"      description:"路由地址"`
	Hidden     bool         `json:"hidden"      description:"是否隐藏"`
	Redirect   string       `json:"redirect"     description:"重定向"`
	Component  string       `json:"component" description:"组件路径"`
	AlwaysShow bool         `json:"alwaysShow"     description:"暂时不知道干啥"`
	IsFrame    string       `json:"isFrame"   description:"是否为外链（0是 1否）"`
	Meta       *DynamicMeta `json:"meta"   description:"配置数据集"`
}

type DynamicMenu struct {
	DynamicBase
	Children []*DynamicBase `json:"children"   description:"子菜单"`
}

type DynamicRes struct {
	List []adminin.MenuRoute `json:"list"   description:"数据列表"`
}

type UpdatePermissionsReq struct {
	g.Meta `path:"/role/updatePermissions" method:"post" tags:"角色" summary:"编辑角色菜单权限"`
	adminin.UpdatePermissionsInp
}

type UpdatePermissionsRes struct{}

type GetPermissionsReq struct {
	g.Meta `path:"/role/getPermissions" method:"get" tags:"角色" summary:"获取指定角色权限"`
	adminin.GetPermissionsInp
}

type GetPermissionsRes struct {
	*adminin.GetPermissionsModel
}

// EditReq 修改/新增角色
type EditReq struct {
	g.Meta `path:"/role/edit" method:"post" tags:"角色" summary:"修改/新增角色"`
	adminin.RoleEditInp
}
type EditRes struct{}

// DeleteReq 删除角色
type DeleteReq struct {
	g.Meta `path:"/role/delete" method:"post" tags:"角色" summary:"删除角色"`
	adminin.RoleDeleteInp
}
type DeleteRes struct{}

// DataScopeSelectReq 获取数据权限选项
type DataScopeSelectReq struct {
	g.Meta `path:"/role/dataScope/select" method:"get" summary:"角色" tags:"获取数据权限选项"`
}
type DataScopeSelectRes struct {
	List form.Selects `json:"list" dc:"数据选项"`
}

// DataScopeEditReq 修改指定角色的数据权限
type DataScopeEditReq struct {
	g.Meta `path:"/role/dataScope/edit" method:"post" tags:"角色" summary:"修改指定角色的数据权限"`
	adminin.DataScopeEditInp
}
type DataScopeEditRes struct{}
