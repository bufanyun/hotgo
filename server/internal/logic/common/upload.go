// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/format"
)

type sCommonUpload struct{}

func NewCommonUpload() *sCommonUpload {
	return &sCommonUpload{}
}

func init() {
	service.RegisterCommonUpload(NewCommonUpload())
}

// UploadFile 上传文件
func (s *sCommonUpload) UploadFile(ctx context.Context, file *ghttp.UploadFile) (res *sysin.AttachmentListModel, err error) {
	attachment, err := storager.DoUpload(ctx, file, consts.UploadTypeFile)
	if err != nil {
		return
	}

	attachment.FileUrl = storager.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
	res = &sysin.AttachmentListModel{
		SysAttachment: *attachment,
		SizeFormat:    format.FileSize(attachment.Size),
	}
	return
}

// UploadImage 上传图片
func (s *sCommonUpload) UploadImage(ctx context.Context, file *ghttp.UploadFile) (res *sysin.AttachmentListModel, err error) {
	attachment, err := storager.DoUpload(ctx, file, consts.UploadTypeImage)
	if err != nil {
		return
	}

	attachment.FileUrl = storager.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
	res = &sysin.AttachmentListModel{
		SysAttachment: *attachment,
		SizeFormat:    format.FileSize(attachment.Size),
	}
	return
}
