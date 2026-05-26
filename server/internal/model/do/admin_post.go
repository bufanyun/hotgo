// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPost is the golang structure of table hg_admin_post for DAO operations like Where/Data.
type AdminPost struct {
	g.Meta    `orm:"table:hg_admin_post, do:true"`
	Id        any         // 岗位ID
	Code      any         // 岗位编码
	Name      any         // 岗位名称
	Remark    any         // 备注
	Sort      any         // 排序
	Status    any         // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
