// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package consts

// 定时任务

const (
	CronArgsKey      = "args" // 上下文变量名称
	CronSplitStr     = ","    // 变量分割符
	CronPolicySame   = 1      // 并行策略
	CronPolicySingle = 2      // 单例策略
	CronPolicyOnce   = 3      // 单次策略
	CronPolicyTimes  = 4      // 多次策略
)
