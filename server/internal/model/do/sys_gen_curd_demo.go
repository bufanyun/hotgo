// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCurdDemo is the golang structure of table hg_sys_gen_curd_demo for DAO operations like Where/Data.
type SysGenCurdDemo struct {
	g.Meta      `orm:"table:hg_sys_gen_curd_demo, do:true"`
	Id          interface{} // ID
	CategoryId  interface{} // 分类ID
	Title       interface{} // 标题
	Description interface{} // 描述
	Content     interface{} // 内容
	Image       interface{} // 单图
	Attachfile  interface{} // 附件
	CityId      interface{} // 所在城市
	Switch      interface{} // 显示开关
	Sort        interface{} // 排序
	Status      interface{} // 状态
	CreatedBy   interface{} // 创建者
	UpdatedBy   interface{} // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
