package adminin

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/utility/encrypt"
)

// RegisterInp 账号注册
type RegisterInp struct {
	Username   string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password   string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Mobile     string `json:"mobile" v:"required|phone-loose#手机号不能为空|手机号格式不正确" dc:"手机号"`
	Code       string `json:"code" v:"required#验证码不能为空"  dc:"验证码"`
	InviteCode string `json:"inviteCode" dc:"邀请码"`
}

func (in *RegisterInp) Filter(ctx context.Context) (err error) {
	// 解密密码
	str, err := gbase64.Decode([]byte(in.Password))
	if err != nil {
		return err
	}

	str, err = encrypt.AesECBDecrypt(str, consts.RequestEncryptKey)
	if err != nil {
		return err
	}

	password := string(str)

	if err = g.Validator().Data(password).Rules("password").Messages("密码长度在6~18之间").Run(ctx); err != nil {
		return
	}

	in.Password = password
	return
}

// LoginModel 统一登录响应
type LoginModel struct {
	Id       int64  `json:"id"              dc:"用户ID"`
	Username string `json:"username"        dc:"用户名"`
	Token    string `json:"token"           dc:"登录token"`
	Expires  int64  `json:"expires"         dc:"登录有效期"`
}

// AccountLoginInp 账号登录
type AccountLoginInp struct {
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Cid      string `json:"cid"  dc:"验证码ID"`
	Code     string `json:"code" dc:"验证码"`
	IsLock   bool   `json:"isLock"  dc:"是否为锁屏状态"`
}

// MobileLoginInp 手机号登录
type MobileLoginInp struct {
	Mobile string `json:"mobile" v:"required|phone-loose#手机号不能为空|手机号格式不正确" dc:"手机号"`
	Code   string `json:"code" v:"required#验证码不能为空"  dc:"验证码"`
}

// MemberLoginPermissions 登录用户角色信息
type MemberLoginPermissions []string

// MemberLoginStatInp 用户登录统计
type MemberLoginStatInp struct {
	MemberId int64
}

type MemberLoginStatModel struct {
	LoginCount  int         `json:"loginCount"  dc:"登录次数"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" dc:"最后登录时间"`
	LastLoginIp string      `json:"lastLoginIp" dc:"最后登录IP"`
}
