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
	"hotgo/internal/service"
)

var (
	CreditsLog = cCreditsLog{}
)

type cCreditsLog struct{}

// List 查看资产变动列表
func (c *cCreditsLog) List(ctx context.Context, req *creditslog.ListReq) (res *creditslog.ListRes, err error) {
	list, totalCount, err := service.AdminCreditsLog().List(ctx, &req.CreditsLogListInp)
	if err != nil {
		return
	}

	res = new(creditslog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出资产变动列表
func (c *cCreditsLog) Export(ctx context.Context, req *creditslog.ExportReq) (res *creditslog.ExportRes, err error) {
	err = service.AdminCreditsLog().Export(ctx, &req.CreditsLogListInp)
	return
}
