// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProvincesDao is the data access object for table hg_sys_provinces.
type SysProvincesDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysProvincesColumns // columns contains all the column names of Table for convenient usage.
}

// SysProvincesColumns defines and stores column names for table hg_sys_provinces.
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

// sysProvincesColumns holds the columns for table hg_sys_provinces.
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
func NewSysProvincesDao() *SysProvincesDao {
	return &SysProvincesDao{
		group:   "default",
		table:   "hg_sys_provinces",
		columns: sysProvincesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProvincesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProvincesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProvincesDao) Columns() SysProvincesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProvincesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProvincesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProvincesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
