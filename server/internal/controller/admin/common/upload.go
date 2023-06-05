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
	"hotgo/internal/service"
)

var Upload = new(cUpload)

type cUpload struct{}

// UploadImage 上传图片
func (c *cUpload) UploadImage(ctx context.Context, _ *common.UploadImageReq) (res common.UploadImageRes, err error) {
	file := g.RequestFromCtx(ctx).GetUploadFile("file")
	if file == nil {
		err = gerror.New("没有找到上传的文件")
		return
	}
	return service.CommonUpload().UploadImage(ctx, file)
}

// UploadFile 上传附件
func (c *cUpload) UploadFile(ctx context.Context, _ *common.UploadFileReq) (res common.UploadFileRes, err error) {
	file := g.RequestFromCtx(ctx).GetUploadFile("file")
	if file == nil {
		err = gerror.New("没有找到上传的文件")
		return
	}
	return service.CommonUpload().UploadFile(ctx, file)
}
