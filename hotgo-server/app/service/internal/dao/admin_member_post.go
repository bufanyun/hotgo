// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/bufanyun/hotgo/app/service/internal/dao/internal"
)

// adminMemberPostDao is the data access object for table hg_admin_member_post.
// You can define custom methods on it to extend its functionality as you wish.
type adminMemberPostDao struct {
	*internal.AdminMemberPostDao
}

var (
	// AdminMemberPost is globally public accessible object for table hg_admin_member_post operations.
	AdminMemberPost = adminMemberPostDao{
		internal.NewAdminMemberPostDao(),
	}
)

// Fill with you ideas below.
