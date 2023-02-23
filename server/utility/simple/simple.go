// Package simple
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package simple

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/utility/encrypt"
)

// CheckPassword 检查密码
func CheckPassword(input, salt, hash string) (err error) {
	// 解密密码
	password, err := gbase64.Decode([]byte(input))
	if err != nil {
		return err
	}

	password, err = encrypt.AesECBDecrypt(password, consts.RequestEncryptKey)
	if err != nil {
		return err
	}

	if hash != gmd5.MustEncryptString(string(password)+salt) {
		err = gerror.New("用户密码不正确")
		return
	}
	return
}

func SafeGo(ctx context.Context, f func(ctx context.Context), level ...interface{}) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				var newLevel = glog.LEVEL_ERRO
				if len(level) > 0 {
					newLevel = gconv.Int(level[0])
				}
				Logf(newLevel, ctx, "SafeGo failed err:%+v", err)
			}
		}()

		f(ctx)
	}()
}

func Logf(level int, ctx context.Context, format string, v ...interface{}) {
	switch level {
	case glog.LEVEL_DEBU:
		g.Log().Debugf(ctx, format, v)
	case glog.LEVEL_INFO:
		g.Log().Infof(ctx, format, v)
	case glog.LEVEL_NOTI:
		g.Log().Noticef(ctx, format, v)
	case glog.LEVEL_WARN:
		g.Log().Warningf(ctx, format, v)
	case glog.LEVEL_ERRO:
		g.Log().Errorf(ctx, format, v)
	case glog.LEVEL_CRIT:
		g.Log().Critical(ctx, format, v)
	case glog.LEVEL_PANI:
		g.Log().Panicf(ctx, format, v)
	case glog.LEVEL_FATA:
		g.Log().Fatalf(ctx, format, v)
	default:
		g.Log().Error(ctx, "Logf level not find")
	}
}

func Log(level int, ctx context.Context, v ...interface{}) {
	switch level {
	case glog.LEVEL_DEBU:
		g.Log().Debug(ctx, v)
	case glog.LEVEL_INFO:
		g.Log().Info(ctx, v)
	case glog.LEVEL_NOTI:
		g.Log().Notice(ctx, v)
	case glog.LEVEL_WARN:
		g.Log().Warning(ctx, v)
	case glog.LEVEL_ERRO:
		g.Log().Error(ctx, v)
	case glog.LEVEL_CRIT:
		g.Log().Critical(ctx, v)
	case glog.LEVEL_PANI:
		g.Log().Panic(ctx, v)
	case glog.LEVEL_FATA:
		g.Log().Fatal(ctx, v)
	default:
		g.Log().Error(ctx, "Logf level not find")
	}
}
