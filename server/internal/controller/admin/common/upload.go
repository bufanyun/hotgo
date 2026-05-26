// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/admin/common"
	"hotgo/internal/library/storager"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var Upload = new(cUpload)

type cUpload struct{}

// UploadFile 上传文件
func (c *cUpload) UploadFile(ctx context.Context, _ *common.UploadFileReq) (res common.UploadFileRes, err error) {
	r := g.RequestFromCtx(ctx)
	uploadType := r.Header.Get("uploadType")
	if uploadType != "default" && !validate.InSlice(storager.KindSlice, uploadType) {
		err = gerror.New("上传类型是无效的")
		return
	}

	file := r.GetUploadFile("file")
	if file == nil {
		err = gerror.New("没有找到上传的文件")
		return
	}
	return service.CommonUpload().UploadFile(ctx, uploadType, file)
}

// CheckMultipart 检查文件分片
func (c *cUpload) CheckMultipart(ctx context.Context, req *common.CheckMultipartReq) (res *common.CheckMultipartRes, err error) {
	data, err := service.CommonUpload().CheckMultipart(ctx, &req.CheckMultipartInp)
	if err != nil {
		return nil, err
	}
	res = new(common.CheckMultipartRes)
	res.CheckMultipartModel = data
	return
}

// UploadPart 上传分片
func (c *cUpload) UploadPart(ctx context.Context, req *common.UploadPartReq) (res *common.UploadPartRes, err error) {
	data, err := service.CommonUpload().UploadPart(ctx, &req.UploadPartInp)
	if err != nil {
		return nil, err
	}
	res = new(common.UploadPartRes)
	res.UploadPartModel = data
	return
}

// ImageTransferStorage 图片链接转存
func (c *cUpload) ImageTransferStorage(ctx context.Context, req *common.ImageTransferStorageReq) (res *common.ImageTransferStorageRes, err error) {
	res = new(common.ImageTransferStorageRes)
	res.ImageTransferStorageModel, err = service.CommonUpload().ImageTransferStorage(ctx, &req.ImageTransferStorageInp)
	return
}
