// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMenu is the golang structure for table admin_menu.
type AdminMenu struct {
	Id             int64       `json:"id"             description:"菜单ID"`
	Pid            int64       `json:"pid"            description:"父菜单ID"`
	Level          int         `json:"level"          description:"关系树等级"`
	Tree           string      `json:"tree"           description:"关系树"`
	Title          string      `json:"title"          description:"菜单名称"`
	Name           string      `json:"name"           description:"名称编码"`
	Path           string      `json:"path"           description:"路由地址"`
	Icon           string      `json:"icon"           description:"菜单图标"`
	Type           int         `json:"type"           description:"菜单类型（1目录 2菜单 3按钮）"`
	Redirect       string      `json:"redirect"       description:"重定向地址"`
	Permissions    string      `json:"permissions"    description:"菜单包含权限集合"`
	PermissionName string      `json:"permissionName" description:"权限名称"`
	Component      string      `json:"component"      description:"组件路径"`
	AlwaysShow     int         `json:"alwaysShow"     description:"取消自动计算根路由模式"`
	ActiveMenu     string      `json:"activeMenu"     description:"高亮菜单编码"`
	IsRoot         int         `json:"isRoot"         description:"是否跟路由"`
	IsFrame        int         `json:"isFrame"        description:"是否内嵌"`
	FrameSrc       string      `json:"frameSrc"       description:"内联外部地址"`
	KeepAlive      int         `json:"keepAlive"      description:"缓存该路由"`
	Hidden         int         `json:"hidden"         description:"是否隐藏"`
	Affix          int         `json:"affix"          description:"是否固定"`
	Sort           int         `json:"sort"           description:"排序"`
	Remark         string      `json:"remark"         description:"备注"`
	Status         int         `json:"status"         description:"菜单状态"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
}
