// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleCasbinDao is the data access object for table hg_admin_role_casbin.
type AdminRoleCasbinDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AdminRoleCasbinColumns // columns contains all the column names of Table for convenient usage.
}

// AdminRoleCasbinColumns defines and stores column names for table hg_admin_role_casbin.
type AdminRoleCasbinColumns struct {
	Id    string //
	PType string //
	V0    string //
	V1    string //
	V2    string //
	V3    string //
	V4    string //
	V5    string //
}

// adminRoleCasbinColumns holds the columns for table hg_admin_role_casbin.
var adminRoleCasbinColumns = AdminRoleCasbinColumns{
	Id:    "id",
	PType: "p_type",
	V0:    "v0",
	V1:    "v1",
	V2:    "v2",
	V3:    "v3",
	V4:    "v4",
	V5:    "v5",
}

// NewAdminRoleCasbinDao creates and returns a new DAO object for table data access.
func NewAdminRoleCasbinDao() *AdminRoleCasbinDao {
	return &AdminRoleCasbinDao{
		group:   "default",
		table:   "hg_admin_role_casbin",
		columns: adminRoleCasbinColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminRoleCasbinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminRoleCasbinDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminRoleCasbinDao) Columns() AdminRoleCasbinColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminRoleCasbinDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminRoleCasbinDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminRoleCasbinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
