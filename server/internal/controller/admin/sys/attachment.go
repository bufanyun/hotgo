// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/attachment"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Attachment = cAttachment{}
)

type cAttachment struct{}

// Delete 删除
func (c *cAttachment) Delete(ctx context.Context, req *attachment.DeleteReq) (res *attachment.DeleteRes, err error) {
	var in sysin.AttachmentDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysAttachment().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cAttachment) Edit(ctx context.Context, req *attachment.EditReq) (res *attachment.EditRes, err error) {
	var in sysin.AttachmentEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysAttachment().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cAttachment) MaxSort(ctx context.Context, req *attachment.MaxSortReq) (res *attachment.MaxSortRes, err error) {
	data, err := service.SysAttachment().MaxSort(ctx, sysin.AttachmentMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(attachment.MaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cAttachment) View(ctx context.Context, req *attachment.ViewReq) (res *attachment.ViewRes, err error) {
	data, err := service.SysAttachment().View(ctx, sysin.AttachmentViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(attachment.ViewRes)
	res.AttachmentViewModel = data
	return
}

// List 查看列表
func (c *cAttachment) List(ctx context.Context, req *attachment.ListReq) (res *attachment.ListRes, err error) {
	var in sysin.AttachmentListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysAttachment().List(ctx, in)
	if err != nil {
		return
	}

	res = new(attachment.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新部门状态
func (c *cAttachment) Status(ctx context.Context, req *attachment.StatusReq) (res *attachment.StatusRes, err error) {
	var in sysin.AttachmentStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysAttachment().Status(ctx, in)
	return
}
