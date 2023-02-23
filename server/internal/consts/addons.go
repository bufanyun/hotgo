package consts

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

var AddonsGroupNameMap = map[int]string{
	AddonsGroupPlug:       "功能扩展",
	AddonsGroupBusiness:   "主要业务",
	AddonsGroupThirdParty: "第三方插件",
	AddonsGroupMiniApp:    "小程序",
	AddonsGroupCustomer:   "客户关系",
	AddonsGroupActivity:   "营销及活动",
	AddonsGroupServices:   "常用服务及工具",
	AddonsGroupBiz:        "行业解决方案",
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

var AddonsInstallStatusNameMap = map[int]string{
	AddonsInstallStatusOk: "已安装",
	AddonsInstallStatusNo: "未安装",
	AddonsInstallStatusUn: "已卸载",
}
