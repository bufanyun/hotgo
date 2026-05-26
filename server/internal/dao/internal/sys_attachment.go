// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAttachmentDao is the data access object for the table hg_sys_attachment.
type SysAttachmentDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SysAttachmentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SysAttachmentColumns defines and stores column names for the table hg_sys_attachment.
type SysAttachmentColumns struct {
	Id        string // 文件ID
	AppId     string // 应用ID
	MemberId  string // 管理员ID
	CateId    string // 上传分类
	Drive     string // 上传驱动
	Name      string // 文件原始名
	Kind      string // 上传类型
	MimeType  string // 扩展类型
	NaiveType string // NaiveUI类型
	Path      string // 本地路径
	FileUrl   string // url
	Size      string // 文件大小
	Ext       string // 扩展名
	Md5       string // md5校验码
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// sysAttachmentColumns holds the columns for the table hg_sys_attachment.
var sysAttachmentColumns = SysAttachmentColumns{
	Id:        "id",
	AppId:     "app_id",
	MemberId:  "member_id",
	CateId:    "cate_id",
	Drive:     "drive",
	Name:      "name",
	Kind:      "kind",
	MimeType:  "mime_type",
	NaiveType: "naive_type",
	Path:      "path",
	FileUrl:   "file_url",
	Size:      "size",
	Ext:       "ext",
	Md5:       "md5",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysAttachmentDao creates and returns a new DAO object for table data access.
func NewSysAttachmentDao(handlers ...gdb.ModelHandler) *SysAttachmentDao {
	return &SysAttachmentDao{
		group:    "default",
		table:    "hg_sys_attachment",
		columns:  sysAttachmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAttachmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAttachmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAttachmentDao) Columns() SysAttachmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAttachmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAttachmentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysAttachmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
