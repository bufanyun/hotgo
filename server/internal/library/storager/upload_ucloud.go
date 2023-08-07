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
	upload "github.com/ufilesdk-dev/ufile-gosdk"
)

// UCloudDrive UCloud对象存储驱动
type UCloudDrive struct {
}

// Upload 上传到UCloud对象存储
func (d *UCloudDrive) Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error) {
	if config.UCloudPath == "" {
		err = gerror.New("UCloud存储驱动必须配置存储路径!")
		return
	}

	client, err := upload.NewFileRequest(&upload.Config{
		PublicKey:       config.UCloudPublicKey,
		PrivateKey:      config.UCloudPrivateKey,
		BucketHost:      config.UCloudBucketHost,
		BucketName:      config.UCloudBucketName,
		FileHost:        config.UCloudFileHost,
		Endpoint:        config.UCloudEndpoint,
		VerifyUploadMD5: false,
	}, nil)
	if err != nil {
		return
	}

	// 流式上传本地小文件
	f2, err := file.Open()
	defer func() { _ = f2.Close() }()
	if err != nil {
		return
	}

	fullPath = GenFullPath(config.UCloudPath, gfile.Ext(file.Filename))
	err = client.IOPut(f2, fullPath, "")
	return
}
