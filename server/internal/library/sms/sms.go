package sms

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/library/sms/aliyun"
	"hotgo/internal/library/sms/tencent"
	"hotgo/internal/model"
	"hotgo/internal/model/input/sysin"
)

// Drive 短信驱动
type Drive interface {
	SendCode(ctx context.Context, in sysin.SendCodeInp, config *model.SmsConfig) (err error)
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
		drive = &aliyun.Handle
	case consts.SmsDriveTencent:
		drive = &tencent.Handle
	default:
		panic(fmt.Sprintf("暂不支持短信驱动:%v", instanceName))
	}
	return drive
}
