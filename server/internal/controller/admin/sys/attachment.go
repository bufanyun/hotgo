// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/api/admin/attachment"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/input/sysin"
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

// ChooserOption 获取选择器选项
func (c *cAttachment) ChooserOption(ctx context.Context, req *attachment.ChooserOptionReq) (res *attachment.ChooserOptionRes, err error) {
	res = new(attachment.ChooserOptionRes)

	res.Drive, err = service.SysDictData().Select(ctx, &sysin.DataSelectInp{Type: "config_upload_drive"})
	if err != nil {
		return
	}

	var kinds = []attachment.KindSelect{
		{
			Label: "全部", Key: "", Value: "",
		},
		{
			Label: "图片", Key: storager.KindImg, Value: storager.KindImg, Icon: "PictureOutlined", Tag: "success",
		},
		{
			Label: "文档", Key: storager.KindDoc, Value: storager.KindDoc, Icon: "FileWordOutlined", Tag: "primary",
		},
		{
			Label: "音频", Key: storager.KindAudio, Value: storager.KindAudio, Icon: "CustomerServiceOutlined", Tag: "info",
		},
		{
			Label: "视频", Key: storager.KindVideo, Value: storager.KindVideo, Icon: "PlaySquareOutlined", Tag: "warning",
		},
		{
			Label: "压缩包", Key: storager.KindZip, Value: storager.KindZip, Icon: "FileZipOutlined", Tag: "error",
		},
		{
			Label: "其他", Key: storager.KindOther, Value: storager.KindOther, Icon: "PlusOutlined", Tag: "default",
		},
	}
	res.Kind = append(res.Kind, kinds...)
	return
}

// ClearKind 清空上传类型
func (c *cAttachment) ClearKind(ctx context.Context, req *attachment.ClearKindReq) (res *attachment.ClearKindRes, err error) {
	err = service.SysAttachment().ClearKind(ctx, &req.AttachmentClearKindInp)
	return
}
