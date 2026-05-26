// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysGenTreeDemoDao is the data access object for the table hg_sys_gen_tree_demo.
type SysGenTreeDemoDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SysGenTreeDemoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SysGenTreeDemoColumns defines and stores column names for the table hg_sys_gen_tree_demo.
type SysGenTreeDemoColumns struct {
	Id          string // ID
	Pid         string // 上级ID
	Level       string // 关系树级别
	Tree        string // 关系树
	CategoryId  string // 分类ID
	Title       string // 标题
	Description string // 描述
	Sort        string // 排序
	Status      string // 状态
	CreatedBy   string // 创建者
	UpdatedBy   string // 更新者
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
}

// sysGenTreeDemoColumns holds the columns for the table hg_sys_gen_tree_demo.
var sysGenTreeDemoColumns = SysGenTreeDemoColumns{
	Id:          "id",
	Pid:         "pid",
	Level:       "level",
	Tree:        "tree",
	CategoryId:  "category_id",
	Title:       "title",
	Description: "description",
	Sort:        "sort",
	Status:      "status",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewSysGenTreeDemoDao creates and returns a new DAO object for table data access.
func NewSysGenTreeDemoDao(handlers ...gdb.ModelHandler) *SysGenTreeDemoDao {
	return &SysGenTreeDemoDao{
		group:    "default",
		table:    "hg_sys_gen_tree_demo",
		columns:  sysGenTreeDemoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysGenTreeDemoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysGenTreeDemoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysGenTreeDemoDao) Columns() SysGenTreeDemoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysGenTreeDemoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysGenTreeDemoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysGenTreeDemoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
