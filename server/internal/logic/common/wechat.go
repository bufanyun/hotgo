// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"hotgo/api/admin/common"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/response"
	"hotgo/internal/library/token"
	"hotgo/internal/library/wechat"
	"hotgo/internal/model/input/commonin"
	"hotgo/internal/service"
	"time"
)

type sCommonWechat struct {
	temp map[string]*AuthorizeCallState
}

// AuthorizeCallState 微信授权回调参数
type AuthorizeCallState struct {
	State        string      `json:"state"         dc:"state"`
	MemberId     int64       `json:"memberId"      dc:"管理员ID"`
	Type         string      `json:"type"          dc:"授权类型"`
	SyncRedirect string      `json:"syncRedirect"  dc:"同步跳转地址"`
	Tick         *gtime.Time `json:"tick"          dc:"标记授权时间"`
}

func NewCommonWechat() *sCommonWechat {
	return &sCommonWechat{
		temp: make(map[string]*AuthorizeCallState),
	}
}

func init() {
	serv := NewCommonWechat()
	service.RegisterCommonWechat(serv)
	_, _ = gcron.Add(gctx.New(), "@every 300s", serv.CleanTempMap, "WechatCleanTempMap")
}

// Authorize 微信用户授权
func (s *sCommonWechat) Authorize(ctx context.Context, in *commonin.WechatAuthorizeInp) (res *commonin.WechatAuthorizeModel, err error) {
	basic, err := service.SysConfig().GetBasic(ctx)
	if err != nil {
		return
	}

	var (
		request     = g.RequestFromCtx(ctx)
		prefix      = g.Cfg().MustGet(ctx, "router.admin.prefix", "/admin").String()
		path        = gmeta.Get(common.WechatAuthorizeCallReq{}, "path").String()
		redirectUri = basic.Domain + prefix + path
		memberId    = contexts.GetUserId(ctx)
		state       = s.GetCacheKey(in.Type, token.GetAuthKey(token.GetAuthorization(request)))
		scope       string
	)

	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	switch in.Type {
	case consts.WechatAuthorizeOpenId: // 设置openid
		scope = consts.WechatScopeBase
	case consts.WechatAuthorizeBindLogin: // 绑定微信登录
		scope = consts.WechatScopeUserinfo
	default:
		err = gerror.New("无效的授权方式！")
		return
	}

	url, err := wechat.GetOauthURL(ctx, redirectUri, scope, state)
	if err != nil {
		return
	}

	s.temp[state] = &AuthorizeCallState{
		State:        state,
		MemberId:     memberId,
		Type:         in.Type,
		SyncRedirect: in.SyncRedirect,
		Tick:         gtime.Now(),
	}
	response.Redirect(g.RequestFromCtx(ctx), url)
	return
}

func (s *sCommonWechat) AuthorizeCall(ctx context.Context, in *commonin.WechatAuthorizeCallInp) (res *commonin.WechatAuthorizeCallModel, err error) {
	data, ok := s.temp[in.State]
	if !ok || data == nil {
		err = gerror.New("授权无效或已过期，请重试")
		return
	}

	defer delete(s.temp, in.State)

	tk, err := wechat.GetUserAccessToken(ctx, in.Code)
	if err != nil {
		return
	}

	switch data.Type {
	case consts.WechatAuthorizeOpenId: // 设置openid
		_ = cache.Instance().Set(ctx, data.State, tk.OpenID, time.Hour*24*7)
	case consts.WechatAuthorizeBindLogin: // 绑定微信登录
		// ...
	default:
		err = gerror.New("无效的授权方式！")
		return
	}

	response.Redirect(g.RequestFromCtx(ctx), data.SyncRedirect)
	return
}

// GetOpenId 从缓存中获取临时openid
func (s *sCommonWechat) GetOpenId(ctx context.Context) (openId string, err error) {
	request := ghttp.RequestFromCtx(ctx)
	key := s.GetCacheKey(consts.WechatAuthorizeOpenId, token.GetAuthKey(token.GetAuthorization(request)))
	date, err := cache.Instance().Get(ctx, key)
	if err != nil {
		err = gerror.Newf("GetOpenId err:%+v", err.Error())
		return
	}
	openId = date.String()
	return
}

func (s *sCommonWechat) GetCacheKey(typ, ak string) string {
	return fmt.Sprintf("%v:%v", typ, ak)
}

// CleanTempMap 清理临时map
func (s *sCommonWechat) CleanTempMap(ctx context.Context) {
	t := gtime.Now().Add(time.Second * 600)
	for _, state := range s.temp {
		if state.Tick.After(t) {
			delete(s.temp, state.State)
		}
	}
}
