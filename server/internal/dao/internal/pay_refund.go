// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayRefundDao is the data access object for table hg_pay_refund.
type PayRefundDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns PayRefundColumns // columns contains all the column names of Table for convenient usage.
}

// PayRefundColumns defines and stores column names for table hg_pay_refund.
type PayRefundColumns struct {
	Id            string // 主键ID
	MemberId      string // 会员ID
	AppId         string // 应用ID
	OrderSn       string // 业务订单号
	RefundTradeNo string // 退款交易号
	RefundMoney   string // 退款金额
	RefundWay     string // 退款方式
	Ip            string // 申请者IP
	Reason        string // 申请退款原因
	Remark        string // 退款备注
	Status        string // 退款状态
	CreatedAt     string // 申请时间
	UpdatedAt     string // 更新时间
}

// payRefundColumns holds the columns for table hg_pay_refund.
var payRefundColumns = PayRefundColumns{
	Id:            "id",
	MemberId:      "member_id",
	AppId:         "app_id",
	OrderSn:       "order_sn",
	RefundTradeNo: "refund_trade_no",
	RefundMoney:   "refund_money",
	RefundWay:     "refund_way",
	Ip:            "ip",
	Reason:        "reason",
	Remark:        "remark",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewPayRefundDao creates and returns a new DAO object for table data access.
func NewPayRefundDao() *PayRefundDao {
	return &PayRefundDao{
		group:   "default",
		table:   "hg_pay_refund",
		columns: payRefundColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PayRefundDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PayRefundDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PayRefundDao) Columns() PayRefundColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PayRefundDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PayRefundDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PayRefundDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
