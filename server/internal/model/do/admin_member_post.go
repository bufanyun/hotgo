// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminMemberPost is the golang structure of table hg_admin_member_post for DAO operations like Where/Data.
type AdminMemberPost struct {
	g.Meta   `orm:"table:hg_admin_member_post, do:true"`
	MemberId any // 管理员ID
	PostId   any // 岗位ID
}
