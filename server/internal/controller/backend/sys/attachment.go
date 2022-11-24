// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/attachment"
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
		return nil, err
	}
	if err = service.SysAttachment().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cAttachment) Edit(ctx context.Context, req *attachment.EditReq) (res *attachment.EditRes, err error) {

	var in sysin.AttachmentEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysAttachment().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// MaxSort 最大排序
func (c *cAttachment) MaxSort(ctx context.Context, req *attachment.MaxSortReq) (*attachment.MaxSortRes, error) {

	data, err := service.SysAttachment().MaxSort(ctx, sysin.AttachmentMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res attachment.MaxSortRes
	res.Sort = data.Sort
	return &res, nil
}

// View 获取指定信息
func (c *cAttachment) View(ctx context.Context, req *attachment.ViewReq) (*attachment.ViewRes, error) {

	data, err := service.SysAttachment().View(ctx, sysin.AttachmentViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res attachment.ViewRes
	res.AttachmentViewModel = data
	return &res, nil
}

// List 查看列表
func (c *cAttachment) List(ctx context.Context, req *attachment.ListReq) (*attachment.ListRes, error) {

	var (
		in  sysin.AttachmentListInp
		res attachment.ListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.SysAttachment().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// Status 更新部门状态
func (c *cAttachment) Status(ctx context.Context, req *attachment.StatusReq) (res *attachment.StatusRes, err error) {

	var in sysin.AttachmentStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysAttachment().Status(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}
