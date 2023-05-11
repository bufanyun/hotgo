// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/common"
	"hotgo/internal/model/input/commonin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var (
	Wechat = cWechat{}
)

type cWechat struct{}

func (c *cWechat) Authorize(ctx context.Context, req *common.WechatAuthorizeReq) (res *common.WechatAuthorizeRes, err error) {
	var in commonin.WechatAuthorizeInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	_, err = service.CommonWechat().Authorize(ctx, in)
	return
}

func (c *cWechat) AuthorizeCall(ctx context.Context, req *common.WechatAuthorizeCallReq) (res *common.WechatAuthorizeCallRes, err error) {
	var in commonin.WechatAuthorizeCallInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	_, err = service.CommonWechat().AuthorizeCall(ctx, in)
	return
}
