// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminDeptDao is the data access object for the table hg_admin_dept.
type AdminDeptDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminDeptColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminDeptColumns defines and stores column names for the table hg_admin_dept.
type AdminDeptColumns struct {
	Id        string // 部门ID
	Pid       string // 父部门ID
	Name      string // 部门名称
	Code      string // 部门编码
	Type      string // 部门类型
	Leader    string // 负责人
	Phone     string // 联系电话
	Email     string // 邮箱
	Level     string // 关系树等级
	Tree      string // 关系树
	Sort      string // 排序
	Status    string // 部门状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// adminDeptColumns holds the columns for the table hg_admin_dept.
var adminDeptColumns = AdminDeptColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	Code:      "code",
	Type:      "type",
	Leader:    "leader",
	Phone:     "phone",
	Email:     "email",
	Level:     "level",
	Tree:      "tree",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminDeptDao creates and returns a new DAO object for table data access.
func NewAdminDeptDao(handlers ...gdb.ModelHandler) *AdminDeptDao {
	return &AdminDeptDao{
		group:    "default",
		table:    "hg_admin_dept",
		columns:  adminDeptColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminDeptDao) Columns() AdminDeptColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminDeptDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
