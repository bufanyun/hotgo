// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminNoticeDao is the data access object for the table hg_admin_notice.
type AdminNoticeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminNoticeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminNoticeColumns defines and stores column names for the table hg_admin_notice.
type AdminNoticeColumns struct {
	Id        string // 公告ID
	Title     string // 公告标题
	Type      string // 公告类型
	Tag       string // 标签
	Content   string // 公告内容
	Receiver  string // 接收者
	Remark    string // 备注
	Sort      string // 排序
	Status    string // 公告状态
	CreatedBy string // 发送人
	UpdatedBy string // 修改人
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// adminNoticeColumns holds the columns for the table hg_admin_notice.
var adminNoticeColumns = AdminNoticeColumns{
	Id:        "id",
	Title:     "title",
	Type:      "type",
	Tag:       "tag",
	Content:   "content",
	Receiver:  "receiver",
	Remark:    "remark",
	Sort:      "sort",
	Status:    "status",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewAdminNoticeDao creates and returns a new DAO object for table data access.
func NewAdminNoticeDao(handlers ...gdb.ModelHandler) *AdminNoticeDao {
	return &AdminNoticeDao{
		group:    "default",
		table:    "hg_admin_notice",
		columns:  adminNoticeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminNoticeDao) Columns() AdminNoticeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminNoticeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
