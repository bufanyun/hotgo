// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import "github.com/gogf/gf/v2/frame/g"

const (
	CreditTypeBalance  = "balance"  // 余额
	CreditTypeIntegral = "integral" // 积分
)

const (
	CreditGroupDecr            = "decr"             // 扣款
	CreditGroupIncr            = "incr"             // 加款
	CreditGroupOpDecr          = "op_decr"          // 操作扣款
	CreditGroupOpIncr          = "op_incr"          // 操作加款
	CreditGroupBalanceRecharge = "balance_recharge" // 余额充值
	CreditGroupBalanceRefund   = "balance_refund"   // 余额退款
	CreditGroupApplyCash       = "apply_cash"       // 申请提现
)

// CreditTypeOptions 变动类型
var CreditTypeOptions = []g.Map{
	{
		"key":       CreditTypeBalance,
		"value":     CreditTypeBalance,
		"label":     "余额",
		"listClass": "success",
	},
	{
		"key":       CreditTypeIntegral,
		"value":     CreditTypeIntegral,
		"label":     "积分",
		"listClass": "info",
	},
}

// CreditGroupOptions 变动分组
var CreditGroupOptions = []g.Map{
	{
		"key":       CreditGroupDecr,
		"value":     CreditGroupDecr,
		"label":     "扣款",
		"listClass": "warning",
	},
	{
		"key":       CreditGroupIncr,
		"value":     CreditGroupIncr,
		"label":     "加款",
		"listClass": "success",
	},
	{
		"key":       CreditGroupOpDecr,
		"value":     CreditGroupOpDecr,
		"label":     "操作扣款",
		"listClass": "warning",
	},
	{
		"key":       CreditGroupOpIncr,
		"value":     CreditGroupOpIncr,
		"label":     "操作加款",
		"listClass": "success",
	},
	{
		"key":       CreditGroupBalanceRefund,
		"value":     CreditGroupBalanceRefund,
		"label":     "余额退款",
		"listClass": "warning",
	},
	{
		"key":       CreditGroupBalanceRecharge,
		"value":     CreditGroupBalanceRecharge,
		"label":     "余额充值",
		"listClass": "success",
	},
	{
		"key":       CreditGroupApplyCash,
		"value":     CreditGroupApplyCash,
		"label":     "申请提现",
		"listClass": "info",
	},
}
