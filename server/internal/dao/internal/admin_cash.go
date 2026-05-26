// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminCashDao is the data access object for the table hg_admin_cash.
type AdminCashDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminCashColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminCashColumns defines and stores column names for the table hg_admin_cash.
type AdminCashColumns struct {
	Id        string // ID
	MemberId  string // 管理员ID
	Money     string // 提现金额
	Fee       string // 手续费
	LastMoney string // 最终到账金额
	Ip        string // 申请人IP
	Status    string // 状态码
	Msg       string // 处理结果
	HandleAt  string // 处理时间
	CreatedAt string // 申请时间
}

// adminCashColumns holds the columns for the table hg_admin_cash.
var adminCashColumns = AdminCashColumns{
	Id:        "id",
	MemberId:  "member_id",
	Money:     "money",
	Fee:       "fee",
	LastMoney: "last_money",
	Ip:        "ip",
	Status:    "status",
	Msg:       "msg",
	HandleAt:  "handle_at",
	CreatedAt: "created_at",
}

// NewAdminCashDao creates and returns a new DAO object for table data access.
func NewAdminCashDao(handlers ...gdb.ModelHandler) *AdminCashDao {
	return &AdminCashDao{
		group:    "default",
		table:    "hg_admin_cash",
		columns:  adminCashColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminCashDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminCashDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminCashDao) Columns() AdminCashColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminCashDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminCashDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminCashDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
