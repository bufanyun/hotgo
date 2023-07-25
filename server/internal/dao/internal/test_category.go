// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestCategoryDao is the data access object for table hg_test_category.
type TestCategoryDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TestCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// TestCategoryColumns defines and stores column names for table hg_test_category.
type TestCategoryColumns struct {
	Id          string // 分类ID
	Name        string // 分类名称
	Description string // 描述
	Sort        string // 排序
	Remark      string // 备注
	Status      string // 状态
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
}

// testCategoryColumns holds the columns for table hg_test_category.
var testCategoryColumns = TestCategoryColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	Sort:        "sort",
	Remark:      "remark",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewTestCategoryDao creates and returns a new DAO object for table data access.
func NewTestCategoryDao() *TestCategoryDao {
	return &TestCategoryDao{
		group:   "default",
		table:   "hg_test_category",
		columns: testCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TestCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TestCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TestCategoryDao) Columns() TestCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TestCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TestCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TestCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
