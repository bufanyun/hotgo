// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAddonsInstallDao is the data access object for the table hg_sys_addons_install.
type SysAddonsInstallDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  SysAddonsInstallColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// SysAddonsInstallColumns defines and stores column names for the table hg_sys_addons_install.
type SysAddonsInstallColumns struct {
	Id        string // 主键
	Name      string // 插件名称
	Version   string // 版本号
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysAddonsInstallColumns holds the columns for the table hg_sys_addons_install.
var sysAddonsInstallColumns = SysAddonsInstallColumns{
	Id:        "id",
	Name:      "name",
	Version:   "version",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysAddonsInstallDao creates and returns a new DAO object for table data access.
func NewSysAddonsInstallDao(handlers ...gdb.ModelHandler) *SysAddonsInstallDao {
	return &SysAddonsInstallDao{
		group:    "default",
		table:    "hg_sys_addons_install",
		columns:  sysAddonsInstallColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAddonsInstallDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAddonsInstallDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAddonsInstallDao) Columns() SysAddonsInstallColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAddonsInstallDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAddonsInstallDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysAddonsInstallDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
