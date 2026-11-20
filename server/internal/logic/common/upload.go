// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"errors"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/file"
	"hotgo/utility/format"
	"net/http"

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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, in.Url, nil)
	if err != nil {
		g.Log().Warningf(ctx, "图片转存创建请求失败 url:%s err:%+v", in.Url, err)
		return nil, gerror.New("图片链接格式不正确")
	}
	resp, err := newTransferClient(false).Do(req)
	if err != nil {
		// 底层错误带内网地址等细节，只记日志
		g.Log().Warningf(ctx, "图片转存请求失败 url:%s err:%+v", in.Url, err)
		switch {
		case errors.Is(err, errPrivateAddr):
			return nil, gerror.New("不允许访问内网地址")
		case errors.Is(err, errTooManyRedirects):
			return nil, gerror.New("重定向次数过多")
		case errors.Is(err, errBadScheme):
			return nil, gerror.New("仅支持 HTTP/HTTPS 协议的图片链接")
		default:
			return nil, gerror.New("请求图片失败，请检查链接是否可访问")
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, gerror.Newf("请求图片资源失败, StatusCode:%v", resp.StatusCode)
	}

	contentType := gstr.ToLower(resp.Header.Get("Content-Type"))
	if !gstr.HasPrefix(contentType, "image/") {
		return nil, gerror.New("资源不是图片类型")
	}

	content, err := readImageBody(resp.Body, imageSizeLimitMB())
	if err != nil {
		if errors.Is(err, errImageTooLarge) {
			return nil, err
		}
		g.Log().Warningf(ctx, "图片转存读取失败 url:%s err:%+v", in.Url, err)
		return nil, gerror.New("读取图片失败，请检查链接是否可访问")
	}
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
