// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysGenCurdDemoDao is the data access object for the table hg_sys_gen_curd_demo.
type SysGenCurdDemoDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SysGenCurdDemoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SysGenCurdDemoColumns defines and stores column names for the table hg_sys_gen_curd_demo.
type SysGenCurdDemoColumns struct {
	Id          string // ID
	CategoryId  string // 分类ID
	Title       string // 标题
	Description string // 描述
	Content     string // 内容
	Image       string // 单图
	Attachfile  string // 附件
	CityId      string // 所在城市
	Switch      string // 显示开关
	Sort        string // 排序
	Status      string // 状态
	CreatedBy   string // 创建者
	UpdatedBy   string // 更新者
	DeletedBy   string // 删除者
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
}

// sysGenCurdDemoColumns holds the columns for the table hg_sys_gen_curd_demo.
var sysGenCurdDemoColumns = SysGenCurdDemoColumns{
	Id:          "id",
	CategoryId:  "category_id",
	Title:       "title",
	Description: "description",
	Content:     "content",
	Image:       "image",
	Attachfile:  "attachfile",
	CityId:      "city_id",
	Switch:      "switch",
	Sort:        "sort",
	Status:      "status",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	DeletedBy:   "deleted_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewSysGenCurdDemoDao creates and returns a new DAO object for table data access.
func NewSysGenCurdDemoDao(handlers ...gdb.ModelHandler) *SysGenCurdDemoDao {
	return &SysGenCurdDemoDao{
		group:    "default",
		table:    "hg_sys_gen_curd_demo",
		columns:  sysGenCurdDemoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysGenCurdDemoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysGenCurdDemoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysGenCurdDemoDao) Columns() SysGenCurdDemoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysGenCurdDemoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysGenCurdDemoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysGenCurdDemoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
