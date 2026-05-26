// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProvinces is the golang structure of table hg_sys_provinces for DAO operations like Where/Data.
type SysProvinces struct {
	g.Meta    `orm:"table:hg_sys_provinces, do:true"`
	Id        any         // 省市区ID
	Title     any         // 栏目名称
	Pinyin    any         // 拼音
	Lng       any         // 经度
	Lat       any         // 纬度
	Pid       any         // 父栏目
	Level     any         // 关系树等级
	Tree      any         // 关系
	Sort      any         // 排序
	Status    any         // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
