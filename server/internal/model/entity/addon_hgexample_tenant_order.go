// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTenantOrder is the golang structure for table addon_hgexample_tenant_order.
type AddonHgexampleTenantOrder struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`
	TenantId    int64       `json:"tenantId"    orm:"tenant_id"    description:"租户ID"`
	MerchantId  int64       `json:"merchantId"  orm:"merchant_id"  description:"商户ID"`
	UserId      int64       `json:"userId"      orm:"user_id"      description:"用户ID"`
	ProductName string      `json:"productName" orm:"product_name" description:"购买产品"`
	OrderSn     string      `json:"orderSn"     orm:"order_sn"     description:"订单号"`
	Money       float64     `json:"money"       orm:"money"        description:"充值金额"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	Status      int         `json:"status"      orm:"status"       description:"订单状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"修改时间"`
}
