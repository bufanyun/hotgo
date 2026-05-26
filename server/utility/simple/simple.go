// Package simple
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package simple

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"hotgo/internal/consts"
	"hotgo/utility/encrypt"
)

// RouterPrefix 获取应用路由前缀
func RouterPrefix(ctx context.Context, app string) string {
	return g.Cfg().MustGet(ctx, "router."+app+".prefix", "/"+app+"").String()
}

// FilterMaskDemo 过滤演示环境下的配置隐藏字段
func FilterMaskDemo(ctx context.Context, src g.Map) g.Map {
	if src == nil {
		return nil
	}

	if !IsDemo(ctx) {
		return src
	}

	for k := range src {
		if _, ok := consts.ConfigMaskDemoField[k]; ok {
			src[k] = consts.DemoTips
		}
	}
	return src
}

// DefaultErrorTplContent 获取默认的错误模板内容
func DefaultErrorTplContent(ctx context.Context) string {
	return gfile.GetContents(g.Cfg().MustGet(ctx, "viewer.paths").String() + "/error/default.html")
}

// DecryptText 解密文本
func DecryptText(text string) (string, error) {
	str, err := gbase64.Decode([]byte(text))
	if err != nil {
		return "", err
	}

	str, err = encrypt.AesECBDecrypt(str, consts.RequestEncryptKey)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

// CheckPassword 检查密码
func CheckPassword(input, salt, hash string) (err error) {
	// 解密密码
	password, err := DecryptText(input)
	if err != nil {
		return err
	}

	if hash != gmd5.MustEncryptString(password+salt) {
		err = gerror.New("用户名或密码错误")
		return
	}
	return
}

// GetHeaderLocale 获取请求头语言设置
// gf支持格式：en/ja/ru/zh-CN/zh-TW
func GetHeaderLocale(ctx context.Context) (lang string) {
	lang = g.Cfg().MustGet(ctx, "system.i18n.defaultLanguage", consts.SysDefaultLanguage).String()
	// 没有开启国际化，使用默认语言
	if !g.Cfg().MustGet(ctx, "system.i18n.switch", true).Bool() {
		return
	}

	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return
	}
	locale := r.Header.Get("Locale")
	// 简体中文
	if locale == "zh-CN" || locale == "zh-Hans" || locale == "zh" || locale == "ZH" {
		lang = "zh-CN"
		return
	}
	// 繁体
	if locale == "zh-TW" || locale == "zh-Hant" {
		lang = "zh-TW"
		return
	}
	// 英文
	if locale == "en" || locale == "EN" {
		lang = "en"
		return
	}
	// 更多语言
	// ...
	return
}

// SafeGo 安全的调用协程，遇到错误时输出错误日志而不是抛出panic
func SafeGo(ctx context.Context, f func(ctx context.Context), lv ...int) {
	g.Go(ctx, f, func(ctx context.Context, err error) {
		var level = glog.LEVEL_ERRO
		if len(lv) > 0 {
			level = lv[0]
		}
		Logf(level, ctx, "SafeGo exec failed:%+v", err)
	})
}

func Logf(level int, ctx context.Context, format string, v ...interface{}) {
	switch level {
	case glog.LEVEL_DEBU:
		g.Log().Debugf(ctx, format, v...)
	case glog.LEVEL_INFO:
		g.Log().Infof(ctx, format, v...)
	case glog.LEVEL_NOTI:
		g.Log().Noticef(ctx, format, v...)
	case glog.LEVEL_WARN:
		g.Log().Warningf(ctx, format, v...)
	case glog.LEVEL_ERRO:
		g.Log().Errorf(ctx, format, v...)
	case glog.LEVEL_CRIT:
		g.Log().Criticalf(ctx, format, v...)
	case glog.LEVEL_PANI:
		g.Log().Panicf(ctx, format, v...)
	case glog.LEVEL_FATA:
		g.Log().Fatalf(ctx, format, v...)
	default:
		g.Log().Errorf(ctx, format, v...)
	}
}
