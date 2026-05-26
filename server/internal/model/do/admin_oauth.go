// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOauth is the golang structure of table hg_admin_oauth for DAO operations like Where/Data.
type AdminOauth struct {
	g.Meta       `orm:"table:hg_admin_oauth, do:true"`
	Id           any         // 主键
	MemberId     any         // 用户ID
	Unionid      any         // 唯一ID
	OauthClient  any         // 授权组别
	OauthOpenid  any         // 授权开放ID
	Sex          any         // 性别
	Nickname     any         // 昵称
	HeadPortrait any         // 头像
	Birthday     *gtime.Time // 生日
	Country      any         // 国家
	Province     any         // 省
	City         any         // 市
	Status       any         // 状态
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 修改时间
}
