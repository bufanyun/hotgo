// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import (
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

func init() {
	dict.RegisterEnums("creditType", "资金变动类型", CreditTypeOptions)
	dict.RegisterEnums("creditGroup", "资金变动分组", CreditGroupOptions)
}

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
var CreditTypeOptions = []*model.Option{
	dict.GenSuccessOption(CreditTypeBalance, "余额"),
	dict.GenInfoOption(CreditTypeIntegral, "积分"),
}

// CreditGroupOptions 变动分组
var CreditGroupOptions = []*model.Option{
	dict.GenWarningOption(CreditGroupDecr, "扣款"),
	dict.GenSuccessOption(CreditGroupIncr, "加款"),
	dict.GenWarningOption(CreditGroupOpDecr, "操作扣款"),
	dict.GenSuccessOption(CreditGroupOpIncr, "操作加款"),
	dict.GenWarningOption(CreditGroupBalanceRefund, "余额退款"),
	dict.GenSuccessOption(CreditGroupBalanceRecharge, "余额充值"),
	dict.GenInfoOption(CreditGroupApplyCash, "申请提现"),
}
