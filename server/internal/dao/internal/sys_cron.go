// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCronDao is the data access object for the table hg_sys_cron.
type SysCronDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysCronColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysCronColumns defines and stores column names for the table hg_sys_cron.
type SysCronColumns struct {
	Id        string // 任务ID
	GroupId   string // 分组ID
	Title     string // 任务标题
	Name      string // 任务方法
	Params    string // 函数参数
	Pattern   string // 表达式
	Policy    string // 策略
	Count     string // 执行次数
	Sort      string // 排序
	Remark    string // 备注
	Status    string // 任务状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysCronColumns holds the columns for the table hg_sys_cron.
var sysCronColumns = SysCronColumns{
	Id:        "id",
	GroupId:   "group_id",
	Title:     "title",
	Name:      "name",
	Params:    "params",
	Pattern:   "pattern",
	Policy:    "policy",
	Count:     "count",
	Sort:      "sort",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysCronDao creates and returns a new DAO object for table data access.
func NewSysCronDao(handlers ...gdb.ModelHandler) *SysCronDao {
	return &SysCronDao{
		group:    "default",
		table:    "hg_sys_cron",
		columns:  sysCronColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCronDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCronDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCronDao) Columns() SysCronColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCronDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCronDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysCronDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
