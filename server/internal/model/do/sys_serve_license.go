// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLicense is the golang structure of table hg_sys_serve_license for DAO operations like Where/Data.
type SysServeLicense struct {
	g.Meta       `orm:"table:hg_sys_serve_license, do:true"`
	Id           any         // 许可ID
	Group        any         // 分组
	Name         any         // 许可名称
	Appid        any         // 应用ID
	SecretKey    any         // 应用秘钥
	RemoteAddr   any         // 最后连接地址
	OnlineLimit  any         // 在线限制
	LoginTimes   any         // 登录次数
	LastLoginAt  *gtime.Time // 最后登录时间
	LastActiveAt *gtime.Time // 最后心跳
	Routes       *gjson.Json // 路由表，空使用默认分组路由
	AllowedIps   any         // IP白名单
	EndAt        *gtime.Time // 授权有效期
	Remark       any         // 备注
	Status       any         // 状态
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 修改时间
}
