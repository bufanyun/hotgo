// Package role
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package role

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
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
	g.Meta `path:"/role/dynamic" method:"get" tags:"角色" summary:"获取动态路由" description:"获取登录用户动态路由"`
}

type DynamicRes struct {
	List []*adminin.MenuRoute `json:"list"   description:"数据列表"`
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
	g.Meta `path:"/role/dataScope/select" method:"get" tags:"角色" summary:"获取数据权限选项"`
}

type DataScopeSelectRes struct {
	List []consts.GroupScopeSelect `json:"list" dc:"数据选项"`
}

// DataScopeEditReq 修改指定角色的数据权限
type DataScopeEditReq struct {
	g.Meta `path:"/role/dataScope/edit" method:"post" tags:"角色" summary:"修改指定角色的数据权限"`
	adminin.DataScopeEditInp
}

type DataScopeEditRes struct{}
