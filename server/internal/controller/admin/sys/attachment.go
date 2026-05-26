// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/attachment"
	"hotgo/internal/service"
)

var (
	Attachment = cAttachment{}
)

type cAttachment struct{}

// Delete 删除附件
func (c *cAttachment) Delete(ctx context.Context, req *attachment.DeleteReq) (res *attachment.DeleteRes, err error) {
	err = service.SysAttachment().Delete(ctx, &req.AttachmentDeleteInp)
	return
}

// View 获取指定附件信息
func (c *cAttachment) View(ctx context.Context, req *attachment.ViewReq) (res *attachment.ViewRes, err error) {
	data, err := service.SysAttachment().View(ctx, &req.AttachmentViewInp)
	if err != nil {
		return
	}

	res = new(attachment.ViewRes)
	res.AttachmentViewModel = data
	return
}

// List 查看附件列表
func (c *cAttachment) List(ctx context.Context, req *attachment.ListReq) (res *attachment.ListRes, err error) {
	list, totalCount, err := service.SysAttachment().List(ctx, &req.AttachmentListInp)
	if err != nil {
		return
	}

	res = new(attachment.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// ClearKind 清空上传类型
func (c *cAttachment) ClearKind(ctx context.Context, req *attachment.ClearKindReq) (res *attachment.ClearKindRes, err error) {
	err = service.SysAttachment().ClearKind(ctx, &req.AttachmentClearKindInp)
	return
}
