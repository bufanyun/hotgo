package apiForm

import (
	"github.com/bufanyun/hotgo/app/model"
	"github.com/gogf/gf/v2/frame/g"
)

//  注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/login/logout" method:"get" tags:"登录" summary:"注销登录"`
}
type LoginLogoutRes struct{}

//  登录效验
type LoginCheckReq struct {
	g.Meta `path:"/login/check" method:"get" tags:"登录" summary:"登录效验"`
}
type LoginCheckRes struct {
	IsValidCodeLogin bool   `json:"isValidCodeLogin"      description:"是否验证码"`
	Message          string `json:"message"      description:"消息"`
	Result           string `json:"result"      description:"响应"`
	// Sessionid        string `json:"sessionid"      description:"sessionid"`
}

//  提交登录
type LoginReq struct {
	g.Meta   `path:"/login/sign" method:"post" tags:"登录" summary:"提交登录"`
	Username string `json:"username" v:"required#用户名不能为空" description:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" description:"密码"`
	Cid      string `json:"cid" v:"required#验证码ID不能为空" description:"验证码ID"`
	Code     string `json:"code" v:"required#验证码不能为空" description:"验证码"`
	Device   string `json:"device"  description:"登录设备"`
}
type LoginRes struct {
	model.Identity
	Token string `json:"token" v:""  description:"登录token"`
}