// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOrder is the golang structure of table hg_admin_order for DAO operations like Where/Data.
type AdminOrder struct {
	g.Meta             `orm:"table:hg_admin_order, do:true"`
	Id                 any         // 主键
	MemberId           any         // 管理员id
	OrderType          any         // 订单类型
	ProductId          any         // 产品id
	OrderSn            any         // 关联订单号
	Money              any         // 充值金额
	Remark             any         // 备注
	RefundReason       any         // 退款原因
	RejectRefundReason any         // 拒绝退款原因
	Status             any         // 状态
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 修改时间
}
