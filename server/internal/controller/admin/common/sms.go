// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/common"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
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
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysSmsLog().SendCode(ctx, in)
	return
}

// SendBindSms 发送换绑短信
func (c *cSms) SendBindSms(ctx context.Context, req *common.SendBindSmsReq) (res *common.SendBindSmsRes, err error) {
	var (
		memberId = contexts.GetUserId(ctx)
		models   *entity.AdminMember
	)

	if memberId <= 0 {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	err = g.Model("admin_member").
		Fields("mobile").
		Where("id", memberId).
		Scan(&models)
	if err != nil {
		return
	}

	if models == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	if models.Mobile == "" {
		err = gerror.New("未绑定手机号无需发送")
		return
	}

	err = service.SysSmsLog().SendCode(ctx, sysin.SendCodeInp{
		Event:  consts.SmsTemplateBind,
		Mobile: models.Mobile,
	})
	return
}
