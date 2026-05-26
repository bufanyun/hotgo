// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAttachment is the golang structure of table hg_sys_attachment for DAO operations like Where/Data.
type SysAttachment struct {
	g.Meta    `orm:"table:hg_sys_attachment, do:true"`
	Id        any         // 文件ID
	AppId     any         // 应用ID
	MemberId  any         // 管理员ID
	CateId    any         // 上传分类
	Drive     any         // 上传驱动
	Name      any         // 文件原始名
	Kind      any         // 上传类型
	MimeType  any         // 扩展类型
	NaiveType any         // NaiveUI类型
	Path      any         // 本地路径
	FileUrl   any         // url
	Size      any         // 文件大小
	Ext       any         // 扩展名
	Md5       any         // md5校验码
	Status    any         // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
