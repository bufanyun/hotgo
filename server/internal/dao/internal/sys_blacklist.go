// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysBlacklistDao is the data access object for table hg_sys_blacklist.
type SysBlacklistDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysBlacklistColumns // columns contains all the column names of Table for convenient usage.
}

// SysBlacklistColumns defines and stores column names for table hg_sys_blacklist.
type SysBlacklistColumns struct {
	Id        string // 黑名单ID
	Ip        string // IP地址
	Remark    string // 备注
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysBlacklistColumns holds the columns for table hg_sys_blacklist.
var sysBlacklistColumns = SysBlacklistColumns{
	Id:        "id",
	Ip:        "ip",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysBlacklistDao creates and returns a new DAO object for table data access.
func NewSysBlacklistDao() *SysBlacklistDao {
	return &SysBlacklistDao{
		group:   "default",
		table:   "hg_sys_blacklist",
		columns: sysBlacklistColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysBlacklistDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysBlacklistDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysBlacklistDao) Columns() SysBlacklistColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysBlacklistDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysBlacklistDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysBlacklistDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
