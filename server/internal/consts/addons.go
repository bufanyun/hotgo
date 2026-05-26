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
	dict.RegisterEnums("addonsGroupOptions", "插件分组", AddonsGroupOptions)
	dict.RegisterEnums("addonsInstallStatus", "插件安装状态", AddonsInstallStatusOptions)
	dict.RegisterEnums("addonsExtend", "插件扩展", AddonsExtendOptions)
}

const (
	AddonsTag = "addons_" // 插件标签前缀
	AddonsDir = "addons"  // 插件路径
)

const (
	AddonsGroupPlug       = 1 // 功能扩展
	AddonsGroupBusiness   = 2 // 主要业务
	AddonsGroupThirdParty = 3 // 第三方插件
	AddonsGroupMiniApp    = 4 // 小程序
	AddonsGroupCustomer   = 5 // 客户关系
	AddonsGroupActivity   = 6 // 营销及活动
	AddonsGroupServices   = 7 // 常用服务及工具
	AddonsGroupBiz        = 8 // 行业解决方案
)

// AddonsGroupOptions 插件分组选项
var AddonsGroupOptions = []*model.Option{
	dict.GenInfoOption(AddonsGroupPlug, "功能扩展"),
	dict.GenInfoOption(AddonsGroupBusiness, "主要业务"),
	dict.GenInfoOption(AddonsGroupThirdParty, "第三方插件"),
	dict.GenInfoOption(AddonsGroupMiniApp, "小程序"),
	dict.GenInfoOption(AddonsGroupCustomer, "客户关系"),
	dict.GenInfoOption(AddonsGroupActivity, "营销及活动"),
	dict.GenInfoOption(AddonsGroupServices, "常用服务及工具"),
	dict.GenInfoOption(AddonsGroupBiz, "行业解决方案"),
}

var AddonsGroupIconMap = map[int]string{
	AddonsGroupPlug:       "AppstoreAddOutlined",
	AddonsGroupBusiness:   "FireOutlined",
	AddonsGroupThirdParty: "ApiOutlined",
	AddonsGroupMiniApp:    "RocketOutlined",
	AddonsGroupCustomer:   "UserSwitchOutlined",
	AddonsGroupActivity:   "TagOutlined",
	AddonsGroupServices:   "ToolOutlined",
	AddonsGroupBiz:        "CheckCircleOutlined",
}

const (
	AddonsInstallStatusOk = 1 // 已安装
	AddonsInstallStatusNo = 2 // 未安装
	AddonsInstallStatusUn = 3 // 已卸载
)

// AddonsInstallStatusOptions 插件安装状态
var AddonsInstallStatusOptions = []*model.Option{
	dict.GenInfoOption(AddonsInstallStatusOk, "已安装"),
	dict.GenInfoOption(AddonsInstallStatusNo, "未安装"),
	dict.GenInfoOption(AddonsInstallStatusUn, "已卸载"),
}

const (
	AddonsExtendResourcePublic   = "resourcePublic"
	AddonsExtendResourceTemplate = "resourceTemplate"
)

// AddonsExtendOptions 插件扩展选项
var AddonsExtendOptions = []*model.Option{
	dict.GenInfoOption(AddonsExtendResourcePublic, "创建静态目录"),
	dict.GenInfoOption(AddonsExtendResourceTemplate, "创建模板目录"),
}
