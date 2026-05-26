// Package sms
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sms

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/sysin"
)

// Drive 短信驱动
type Drive interface {
	SendCode(ctx context.Context, in *sysin.SendCodeInp) (err error)
}

func New(name ...string) Drive {
	var (
		instanceName = consts.SmsDriveAliYun
		drive        Drive
	)

	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}

	switch instanceName {
	case consts.SmsDriveAliYun:
		drive = &AliYunDrive{}
	case consts.SmsDriveTencent:
		drive = &TencentDrive{}
	default:
		panic(fmt.Sprintf("暂不支持短信驱动:%v", instanceName))
	}
	return drive
}
