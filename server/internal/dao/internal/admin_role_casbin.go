// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleCasbinDao is the data access object for the table hg_admin_role_casbin.
type AdminRoleCasbinDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  AdminRoleCasbinColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// AdminRoleCasbinColumns defines and stores column names for the table hg_admin_role_casbin.
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

// adminRoleCasbinColumns holds the columns for the table hg_admin_role_casbin.
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
func NewAdminRoleCasbinDao(handlers ...gdb.ModelHandler) *AdminRoleCasbinDao {
	return &AdminRoleCasbinDao{
		group:    "default",
		table:    "hg_admin_role_casbin",
		columns:  adminRoleCasbinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminRoleCasbinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminRoleCasbinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminRoleCasbinDao) Columns() AdminRoleCasbinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminRoleCasbinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminRoleCasbinDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AdminRoleCasbinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
