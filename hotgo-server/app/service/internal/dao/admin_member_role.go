// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/bufanyun/hotgo/app/service/internal/dao/internal"
)

// adminMemberRoleDao is the data access object for table hg_admin_member_role.
// You can define custom methods on it to extend its functionality as you wish.
type adminMemberRoleDao struct {
	*internal.AdminMemberRoleDao
}

var (
	// AdminMemberRole is globally public accessible object for table hg_admin_member_role operations.
	AdminMemberRole = adminMemberRoleDao{
		internal.NewAdminMemberRoleDao(),
	}
)

// Fill with you ideas below.
