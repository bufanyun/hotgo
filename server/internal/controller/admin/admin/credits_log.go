// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.5.3
// @AutoGenerate Date 2023-04-15 15:59:58
package admin

import (
	"context"
	"hotgo/api/admin/creditslog"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/util/gconv"
)

var (
	CreditsLog = cCreditsLog{}
)

type cCreditsLog struct{}

// List 查看资产变动列表
func (c *cCreditsLog) List(ctx context.Context, req *creditslog.ListReq) (res *creditslog.ListRes, err error) {
	var in adminin.CreditsLogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	list, totalCount, err := service.AdminCreditsLog().List(ctx, in)
	if err != nil {
		return
	}

	res = new(creditslog.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Export 导出资产变动列表
func (c *cCreditsLog) Export(ctx context.Context, req *creditslog.ExportReq) (res *creditslog.ExportRes, err error) {
	var in adminin.CreditsLogListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.AdminCreditsLog().Export(ctx, in)
	return
}

// Option 获取变动状态选项
func (c *cCreditsLog) Option(_ context.Context, _ *creditslog.OptionReq) (res *creditslog.OptionRes, err error) {
	res = &creditslog.OptionRes{
		CreditType:  consts.CreditTypeOptions,
		CreditGroup: consts.CreditGroupOptions,
	}
	return
}
