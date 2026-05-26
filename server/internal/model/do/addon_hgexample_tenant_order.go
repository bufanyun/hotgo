// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTenantOrder is the golang structure of table hg_addon_hgexample_tenant_order for DAO operations like Where/Data.
type AddonHgexampleTenantOrder struct {
	g.Meta      `orm:"table:hg_addon_hgexample_tenant_order, do:true"`
	Id          any         // 主键
	TenantId    any         // 租户ID
	MerchantId  any         // 商户ID
	UserId      any         // 用户ID
	ProductName any         // 购买产品
	OrderSn     any         // 订单号
	Money       any         // 充值金额
	Remark      any         // 备注
	Status      any         // 订单状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
