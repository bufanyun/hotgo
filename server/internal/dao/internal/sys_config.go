// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigDao is the data access object for the table hg_sys_config.
type SysConfigDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysConfigColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysConfigColumns defines and stores column names for the table hg_sys_config.
type SysConfigColumns struct {
	Id           string // 配置ID
	Group        string // 配置分组
	Name         string // 参数名称
	Type         string // 键值类型:string,int,uint,bool,TIMESTAMP,date
	Key          string // 参数键名
	Value        string // 参数键值
	DefaultValue string // 默认值
	Sort         string // 排序
	Tip          string // 变量描述
	IsDefault    string // 是否为系统默认
	Status       string // 状态
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// sysConfigColumns holds the columns for the table hg_sys_config.
var sysConfigColumns = SysConfigColumns{
	Id:           "id",
	Group:        "group",
	Name:         "name",
	Type:         "type",
	Key:          "key",
	Value:        "value",
	DefaultValue: "default_value",
	Sort:         "sort",
	Tip:          "tip",
	IsDefault:    "is_default",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSysConfigDao creates and returns a new DAO object for table data access.
func NewSysConfigDao(handlers ...gdb.ModelHandler) *SysConfigDao {
	return &SysConfigDao{
		group:    "default",
		table:    "hg_sys_config",
		columns:  sysConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysConfigDao) Columns() SysConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
