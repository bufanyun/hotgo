// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminOrderDao is the data access object for table hg_admin_order.
type AdminOrderDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AdminOrderColumns // columns contains all the column names of Table for convenient usage.
}

// AdminOrderColumns defines and stores column names for table hg_admin_order.
type AdminOrderColumns struct {
	Id                 string // 主键
	MemberId           string // 管理员id
	OrderType          string // 订单类型
	ProductId          string // 产品id
	OrderSn            string // 关联订单号
	Money              string // 充值金额
	Remark             string // 备注
	RefundReason       string // 退款原因
	RejectRefundReason string // 拒绝退款原因
	Status             string // 状态
	CreatedAt          string // 创建时间
	UpdatedAt          string // 修改时间
}

// adminOrderColumns holds the columns for table hg_admin_order.
var adminOrderColumns = AdminOrderColumns{
	Id:                 "id",
	MemberId:           "member_id",
	OrderType:          "order_type",
	ProductId:          "product_id",
	OrderSn:            "order_sn",
	Money:              "money",
	Remark:             "remark",
	RefundReason:       "refund_reason",
	RejectRefundReason: "reject_refund_reason",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewAdminOrderDao creates and returns a new DAO object for table data access.
func NewAdminOrderDao() *AdminOrderDao {
	return &AdminOrderDao{
		group:   "default",
		table:   "hg_admin_order",
		columns: adminOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminOrderDao) Columns() AdminOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
