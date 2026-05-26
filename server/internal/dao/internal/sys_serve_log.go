// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysServeLogDao is the data access object for the table hg_sys_serve_log.
type SysServeLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysServeLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysServeLogColumns defines and stores column names for the table hg_sys_serve_log.
type SysServeLogColumns struct {
	Id          string // 日志ID
	TraceId     string // 链路ID
	LevelFormat string // 日志级别
	Content     string // 日志内容
	Stack       string // 打印堆栈
	Line        string // 调用行
	TriggerNs   string // 触发时间(ns)
	Status      string // 状态
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// sysServeLogColumns holds the columns for the table hg_sys_serve_log.
var sysServeLogColumns = SysServeLogColumns{
	Id:          "id",
	TraceId:     "trace_id",
	LevelFormat: "level_format",
	Content:     "content",
	Stack:       "stack",
	Line:        "line",
	TriggerNs:   "trigger_ns",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSysServeLogDao creates and returns a new DAO object for table data access.
func NewSysServeLogDao(handlers ...gdb.ModelHandler) *SysServeLogDao {
	return &SysServeLogDao{
		group:    "default",
		table:    "hg_sys_serve_log",
		columns:  sysServeLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysServeLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysServeLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysServeLogDao) Columns() SysServeLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysServeLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysServeLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysServeLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
