// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/backend/common"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var Ems = new(cEms)

type cEms struct{}

// SendTest 发送测试邮件
func (c *cEms) SendTest(ctx context.Context, req *common.SendTestEmailReq) (res *common.SendTestEmailRes, err error) {
	err = service.SysEmsLog().Send(ctx, sysin.SendEmsInp{
		Event: consts.EmsTemplateText,
		Email: req.To,
		Content: `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="iso-8859-15">
				<title>这是一封来自HotGo的测试邮件</title>
			</head>
			<body>
				这是您通过HotGo后台发送的测试邮件。当你收到这封邮件的时候，说明已经联调成功了，恭喜你！
			</body>
			</html>`,
	})
	return
}

// SendBindEms 发送换绑邮件
func (c *cSms) SendBindEms(ctx context.Context, req *common.SendBindEmsReq) (res *common.SendBindEmsRes, err error) {
	var (
		memberId = contexts.GetUserId(ctx)
		models   *entity.AdminMember
	)
	if memberId <= 0 {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	err = g.Model("admin_member").
		Fields("email").
		Where("id", memberId).
		Scan(&models)
	if err != nil {
		return
	}

	if models == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	if models.Email == "" {
		err = gerror.New("未绑定邮箱无需发送")
		return
	}

	err = service.SysEmsLog().Send(ctx, sysin.SendEmsInp{
		Event: consts.EmsTemplateBind,
		Email: models.Email,
	})
	return
}
