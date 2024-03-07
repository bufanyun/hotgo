// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
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
func (s *sCommonUpload) UploadFile(ctx context.Context, uploadType string, file *ghttp.UploadFile) (res *sysin.AttachmentListModel, err error) {
	attachment, err := storager.DoUpload(ctx, uploadType, file)
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

// CheckMultipart 检查文件分片
func (s *sCommonUpload) CheckMultipart(ctx context.Context, in *sysin.CheckMultipartInp) (res *sysin.CheckMultipartModel, err error) {
	data, err := storager.CheckMultipart(ctx, in.CheckMultipartParams)
	if err != nil {
		return nil, err
	}
	res = new(sysin.CheckMultipartModel)
	res.CheckMultipartModel = data
	return
}

// UploadPart 上传分片
func (s *sCommonUpload) UploadPart(ctx context.Context, in *sysin.UploadPartInp) (res *sysin.UploadPartModel, err error) {
	data, err := storager.UploadPart(ctx, in.UploadPartParams)
	if err != nil {
		return nil, err
	}
	res = new(sysin.UploadPartModel)
	res.UploadPartModel = data
	return
}
