package commonin

// WechatAuthorizeInp 微信用户授权
type WechatAuthorizeInp struct {
	Type         string `json:"type"  v:"required#授权类型不能为空" dc:"授权类型"`
	SyncRedirect string `json:"syncRedirect"  v:"required#同步跳转地址不能为空" dc:"同步跳转地址"`
}

type WechatAuthorizeModel struct{}

// WechatAuthorizeCallInp 微信用户授权回调
type WechatAuthorizeCallInp struct {
	Code  string `json:"code"   dc:"code作为换取access_token的票据"`
	State string `json:"state"   dc:"state"`
}

type WechatAuthorizeCallModel struct{}
