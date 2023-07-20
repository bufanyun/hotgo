// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAddonsConfigDao is the data access object for table hg_sys_addons_config.
type SysAddonsConfigDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SysAddonsConfigColumns // columns contains all the column names of Table for convenient usage.
}

// SysAddonsConfigColumns defines and stores column names for table hg_sys_addons_config.
type SysAddonsConfigColumns struct {
	Id           string // 配置ID
	AddonName    string // 插件名称
	Group        string // 分组
	Name         string // 参数名称
	Type         string // 键值类型:string,int,uint,bool,datetime,date
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

// sysAddonsConfigColumns holds the columns for table hg_sys_addons_config.
var sysAddonsConfigColumns = SysAddonsConfigColumns{
	Id:           "id",
	AddonName:    "addon_name",
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

// NewSysAddonsConfigDao creates and returns a new DAO object for table data access.
func NewSysAddonsConfigDao() *SysAddonsConfigDao {
	return &SysAddonsConfigDao{
		group:   "default",
		table:   "hg_sys_addons_config",
		columns: sysAddonsConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAddonsConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAddonsConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAddonsConfigDao) Columns() SysAddonsConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAddonsConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAddonsConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAddonsConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
