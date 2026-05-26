// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayRefund is the golang structure of table hg_pay_refund for DAO operations like Where/Data.
type PayRefund struct {
	g.Meta        `orm:"table:hg_pay_refund, do:true"`
	Id            any         // 主键ID
	MemberId      any         // 会员ID
	AppId         any         // 应用ID
	OrderSn       any         // 业务订单号
	RefundTradeNo any         // 退款交易号
	RefundMoney   any         // 退款金额
	RefundWay     any         // 退款方式
	Ip            any         // 申请者IP
	Reason        any         // 申请退款原因
	Remark        any         // 退款备注
	Status        any         // 退款状态
	CreatedAt     *gtime.Time // 申请时间
	UpdatedAt     *gtime.Time // 更新时间
}
