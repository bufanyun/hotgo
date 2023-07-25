// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestCategory is the golang structure of table hg_test_category for DAO operations like Where/Data.
type TestCategory struct {
	g.Meta      `orm:"table:hg_test_category, do:true"`
	Id          interface{} // 分类ID
	Name        interface{} // 分类名称
	Description interface{} // 描述
	Sort        interface{} // 排序
	Remark      interface{} // 备注
	Status      interface{} // 状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
