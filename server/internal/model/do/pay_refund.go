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
	Id            interface{} // 主键ID
	MemberId      interface{} // 会员ID
	AppId         interface{} // 应用ID
	OrderSn       interface{} // 业务订单号
	RefundTradeNo interface{} // 退款交易号
	RefundMoney   interface{} // 退款金额
	RefundWay     interface{} // 退款方式
	Ip            interface{} // 申请者IP
	Reason        interface{} // 申请退款原因
	Remark        interface{} // 退款备注
	Status        interface{} // 退款状态
	CreatedAt     *gtime.Time // 申请时间
	UpdatedAt     *gtime.Time // 更新时间
}
