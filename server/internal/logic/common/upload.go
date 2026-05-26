// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/file"
	"hotgo/utility/format"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
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

// ImageTransferStorage 图片链接转存
func (s *sCommonUpload) ImageTransferStorage(ctx context.Context, in *sysin.ImageTransferStorageInp) (res *sysin.ImageTransferStorageModel, err error) {
	if !gstr.HasPrefix(in.Url, "http://") && !gstr.HasPrefix(in.Url, "https://") {
		return nil, gerror.New("仅支持 HTTP/HTTPS 协议的图片链接")
	}

	resp, err := g.Client().SetTimeout(time.Second*30).Get(ctx, in.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		return nil, gerror.Newf("请求图片资源失败, StatusCode:%v", resp.StatusCode)
	}

	contentType := gstr.ToLower(resp.Header.Get("Content-Type"))
	if !gstr.HasPrefix(contentType, "image/") {
		return nil, gerror.New("资源不是图片类型")
	}

	content := resp.ReadAll()
	if len(content) == 0 {
		return nil, gerror.New("图片内容为空")
	}

	res = new(sysin.ImageTransferStorageModel)
	fileHeader, err := file.NewMultipartFileHeader("its-"+grand.Letters(8)+".png", content)
	if err != nil {
		return nil, gerror.Newf("创建文件头失败：%v", err)
	}
	res.AttachmentListModel, err = s.UploadFile(ctx, storager.KindImg, &ghttp.UploadFile{FileHeader: fileHeader})
	return
}
