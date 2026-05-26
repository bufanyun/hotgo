// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProvincesDao is the data access object for the table hg_sys_provinces.
type SysProvincesDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysProvincesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysProvincesColumns defines and stores column names for the table hg_sys_provinces.
type SysProvincesColumns struct {
	Id        string // 省市区ID
	Title     string // 栏目名称
	Pinyin    string // 拼音
	Lng       string // 经度
	Lat       string // 纬度
	Pid       string // 父栏目
	Level     string // 关系树等级
	Tree      string // 关系
	Sort      string // 排序
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysProvincesColumns holds the columns for the table hg_sys_provinces.
var sysProvincesColumns = SysProvincesColumns{
	Id:        "id",
	Title:     "title",
	Pinyin:    "pinyin",
	Lng:       "lng",
	Lat:       "lat",
	Pid:       "pid",
	Level:     "level",
	Tree:      "tree",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysProvincesDao creates and returns a new DAO object for table data access.
func NewSysProvincesDao(handlers ...gdb.ModelHandler) *SysProvincesDao {
	return &SysProvincesDao{
		group:    "default",
		table:    "hg_sys_provinces",
		columns:  sysProvincesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysProvincesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysProvincesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysProvincesDao) Columns() SysProvincesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysProvincesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysProvincesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysProvincesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
