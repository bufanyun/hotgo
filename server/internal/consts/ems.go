// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import "github.com/gogf/gf/v2/frame/g"

// 邮件内置模板
const (
	EmsTemplateText     = "text"     // 通用文本
	EmsTemplateCode     = "code"     // 通用验证码
	EmsTemplateLogin    = "login"    // 登录
	EmsTemplateRegister = "register" // 注册
	EmsTemplateResetPwd = "resetPwd" // 重置密码
	EmsTemplateBind     = "bind"     // 绑定邮箱
	EmsTemplateCash     = "cash"     // 申请提现
)

var EmsSubjectMap = g.MapStrStr{
	EmsTemplateText:     "这是一封来自HotGo的邮件",
	EmsTemplateCode:     "验证码",
	EmsTemplateLogin:    "登录验证码",
	EmsTemplateRegister: "注册验证码",
	EmsTemplateResetPwd: "重置密码",
	EmsTemplateBind:     "绑定邮箱验证码",
	EmsTemplateCash:     "申请提现验证码",
}

// IsCodeEmsTemplate 是否是验证码类型的模板
func IsCodeEmsTemplate(template string) bool {
	return template == EmsTemplateCode ||
		template == EmsTemplateLogin ||
		template == EmsTemplateRegister ||
		template == EmsTemplateBind ||
		template == EmsTemplateCash
}
