// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysGenCodesDao is the data access object for the table hg_sys_gen_codes.
type SysGenCodesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysGenCodesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysGenCodesColumns defines and stores column names for the table hg_sys_gen_codes.
type SysGenCodesColumns struct {
	Id            string // 生成ID
	GenType       string // 生成类型
	GenTemplate   string // 生成模板
	VarName       string // 实体命名
	Options       string // 配置选项
	DbName        string // 数据库名称
	TableName     string // 主表名称
	TableComment  string // 主表注释
	DaoName       string // 主表dao模型
	MasterColumns string // 主表字段
	AddonName     string // 插件名称
	Status        string // 生成状态
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// sysGenCodesColumns holds the columns for the table hg_sys_gen_codes.
var sysGenCodesColumns = SysGenCodesColumns{
	Id:            "id",
	GenType:       "gen_type",
	GenTemplate:   "gen_template",
	VarName:       "var_name",
	Options:       "options",
	DbName:        "db_name",
	TableName:     "table_name",
	TableComment:  "table_comment",
	DaoName:       "dao_name",
	MasterColumns: "master_columns",
	AddonName:     "addon_name",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSysGenCodesDao creates and returns a new DAO object for table data access.
func NewSysGenCodesDao(handlers ...gdb.ModelHandler) *SysGenCodesDao {
	return &SysGenCodesDao{
		group:    "default",
		table:    "hg_sys_gen_codes",
		columns:  sysGenCodesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysGenCodesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysGenCodesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysGenCodesDao) Columns() SysGenCodesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysGenCodesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysGenCodesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysGenCodesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
