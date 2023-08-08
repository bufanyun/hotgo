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
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// QiNiuDrive 七牛云对象存储驱动
type QiNiuDrive struct {
}

// Upload 上传到七牛云对象存储
func (d *QiNiuDrive) Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error) {
	if config.QiNiuPath == "" {
		err = gerror.New("七牛云存储驱动必须配置存储路径!")
		return
	}

	// 流式上传本地小文件
	f2, err := file.Open()
	defer func() { _ = f2.Close() }()
	if err != nil {
		return
	}

	putPolicy := storage.PutPolicy{
		Scope: config.QiNiuBucket,
	}
	token := putPolicy.UploadToken(qbox.NewMac(config.QiNiuAccessKey, config.QiNiuSecretKey))

	cfg := storage.Config{}

	// 是否使用https域名
	cfg.UseHTTPS = true

	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 空间对应的机房
	cfg.Region, err = storage.GetRegion(config.QiNiuAccessKey, config.QiNiuBucket)
	if err != nil {
		return
	}

	fullPath = GenFullPath(config.QiNiuPath, gfile.Ext(file.Filename))
	err = storage.NewFormUploader(&cfg).Put(ctx, &storage.PutRet{}, token, fullPath, f2, file.Size, &storage.PutExtra{})
	return
}
