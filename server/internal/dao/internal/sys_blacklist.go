// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysBlacklistDao is the data access object for the table hg_sys_blacklist.
type SysBlacklistDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysBlacklistColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysBlacklistColumns defines and stores column names for the table hg_sys_blacklist.
type SysBlacklistColumns struct {
	Id        string // 黑名单ID
	Ip        string // IP地址
	Remark    string // 备注
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysBlacklistColumns holds the columns for the table hg_sys_blacklist.
var sysBlacklistColumns = SysBlacklistColumns{
	Id:        "id",
	Ip:        "ip",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysBlacklistDao creates and returns a new DAO object for table data access.
func NewSysBlacklistDao(handlers ...gdb.ModelHandler) *SysBlacklistDao {
	return &SysBlacklistDao{
		group:    "default",
		table:    "hg_sys_blacklist",
		columns:  sysBlacklistColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysBlacklistDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysBlacklistDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysBlacklistDao) Columns() SysBlacklistColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysBlacklistDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysBlacklistDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysBlacklistDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
