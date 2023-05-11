package global

import (
	"hotgo/internal/consts"
	"hotgo/internal/library/payment"
	"hotgo/internal/service"
)

// 注册支付成功回调方法
func payNotifyCall() {
	payment.RegisterNotifyCall(consts.OrderGroupAdminOrder, service.AdminOrder().PayNotify) // 后台充值订单
	// ...
}
