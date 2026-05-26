// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMenu is the golang structure for table admin_menu.
type AdminMenu struct {
	Id             int64       `json:"id"             orm:"id"              description:"菜单ID"`
	Pid            int64       `json:"pid"            orm:"pid"             description:"父菜单ID"`
	Level          int         `json:"level"          orm:"level"           description:"关系树等级"`
	Tree           string      `json:"tree"           orm:"tree"            description:"关系树"`
	Title          string      `json:"title"          orm:"title"           description:"菜单名称"`
	Name           string      `json:"name"           orm:"name"            description:"名称编码"`
	Path           string      `json:"path"           orm:"path"            description:"路由地址"`
	Icon           string      `json:"icon"           orm:"icon"            description:"菜单图标"`
	Type           int         `json:"type"           orm:"type"            description:"菜单类型（1目录 2菜单 3按钮）"`
	Redirect       string      `json:"redirect"       orm:"redirect"        description:"重定向地址"`
	Permissions    string      `json:"permissions"    orm:"permissions"     description:"菜单包含权限集合"`
	PermissionName string      `json:"permissionName" orm:"permission_name" description:"权限名称"`
	Component      string      `json:"component"      orm:"component"       description:"组件路径"`
	AlwaysShow     int         `json:"alwaysShow"     orm:"always_show"     description:"取消自动计算根路由模式"`
	ActiveMenu     string      `json:"activeMenu"     orm:"active_menu"     description:"高亮菜单编码"`
	IsRoot         int         `json:"isRoot"         orm:"is_root"         description:"是否跟路由"`
	IsFrame        int         `json:"isFrame"        orm:"is_frame"        description:"是否内嵌"`
	FrameSrc       string      `json:"frameSrc"       orm:"frame_src"       description:"内联外部地址"`
	KeepAlive      int         `json:"keepAlive"      orm:"keep_alive"      description:"缓存该路由"`
	Hidden         int         `json:"hidden"         orm:"hidden"          description:"是否隐藏"`
	Affix          int         `json:"affix"          orm:"affix"           description:"是否固定"`
	Sort           int         `json:"sort"           orm:"sort"            description:"排序"`
	Remark         string      `json:"remark"         orm:"remark"          description:"备注"`
	Status         int         `json:"status"         orm:"status"          description:"菜单状态"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
}
