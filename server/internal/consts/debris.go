// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

// 碎片

const (
	DemoTips              = "演示系统已隐藏" // 演示系统敏感数据打码
	NilJsonToString       = "{}"      // 空json初始化值
	RegionSpilt           = " / "     // 地区分隔符
	Unknown               = "Unknown" // Unknown
	SuperRoleKey          = "super"   // 超管角色唯一标识符，通过角色验证超管
	MaxServeLogContentLen = 2048      // 最大保留服务日志内容大小
	SysDefaultLanguage    = "zh_CN"   // 系统默认语言，当配置文件没有国际化配置时生效
)

// curd.
const (
	DefaultPage     = 10 // 默认列表分页加载数量
	DefaultPageSize = 1  // 默认列表分页加载页码
	MaxSortIncr     = 10 // 最大排序值增量
)

// TenantField 租户字段
const (
	TenantId   = "tenant_id"   // 租户ID
	MerchantId = "merchant_id" // 商户ID
	UserId     = "user_id"     // 用户ID
)

const (
	DBMysql = "mysql"
	DBPgsql = "pgsql"
)
