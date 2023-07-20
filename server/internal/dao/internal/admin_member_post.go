// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminMemberPostDao is the data access object for table hg_admin_member_post.
type AdminMemberPostDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AdminMemberPostColumns // columns contains all the column names of Table for convenient usage.
}

// AdminMemberPostColumns defines and stores column names for table hg_admin_member_post.
type AdminMemberPostColumns struct {
	MemberId string // 管理员ID
	PostId   string // 岗位ID
}

// adminMemberPostColumns holds the columns for table hg_admin_member_post.
var adminMemberPostColumns = AdminMemberPostColumns{
	MemberId: "member_id",
	PostId:   "post_id",
}

// NewAdminMemberPostDao creates and returns a new DAO object for table data access.
func NewAdminMemberPostDao() *AdminMemberPostDao {
	return &AdminMemberPostDao{
		group:   "default",
		table:   "hg_admin_member_post",
		columns: adminMemberPostColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminMemberPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminMemberPostDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminMemberPostDao) Columns() AdminMemberPostColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminMemberPostDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminMemberPostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminMemberPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
