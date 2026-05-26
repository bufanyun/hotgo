// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

// CosDrive 腾讯云cos驱动
type CosDrive struct {
}

// Upload 上传到腾讯云cos对象存储
func (d *CosDrive) Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error) {
	if config.CosPath == "" {
		err = gerror.New("COS存储驱动必须配置存储路径!")
		return
	}

	// 流式上传本地小文件
	f2, err := file.Open()
	defer func() { _ = f2.Close() }()
	if err != nil {
		return
	}

	URL, _ := url.Parse(config.CosBucketURL)
	client := cos.NewClient(&cos.BaseURL{BucketURL: URL}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.CosSecretId,
			SecretKey: config.CosSecretKey,
		},
	})

	fullPath = GenFullPath(config.CosPath, gfile.Ext(file.Filename))
	_, err = client.Object.Put(ctx, fullPath, f2, nil)
	return
}

// CreateMultipart 创建分片事件
func (d *CosDrive) CreateMultipart(ctx context.Context, in *CheckMultipartParams) (res *MultipartProgress, err error) {
	err = gerror.New("当前驱动暂不支持分片上传！")
	return
}

// UploadPart 上传分片
func (d *CosDrive) UploadPart(ctx context.Context, in *UploadPartParams) (res *UploadPartModel, err error) {
	err = gerror.New("当前驱动暂不支持分片上传！")
	return
}
