// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"hotgo/api/backend/common"
	"hotgo/internal/library/ems"
	"hotgo/internal/service"
)

var Ems = new(cEms)

type cEms struct{}

// SendTest 发送测试邮件
func (c *cEms) SendTest(ctx context.Context, req *common.SendTestEmailReq) (res *common.SendTestEmailRes, err error) {

	conf, err := service.SysConfig().GetSmtp(ctx)
	if err != nil {
		return
	}

	if err = ems.SendTestMail(conf, req.To); err != nil {
		return nil, err
	}
	return
}
