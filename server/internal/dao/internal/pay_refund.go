// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayRefundDao is the data access object for the table hg_pay_refund.
type PayRefundDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PayRefundColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PayRefundColumns defines and stores column names for the table hg_pay_refund.
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

// payRefundColumns holds the columns for the table hg_pay_refund.
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
func NewPayRefundDao(handlers ...gdb.ModelHandler) *PayRefundDao {
	return &PayRefundDao{
		group:    "default",
		table:    "hg_pay_refund",
		columns:  payRefundColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PayRefundDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PayRefundDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PayRefundDao) Columns() PayRefundColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PayRefundDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PayRefundDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PayRefundDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
