// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAddonsInstall is the golang structure of table hg_sys_addons_install for DAO operations like Where/Data.
type SysAddonsInstall struct {
	g.Meta    `orm:"table:hg_sys_addons_install, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 插件名称
	Version   interface{} // 版本号
	Status    interface{} // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
