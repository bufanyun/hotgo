package consts

// 授权类型
const (
	WechatAuthorizeOpenId    = "openId"    // 设置openid
	WechatAuthorizeBindLogin = "bindLogin" // 绑定微信登录
)

// 应用授权作用域
const (
	WechatScopeBase     = "snsapi_base"     // 只获取openid，无需用户授权
	WechatScopeUserinfo = "snsapi_userinfo" // 获取用户信息，需要授权
)
