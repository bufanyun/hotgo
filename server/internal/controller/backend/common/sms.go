// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/common"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var Sms = new(cSms)

type cSms struct{}

// SendTest 发送测试短信
func (c *cSms) SendTest(ctx context.Context, req *common.SendTestSmsReq) (res *common.SendTestSmsRes, err error) {
	var in sysin.SendCodeInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return nil, err
	}

	if err = service.SysSmsLog().SendCode(ctx, in); err != nil {
		return nil, err
	}
	return
}
