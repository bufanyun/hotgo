// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLogDao is the data access object for table hg_sys_log.
type SysLogDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SysLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysLogColumns defines and stores column names for table hg_sys_log.
type SysLogColumns struct {
	Id         string //
	AppId      string // 应用id
	MerchantId string // 商户id
	MemberId   string // 用户id
	Method     string // 提交类型
	Module     string // 模块
	Url        string // 提交url
	GetData    string // get数据
	PostData   string // post数据
	HeaderData string // header数据
	Ip         string // ip地址
	ProvinceId string // 省编码
	CityId     string // 市编码
	ErrorCode  string // 报错code
	ErrorMsg   string // 报错信息
	ErrorData  string // 报错日志
	ReqId      string // 对外id
	Timestamp  string // 响应时间
	UserAgent  string // UA信息
	TakeUpTime string // 请求耗时
	Status     string // 状态
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

//  sysLogColumns holds the columns for table hg_sys_log.
var sysLogColumns = SysLogColumns{
	Id:         "id",
	AppId:      "app_id",
	MerchantId: "merchant_id",
	MemberId:   "member_id",
	Method:     "method",
	Module:     "module",
	Url:        "url",
	GetData:    "get_data",
	PostData:   "post_data",
	HeaderData: "header_data",
	Ip:         "ip",
	ProvinceId: "province_id",
	CityId:     "city_id",
	ErrorCode:  "error_code",
	ErrorMsg:   "error_msg",
	ErrorData:  "error_data",
	ReqId:      "req_id",
	Timestamp:  "timestamp",
	UserAgent:  "user_agent",
	TakeUpTime: "take_up_time",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewSysLogDao creates and returns a new DAO object for table data access.
func NewSysLogDao() *SysLogDao {
	return &SysLogDao{
		group:   "default",
		table:   "hg_sys_log",
		columns: sysLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysLogDao) Columns() SysLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
