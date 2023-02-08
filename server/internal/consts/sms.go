package consts

// 短信驱动

const (
	SmsDriveAliYun  = "aliyun"  // 阿里云
	SmsDriveTencent = "tencent" // 腾讯云
)

// 短信内置模板
const (
	SmsTemplateCode     = "code"     // 通用验证码
	SmsTemplateLogin    = "login"    // 登录
	SmsTemplateRegister = "register" // 注册
	SmsTemplateResetPwd = "resetPwd" // 重置密码
	SmsTemplateBind     = "bind"     // 绑定手机号
	SmsTemplateCash     = "cash"     // 申请提现
)

// 验证码状态
const (
	SmsStatusNotUsed = 1 // 未使用
	SmsStatusUsed    = 2 // 已使用
)
