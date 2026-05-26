// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCronGroupDao is the data access object for the table hg_sys_cron_group.
type SysCronGroupDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysCronGroupColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysCronGroupColumns defines and stores column names for the table hg_sys_cron_group.
type SysCronGroupColumns struct {
	Id        string // 任务分组ID
	Pid       string // 父类任务分组ID
	Name      string // 分组名称
	IsDefault string // 是否默认
	Sort      string // 排序
	Remark    string // 备注
	Status    string // 分组状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysCronGroupColumns holds the columns for the table hg_sys_cron_group.
var sysCronGroupColumns = SysCronGroupColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	IsDefault: "is_default",
	Sort:      "sort",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysCronGroupDao creates and returns a new DAO object for table data access.
func NewSysCronGroupDao(handlers ...gdb.ModelHandler) *SysCronGroupDao {
	return &SysCronGroupDao{
		group:    "default",
		table:    "hg_sys_cron_group",
		columns:  sysCronGroupColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCronGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCronGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCronGroupDao) Columns() SysCronGroupColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCronGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCronGroupDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysCronGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
