// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminPostDao is the data access object for the table hg_admin_post.
type AdminPostDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminPostColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminPostColumns defines and stores column names for the table hg_admin_post.
type AdminPostColumns struct {
	Id        string // 岗位ID
	Code      string // 岗位编码
	Name      string // 岗位名称
	Remark    string // 备注
	Sort      string // 排序
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// adminPostColumns holds the columns for the table hg_admin_post.
var adminPostColumns = AdminPostColumns{
	Id:        "id",
	Code:      "code",
	Name:      "name",
	Remark:    "remark",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminPostDao creates and returns a new DAO object for table data access.
func NewAdminPostDao(handlers ...gdb.ModelHandler) *AdminPostDao {
	return &AdminPostDao{
		group:    "default",
		table:    "hg_admin_post",
		columns:  adminPostColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminPostDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminPostDao) Columns() AdminPostColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminPostDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminPostDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
